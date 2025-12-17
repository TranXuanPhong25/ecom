// products/src/main/java/com/ecom/products/controllers/BrandController.java

package com.ecom.products.controllers;

import com.ecom.products.dtos.BrandDTO;
import com.ecom.products.services.BrandService;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequiredArgsConstructor
@RequestMapping("/api/brands")
public class BrandController {
    private final BrandService brandService;

    @GetMapping
    public List<BrandDTO> getBrands() {
        return brandService.getBrands();
    }

    @PostMapping
    public ResponseEntity<BrandDTO> createBrand(@RequestBody BrandDTO brand) {
        BrandDTO createdBrand = brandService.createBrand(brand);
        return new ResponseEntity<>(createdBrand, HttpStatus.CREATED);
    }

    @PutMapping("/{id}")
    public ResponseEntity<BrandDTO> updateBrand(@Valid @NotNull @PathVariable Long id, @RequestBody BrandDTO brand) {
        BrandDTO updatedBrand = brandService.updateBrand(id, brand);
        return updatedBrand != null ? ResponseEntity.ok(updatedBrand) : ResponseEntity.notFound().build();
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteBrand(@PathVariable Long id) {
        brandService.deleteBrand(id);
        return ResponseEntity.noContent().build();
    }
}