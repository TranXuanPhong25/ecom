package com.ecom.products.dtos;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class VariantImageDTO {
    private String imageUrl;
    private boolean isPrimary;
}
