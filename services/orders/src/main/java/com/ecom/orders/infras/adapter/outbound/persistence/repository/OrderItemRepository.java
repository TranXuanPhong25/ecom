package com.ecom.orders.infras.adapter.outbound.persistence.repository;

import com.ecom.orders.core.domain.model.OrderItem;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface OrderItemRepository extends JpaRepository<OrderItem, Long> {

   List<OrderItem> findByOrderId(Long orderId);

   List<OrderItem> findByProductId(String productId);

   @Query("SELECT oi FROM OrderItem oi WHERE oi.order.userId = :userId")
   List<OrderItem> findByUserId(@Param("userId") String userId);

   @Query("SELECT SUM(oi.quantity) FROM OrderItem oi WHERE oi.productId = :productId")
   Long getTotalQuantitySoldByProduct(@Param("productId") String productId);
}
