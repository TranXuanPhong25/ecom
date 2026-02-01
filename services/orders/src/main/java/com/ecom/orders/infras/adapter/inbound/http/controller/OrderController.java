package com.ecom.orders.infras.adapter.inbound.http.controller;

import com.ecom.orders.core.app.dto.ConfirmOrdersRequest;
import com.ecom.orders.core.app.dto.ConfirmOrdersResponse;
import com.ecom.orders.core.app.dto.OrderDTO;
import com.ecom.orders.core.app.dto.OrderListItemDTO;
import com.ecom.orders.core.app.dto.PageResponse;
import com.ecom.orders.core.app.mapper.OrderMapper;
import com.ecom.orders.core.domain.model.Order;
import com.ecom.orders.core.domain.model.OrderStatus;
import com.ecom.orders.core.domain.service.OrderService;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Slf4j
@RestController
@RequestMapping("/api/orders")
@RequiredArgsConstructor
public class OrderController {

   private final OrderService orderService;
   private final OrderMapper orderMapper;

   @GetMapping("/{id}")
   public ResponseEntity<OrderDTO> getOrderById(@PathVariable Long id) {
      log.info("Fetching order by ID: {}", id);

      return orderService.findById(id)
            .map(orderMapper::toDTO)
            .map(ResponseEntity::ok)
            .orElse(ResponseEntity.notFound().build());
   }

   /**
    * Tìm kiếm orders với các filter tùy chọn
    * 
    * @param userId - lấy từ header X-User-Id
    * @param shopId - lọc theo shop (optional)
    * @param status - lọc theo status (optional)
    */
   @GetMapping
   public PageResponse<OrderListItemDTO> searchOrders(
         @RequestHeader(value = "X-User-Id", required = true) String userId,
         @RequestParam(required = false) String shopId,
         @RequestParam(required = false) String status,
         Pageable pageable) {
      log.info("Searching orders - userId: {}, shopId: {}, status: {}", userId, shopId, status);

      try {
         OrderStatus orderStatus = null;
         if (status != null && !status.isBlank()) {
            orderStatus = OrderStatus.valueOf(status.toUpperCase());
         }

         Page<Order> pageOrders = orderService.searchOrders(userId, shopId, orderStatus, pageable);
         Page<OrderListItemDTO> response = pageOrders.map(orderMapper::toListItemDTO);
         return new PageResponse<>(response);
      } catch (IllegalArgumentException e) {
         log.error("Invalid order status: {}", status);
         return new PageResponse<>(null);
      }
   }

   // TODO: seller only
   @PatchMapping("/{id}/status")
   public ResponseEntity<OrderDTO> updateOrderStatus(
         @PathVariable Long id,
         @RequestParam String status) {
      log.info("Updating order {} status to: {}", id, status);

      try {
         OrderStatus newStatus = OrderStatus.valueOf(status.toUpperCase());
         Order updatedOrder = orderService.updateOrderStatus(id, newStatus);
         OrderDTO response = orderMapper.toDTO(updatedOrder);

         return ResponseEntity.ok(response);
      } catch (IllegalArgumentException e) {
         log.error("Invalid order status: {}", status);
         return ResponseEntity.badRequest().build();
      } catch (RuntimeException e) {
         log.error("Error updating order status", e);
         return ResponseEntity.notFound().build();
      }
   }

   /**
    * Xác nhận đơn hàng hàng loạt (Seller only)
    * Chỉ xác nhận được các đơn hàng có status = CREATED
    */
   @PostMapping("/confirm")
   public ResponseEntity<ConfirmOrdersResponse> confirmOrders(
         @Valid @RequestBody ConfirmOrdersRequest request) {
      log.info("Confirming {} orders", request.getOrderIds().size());

      List<Order> confirmedOrders = orderService.confirmOrders(request.getOrderIds());

      // Tìm các order không tìm thấy hoặc không đủ điều kiện
      Set<Long> confirmedIds = confirmedOrders.stream()
            .filter(o -> o.getStatus() == OrderStatus.CONFIRMED)
            .map(Order::getId)
            .collect(Collectors.toSet());

      List<ConfirmOrdersResponse.FailedOrder> failedOrders = new ArrayList<>();
      for (Long orderId : request.getOrderIds()) {
         if (!confirmedIds.contains(orderId)) {
            // Kiểm tra lý do thất bại
            Order order = confirmedOrders.stream()
                  .filter(o -> o.getId().equals(orderId))
                  .findFirst()
                  .orElse(null);

            String reason = order == null
                  ? "Order not found"
                  : "Order status is " + order.getStatus() + ", only CREATED orders can be confirmed";

            failedOrders.add(ConfirmOrdersResponse.FailedOrder.builder()
                  .orderId(orderId)
                  .reason(reason)
                  .build());
         }
      }

      List<OrderDTO> confirmedDTOs = confirmedOrders.stream()
            .filter(o -> o.getStatus() == OrderStatus.CONFIRMED)
            .map(orderMapper::toDTO)
            .collect(Collectors.toList());

      ConfirmOrdersResponse response = ConfirmOrdersResponse.builder()
            .totalRequested(request.getOrderIds().size())
            .successCount(confirmedDTOs.size())
            .failedCount(failedOrders.size())
            .confirmedOrders(confirmedDTOs)
            .failedOrders(failedOrders)
            .build();

      return ResponseEntity.ok(response);
   }
}
