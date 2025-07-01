package com.ecom.productcategory.repositories;

import com.ecom.productcategory.entities.ProductCategoryEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ProductCategoryRepository  extends JpaRepository<ProductCategoryEntity, Integer> {
    @Query(value="SELECT pc.id, pc.name, pc.image_url " +
            "FROM product_category pc " +
            "JOIN product_category_closure pcc " +
            "ON pc.id = pcc.descendant_id AND depth = 1 AND pcc.ancestor_id = :id",
            nativeQuery = true)
    List<ProductCategoryEntity> findAllChildrenById(Integer id);

    @Query(value="SELECT pc.id, pc.name, pc.image_url " +
            "FROM product_category pc LEFT JOIN product_category_closure pcc ON pc.id = pcc.descendant_id AND depth !=0 " +
            "WHERE pcc.ancestor_id IS NULL",
            nativeQuery = true)
    List<ProductCategoryEntity> findAllRootCategories();
}
