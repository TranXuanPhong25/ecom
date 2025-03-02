package com.ecom.productvariant.service;

import com.ecom.productvariant.entity.ProductVariant;
import com.ecom.productvariant.repository.ProductVariantRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProductVariantService {

    @Autowired
    private ProductVariantRepository productVariantRepository;

    public List<ProductVariant> getAllVariants() {
        return productVariantRepository.findAll();
    }

    public List<ProductVariant> getVariantByProductId(Long id) {
        return productVariantRepository.findByProductId(id);
    }

    public ProductVariant createVariant(ProductVariant variant) {
        return productVariantRepository.save(variant);
    }

    public ProductVariant updateVariant(Long id, ProductVariant variantDetails) {
        ProductVariant variant = productVariantRepository.findById(id)
                .orElseThrow(() -> new RuntimeException("Variant not found with id " + id));

        variant.setProductId(variantDetails.getProductId());
        variant.setPrice(variantDetails.getPrice());
        variant.setQuantityInStock(variantDetails.getQuantityInStock());
        // Xóa hết attribute cũ và thêm attribute mới
        variant.getAttributes().clear();
        if (variantDetails.getAttributes() != null) {
            for (var attr : variantDetails.getAttributes()) {
                variant.addAttribute(attr);
            }
        }
        return productVariantRepository.save(variant);
    }

    public void deleteVariant(Long id) {
        productVariantRepository.deleteById(id);
    }


}
