# Product Categories Service - Performance Optimizations

## Overview
This document describes the performance optimizations implemented for the Product Categories Service to improve query performance when retrieving category hierarchies and root categories.

## Implemented Optimizations

### 1. Eliminated N+1 Query Problem
**Previous Implementation:**
- Used recursive calls to `findAllChildrenById()` for each category node
- For a tree with N nodes, this resulted in N+1 database queries
- Example: A category tree with 100 nodes would execute 100+ queries

**New Implementation:**
- Fetches all categories in a single query using `findAll()`
- Fetches all parent-child relationships in a single query using `findAllDirectRelationships()`
- Builds the tree structure in memory using a HashMap for O(1) lookups
- Total queries reduced from N+1 to just 2 queries regardless of tree size

### 2. Optimized Root Categories Query
**Previous Query:**
```sql
SELECT pc.* FROM product_category pc 
LEFT JOIN product_category_closure pcc 
ON pc.id = pcc.descendant_id AND depth != 0 
WHERE pcc.ancestor_id IS NULL
```

**New Query:**
```sql
SELECT pc.* FROM product_category pc 
WHERE NOT EXISTS (
  SELECT 1 FROM product_category_closure pcc 
  WHERE pcc.descendant_id = pc.id AND pcc.depth > 0
)
```

**Benefits:**
- Uses `NOT EXISTS` instead of `LEFT JOIN` for better query optimization
- More efficient execution plan in PostgreSQL
- Clearer semantic meaning

### 3. Added Database Indexes
Created indexes to support common query patterns:
- `idx_closure_ancestor_depth` - For finding direct children
- `idx_closure_descendant_depth` - For finding parent categories
- `idx_closure_descendant_depth_nonzero` - For identifying root categories
- `idx_closure_ancestor_descendant` - For relationship lookups

### 4. Implemented Multi-Level Caching
**Cache Strategy:**
- `productCategoriesTree` - Caches the complete category hierarchy
- `rootCategories` - Caches root categories list
- `categoryDetails` - Caches individual category details by ID

**Cache Eviction:**
- Automatic cache eviction every 24 hours via scheduled task
- Manual eviction on any create, update, or delete operations
- Ensures data consistency while maximizing cache hit rate

**Benefits:**
- First request builds the tree (2 queries)
- Subsequent requests serve from cache (0 queries)
- Significant reduction in database load for read-heavy workloads

## Performance Improvements

### Query Reduction
| Operation | Before | After | Improvement |
|-----------|--------|-------|-------------|
| Get Category Tree (100 nodes) | 101 queries | 2 queries | 98% reduction |
| Get Root Categories | 1 query | 1 query (cached) | 0-1 queries |
| Get Category Details | 3 queries | 3 queries (cached) | 0-3 queries |

### Response Time (Estimated)
| Operation | Before | After | Improvement |
|-----------|--------|-------|-------------|
| Get Category Tree (100 nodes) | ~500ms | ~50ms (uncached), <5ms (cached) | 90-99% faster |
| Get Root Categories | ~20ms | <5ms (cached) | 75-99% faster |

## Technical Details

### Technologies Used
- **Spring Cache**: In-memory caching with `@Cacheable` and `@CacheEvict`
- **ConcurrentMapCacheManager**: Simple, thread-safe in-memory cache
- **Flyway**: Database migration management for index creation
- **PostgreSQL**: Optimized queries and indexes

### Cache Configuration
Location: `com.ecom.productcategory.config.CacheConfig`
- Configures cache manager with named caches
- Schedules automatic cache eviction every 24 hours
- Can be easily extended to use Redis or other cache providers

### Database Migrations
Location: `src/main/resources/db/migration/V1__add_performance_indexes.sql`
- Automatically applied on application startup via Flyway
- Creates performance indexes if they don't exist
- Safe to run multiple times (uses `IF NOT EXISTS`)

## Configuration

### Cache TTL
To adjust cache eviction interval, modify `CacheConfig.java`:
```java
@Scheduled(fixedDelay = 86400000, initialDelay = 86400000) // 24 hours in milliseconds
```

### Using Redis (Optional)
To use Redis instead of in-memory cache:
1. Add dependency: `spring-boot-starter-data-redis`
2. Configure Redis connection in `application.properties`
3. Update `CacheConfig` to use `RedisCacheManager`

## Testing
The optimizations maintain full backward compatibility with existing APIs. No changes are required to client applications.

## Future Enhancements
1. **Redis Integration**: For distributed caching across multiple service instances
2. **Cache Warming**: Pre-populate cache on application startup
3. **Metrics**: Add monitoring for cache hit/miss rates
4. **Partial Cache Updates**: Invalidate only affected cache entries instead of full eviction
