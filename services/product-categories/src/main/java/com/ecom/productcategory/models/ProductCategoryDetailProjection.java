package com.ecom.productcategory.models;

public interface ProductCategoryDetailProjection {
   Integer getId();

   String getName();

   String getImageUrl();

   Integer getChildrenId();

   String getChildrenName();

   String getChildrenImageUrl();
}