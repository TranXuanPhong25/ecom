package com.ecom.products.repositories;

import com.ecom.products.entities.VariantImage;
import org.springframework.data.jpa.repository.JpaRepository;

public interface VariantImageRepository extends JpaRepository<VariantImage, Long> {
}