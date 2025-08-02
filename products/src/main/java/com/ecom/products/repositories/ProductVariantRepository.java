package com.ecom.products.repositories;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.entities.ProductVariant;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface ProductVariantRepository extends JpaRepository<ProductVariant, Long> {
    List<VariantDTO> findByProductId(Long productId);
}