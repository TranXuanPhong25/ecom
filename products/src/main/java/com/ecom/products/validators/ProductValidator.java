package com.ecom.products.validators;

import com.ecom.products.dtos.ProductDTO;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.validation.Errors;
import org.springframework.validation.Validator;

public class ProductValidator implements Validator {

    public ProductValidator productValidator() {
        return new ProductValidator();
    }

    @Override
    public boolean supports(Class clazz) {
        return ProductDTO.class.isAssignableFrom(clazz);
    }

    @Override
    public void validate(Object target, Errors errors) {
        ProductDTO product = (ProductDTO) target;
        if (product.getName() == null || product.getName().isEmpty()) {
            errors.rejectValue("name", "name.empty", "Name cannot be empty");
        }

        if (product.getDescription() == null || product.getDescription().isEmpty()) {
            errors.rejectValue("description", "description.empty", "Description cannot be empty");
        }

        if (product.getCategoryId() == null) {
            errors.rejectValue("categoryId", "categoryId.null", "Category ID cannot be null");
        }

    }


}
