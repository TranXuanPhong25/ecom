package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderItemDTO {
   private Long id;
   private String productId;
   private String productName;
   private String productSku;
   private String imageUrl;
   private String variantId;
   private String variantName;
   private Long originalPrice;
   private Long salePrice;
   private Integer quantity;
   private Long subtotal;
}
