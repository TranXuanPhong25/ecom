package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.dto.ProductCategoryPathNode;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.models.ProductCategoryUpdateModel;
import org.springframework.stereotype.Service;

import java.util.List;
@Service
public interface ProductCategoryService {
    List<ProductCategoryEntity> getALlRootProductCategories();

    List<ProductCategoryNodeDTO> getProductCategoriesTree();


    List<ProductCategoryEntity> updateProductCategories(ProductCategoryUpdateModel productCategoryUpdateModel);

    ProductCategoryDTO getProductCategoryById(Integer id) ;

    ProductCategoryEntity updateProductCategory( ProductCategoryEntity productCategoryDetails) ;

    ProductCategoryEntity createProductCategory(ProductCategoryDTO productCategoryDTO);

    void deleteProductCategory(Integer id);

    List<ProductCategoryPathNode> getCategoryPath(Integer id);
}
