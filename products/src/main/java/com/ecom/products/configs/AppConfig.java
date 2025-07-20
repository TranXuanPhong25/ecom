package com.ecom.products.configs;

import com.ecom.products.validators.ProductValidator;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfig {
    @Bean
    public ProductValidator productValidator() {
        return new ProductValidator();
    }
}
