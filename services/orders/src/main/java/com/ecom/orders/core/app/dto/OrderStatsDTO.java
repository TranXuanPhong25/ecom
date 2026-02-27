package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderStatsDTO {
   private Long totalOrders;
   private Long unconfirmedCount;
   private Long confirmedCount;
   private Long readyToShipCount;
   private Long shippingCount;
   private Long deliveredCount;
   private Long refundedCount;
   private Long cancelledCount;
   private Long totalRevenue;
   private Long pendingRevenue;
}
