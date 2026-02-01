package com.ecom.products.dtos;

import jakarta.validation.constraints.NotNull;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;

@NoArgsConstructor
@Data
public class VariantDTO {
    private Long id;

    @NotNull(message = "Please provide original price for the variant")
    private BigDecimal originalPrice;

    @NotNull(message = "Please provide sale price for the variant")
    private BigDecimal salePrice;

    private Map<String, String> attributes;
    private boolean isActive;

    private Integer stockQuantity;
    private List<String> images;
    private String sku;

    public VariantDTO(Long id, BigDecimal originalPrice, BigDecimal salePrice,
            Map<String, String> attributes, boolean isActive,
            Integer stockQuantity, String sku, String[] images) {
        this.id = id;
        this.originalPrice = originalPrice;
        this.salePrice = salePrice;
        this.attributes = attributes;
        this.isActive = isActive;
        this.stockQuantity = stockQuantity;
        this.sku = sku;
        this.images = List.of(images);
    }
}