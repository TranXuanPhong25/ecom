package com.ecom.products.services;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;

@Service
public class CategoryService {
    private final WebClient categoryServiceClient;

    public CategoryService(@Qualifier("ProductCategoryServiceClient") WebClient categoryServiceClient) {
        this.categoryServiceClient = categoryServiceClient;
    }

    public Mono<String> getCategoryPath(Long categoryId) {
        return categoryServiceClient.get()
                .uri("/path?id={categoryId}", categoryId)
                .retrieve()
                .bodyToMono(String.class);
    }
}
