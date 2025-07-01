package com.ecom.productcategory.services;

import com.ecom.productcategory.entities.ProductCategoryClosureEntity;
import com.ecom.productcategory.repositories.ProductCategoryClosureRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class ProductCategoryClosureServiceImpl implements ProductCategoryClosureService {
    @Autowired
    private ProductCategoryClosureRepository productCategoryClosureRepository;
    @Transactional
    @Override
    public void deleteByCategoryIds(List<Integer> categoryIds) {
        productCategoryClosureRepository.deleteByCategoryIds(categoryIds);
    }

    @Transactional
    @Override
    public void createProductCategory(Integer parentId, Integer id) {
       if(id==null){
           throw new IllegalArgumentException("id cannot be null");
       }
       ProductCategoryClosureEntity productCategoryClosureEntity = new ProductCategoryClosureEntity(id,id,0);
       productCategoryClosureRepository.save(productCategoryClosureEntity);
       if(parentId!=null){
           productCategoryClosureRepository.createSubCategory(parentId, id);
       }
    }

    @Override
    public List<ProductCategoryClosureEntity> getProductCategoryHierachyById(Integer id) {
        return productCategoryClosureRepository.getProductCategoryHierachyById(id);
    }

}
