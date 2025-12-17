package com.ecom.products.models;

import java.util.List;


public record DeleteProductsPayload (
     List<Long> ids
){}
