package com.ecom.orders.core.app.dto;

import jakarta.validation.constraints.NotEmpty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class ConfirmOrdersRequest {

   @NotEmpty(message = "Order IDs cannot be empty")
   private List<Long> orderIds;
}
