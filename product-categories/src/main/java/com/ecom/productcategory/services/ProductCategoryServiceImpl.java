package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.exceptions.ResourceNotFoundException;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import com.ecom.productcategory.repositories.ProductCategoryRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class ProductCategoryServiceImpl implements ProductCategoryService {
    private final ProductCategoryRepository productCategoryRepository;
    private final ProductCategoryClosureService productCategoryClosureService;

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
}
