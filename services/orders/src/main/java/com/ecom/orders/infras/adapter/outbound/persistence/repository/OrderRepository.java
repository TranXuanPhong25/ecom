package com.ecom.orders.infras.adapter.outbound.persistence.repository;

import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderStatus;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.Instant;
import java.util.List;
import java.util.Optional;

@Repository
public interface OrderRepository extends JpaRepository<Order, Long>, JpaSpecificationExecutor<Order> {

   List<Order> findByUserId(String userId);

   List<Order> findByStatus(OrderStatus status);

   List<Order> findByUserIdAndStatus(String userId, OrderStatus status);

   @Query("SELECT o FROM Order o LEFT JOIN FETCH o.orderItems WHERE o.id = :id")
   Optional<Order> findByIdWithItems(@Param("id") Long id);

   @Query("SELECT o FROM Order o LEFT JOIN FETCH o.orderItems WHERE o.userId = :userId")
   List<Order> findByUserIdWithItems(@Param("userId") String userId);

   @Query("SELECT COUNT(o) FROM Order o WHERE o.userId = :userId AND o.status = :status")
   Long countByUserIdAndStatus(@Param("userId") String userId, @Param("status") OrderStatus status);

   List<Order> findByCreatedAtBetween(Instant startDate, Instant endDate);

   @Query("SELECT o FROM Order o WHERE o.id IN :ids")
   List<Order> findByIdIn(@Param("ids") List<Long> ids);

   @Query("SELECT o FROM Order o WHERE o.shopId = :shopId AND o.status = :status")
   List<Order> findByShopIdAndStatus(@Param("shopId") String shopId, @Param("status") OrderStatus status);
}
