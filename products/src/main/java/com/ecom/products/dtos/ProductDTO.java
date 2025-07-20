package com.ecom.products.dtos;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

@Data
public class ProductDTO {
    private Long id;

    @NotNull(message = "Product name is required")
    @NotBlank(message = "Product name cannot be empty")
    private String name;

    @NotNull(message = "Product description is required")
    @NotBlank(message = "Product description cannot be empty")
    private String description;

    @NotNull(message = "Category is required")
    private Long categoryId;
    private Long brandId;
    private boolean isActive;
}