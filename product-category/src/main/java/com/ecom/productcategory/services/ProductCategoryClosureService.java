package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import com.ecom.productcategory.entities.ProductCategoryClosureEntity;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface ProductCategoryClosureService {

    void deleteByCategoryIds(List<Integer> categoryIds);

    void createProductCategory(Integer parentId, Integer id);

    List<Integer> getAllDescendantById(Integer id);
}
