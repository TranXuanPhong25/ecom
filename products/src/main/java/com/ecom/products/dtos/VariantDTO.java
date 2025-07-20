package com.ecom.products.dtos;

import lombok.Data;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@Data
public class VariantDTO {
    private Long id;
    private Long productId;
    private BigDecimal price;
    private Map<String, String> attributes;
    private boolean isActive;
    private Integer stockQuantity;
    private List<VariantImageDTO> images;
    private String sku;
}