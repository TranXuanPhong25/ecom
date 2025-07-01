package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryDTO;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public interface ProductCategoryClosureService {

    void deleteByCategoryIds(List<Integer> categoryIds);

    void createProductCategory(Integer parentId, Integer id);
}
