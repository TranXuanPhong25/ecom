package com.ecom.productvariant.controller;


import com.ecom.productvariant.dto.ProductVariantDTO;
import com.ecom.productvariant.entity.ProductVariant;
import com.ecom.productvariant.service.ProductVariantService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/api/")
public class ProductVariantController {

    @Autowired
    private ProductVariantService variantService;

    @GetMapping
    public List<ProductVariant> getAllVariants() {
        return variantService.getAllVariants();
    }

    @GetMapping("/product/{id}")
    public ResponseEntity<List<ProductVariantDTO>> getVariantByProductId(@PathVariable Long id) {
        List<ProductVariant> variant = variantService.getVariantByProductId(id);
        if(variant.isEmpty()){
            return ResponseEntity.notFound().build();
        }
        List<ProductVariantDTO> productVariantDTOs = variant.stream().map(ProductVariantDTO::new).collect(Collectors.toList());
        return ResponseEntity.ok().body(productVariantDTOs);
    }

    @PostMapping
    public ProductVariant createVariant(@RequestBody ProductVariant variant) {
        return variantService.createVariant(variant);
    }

    @PutMapping("/{id}")
    public ProductVariant updateVariant(@PathVariable Long id, @RequestBody ProductVariant variant) {
        return variantService.updateVariant(id, variant);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteVariant(@PathVariable Long id) {
        variantService.deleteVariant(id);
        return ResponseEntity.noContent().build();
    }
}
