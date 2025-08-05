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
    private Long productId;

    @NotNull(message = "Please provide a price for the variant")
    private BigDecimal price;

    private Map<String, String> attributes;
    private boolean isActive;

    private Integer stockQuantity;
    private List<VariantImageDTO> images;
    private String sku;

    public VariantDTO(Long id, BigDecimal price, Map<String, String> attributes, boolean isActive,
                      Integer stockQuantity, String sku) {
        this.id = id;
        this.price = price;
        this.attributes = attributes;
        this.isActive = isActive;
        this.stockQuantity = stockQuantity;
        this.sku = sku;
    }
}