package com.ecom.products.dtos;

import com.ecom.products.enums.ProductStatus;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Pattern;
import lombok.Data;

import java.time.ZonedDateTime;
import java.util.List;
import java.util.Map;

@Data
public class ProductDTO {
    private Long id;

    @Pattern(regexp = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$",
            message = "Shop ID must be a valid UUID")
    private String shopId;
    @NotNull(message = "Product name is required")
    @NotBlank(message = "Product name cannot be empty")
    private String name;

    @NotNull(message = "Product description is required")
    @NotBlank(message = "Product description cannot be empty")
    private String description;

    @NotNull(message = "Category is required")
    private Long categoryId;

    private List<String> images;

    private String categoryPath;
    private BrandDTO brand;
    private boolean isActive;

    private Map<String, String> specs;
    private List<VariantDTO> variants;
    private ProductStatus status;

    private ZonedDateTime createdAt;
    private ZonedDateTime updatedAt;
}