package com.ecom.products.models;

import com.ecom.products.dtos.VariantWithNameDTO;

import java.util.List;

public record GetProductVariantsResponse (
    List<VariantWithNameDTO> variants,
    List<Long> notFoundIds
) {}
