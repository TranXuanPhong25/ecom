package com.ecom.orders.core.app.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ReadyToShipResponse {
    
    private Long orderId;
    private String orderNumber;
    private String packageNumber;
    private String pickupScheduledAt;
    private String estimatedDelivery;
    private String message;
}
