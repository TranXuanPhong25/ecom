package com.ecom.productvariant.entity;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(
        name = "product_variants",
        indexes = {
    @Index(name = "idx_product_id", columnList = "product_id"),
})
public class ProductVariant {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private Long productId;

    private BigDecimal price;

    private Integer quantityInStock;

    private Integer soldQuantity;

    @OneToMany(mappedBy = "variant",cascade = CascadeType.ALL, orphanRemoval = true,fetch = FetchType.EAGER)
    private List<ProductVariantAttribute> attributes = new ArrayList<>();


    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getProductId() {
        return productId;
    }

    public void setProductId(Long productId) {
        this.productId = productId;
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

    public List<ProductVariantAttribute> getAttributes() {
        return attributes;
    }

    public void setAttributes(List<ProductVariantAttribute> attributes) {
        this.attributes.clear();
        if(attributes != null){
            this.attributes.addAll(attributes);
        }
    }

    // Phương thức hỗ trợ thêm/xóa attribute
    public void addAttribute(ProductVariantAttribute attribute) {
        attributes.add(attribute);
        attribute.setVariant(this);
    }

    public void removeAttribute(ProductVariantAttribute attribute) {
        attributes.remove(attribute);
        attribute.setVariant(null);
    }

    public Integer getSoldQuantity() {
        return soldQuantity;
    }

    public void setSoldQuantity(Integer soldQuantity) {
        this.soldQuantity = soldQuantity;
    }
}
