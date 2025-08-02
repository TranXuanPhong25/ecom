package com.ecom.products.repositories;

import com.ecom.products.entities.VariantImage;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Collection;

public interface VariantImageRepository extends JpaRepository<VariantImage, Long> {
    Collection<Object> findVariantImagesByVariantId(Long variantId);
}