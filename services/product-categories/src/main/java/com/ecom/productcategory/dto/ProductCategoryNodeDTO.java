package com.ecom.productcategory.dto;

import com.ecom.productcategory.entities.ProductCategoryEntity;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

import java.util.ArrayList;
import java.util.List;

@EqualsAndHashCode(callSuper = true)
@Data
@NoArgsConstructor
@AllArgsConstructor
public class ProductCategoryNodeDTO extends ProductCategoryEntity {
    private List<ProductCategoryNodeDTO> children = new ArrayList<>();
    public ProductCategoryNodeDTO(ProductCategoryEntity productCategoryEntity) {
        super(productCategoryEntity.getId(), productCategoryEntity.getName(), productCategoryEntity.getImageUrl());
    }
}
