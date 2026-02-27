package com.ecom.orders.infras.adapter.inbound.event.listener;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.support.Acknowledgment;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import com.ecom.orders.core.domain.model.OrderStatus;
import com.ecom.orders.infras.adapter.inbound.event.FulfillmentEvent;
import com.ecom.orders.infras.adapter.outbound.persistence.repository.OrderRepository;
import com.fasterxml.jackson.databind.ObjectMapper;

import lombok.extern.slf4j.Slf4j;

/**
 * Kafka listener để nhận events từ Fulfillment Service
 */
@Slf4j
@Component
public class FulfillmentEventListener {

    private OrderRepository orderRepository;
    private ObjectMapper snakeCaseMapper;

    FulfillmentEventListener(
            OrderRepository orderRepository,
            @Qualifier("snakeCaseObjectMapper") ObjectMapper snakeCaseMapper) {
        this.orderRepository = orderRepository;
        this.snakeCaseMapper = snakeCaseMapper;
    }

    /**
     * Khi package được picked up từ seller
     */
    @KafkaListener(topics = "fulfillment.picked_up", groupId = "${spring.kafka.consumer.group-id}", containerFactory = "kafkaListenerContainerFactory")
    public void onPackagePickedUp(
            @Payload String message,
            @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
            @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
            @Header(KafkaHeaders.OFFSET) long offset,
            Acknowledgment acknowledgment) {
        log.info("Received message from topic: {}, partition: {}, offset: {}", topic, partition, offset);
        try {
            FulfillmentEvent event = snakeCaseMapper.readValue(message, FulfillmentEvent.class);
            orderRepository.findById(event.getOrderId()).ifPresent(order -> {
                order.setStatus(OrderStatus.SHIPPING);
                orderRepository.save(order);
                log.info("✅ Order {} status updated to SHIPPING", order.getId());
            });
            acknowledgment.acknowledge();
        } catch (Exception e) {
            log.error("Error processing fulfillment event from topic: {}, partition: {}, offset: {}",
                    topic, partition, offset, e);
        }
    }

    /**
     * Khi package được delivered
     */
    @KafkaListener(topics = "fulfillment.delivered", groupId = "${spring.kafka.consumer.group-id}", containerFactory = "kafkaListenerContainerFactory")
    public void onPackageDelivered(
            @Payload String message,
            @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
            @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
            @Header(KafkaHeaders.OFFSET) long offset,
            Acknowledgment acknowledgment) {
        log.info("Received message from topic: {}, partition: {}, offset: {}", topic, partition, offset);
        try {
            FulfillmentEvent event = snakeCaseMapper.readValue(message, FulfillmentEvent.class);
            orderRepository.findById(event.getOrderId()).ifPresent(order -> {
                order.setStatus(OrderStatus.DELIVERED);
                order.setDeliveredAt(event.getDeliveredAt());
                orderRepository.save(order);
                log.info("✅ Order {} status updated to DELIVERED", order.getId());
            });
            acknowledgment.acknowledge();
        } catch (Exception e) {
            log.error("Error processing fulfillment event from topic: {}, partition: {}, offset: {}",
                    topic, partition, offset, e);
        }
    }

    /**
     * Khi package đang trong quá trình vận chuyển
     */
    @KafkaListener(topics = "fulfillment.in_transit", groupId = "${spring.kafka.consumer.group-id}", containerFactory = "kafkaListenerContainerFactory")
    public void onPackageInTransit(
            @Payload String message,
            @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
            @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
            @Header(KafkaHeaders.OFFSET) long offset,
            Acknowledgment acknowledgment) {
        log.info("Received message from topic: {}, partition: {}, offset: {}", topic, partition, offset);
        try {
            // FulfillmentEvent event = snakeCaseMapper.readValue(message,
            // FulfillmentEvent.class);
            // orderRepository.findById(event.getOrderId()).ifPresent(order -> {
            // order.setStatus(OrderStatus.);
            // orderRepository.save(order);
            // log.info("✅ Order {} status updated to IN_TRANSIT", order.getId());
            // });
            // log.info("Received in_transit event for order {}, but no status update
            // needed", event.getOrderId());
            acknowledgment.acknowledge();
        } catch (Exception e) {
            log.error("Error processing fulfillment event from topic: {}, partition: {}, offset: {}",
                    topic, partition, offset, e);
        }
    }
}
