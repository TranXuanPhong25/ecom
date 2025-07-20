package com.ecom.products.controllers;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.services.ProductVariantService;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/product-variants")
@RequiredArgsConstructor
public class ProductVariantController {

    private final ProductVariantService ProductVariantService;

    @PostMapping
    public ResponseEntity<VariantDTO> createVariant(@RequestBody VariantDTO variantDTO) {
        return new ResponseEntity<>(ProductVariantService.createVariant(variantDTO), HttpStatus.CREATED);
    }

    @GetMapping("/product/{productId}")
    public List<VariantDTO> getVariantsByProductId(@PathVariable Long productId) {
        return ProductVariantService.getVariantsByProductId(productId);
    }

    @GetMapping("/{id}")
    public ResponseEntity<VariantDTO> getVariantById(@PathVariable Long id) {
        VariantDTO variantDTO = ProductVariantService.getVariantById(id);
        return variantDTO != null ? ResponseEntity.ok(variantDTO) : ResponseEntity.notFound().build();
    }

    @PutMapping("/{id}")
    public ResponseEntity<VariantDTO> updateVariant(@PathVariable Long id, @RequestBody VariantDTO variantDTO) {
        VariantDTO updatedVariant = ProductVariantService.updateVariant(id, variantDTO);
        return updatedVariant != null ? ResponseEntity.ok(updatedVariant) : ResponseEntity.notFound().build();
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteVariant(@PathVariable Long id) {
        ProductVariantService.deleteVariant(id);
        return ResponseEntity.noContent().build();
    }
}