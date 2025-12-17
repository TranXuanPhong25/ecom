package com.ecom.productcategory.dto;

import lombok.Data;

@Data
public class ProductCategoryPathNode {
    private Integer id;
    private String name;

    public ProductCategoryPathNode(Integer id, String name) {
        this.id = id;
        this.name = name;
    }
}
