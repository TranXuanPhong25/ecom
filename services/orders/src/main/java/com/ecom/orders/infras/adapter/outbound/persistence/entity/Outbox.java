package com.ecom.orders.infras.adapter.outbound.persistence.entity;

import java.time.Instant;
import java.util.Map;
import java.util.UUID;

import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.type.SqlTypes;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

// Entity
@Entity
@Table(name = "outboxes")
public class Outbox {

   @Id
   @GeneratedValue
   private UUID id;

   @Column(name = "aggregate_type", nullable = false)
   private String aggregateType;

   @Column(name = "aggregate_id", nullable = false)
   private String aggregateId;

   @Column(name = "type", nullable = false)
   private String type;

   @JdbcTypeCode(SqlTypes.JSON)
   @Column(name = "payload", columnDefinition = "jsonb", nullable = false)
   private Map<String, String> payload;

   @Column(name = "timestamp", nullable = false)
   private Instant timestamp;

   @Column(name = "tracing_span_id")
   private String tracingSpanId;

   public Outbox(
         String aggregateType,
         String aggregateId,
         String type,
         Map<String, String> payload) {
      this.id = UUID.randomUUID();
      this.aggregateType = aggregateType;
      this.aggregateId = aggregateId;
      this.type = type;
      this.payload = payload;
      this.timestamp = Instant.now();
      this.tracingSpanId = getCurrentTraceId(); // từ distributed tracing
   }

   private String getCurrentTraceId() {
      // Integration với tracing system (Jaeger, Zipkin, etc.)
      return null; // hoặc lấy từ MDC/trace context
   }

   // Getters only (immutable)
}
