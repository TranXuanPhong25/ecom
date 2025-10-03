package com.ecom.products.repositories;

import com.ecom.products.entities.ProductVariant;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;
import java.util.Optional;

public interface ProductVariantRepository extends JpaRepository<ProductVariant, Long> {
    @Query("SELECT v FROM ProductVariant v WHERE v.productId = :productId")
    List<ProductVariant> findByProductId(Long productId);

    @Query("SELECT v FROM ProductVariant v JOIN FETCH v.product p WHERE v.productId = :productId")
    List<ProductVariant> findByProductIdWithProduct(@Param("productId") Long productId);

    @Query("SELECT v FROM ProductVariant v JOIN FETCH v.product p WHERE v.id = :id")
    Optional<ProductVariant> findByIdWithProduct(@Param("id") Long id);

    @Query("SELECT v FROM ProductVariant v JOIN FETCH v.product p WHERE v.id IN :ids")
    List<ProductVariant> findAllByIdWithProduct(@Param("ids") List<Long> ids);
}