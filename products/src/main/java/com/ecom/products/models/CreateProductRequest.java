package com.ecom.products.models;

import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.dtos.VariantDTO;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;

import java.util.List;
import java.util.UUID;

public record CreateProductRequest(
        @NotNull(message = "Shop id is required")
        @NotEmpty(message = "Shop id is required")
        UUID shopId,
        @Valid
        @NotNull(message = "Product details are required")
        ProductDTO product,
        List<VariantDTO> variants
) {}
