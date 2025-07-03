package com.ecom.productcategory.dto;

import com.ecom.productcategory.entities.ProductCategoryEntity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

import java.util.List;

@EqualsAndHashCode(callSuper = true)
@Data
@AllArgsConstructor
@NoArgsConstructor
public class ProductCategoryDTO extends ProductCategoryEntity {
    private ProductCategoryEntity parent;
    private Integer parentId;
    private List<ProductCategoryEntity> children;
    public ProductCategoryDTO(ProductCategoryEntity productCategoryEntity) {
        super(productCategoryEntity.getId(), productCategoryEntity.getName(), productCategoryEntity.getImageUrl());
    }
}
