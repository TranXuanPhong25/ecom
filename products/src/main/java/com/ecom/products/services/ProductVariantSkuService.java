package com.ecom.products.services;

import com.ecom.products.entities.ProductVariantSku;
import com.ecom.products.repositories.ProductVariantSkuRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ProductVariantSkuService {
    private final ProductVariantSkuRepository productVariantSkuRepository;

    public ProductVariantSku createProductVariantSku(ProductVariantSku productVariantSku) {
        return productVariantSkuRepository.save(productVariantSku);
    }
}
