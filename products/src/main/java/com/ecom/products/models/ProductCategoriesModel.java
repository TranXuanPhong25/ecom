package com.ecom.products.models;

import lombok.Data;

import java.util.List;

@Data
public class ProductCategoriesModel {
    private List<ProductCategoryPathNode> categoryPath;
}

