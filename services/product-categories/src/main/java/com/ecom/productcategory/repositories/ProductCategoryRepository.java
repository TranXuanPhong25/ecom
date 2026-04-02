package com.ecom.productcategory.repositories;

import com.ecom.productcategory.dto.ProductCategoryPathNode;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.models.ProductCategoryDetailProjection;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface ProductCategoryRepository extends JpaRepository<ProductCategoryEntity, Integer> {
        @Query(value = "SELECT pc.id, pc.name, pc.image_url " +
                        "FROM product_category pc " +
                        "JOIN product_category_closure pcc " +
                        "ON pc.id = pcc.descendant_id AND depth = 1 AND pcc.ancestor_id = :id", nativeQuery = true)
        List<ProductCategoryEntity> findAllChildrenById(Integer id);

        @Query(value = "SELECT pc.id, pc.name, pc.image_url " +
                        "FROM product_category pc " +
                        "JOIN product_category_closure pcc ON pc.id = pcc.ancestor_id AND pcc.descendant_id = :id AND depth = 1 "
                        +
                        "LIMIT 1", nativeQuery = true)
        ProductCategoryEntity findAncestorById(Integer id);

        @Query(value = "SELECT pc.id, pc.name, pc.image_url " +
                        "FROM product_category pc LEFT JOIN product_category_closure pcc ON pc.id = pcc.descendant_id AND depth !=0 "
                        +
                        "WHERE pcc.ancestor_id IS NULL", nativeQuery = true)
        List<ProductCategoryEntity> findAllRootCategories();

        @Query(value = "SELECT pc.id, pc.name " +
                        "FROM product_category pc " +
                        "JOIN product_category_closure pcc ON pc.id = pcc.ancestor_id AND pcc.descendant_id = :id " +
                        "ORDER BY depth desc", nativeQuery = true)
        List<ProductCategoryPathNode> getCategoryPath(Integer id);

        @Query(value = "select " +
                        "pc.id as id,pc.name as name,  pc.image_url as imageUrl," +
                        "c.id as childrenId, c.name as childrenName, c.image_url as childrenImageUrl " +
                        "from product_category pc " +
                        "JOIN product_category_closure pcc ON pc.id = pcc.ancestor_id and pcc.depth =1 " +
                        "JOIN product_category c ON c.id = pcc.descendant_id " +
                        "where pc.id=1;", nativeQuery = true)
        List<ProductCategoryDetailProjection> findCategoryWithChildren(Integer id);

}
