package com.ecom.orders.core.domain.model;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Table(name = "order_items")
@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderItem {

   @Id
   @GeneratedValue(strategy = GenerationType.IDENTITY)
   private Long id;

   @ManyToOne(fetch = FetchType.LAZY)
   @JoinColumn(name = "order_id", nullable = false)
   private Order order;

   // Product information (snapshot tại thời điểm đặt hàng)
   @Column(name = "product_id", nullable = false)
   private String productId;

   @Column(name = "product_name", nullable = false, length = 500)
   private String productName;

   @Column(name = "product_sku", length = 100)
   private String productSku;

   @Column(name = "image_url", nullable = false, columnDefinition = "TEXT")
   private String imageUrl;

   // Variant information
   @Column(name = "variant_id")
   private String variantId;

   @Column(name = "variant_name")
   private String variantName;

   // Pricing
   @Column(name = "original_price", nullable = false)
   private Long originalPrice;

   @Column(name = "sale_price", nullable = false)
   private Long salePrice;

   @Column(name = "quantity", nullable = false)
   private Integer quantity;

   // Business logic methods
   public Long getSubtotal() {
      return salePrice * quantity;
   }

   @Override
   public boolean equals(Object o) {
      if (this == o)
         return true;
      if (!(o instanceof OrderItem))
         return false;
      OrderItem that = (OrderItem) o;
      return id != null && id.equals(that.getId());
   }

   @Override
   public int hashCode() {
      return getClass().hashCode();
   }
}
