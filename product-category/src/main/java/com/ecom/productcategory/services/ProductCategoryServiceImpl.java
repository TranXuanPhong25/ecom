package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.entities.ProductCategoryClosureEntity;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.exceptions.ResourceNotFoundException;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import com.ecom.productcategory.repositories.ProductCategoryRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.crossstore.ChangeSetPersister;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import java.util.*;
import java.util.stream.Collectors;

@Service
public class ProductCategoryServiceImpl implements ProductCategoryService {
    @Autowired
    ProductCategoryRepository productCategoryRepository;
    @Autowired
    ProductCategoryClosureService productCategoryClosureService;

    @Override
    public List<ProductCategoryEntity> getALlRootProductCategories() {
        return productCategoryRepository.findAllRootCategories();
    }

    private void constructTree(ProductCategoryNodeDTO category) {
        List<ProductCategoryEntity> childEntities = productCategoryRepository.findAllChildrenById(category.getId());
        if (!childEntities.isEmpty()) {
            List<ProductCategoryNodeDTO> children = childEntities.stream()
                                                                .map(ProductCategoryNodeDTO::new)
                                                                .toList();
            category.setChildren(children);
            for (ProductCategoryNodeDTO child : children) {
                constructTree(child);
            }
        }
    }

    @Override
    public List<ProductCategoryNodeDTO> getProductCategoriesTree() {
        List<ProductCategoryEntity> rootCategories = productCategoryRepository.findAllRootCategories();
        List<ProductCategoryNodeDTO> productCategoryDTOs = rootCategories.stream()
                                                                        .map(ProductCategoryNodeDTO::new)
                                                                        .collect(Collectors.toList());
        for (ProductCategoryNodeDTO rootCategory : productCategoryDTOs) {
            constructTree(rootCategory);
        }
        return productCategoryDTOs;
    }

    @Override
    public List<ProductCategoryEntity> updateProductCategories(ProductCategoryUpdateModel productCategoryUpdateModel) {
        List<Integer> deletedProductCategoryIds = productCategoryUpdateModel.getDeletedProductCategoryIds();
        if (deletedProductCategoryIds != null && !deletedProductCategoryIds.isEmpty()) {
            productCategoryRepository.deleteAllById(deletedProductCategoryIds);
            productCategoryClosureService.deleteByCategoryIds(deletedProductCategoryIds);
        }

        productCategoryRepository.saveAll(productCategoryUpdateModel.getProductCategories());
        return productCategoryRepository.findAll();
    }

    @Override
    public ProductCategoryDTO getProductCategoryById(Integer id) {
        ProductCategoryEntity productCategoryEntity = productCategoryRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Product category not found with id: " + id));
        ProductCategoryDTO productCategoryDTO = new ProductCategoryDTO(productCategoryEntity);
        List<ProductCategoryEntity> children = productCategoryRepository.findAllChildrenById(id);
        productCategoryDTO.setChildren(children);
        return productCategoryDTO;
    }

    @Override
    public ProductCategoryEntity createProductCategory(ProductCategoryDTO productCategoryDTO) {
        if (productCategoryDTO == null
                || productCategoryDTO.getName() == null
                || productCategoryDTO.getName().isEmpty()
        ) {
            throw new IllegalArgumentException("Product category name cannot be null or empty");
        }
        ProductCategoryEntity productCategoryEntity = new ProductCategoryEntity();
        productCategoryEntity.setName(productCategoryDTO.getName());
        productCategoryEntity.setImageUrl(productCategoryDTO.getImageUrl());
        productCategoryEntity = productCategoryRepository.save(productCategoryEntity);
        Integer parentCategoryId = productCategoryDTO.getParent() != null
                ? productCategoryDTO.getParent().getId()
                : null;
        productCategoryClosureService.createProductCategory(
                parentCategoryId,
                productCategoryEntity.getId()
        );
        return productCategoryEntity;
    }

    @Override
    public ProductCategoryEntity updateProductCategory( ProductCategoryEntity productCategoryDetails) {
        ProductCategoryEntity productCategory = getProductCategoryById(productCategoryDetails.getId());

        productCategory.setName(productCategoryDetails.getName());
        productCategory.setImageUrl(productCategoryDetails.getImageUrl());

        return productCategoryRepository.save(productCategory);
    }
    @Override
    public void deleteProductCategory(Integer id) {
        if (!productCategoryRepository.existsById(id)) {
            throw new ResourceNotFoundException("Product category not found with id: " + id);
        }
        productCategoryRepository.deleteById(id);
        // Delete the closure entries for this category
        List<ProductCategoryEntity> closureEntities = productCategoryRepository.findAllChildrenById(id);
        List<Integer> categoryIds = closureEntities.stream()
                .map(ProductCategoryEntity::getId)
                .collect(Collectors.toList());
        if (categoryIds.isEmpty()) {
            throw new ResourceNotFoundException("No closure entries found for category with id: " + id);
        }
        productCategoryClosureService.deleteByCategoryIds(categoryIds);
    }
}
