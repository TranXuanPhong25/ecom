package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ConfirmOrdersResponse {

   private int totalRequested;
   private int successCount;
   private int failedCount;
   private List<OrderDTO> confirmedOrders;
   private List<FailedOrder> failedOrders;

   @Data
   @Builder
   @NoArgsConstructor
   @AllArgsConstructor
   public static class FailedOrder {
      private Long orderId;
      private String reason;
   }
}
