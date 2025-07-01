package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import org.springframework.data.crossstore.ChangeSetPersister;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public interface ProductCategoryService {
    List<ProductCategoryEntity> getAllProductCategories();

    List<ProductCategoryEntity> updateProductCategories(ProductCategoryUpdateModel productCategoryUpdateModel);

    ProductCategoryDTO getProductCategoryById(Integer id) ;

    ProductCategoryEntity updateProductCategory( ProductCategoryEntity productCategoryDetails) ;

    ProductCategoryEntity createProductCategory(ProductCategoryDTO productCategoryDTO);

    void deleteProductCategory(Integer id);
}
