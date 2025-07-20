package com.ecom.products.controllers;

import com.ecom.products.dtos.PageResponse;
import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.model.CreateProductRequest;
import com.ecom.products.services.ProductService;
import com.ecom.products.validators.ProductValidator;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/products")
@RequiredArgsConstructor
public class ProductController {
    private final ProductService productService;
    private final ProductValidator productValidator;


    @PostMapping
    public ResponseEntity<?> createProduct(@RequestBody @Valid CreateProductRequest createProductRequest) {
        ProductDTO productDTO = productService.createProduct(createProductRequest);
        return new ResponseEntity<>(productDTO, HttpStatus.CREATED);
    }

    @GetMapping
    public PageResponse<ProductDTO> getProducts(Pageable pageable) {
        return new PageResponse<>(productService.getProducts(pageable));
    }

    @GetMapping("/{id}")
    public ResponseEntity<ProductDTO> getProductById(@PathVariable Long id) {
        ProductDTO productDTO = productService.getProductById(id);
        return productDTO != null ? ResponseEntity.ok(productDTO) : ResponseEntity.notFound().build();
    }

    @PutMapping("/{id}")
    public ResponseEntity<ProductDTO> updateProduct(@PathVariable Long id, @RequestBody ProductDTO productDTO) {
        ProductDTO updatedProduct = productService.updateProduct(id, productDTO);
        return updatedProduct != null ? ResponseEntity.ok(updatedProduct) : ResponseEntity.notFound().build();
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteProduct(@PathVariable Long id) {
        productService.deleteProduct(id);
        return ResponseEntity.noContent().build();
    }
}