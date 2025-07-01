package com.ecom.productcategory.repositories;

import com.ecom.productcategory.entities.ProductCategoryEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ProductCategoryRepository  extends JpaRepository<ProductCategoryEntity, Integer> {
}
