package com.ecom.productcategory.config;

import org.springframework.cache.CacheManager;
import org.springframework.cache.annotation.CacheEvict;
import org.springframework.cache.concurrent.ConcurrentMapCacheManager;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;

@Configuration
@EnableScheduling
public class CacheConfig {

    @Bean
    public CacheManager cacheManager() {
        return new ConcurrentMapCacheManager("productCategoriesTree", "rootCategories", "categoryDetails");
    }

    // Evict all caches every 24 hours (86400000 ms) to ensure freshness
    // This can be adjusted based on how frequently categories are updated
    @CacheEvict(value = {"productCategoriesTree", "rootCategories", "categoryDetails"}, allEntries = true)
    @Scheduled(fixedDelay = 86400000, initialDelay = 86400000)
    public void evictAllCaches() {
        // This method will be called automatically by Spring to evict caches
    }
}
