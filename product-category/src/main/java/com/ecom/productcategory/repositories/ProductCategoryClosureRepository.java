package com.ecom.productcategory.repositories;

import com.ecom.productcategory.entities.ProductCategoryClosureEntity;
import com.ecom.productcategory.entities.ProductCategoryClosureId;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Repository
public interface ProductCategoryClosureRepository extends JpaRepository<ProductCategoryClosureEntity, ProductCategoryClosureId> {
    @Modifying
    @Query("DELETE FROM ProductCategoryClosureEntity pcs " +
            "WHERE pcs.ancestorId IN :categoryIds OR pcs.descendantId IN :categoryIds")
    void deleteByCategoryIds(@Param("categoryIds") List<Integer> categoryIds);

    @Modifying
    @Query(value = "insert into product_category_closure(ancestor_id, descendant_id, depth) " +
            "select pcc.ancestor_id, :id,  (depth+1) " +
            "from product_category_closure pcc " +
            "where pcc.descendant_id =:parentId",
    nativeQuery = true)
    void createSubCategory(Integer parentId, Integer id);

    @Query(value = "WITH recursive cte AS (" +
            "select * from product_category_closure " +
            "where product_category_closure.ancestor_id  = :id and depth!=0 " +
            "union " +
            "select pcc.ancestor_id , pcc.descendant_id ,pcc.depth " +
            "from product_category_closure pcc " +
            "inner join cte c on c.descendant_id = pcc.ancestor_id and pcc.depth!=0 " +
            ") select * from cte",
            nativeQuery = true)
    List<ProductCategoryClosureEntity> getProductCategoryHierachyById(Integer id);

}
