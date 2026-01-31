package com.ecom.orders.infras.adapter.inbound.event.listener;

import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.service.OrderService;
import com.ecom.orders.infras.adapter.inbound.event.OrderCreatedEvent;
import com.ecom.orders.infras.adapter.inbound.event.mapper.OrderEventMapper;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.support.Acknowledgment;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

@Slf4j
@Component
@RequiredArgsConstructor
public class OrderCreatedEventListener {

   private final ObjectMapper objectMapper;
   private final OrderEventMapper orderEventMapper;
   private final OrderService orderService;

   @KafkaListener(topics = "${spring.kafka.topics.order-created}", groupId = "${spring.kafka.consumer.group-id}", containerFactory = "kafkaListenerContainerFactory")
   public void handleOrderCreatedEvent(
         @Payload String message,
         @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
         @Header(KafkaHeaders.RECEIVED_PARTITION) int partition,
         @Header(KafkaHeaders.OFFSET) long offset,
         Acknowledgment acknowledgment) {
      try {
         log.info("Received message from topic: {}, partition: {}, offset: {}", topic, partition, offset);
         log.info("Message: {}", message);

         // Parse the event
         OrderCreatedEvent event = objectMapper.readValue(message, OrderCreatedEvent.class);

         // Map to domain entity and process
         processOrderCreatedEvent(event);

         // Manually acknowledge the message
         acknowledgment.acknowledge();

      } catch (Exception e) {
         log.error("Error processing order event from topic: {}, partition: {}, offset: {}",
               topic, partition, offset, e);
         // Handle error - could implement retry logic or send to DLQ
      }
   }

   private void processOrderCreatedEvent(OrderCreatedEvent event) {
      Order order = orderEventMapper.toEntity(event);

      // Save order to database or perform other business logic
      orderService.createOrder(order);
      log.info("Created order: orderId={}, userId={}, amount={}",
            order.getId(), order.getUserId(), order.getTotalAmount());
   }
}