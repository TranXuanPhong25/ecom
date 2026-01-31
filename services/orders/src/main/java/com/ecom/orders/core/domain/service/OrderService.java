package com.ecom.orders.core.domain.service;

import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderStatus;
import com.ecom.orders.infras.adapter.outbound.persistence.repository.OrderRepository;
import jakarta.persistence.criteria.Predicate;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Slf4j
@Service
@RequiredArgsConstructor
public class OrderService {

   private final OrderRepository orderRepository;

   @Transactional
   public Order createOrder(Order order) {
      log.info("Creating order for user: {}", order.getUserId());

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
         if (order.getStatus() == OrderStatus.CREATED) {
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
   public List<Order> searchOrders(String userId, String shopId, OrderStatus status) {
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

      return orderRepository.findAll(spec);
   }
}
