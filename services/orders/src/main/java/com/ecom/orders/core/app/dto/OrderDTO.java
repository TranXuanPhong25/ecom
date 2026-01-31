package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.Instant;
import java.util.List;
import java.util.Map;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderDTO {
   private Long id;
   private String orderNumber;
   private String userId;
   private String shopId;

   // Shipping information
   private String recipientName;
   private String recipientPhone;
   private String deliveryAddress;

   // Status
   private String status;
   private String paymentMethod;
   private String paymentStatus;
   private Instant paidAt;

   // Pricing
   private Long subtotal;
   private Long shippingFee;
   private Map<String, Object> discount;
   private Long totalAmount;

   // Shipping details
   private String shippingMethod;
   private String shippingProvider;
   private String trackingNumber;
   private Instant estimatedDelivery;
   private Instant actualDelivery;

   // Notes
   private String customerNote;
   private String sellerNote;
   private String cancelReason;

   // Timestamps
   private Instant confirmedAt;
   private Instant processingAt;
   private Instant shippedAt;
   private Instant deliveredAt;
   private Instant completedAt;
   private Instant cancelledAt;
   private Instant createdAt;
   private Instant updatedAt;

   private List<OrderItemDTO> orderItems;
}
