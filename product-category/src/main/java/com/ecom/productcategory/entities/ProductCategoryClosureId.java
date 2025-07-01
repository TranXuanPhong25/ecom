package com.ecom.productcategory.entities;

import jakarta.persistence.Embeddable;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor

public class ProductCategoryClosureId{
    private Integer ancestorId;
    private Integer descendantId;
    public ProductCategoryClosureId(Integer ancestorId, Integer descendantId) {
        this.ancestorId = ancestorId;
        this.descendantId = descendantId;
    }
}
