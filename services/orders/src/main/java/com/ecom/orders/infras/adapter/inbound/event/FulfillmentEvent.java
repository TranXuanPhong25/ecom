package com.ecom.orders.infras.adapter.inbound.event;

import java.time.Instant;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class FulfillmentEvent {
   private String eventType;
   private String packageNumber;
   private Long orderId;
   private String shopId;
   private Instant pickedUpAt;
   private Instant deliveredAt;
   private String currentLocation;

}
