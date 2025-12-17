package com.ecom.products.controllers;

import com.ecom.products.models.GetProductVariantsResponse;
import com.ecom.products.services.ProductVariantService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
@RestController
@RequestMapping("/api/product-variants")
@RequiredArgsConstructor
public class ProductVariantController {
   private final ProductVariantService productVariantService;
   @GetMapping
   public GetProductVariantsResponse getProductVariantByIds(@RequestParam List<Long> ids) {
       return productVariantService.getVariantByIds(ids);
   }
}
