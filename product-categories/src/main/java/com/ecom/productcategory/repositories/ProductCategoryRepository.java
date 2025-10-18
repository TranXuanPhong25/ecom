package com.ecom.productcategory.repositories;

import com.ecom.productcategory.dto.ProductCategoryPathNode;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Map;

@Repository
public interface ProductCategoryRepository  extends JpaRepository<ProductCategoryEntity, Integer> {
    @Query(value="SELECT pc.id, pc.name, pc.image_url " +
            "FROM product_category pc " +
            "JOIN product_category_closure pcc " +
            "ON pc.id = pcc.descendant_id AND depth = 1 AND pcc.ancestor_id = :id",
            nativeQuery = true)
    List<ProductCategoryEntity> findAllChildrenById(Integer id);

    @Query(value="SELECT pc.id, pc.name, pc.image_url " +
            "FROM product_category pc " +
            "JOIN product_category_closure pcc ON pc.id = pcc.ancestor_id AND pcc.descendant_id = :id AND depth = 1 " +
            "LIMIT 1",
            nativeQuery = true)
    ProductCategoryEntity findAncestorById(Integer id);

    @Query(value="SELECT pc.id, pc.name, pc.image_url " +
            "FROM product_category pc " +
            "WHERE NOT EXISTS (" +
            "  SELECT 1 FROM product_category_closure pcc " +
            "  WHERE pcc.descendant_id = pc.id AND pcc.depth > 0" +
            ")",
            nativeQuery = true)
    List<ProductCategoryEntity> findAllRootCategories();

    @Query(value = "SELECT pc.id, pc.name " +
            "FROM product_category pc " +
            "JOIN product_category_closure pcc ON pc.id = pcc.ancestor_id AND pcc.descendant_id = :id " +
            "ORDER BY depth desc",
            nativeQuery = true)
    List<ProductCategoryPathNode> getCategoryPath(Integer id);

    // Batch fetch all parent-child relationships with depth 1 (direct children)
    @Query(value = "SELECT pcc.ancestor_id as parentId, pcc.descendant_id as childId " +
            "FROM product_category_closure pcc " +
            "WHERE pcc.depth = 1",
            nativeQuery = true)
    List<Map<String, Object>> findAllDirectRelationships();
}
