package com.ecom.products.dtos;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Pattern;
import lombok.Data;

import java.util.UUID;

@Data
public class ProductDTO {
    private Long id;

    @Pattern(regexp = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$",
            message = "Shop ID must be a valid UUID")
    private UUID shopId;
    @NotNull(message = "Product name is required")
    @NotBlank(message = "Product name cannot be empty")
    private String name;

    @NotNull(message = "Product description is required")
    @NotBlank(message = "Product description cannot be empty")
    private String description;

    @NotNull(message = "Category is required")
    private Long categoryId;

    private BrandDTO brand;
    private boolean isActive;
}