package com.ecom.products.services;

import com.ecom.products.models.ProductCategoryPathNode;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.core.ParameterizedTypeReference;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;

import java.util.List;

@Service
public class CategoryService {
    private final WebClient categoryServiceClient;

    public CategoryService(@Qualifier("ProductCategoryServiceClient") WebClient categoryServiceClient) {
        this.categoryServiceClient = categoryServiceClient;
    }

    public Mono<List<ProductCategoryPathNode>> getCategoryPath(Long categoryId) {
        return categoryServiceClient.get()
                .uri("/path?id={categoryId}", categoryId)
                .retrieve()
                .bodyToMono(new ParameterizedTypeReference<List<ProductCategoryPathNode>>() {});
    }
}
