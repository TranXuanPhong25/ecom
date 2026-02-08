package com.ecom.orders.core.domain.model;

import com.ecom.orders.infras.adapter.outbound.persistence.converter.JsonBConverter;
import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.type.SqlTypes;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Entity
@Table(name = "orders")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class Order {

   @Id
   private Long id;

   @Column(name = "order_number", unique = true, nullable = false, length = 50)
   private String orderNumber;

   @Column(name = "user_id", nullable = false)
   private String userId;

   @Column(name = "shop_id", nullable = false)
   private String shopId;

   // Shipping information
   @Column(name = "recipient_name", nullable = false)
   private String recipientName;

   @Column(name = "recipient_phone", nullable = false, length = 20)
   private String recipientPhone;

   @Column(name = "delivery_address", nullable = false, columnDefinition = "TEXT")
   private String deliveryAddress;

   // Order status
   @Column(name = "status", nullable = false, length = 50)
   @Enumerated(EnumType.STRING)
   @Builder.Default
   private OrderStatus status = OrderStatus.UNCONFIRMED;

   // Payment information
   @Column(name = "payment_method", nullable = false, length = 50)
   private String paymentMethod;

   @Column(name = "payment_status", nullable = false, length = 50)
   @Builder.Default
   private String paymentStatus = "UNPAID";

   @Column(name = "paid_at")
   private Instant paidAt;

   // Pricing
   @Column(name = "subtotal", nullable = false)
   @Builder.Default
   private Long subtotal = 0L;

   @Column(name = "shipping_fee", nullable = false)
   @Builder.Default
   private Long shippingFee = 0L;

   @Column(name = "discount", columnDefinition = "jsonb")
   @JdbcTypeCode(SqlTypes.JSON)
   @Builder.Default
   private Map<String, Object> discount = Map.of();

   @Column(name = "total_amount", nullable = false)
   @Builder.Default
   private Long totalAmount = 0L;

   // Shipping
   @Column(name = "shipping_method", length = 50)
   private String shippingMethod;

   @Column(name = "shipping_provider", length = 100)
   private String shippingProvider;

   @Column(name = "tracking_number", length = 100)
   private String trackingNumber;

   @Column(name = "estimated_delivery")
   private Instant estimatedDelivery;

   @Column(name = "actual_delivery")
   private Instant actualDelivery;

   // Notes
   @Column(name = "customer_note", columnDefinition = "TEXT")
   private String customerNote;

   @Column(name = "seller_note", columnDefinition = "TEXT")
   private String sellerNote;

   @Column(name = "cancel_reason", columnDefinition = "TEXT")
   private String cancelReason;

   // Timestamps
   @Column(name = "confirmed_at")
   private Instant confirmedAt;

   @Column(name = "processing_at")
   private Instant processingAt;

   @Column(name = "shipped_at")
   private Instant shippedAt;

   @Column(name = "delivered_at")
   private Instant deliveredAt;

   @Column(name = "completed_at")
   private Instant completedAt;

   @Column(name = "cancelled_at")
   private Instant cancelledAt;

   @Column(name = "created_at", nullable = false, updatable = false)
   @Builder.Default
   private Instant createdAt = Instant.now();

   @Column(name = "updated_at", nullable = false)
   @Builder.Default
   private Instant updatedAt = Instant.now();

   @OneToMany(mappedBy = "order", cascade = CascadeType.ALL, orphanRemoval = true, fetch = FetchType.LAZY)
   @Builder.Default
   private List<OrderItem> orderItems = new ArrayList<>();

   @PreUpdate
   public void preUpdate() {
      this.updatedAt = Instant.now();
   }

   // Helper methods for managing bidirectional relationship
   public void addOrderItem(OrderItem orderItem) {
      orderItems.add(orderItem);
      orderItem.setOrder(this);
   }

   public void removeOrderItem(OrderItem orderItem) {
      orderItems.remove(orderItem);
      orderItem.setOrder(null);
   }

   public void calculateTotalAmount() {
      this.subtotal = orderItems.stream()
            .mapToLong(OrderItem::getSubtotal)
            .sum();
      this.totalAmount = this.subtotal + this.shippingFee;
   }
}
