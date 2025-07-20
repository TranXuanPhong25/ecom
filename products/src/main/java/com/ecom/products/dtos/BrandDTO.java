package com.ecom.products.dtos;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class BrandDTO {
    private Long id;
    private String name;
    private String description;
}
