package com.ecom.products.models;

import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.dtos.VariantDTO;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;

import java.util.List;

public record CreateProductRequest(
        @Valid
        @NotNull(message = "Product details are required")
        ProductDTO product,
        List<VariantDTO> variants
) {}
