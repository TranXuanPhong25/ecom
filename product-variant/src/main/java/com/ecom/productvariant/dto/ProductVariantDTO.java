package com.ecom.productvariant.dto;

import com.ecom.productvariant.entity.ProductVariant;
import com.ecom.productvariant.entity.ProductVariantAttribute;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
public class ProductVariantDTO {
    private Long id;
    private BigDecimal price;
    private Integer quantityInStock;
    private Integer soldQuantity;
    private Map<String, String> attributes; // Flatten attributes

    public ProductVariantDTO(ProductVariant variant) {
        this.id = variant.getId();
        this.price = variant.getPrice();
        this.quantityInStock = variant.getQuantityInStock();
        this.soldQuantity = variant.getSoldQuantity();
        this.attributes = variant.getAttributes().stream()
                .collect(Collectors.toMap(ProductVariantAttribute::getAttributeName, ProductVariantAttribute::getAttributeValue));
    }

    public void setAttributes(List<ProductVariantAttribute> attributes) {
        this.attributes = attributes.stream()
                .collect(Collectors.toMap(ProductVariantAttribute::getAttributeName, ProductVariantAttribute::getAttributeValue));
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public BigDecimal getPrice() {
        return price;
    }

    public void setPrice(BigDecimal price) {
        this.price = price;
    }

    public Integer getQuantityInStock() {
        return quantityInStock;
    }

    public void setQuantityInStock(Integer quantityInStock) {
        this.quantityInStock = quantityInStock;
    }

    public Integer getSoldQuantity() {
        return soldQuantity;
    }

    public void setSoldQuantity(Integer soldQuantity) {
        this.soldQuantity = soldQuantity;
    }

    public Map<String, String> getAttributes() {
        return attributes;
    }
}
