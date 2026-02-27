package com.ecom.orders.infras.adapter.outbound.client;

import lombok.Data;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

import java.time.Instant;

/**
 * Client để gọi Fulfillment Service REST API
 */
@Slf4j
@Component
public class FulfillmentClient {

    private final RestTemplate restTemplate;
    private final String fulfillmentServiceUrl;

    public FulfillmentClient(
            RestTemplate restTemplate,
            @Value("${fulfillment.service.url}") String fulfillmentServiceUrl) {
        this.restTemplate = restTemplate;
        this.fulfillmentServiceUrl = fulfillmentServiceUrl;
    }

    /**
     * Schedule pickup từ seller
     * Call khi seller marks order as "Ready to Ship"
     */
    public FulfillmentResponse schedulePickup(SchedulePickupRequest request) {
        String url = fulfillmentServiceUrl + "/api/fulfillment/pickup/schedule";

        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);

        HttpEntity<SchedulePickupRequest> entity = new HttpEntity<>(request, headers);

        try {
            log.info("Calling Fulfillment Service to schedule pickup for order {}", request.getOrderId());

            ResponseEntity<FulfillmentResponse> response = restTemplate.postForEntity(
                    url,
                    entity,
                    FulfillmentResponse.class);

            log.info("Pickup scheduled successfully. Package number: {}",
                    response.getBody().getPackageNumber());

            return response.getBody();

        } catch (Exception e) {
            log.error("Failed to schedule pickup for order {}: {}", request.getOrderId(), e.getMessage());
            throw new RuntimeException("Failed to schedule fulfillment pickup", e);
        }
    }

    /**
     * Get package tracking info
     */
    public PackageTrackingResponse getTracking(String packageNumber) {
        String url = fulfillmentServiceUrl + "/api/fulfillment/tracking/" + packageNumber;

        try {
            ResponseEntity<PackageTrackingResponse> response = restTemplate.getForEntity(
                    url,
                    PackageTrackingResponse.class);

            return response.getBody();

        } catch (Exception e) {
            log.error("Failed to get tracking for package {}: {}", packageNumber, e.getMessage());
            return null;
        }
    }

    // ===== DTOs =====

    @Data
    public static class SchedulePickupRequest {
        private Long orderId;
        private String shopId;
        private String pickupAddress;
        private String pickupContactName;
        private String pickupContactPhone;
        private String deliveryAddress;
        private String deliveryContactName;
        private String deliveryContactPhone;
        private Integer weightGrams;
        private String specialInstructions;
    }

    @Data
    public static class FulfillmentResponse {
        private String packageNumber;
        private Instant pickupScheduledAt;
        private Instant estimatedDelivery;
        private String message;
    }

    @Data
    public static class PackageTrackingResponse {
        private String packageNumber;
        private Long orderId;
        private String status;
        private String currentLocation;
        private Instant lastScanAt;
        private Instant estimatedDelivery;
    }
}
