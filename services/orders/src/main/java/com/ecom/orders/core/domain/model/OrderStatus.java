package com.ecom.orders.core.domain.model;

public enum OrderStatus {
   UNCONFIRMED,
   CONFIRMED,
   READY_TO_SHIP,
   SHIPPING,
   DELIVERED,
   CANCELLED,
   REFUNDED
}
