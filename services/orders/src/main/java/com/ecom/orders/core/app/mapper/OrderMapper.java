package com.ecom.orders.core.app.mapper;

import com.ecom.orders.core.app.dto.CreateOrderItemRequest;
import com.ecom.orders.core.app.dto.CreateOrderRequest;
import com.ecom.orders.core.app.dto.OrderDTO;
import com.ecom.orders.core.app.dto.OrderItemDTO;
import com.ecom.orders.core.app.dto.OrderListItemDTO;
import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderItem;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Map;
import java.util.UUID;
import java.util.stream.Collectors;

@Component
public class OrderMapper {

      public Order toEntity(CreateOrderRequest request) {
            Order order = Order.builder()
                        .orderNumber(generateOrderNumber())
                        .userId(request.getUserId())
                        .shopId(request.getShopId())
                        .recipientName(request.getRecipientName())
                        .recipientPhone(request.getRecipientPhone())
                        .deliveryAddress(request.getDeliveryAddress())
                        .paymentMethod(request.getPaymentMethod())
                        .shippingMethod(request.getShippingMethod())
                        .shippingFee(request.getShippingFee() != null ? request.getShippingFee() : 0L)
                        .discount(request.getDiscount() != null ? request.getDiscount() : Map.of())
                        .customerNote(request.getCustomerNote())
                        .build();

            List<OrderItem> items = request.getItems().stream()
                        .map(this::toItemEntity)
                        .collect(Collectors.toList());

            items.forEach(order::addOrderItem);
            order.calculateTotalAmount();

            return order;
      }

      public OrderItem toItemEntity(CreateOrderItemRequest request) {
            return OrderItem.builder()
                        .productId(request.getProductId())
                        .productName(request.getProductName())
                        .productSku(request.getProductSku())
                        .imageUrl(request.getImageUrl())
                        .variantId(request.getVariantId())
                        .variantName(request.getVariantName())
                        .originalPrice(request.getOriginalPrice())
                        .salePrice(request.getSalePrice())
                        .quantity(request.getQuantity())
                        .build();
      }

      public OrderDTO toDTO(Order order) {
            return OrderDTO.builder()
                        .id(order.getId())
                        .orderNumber(order.getOrderNumber())
                        .userId(order.getUserId())
                        .shopId(order.getShopId())
                        .recipientName(order.getRecipientName())
                        .recipientPhone(order.getRecipientPhone())
                        .deliveryAddress(order.getDeliveryAddress())
                        .status(order.getStatus().name())
                        .paymentMethod(order.getPaymentMethod())
                        .paymentStatus(order.getPaymentStatus())
                        .paidAt(order.getPaidAt())
                        .subtotal(order.getSubtotal())
                        .shippingFee(order.getShippingFee())
                        .discount(order.getDiscount())
                        .totalAmount(order.getTotalAmount())
                        .shippingMethod(order.getShippingMethod())
                        .shippingProvider(order.getShippingProvider())
                        .trackingNumber(order.getTrackingNumber())
                        .estimatedDelivery(order.getEstimatedDelivery())
                        .actualDelivery(order.getActualDelivery())
                        .customerNote(order.getCustomerNote())
                        .sellerNote(order.getSellerNote())
                        .cancelReason(order.getCancelReason())
                        .confirmedAt(order.getConfirmedAt())
                        .processingAt(order.getProcessingAt())
                        .shippedAt(order.getShippedAt())
                        .deliveredAt(order.getDeliveredAt())
                        .completedAt(order.getCompletedAt())
                        .cancelledAt(order.getCancelledAt())
                        .createdAt(order.getCreatedAt())
                        .updatedAt(order.getUpdatedAt())
                        .orderItems(toItemDTOs(order.getOrderItems()))
                        .build();
      }

      public OrderItemDTO toItemDTO(OrderItem item) {
            return OrderItemDTO.builder()
                        .id(item.getId())
                        .productId(item.getProductId())
                        .productName(item.getProductName())
                        .productSku(item.getProductSku())
                        .imageUrl(item.getImageUrl())
                        .variantId(item.getVariantId())
                        .variantName(item.getVariantName())
                        .originalPrice(item.getOriginalPrice())
                        .salePrice(item.getSalePrice())
                        .quantity(item.getQuantity())
                        .subtotal(item.getSubtotal())
                        .build();
      }

      public List<OrderItemDTO> toItemDTOs(List<OrderItem> items) {
            return items.stream()
                        .map(this::toItemDTO)
                        .collect(Collectors.toList());
      }

      public List<OrderDTO> toDTOs(List<Order> orders) {
            return orders.stream()
                        .map(this::toDTO)
                        .collect(Collectors.toList());
      }

      public OrderListItemDTO toListItemDTO(Order order) {
            List<OrderItem> items = order.getOrderItems() != null ? order.getOrderItems() : List.of();
            OrderItem firstItem = items.isEmpty() ? null : items.get(0);

            return OrderListItemDTO.builder()
                        .id(order.getId())
                        .orderNumber(order.getOrderNumber())
                        .status(order.getStatus().name())
                        .paymentStatus(order.getPaymentStatus())
                        .totalAmount(order.getTotalAmount())
                        .itemCount(items.size())
                        .firstItemName(firstItem != null ? firstItem.getProductName() : null)
                        .firstItemImage(firstItem != null ? firstItem.getImageUrl() : null)
                        .createdAt(order.getCreatedAt())
                        .build();
      }

      public List<OrderListItemDTO> toListItemDTOs(List<Order> orders) {
            return orders.stream()
                        .map(this::toListItemDTO)
                        .collect(Collectors.toList());
      }

      private String generateOrderNumber() {
            return "ORD-" + System.currentTimeMillis() + "-"
                        + UUID.randomUUID().toString().substring(0, 8).toUpperCase();
      }
}
