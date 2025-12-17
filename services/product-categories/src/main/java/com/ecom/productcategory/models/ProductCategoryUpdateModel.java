package com.ecom.productcategory.models;

import com.ecom.productcategory.entities.ProductCategoryEntity;
import lombok.Data;

import java.util.List;
@Data
public class ProductCategoryUpdateModel {
    List<ProductCategoryEntity> productCategories;
    List<Integer> deletedProductCategoryIds;
}
