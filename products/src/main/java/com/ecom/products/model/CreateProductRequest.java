package com.ecom.products.model;

import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.dtos.VariantDTO;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;
import org.springframework.validation.annotation.Validated;

import java.util.List;

public record CreateProductRequest(
        @Valid
        @NotNull(message = "Product details cannot be null")
        ProductDTO product,
        List<VariantDTO> variants
) {}
