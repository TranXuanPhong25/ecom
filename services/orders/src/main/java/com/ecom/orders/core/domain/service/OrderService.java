package com.ecom.orders.core.domain.service;

import com.ecom.orders.core.app.dto.OrderStatsDTO;
import com.ecom.orders.core.app.dto.ReadyToShipRequest;
import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderStatus;
import com.ecom.orders.infras.adapter.outbound.client.FulfillmentClient;
import com.ecom.orders.infras.adapter.outbound.persistence.repository.OrderRepository;
import jakarta.persistence.criteria.Predicate;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.Instant;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Optional;

@Slf4j
@Service
@RequiredArgsConstructor
public class OrderService {

   private final OrderRepository orderRepository;
   private final FulfillmentClient fulfillmentClient;

   @Transactional
   public Order createOrder(Order order) {
      log.info("Creating order for user: {}", order.getUserId());
      order.setStatus(OrderStatus.UNCONFIRMED);
      // Calculate total amount from order items
      order.calculateTotalAmount();

      // Set bidirectional relationship
      order.getOrderItems().forEach(item -> item.setOrder(order));

      Order savedOrder = orderRepository.save(order);
      log.info("Order created with ID: {}", savedOrder.getId());

      return savedOrder;
   }

   @Transactional(readOnly = true)
   public Optional<Order> findById(Long id) {
      return orderRepository.findByIdWithItems(id);
   }

   @Transactional(readOnly = true)
   public List<Order> findByUserId(String userId) {
      return orderRepository.findByUserIdWithItems(userId);
   }

   @Transactional(readOnly = true)
   public List<Order> findByStatus(OrderStatus status) {
      return orderRepository.findByStatus(status);
   }

   @Transactional
   public Order updateOrderStatus(Long orderId, OrderStatus newStatus) {
      Order order = orderRepository.findById(orderId)
            .orElseThrow(() -> new RuntimeException("Order not found: " + orderId));

      log.info("Updating order {} status from {} to {}", orderId, order.getStatus(), newStatus);
      order.setStatus(newStatus);

      return orderRepository.save(order);
   }

   @Transactional
   public List<Order> confirmOrders(List<Long> orderIds) {
      List<Order> orders = orderRepository.findByIdIn(orderIds);

      for (Order order : orders) {
         if (order.getStatus() == OrderStatus.UNCONFIRMED) {
            order.setStatus(OrderStatus.CONFIRMED);
            order.setConfirmedAt(java.time.Instant.now());
         }
      }

      return orderRepository.saveAll(orders);
   }

   @Transactional(readOnly = true)
   public List<Order> findByShopIdAndStatus(String shopId, OrderStatus status) {
      return orderRepository.findByShopIdAndStatus(shopId, status);
   }

   /**
    * Tìm kiếm orders với các filter tùy chọn
    */
   @Transactional(readOnly = true)
   public Page<Order> searchOrders(String userId, String shopId, OrderStatus status, Pageable pageable) {
      Specification<Order> spec = (root, query, cb) -> {
         List<Predicate> predicates = new ArrayList<>();

         if (userId != null && !userId.isBlank()) {
            predicates.add(cb.equal(root.get("userId"), userId));
         }
         if (shopId != null && !shopId.isBlank()) {
            predicates.add(cb.equal(root.get("shopId"), shopId));
         }
         if (status != null) {
            predicates.add(cb.equal(root.get("status"), status));
         }

         // Order by createdAt DESC
         query.orderBy(cb.desc(root.get("createdAt")));

         return cb.and(predicates.toArray(new Predicate[0]));
      };

      return orderRepository.findAll(spec, pageable);
   }

