package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.dto.ProductCategoryPathNode;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.exceptions.ResourceNotFoundException;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import com.ecom.productcategory.repositories.ProductCategoryRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

import java.util.*;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class ProductCategoryServiceImpl implements ProductCategoryService {
    private final ProductCategoryRepository productCategoryRepository;
    private final ProductCategoryClosureService productCategoryClosureService;

    @Override
    @Cacheable(value = "rootCategories", key = "'all'")
    public List<ProductCategoryEntity> getALlRootProductCategories() {
        return productCategoryRepository.findAllRootCategories();
    }

    /**
     * Optimized tree construction using batch fetching to avoid N+1 queries
     * Fetches all categories and relationships once, then builds tree in memory
     */
    @Override
    @Cacheable(value = "productCategoriesTree", key = "'tree'")
    public List<ProductCategoryNodeDTO> getProductCategoriesTree() {
        // Fetch all categories once
        List<ProductCategoryEntity> allCategories = productCategoryRepository.findAll();
        
        // Create a map for quick lookup
        Map<Integer, ProductCategoryNodeDTO> categoryMap = allCategories.stream()
                .collect(Collectors.toMap(
                        ProductCategoryEntity::getId,
                        ProductCategoryNodeDTO::new
                ));
        
        // Fetch all parent-child relationships once
        List<Map<String, Object>> relationships = productCategoryRepository.findAllDirectRelationships();
        
        // Build the tree structure
        Set<Integer> childIds = new HashSet<>();
        for (Map<String, Object> rel : relationships) {
            Integer parentId = (Integer) rel.get("parentId");
            Integer childId = (Integer) rel.get("childId");
            
            ProductCategoryNodeDTO parent = categoryMap.get(parentId);
            ProductCategoryNodeDTO child = categoryMap.get(childId);
            
            if (parent != null && child != null) {
                parent.getChildren().add(child);
                childIds.add(childId);
            }
        }
        
        // Return only root categories (those without parents)
        return categoryMap.values().stream()
                .filter(cat -> !childIds.contains(cat.getId()))
                .collect(Collectors.toList());
    }

    @Override
    @CacheEvict(value = {"productCategoriesTree", "rootCategories", "categoryDetails"}, allEntries = true)
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
    @Cacheable(value = "categoryDetails", key = "#id")
    public ProductCategoryDTO getProductCategoryById(Integer id) {
        ProductCategoryEntity productCategoryEntity = productCategoryRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Product category not found with id: " + id));
        ProductCategoryDTO productCategoryDTO = new ProductCategoryDTO(productCategoryEntity);
        List<ProductCategoryEntity> children = productCategoryRepository.findAllChildrenById(id);
        productCategoryDTO.setChildren(children.stream().map(ProductCategoryDTO::new).collect(Collectors.toList()));

        ProductCategoryEntity parentCategory = productCategoryRepository.findAncestorById(id);
        if( parentCategory != null) {
            productCategoryDTO.setParent(parentCategory);
            productCategoryDTO.setParentId(parentCategory.getId());
        } else {
            productCategoryDTO.setParent(null);
            productCategoryDTO.setParentId(null);
        }
        return productCategoryDTO;
    }

    @Override
    @CacheEvict(value = {"productCategoriesTree", "rootCategories", "categoryDetails"}, allEntries = true)
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
        Integer parentCategoryId = productCategoryDTO.getParentId();
        productCategoryClosureService.createProductCategory(
                parentCategoryId,
                productCategoryEntity.getId()
        );
        if (productCategoryDTO.getChildren() != null && !productCategoryDTO.getChildren().isEmpty()) {
            for (ProductCategoryDTO child : productCategoryDTO.getChildren()) {
                child.setParentId(productCategoryEntity.getId());
                createProductCategory(child);
            }
        }
        return productCategoryEntity;
    }

    @Override
    @CacheEvict(value = {"productCategoriesTree", "rootCategories", "categoryDetails"}, allEntries = true)
    public ProductCategoryEntity updateProductCategory( ProductCategoryEntity productCategoryDetails) {
        Optional<ProductCategoryEntity> productCategory = productCategoryRepository.findById(productCategoryDetails.getId());
        if( !productCategory.isPresent()) {
            throw new ResourceNotFoundException("Product category not found with id: " + productCategoryDetails.getId());
        }
        productCategory.get().setName(productCategoryDetails.getName());
        productCategory.get().setImageUrl(productCategoryDetails.getImageUrl());

        return productCategoryRepository.save(productCategory.get());
    }
    
    @Override
    @CacheEvict(value = {"productCategoriesTree", "rootCategories", "categoryDetails"}, allEntries = true)
    public void deleteProductCategory(Integer id) {
        if (!productCategoryRepository.existsById(id)) {
            throw new ResourceNotFoundException("Product category not found with id: " + id);
        }
        // Get all descendant categories first
        List<Integer> productCategoriesNeedDeleted = productCategoryClosureService.getAllDescendantById(id);
        if (productCategoriesNeedDeleted.isEmpty()) {
            productCategoriesNeedDeleted.add(id);
        } else {
            productCategoriesNeedDeleted.addFirst(id); // Add the current category ID to the list
        }
        productCategoryRepository.deleteAllById(productCategoriesNeedDeleted);
        productCategoryClosureService.deleteByCategoryIds(productCategoriesNeedDeleted);
    }

    @Override
    public List<ProductCategoryPathNode> getCategoryPath(Integer id) {
        List<ProductCategoryPathNode> productCategoryEntities = productCategoryRepository.getCategoryPath(id);
        if (productCategoryEntities.isEmpty()) {
            throw new ResourceNotFoundException("Product category not found with id: " + id);
        }
        return productCategoryEntities;
    }
}
