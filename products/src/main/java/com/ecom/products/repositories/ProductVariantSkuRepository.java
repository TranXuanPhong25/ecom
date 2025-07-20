package com.ecom.products.repositories;

import com.ecom.products.entities.ProductVariantSku;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ProductVariantSkuRepository extends JpaRepository<ProductVariantSku, Long> {
}