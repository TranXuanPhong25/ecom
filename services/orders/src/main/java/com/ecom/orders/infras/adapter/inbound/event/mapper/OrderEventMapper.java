package com.ecom.orders.infras.adapter.inbound.event.mapper;

import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderItem;
import com.ecom.orders.core.domain.model.OrderStatus;
import com.ecom.orders.infras.adapter.inbound.event.OrderCreatedEvent;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Component;

import java.util.Collections;
import java.util.List;
import java.util.Map;

@Slf4j
@Component
public class OrderEventMapper {

   @Qualifier("camelCaseObjectMapper")
   private ObjectMapper camelCaseMapper;

   OrderEventMapper(@Qualifier("camelCaseObjectMapper") ObjectMapper camelCaseMapper) {
      this.camelCaseMapper = camelCaseMapper;
   }

   public Order toEntity(OrderCreatedEvent event) {
      Order order = Order.builder()
            .id(event.getId())
            .orderNumber(event.getOrderNumber()) // Dùng orderNumber làm business key
            .userId(event.getUserId())
            .shopId(event.getShopId())
            .recipientName(event.getRecipientName())
            .recipientPhone(event.getRecipientPhone())
            .deliveryAddress(event.getDeliveryAddress())
            .status(event.getStatus() != null ? event.getStatus() : OrderStatus.UNCONFIRMED)
            .paymentMethod(event.getPaymentMethod())
            .paymentStatus(event.getPaymentStatus() != null ? event.getPaymentStatus() : "UNPAID")
            .paidAt(event.getPaidAt())
            .subtotal(event.getSubtotal() != null ? event.getSubtotal() : 0L)
            .shippingFee(event.getShippingFee() != null ? event.getShippingFee() : 0L)
            .discount(parseDiscount(event.getDiscount()))
            .totalAmount(event.getTotalAmount() != null ? event.getTotalAmount() : 0L)
            .shippingMethod(event.getShippingMethod())
            .shippingProvider(event.getShippingProvider())
            .trackingNumber(event.getTrackingNumber())
            .estimatedDelivery(event.getEstimatedDelivery())
            .actualDelivery(event.getActualDelivery())
            .customerNote(event.getCustomerNote())
            .sellerNote(event.getSellerNote())
            .cancelReason(event.getCancelReason())
            .confirmedAt(event.getConfirmedAt())
            .processingAt(event.getProcessingAt())
            .shippedAt(event.getShippedAt())
            .deliveredAt(event.getDeliveredAt())
            .completedAt(event.getCompletedAt())
            .cancelledAt(event.getCancelledAt())
            .createdAt(event.getCreatedAt())
            .updatedAt(event.getUpdatedAt())
            .build();

      // Parse and add order items
      List<OrderItem> items = parseOrderItems(event.getItems());
      items.forEach(order::addOrderItem);

      return order;
   }

   private Map<String, Object> parseDiscount(String discountJson) {
      if (discountJson == null || discountJson.isBlank()) {
         return Map.of();
      }
      try {
         return camelCaseMapper.readValue(discountJson, new TypeReference<>() {
         });
      } catch (JsonProcessingException e) {
         log.warn("Failed to parse discount JSON: {}", discountJson, e);
         return Map.of();
      }
   }

   private List<OrderItem> parseOrderItems(String itemsJson) {
      if (itemsJson == null || itemsJson.isBlank()) {
         return Collections.emptyList();
      }
      try {
         return camelCaseMapper.readValue(itemsJson, new TypeReference<>() {
         });
      } catch (JsonProcessingException e) {
         log.warn("Failed to parse order items JSON: {}", itemsJson, e);
         return Collections.emptyList();
      }
   }
}
