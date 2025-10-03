package com.ecom.products.dtos;

import java.math.BigDecimal;
import java.util.Map;

public class VariantWithNameDTO extends VariantDTO {
    private String name;

    public VariantWithNameDTO(Long id, String name, BigDecimal price, Map<String, String> attributes,
                              boolean isActive, Integer stockQuantity, String sku, String[] images) {
        super(id, price, attributes, isActive, stockQuantity, sku, images);
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {  
        this.name = name;
    }
   
}
