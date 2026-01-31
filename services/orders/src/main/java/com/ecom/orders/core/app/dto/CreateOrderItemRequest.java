package com.ecom.orders.core.app.dto;

import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateOrderItemRequest {

   @NotBlank(message = "Product ID is required")
   private String productId;

   @NotBlank(message = "Product name is required")
   private String productName;

   private String productSku;

   @NotBlank(message = "Image URL is required")
   private String imageUrl;

   private String variantId;
   private String variantName;

   @NotNull(message = "Original price is required")
   @Min(value = 0, message = "Original price must be >= 0")
   private Long originalPrice;

   @NotNull(message = "Sale price is required")
   @Min(value = 0, message = "Sale price must be >= 0")
   private Long salePrice;

   @NotNull(message = "Quantity is required")
   @Min(value = 1, message = "Quantity must be at least 1")
   private Integer quantity;
}
