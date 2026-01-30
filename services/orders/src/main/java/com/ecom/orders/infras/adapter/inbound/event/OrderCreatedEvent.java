package com.ecom.orders.infras.adapter.inbound.event;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.Instant;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class OrderCreatedEvent {
   private String orderId;
   private String userId;
   private String status;
   private Double totalAmount;
   private Instant createdAt;
}