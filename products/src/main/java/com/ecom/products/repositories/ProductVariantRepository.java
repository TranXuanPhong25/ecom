package com.ecom.products.repositories;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.entities.ProductVariant;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface ProductVariantRepository extends JpaRepository<ProductVariant, Long> {
    @Query("SELECT new com.ecom.products.dtos.VariantDTO(v.id, v.price, v.attributes, v.isActive, v.stockQuantity, v.sku) " +
            "FROM ProductVariant v " +
            "WHERE v.productId = :productId")
    List<VariantDTO> findByProductId(Long productId);
}