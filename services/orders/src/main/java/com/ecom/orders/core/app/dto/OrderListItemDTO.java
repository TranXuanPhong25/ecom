package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.Instant;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderListItemDTO {
   private Long id;
   private String orderNumber;
   private String status;
   private String paymentStatus;
   private Long totalAmount;
   private Integer itemCount;
   private String firstItemName;
   private String firstItemImage;
   private Instant createdAt;
}
