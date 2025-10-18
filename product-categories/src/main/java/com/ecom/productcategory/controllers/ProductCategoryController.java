package com.ecom.productcategory.controllers;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.dto.ProductCategoryPathNode;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.services.ProductCategoryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/product-categories")
public class ProductCategoryController {
    @Autowired
    private ProductCategoryService productCategoryService;

    @GetMapping
    public List<ProductCategoryEntity> getAllProductCategories() {
        return productCategoryService.getALlRootProductCategories();
    }


    @GetMapping(value = "/path", params = "id")
    public List<ProductCategoryPathNode> getAllProductCategories(@RequestParam(value = "id") Integer id) {
        return productCategoryService.getCategoryPath(id);
    }

    @GetMapping("/hierarchy")
    public List<ProductCategoryNodeDTO> getProductCategoriesTree() {
        return productCategoryService.getProductCategoriesTree();
    }
//    @PutMapping
//    public List<ProductCategoryEntity> updateProductCategories(@RequestBody ProductCategoryUpdateModel productCategoryUpdateModel) {
//        return productCategoryService.updateProductCategories(productCategoryUpdateModel);
//    }


    @GetMapping("/{id}")
    public ResponseEntity<ProductCategoryDTO> getProductCategoryById(
            @PathVariable(value = "id") Integer id) {
        ProductCategoryDTO productCategory = productCategoryService.getProductCategoryById(id);
        return ResponseEntity.ok().body(productCategory);
    }

    @PutMapping
    public ResponseEntity<ProductCategoryEntity> updateProductCategory(
            @RequestBody ProductCategoryEntity productCategoryEntity) {
        ProductCategoryEntity updatedProductCategory =
                productCategoryService.updateProductCategory(productCategoryEntity);
        return ResponseEntity.ok(updatedProductCategory);
    }

    @PostMapping
    public ResponseEntity<ProductCategoryEntity> createProductCategory(
            @RequestBody ProductCategoryDTO productCategoryDTO) {
        ProductCategoryEntity createdProductCategory = productCategoryService.createProductCategory(productCategoryDTO);
        return ResponseEntity.status(201).body(createdProductCategory);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteProductCategory(
            @PathVariable(value = "id") Integer id) {
        productCategoryService.deleteProductCategory(id);
        return ResponseEntity.noContent().build();
    }
}
