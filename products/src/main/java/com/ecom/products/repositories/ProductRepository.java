package com.ecom.products.repositories;

import com.ecom.products.entities.Product;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface ProductRepository extends JpaRepository<Product, Long> {

    Page<Product> findByShopId(UUID shopId, Pageable pageable);
}