   @Transactional(readOnly = true)
   public OrderStatsDTO getOrderStatsByShopId(String shopId) {
      log.info("Getting order stats for shop: {}", shopId);

      Long totalOrders = orderRepository.countByShopId(shopId);
      Long unconfirmedCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.UNCONFIRMED);
      Long confirmedCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.CONFIRMED);
      Long readyToShipCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.READY_TO_SHIP);
      Long shippingCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.SHIPPING);
      Long deliveredCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.DELIVERED);
      Long cancelledCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.CANCELLED);
      Long refundedCount = orderRepository.countByShopIdAndStatus(shopId, OrderStatus.REFUNDED);
      // Total revenue from completed orders
      Long totalRevenue = orderRepository.sumTotalAmountByShopIdAndStatusIn(
            shopId, Arrays.asList(OrderStatus.DELIVERED));

      // Pending revenue from non-cancelled orders
      Long pendingRevenue = orderRepository.sumTotalAmountByShopIdAndStatusIn(
            shopId, Arrays.asList(
                  OrderStatus.UNCONFIRMED,
                  OrderStatus.CONFIRMED,
                  OrderStatus.SHIPPING,
                  OrderStatus.DELIVERED));

      return OrderStatsDTO.builder()
            .totalOrders(totalOrders != null ? totalOrders : 0L)
            .unconfirmedCount(unconfirmedCount != null ? unconfirmedCount : 0L)
            .confirmedCount(confirmedCount != null ? confirmedCount : 0L)
            .readyToShipCount(readyToShipCount != null ? readyToShipCount : 0L)
            .shippingCount(shippingCount != null ? shippingCount : 0L)
            .deliveredCount(deliveredCount != null ? deliveredCount : 0L)
            .refundedCount(refundedCount != null ? refundedCount : 0L)
            .cancelledCount(cancelledCount != null ? cancelledCount : 0L)
            .totalRevenue(totalRevenue != null ? totalRevenue : 0L)
            .pendingRevenue(pendingRevenue != null ? pendingRevenue : 0L)
            .build();
   }

   /**
    * Mark order as ready to ship và schedule pickup với Fulfillment Service
    * Pickup info sẽ được lấy từ shop data (không cần input từ seller)
    */
   @Transactional
   public FulfillmentClient.FulfillmentResponse markReadyToShip(Long orderId, ReadyToShipRequest request) {
      log.info("Marking order {} as ready to ship", orderId);

      Order order = orderRepository.findByIdWithItems(orderId)
            .orElseThrow(() -> new RuntimeException("Order not found: " + orderId));

      // Validate order status
      if (order.getStatus() != OrderStatus.CONFIRMED) {
         throw new RuntimeException(
               "Order must be CONFIRMED before ready to ship. Current status: " + order.getStatus());
      }

      // Call Fulfillment Service to schedule pickup
      // TODO: Get shop pickup info from Shop Service
      FulfillmentClient.SchedulePickupRequest pickupRequest = new FulfillmentClient.SchedulePickupRequest();
      pickupRequest.setOrderId(order.getId());
      pickupRequest.setShopId(order.getShopId());
      pickupRequest.setPickupAddress("Shop Address - " + order.getShopId()); // TODO: Get from Shop Service
      pickupRequest.setPickupContactName("Shop Owner"); // TODO: Get from Shop Service
      pickupRequest.setPickupContactPhone("0900000000"); // TODO: Get from Shop Service
      pickupRequest.setDeliveryAddress(order.getDeliveryAddress());
      pickupRequest.setDeliveryContactName(order.getRecipientName());
      pickupRequest.setDeliveryContactPhone(order.getRecipientPhone());
      pickupRequest.setWeightGrams(1000); // Default weight
      pickupRequest.setSpecialInstructions(request.getSpecialInstructions());

      FulfillmentClient.FulfillmentResponse response = fulfillmentClient.schedulePickup(pickupRequest);

      // Update order status
      order.setStatus(OrderStatus.READY_TO_SHIP);
      order.setShippedAt(Instant.now());
      order.setTrackingNumber(response.getPackageNumber());
      order.setEstimatedDelivery(response.getEstimatedDelivery());
      orderRepository.save(order);

      log.info("Order {} marked as ready to ship. Package: {}", orderId, response.getPackageNumber());

      return response;
   }
}
