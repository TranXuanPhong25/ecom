package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.exceptions.ResourceNotFoundException;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import com.ecom.productcategory.repositories.ProductCategoryRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.crossstore.ChangeSetPersister;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProductCategoryServiceImpl implements ProductCategoryService {
    @Autowired
    ProductCategoryRepository productCategoryRepository;
    @Autowired
    ProductCategoryClosureService productCategoryClosureService;

    @Override
    public List<ProductCategoryEntity> getAllProductCategories() {
        return productCategoryRepository.findAll();
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
        return productCategoryDTO;
    }

    @Override
    public ProductCategoryEntity createProductCategory(ProductCategoryDTO productCategoryDTO) {
        if (productCategoryDTO == null || productCategoryDTO.getName() == null || productCategoryDTO.getName().isEmpty()) {
            throw new IllegalArgumentException("Product category name cannot be null or empty");
        }
        ProductCategoryEntity productCategoryEntity = new ProductCategoryEntity();
        productCategoryEntity.setName(productCategoryDTO.getName());
        productCategoryEntity.setImageUrl(productCategoryDTO.getImageUrl());
        productCategoryEntity = productCategoryRepository.save(productCategoryEntity);
        Integer parentCategoryId = productCategoryDTO.getParentCategory() != null ? productCategoryDTO.getParentCategory().getId() : null;
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
        // Check if the product category exists
        if (!productCategoryRepository.existsById(id)) {
            throw new ResourceNotFoundException("Product category not found with id: " + id);
        }
        productCategoryRepository.deleteById(id);
    }
}
