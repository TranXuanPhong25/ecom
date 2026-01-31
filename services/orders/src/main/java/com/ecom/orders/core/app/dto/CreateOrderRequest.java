package com.ecom.orders.core.app.dto;

import jakarta.validation.Valid;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotEmpty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.Map;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateOrderRequest {

   @NotBlank(message = "User ID is required")
   private String userId;

   @NotBlank(message = "Shop ID is required")
   private String shopId;

   // Shipping information
   @NotBlank(message = "Recipient name is required")
   private String recipientName;

   @NotBlank(message = "Recipient phone is required")
   private String recipientPhone;

   @NotBlank(message = "Delivery address is required")
   private String deliveryAddress;

   // Payment
   @NotBlank(message = "Payment method is required")
   private String paymentMethod;

   // Shipping
   private String shippingMethod;
   private Long shippingFee;

   // Discount
   private Map<String, Object> discount;

   // Notes
   private String customerNote;

   @NotEmpty(message = "Order items cannot be empty")
   @Valid
   private List<CreateOrderItemRequest> items;
}
