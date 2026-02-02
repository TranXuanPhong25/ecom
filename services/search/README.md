# Search Service - Hexagonal Architecture

A dedicated search service for product search in the Shopiew e-commerce platform, built with **Hexagonal Architecture** (Clean Architecture) principles.

## Architecture Overview

This service follows the **Hexagonal Architecture** pattern (also known as Ports and Adapters):

```
┌─────────────────────────────────────────────────────────┐
│                    Adapters (IN)                         │
│              HTTP Handler (Echo)                         │
└──────────────────┬──────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────┐
│                  Ports (IN)                              │
│            SearchServicePort Interface                   │
└──────────────────┬──────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────┐
│              Service Layer (Core)                        │
│           Business Logic / Use Cases                     │
└──────────────────┬──────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────┐
│                  Ports (OUT)                             │
│        ProductRepositoryPort Interface                   │
└──────────────────┬──────────────────────────────────────┘
                   │
┌──────────────────▼──────────────────────────────────────┐
│                 Adapters (OUT)                           │
│          PostgreSQL Repository                           │
└─────────────────────────────────────────────────────────┘
```

### Directory Structure

```
search-service/
├── cmd/
│   └── server.go              # Application entry point (DI wiring)
├── internal/
│   ├── adapter/
│   │   ├── handler/           # HTTP handlers (Input Adapter)
│   │   │   ├── search_handler.go
│   │   │   └── routes.go
│   │   └── storage/           # PostgreSQL repository (Output Adapter)
│   │       └── postgres_product_repository.go
│   ├── core/
│   │   ├── entity/            # Domain models (pure business objects)
│   │   │   └── product.go
│   │   ├── dto/               # Data Transfer Objects
│   │   │   └── search_dto.go
│   │   └── port/              # Interfaces (contracts)
│   │       ├── in/            # Input ports (what service provides)
│   │       │   └── search_service_port.go
│   │       └── out/           # Output ports (what service needs)
│   │           └── product_repository_port.go
│   ├── service/               # Business logic implementation
│   │   └── search_service.go
│   └── config/                # Configuration
│       ├── config.go
│       └── database.go
└── validators/                # Input validation
    └── search_validator.go
```

## Features

- ✅ **PostgreSQL Full-Text Search** using tsvector
- ✅ **Read-only** service (no data modification)
- ✅ **Hexagonal Architecture** (easy to test and swap implementations)
- ✅ **Clean separation** of concerns (Domain, Ports, Adapters)
- ✅ **Dependency Injection** for loose coupling
- ✅ **Pagination** support with relevance ranking
- ✅ **Health check** endpoint

## API Endpoints

### Search Products
```http
GET /api/search?q=laptop&page=1&limit=20
```

**Query Parameters:**
- `q` (required): Search keyword
- `page` (optional): Page number (default: 1)
- `limit` (optional): Results per page (default: 20, max: 100)

**Response:**
```json
{
  "success": true,
  "message": "Products found successfully",
  "data": {
    "products": [
      {
        "id": 1,
        "name": "Gaming Laptop",
        "description": "High-performance laptop",
        "price": 1299.99,
        "category_id": 5,
        "stock_quantity": 10,
        "image_url": "https://..."
      }
    ],
    "total": 42,
    "page": 1,
    "limit": 20
  }
}
```

### Get Single Product
```http
GET /api/search/products/:id
```

**Response:**
```json
{
  "success": true,
  "message": "Product retrieved successfully",
  "data": {
    "id": 1,
    "name": "Gaming Laptop",
    "description": "High-performance laptop",
    "price": 1299.99,
    "category_id": 5,
    "stock_quantity": 10,
    "image_url": "https://..."
  }
}
```

### Health Check
```http
GET /api/search/health
```

## PostgreSQL Full-Text Search

The service uses PostgreSQL's built-in full-text search with `tsvector` and `tsquery`:

```sql
SELECT id, name, description, price, category_id, stock_quantity, image_url
FROM products
WHERE to_tsvector('english', name || ' ' || COALESCE(description, '')) 
      @@ plainto_tsquery('english', $1)
ORDER BY ts_rank(to_tsvector('english', name || ' ' || description), 
                  plainto_tsquery('english', $1)) DESC
LIMIT $2 OFFSET $3
```

**Features:**
- Full-text search on product name and description
- Relevance ranking using `ts_rank`
- Case-insensitive search
- Language support (English)

## Hexagonal Architecture Benefits

### 1. **Testability**
```go
// Easy to test with mocks
mockRepo := &MockProductRepository{}
service := service.NewSearchService(mockRepo)
```

### 2. **Flexibility**
```go
// Easy to swap PostgreSQL with Elasticsearch
elasticRepo := elasticsearch.NewProductRepository()
service := service.NewSearchService(elasticRepo)
```

### 3. **Independence**
- Core domain has **zero external dependencies**
- Business logic is **framework-agnostic**
- Can run without HTTP server or database

### 4. **Clear Boundaries**
- **Ports** define contracts between layers
- **Adapters** implement the contracts
- **Service** contains pure business logic

## Environment Variables

```bash
SERVER_PORT=:8080
DB_HOST=products-pg-db.storages
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=mydatabase
```

**Note:** The service connects to the existing products database (no new database needed).

## Dependencies

```go
require (
    github.com/labstack/echo/v4
    gorm.io/driver/postgres
    gorm.io/gorm
)
```

## Build & Deploy

### Local Build
```bash
cd /home/phong/dev/shopiew/be/services/search-service
go build -o server ./cmd/server.go
./server
```

### Docker Build
```bash
cd /home/phong/dev/shopiew/be/services/search-service
./build.sh
```

### Kubernetes Deployment

#### Deploy Service
```bash
kubectl apply -f /home/phong/dev/shopiew/be/k8s/services/search-svc.yaml
```

#### Apply Route
```bash
kubectl apply -f /home/phong/dev/shopiew/be/k8s/envoy-gateway/routes/search.yaml
```

#### Verify Deployment
```bash
# Check pods
kubectl get pods -n services | grep search

# Check logs
kubectl logs -n services -l app=search --tail=50

# Check route
kubectl get httproute search-route -o yaml
```

## Testing

### Test Search
```bash
curl "http://localhost:8000/api/search?q=laptop&page=1&limit=10"
```

### Test Get Product
```bash
curl "http://localhost:8000/api/search/products/1"
```

### Test Health Check
```bash
curl "http://localhost:8000/api/search/health"
```

## Hexagonal Architecture Layers Explained

### Core Layer (Domain)

**Location:** `internal/core/entity/`

Pure business objects with no external dependencies:
```go
type Product struct {
    ID          uint
    Name        string
    Description string
    Price       float64
    // ... no GORM tags, no JSON tags in core
}
```

### Ports Layer (Interfaces)

**Input Ports** (`internal/core/port/in/`): What the service provides
```go
type SearchServicePort interface {
    SearchProducts(ctx, request) (response, error)
    GetProduct(ctx, id) (product, error)
}
```

**Output Ports** (`internal/core/port/out/`): What the service needs
```go
type ProductRepositoryPort interface {
    SearchByKeyword(ctx, keyword, limit, offset) ([]Product, int64, error)
    FindByID(ctx, id) (*Product, error)
}
```

### Service Layer (Use Cases)

**Location:** `internal/service/`

Business logic that:
- Validates input
- Coordinates operations
- Calls repository through port (interface)
- Returns DTOs to handlers

### Adapters Layer

**Input Adapter** (`internal/adapter/handler/`):
- Receives HTTP requests
- Calls service through input port
- Returns HTTP responses

**Output Adapter** (`internal/adapter/storage/`):
- Implements repository port
- Communicates with PostgreSQL
- Returns domain entities

## Adding New Search Features

### Example: Add Category Filter

1. **Update DTO** (`internal/core/dto/search_dto.go`):
```go
type SearchRequest struct {
    Query      string
    CategoryID *uint  // NEW
    Page       int
    Limit      int
}
```

2. **Update Port** (`internal/core/port/out/product_repository_port.go`):
```go
SearchByKeywordAndCategory(ctx, keyword string, categoryID *uint, limit, offset int) ([]Product, int64, error)
```

3. **Implement in Adapter** (`internal/adapter/storage/`):
```go
func (r *PostgresProductRepository) SearchByKeywordAndCategory(...) {
    // Add category filter to SQL query
}
```

4. **Update Service** (`internal/service/search_service.go`):
```go
func (s *SearchService) SearchProducts(...) {
    // Call new repository method
}
```

**No changes needed** in HTTP handler - it just passes the request!

## Future Enhancements

### Easy to Implement (Swap Adapters)
- [ ] **Elasticsearch Integration** - Just create new adapter
- [ ] **Redis Caching** - Add caching adapter
- [ ] **gRPC Support** - Add gRPC handler adapter

### Core Features
- [ ] Advanced filters (price range, ratings)
- [ ] Search suggestions/autocomplete
- [ ] Search analytics
- [ ] Fuzzy matching
- [ ] Multi-language support

### Performance
- [ ] Add tsvector indexes to database
- [ ] Query result caching
- [ ] Search query optimization

## Troubleshooting

### Pod Not Starting
```bash
kubectl describe pod -n services -l app=search
kubectl logs -n services -l app=search
```

### Database Connection Issues
```bash
# Verify products-pg-db is running
kubectl get pods -n storages | grep products-pg

# Check connection from search pod
kubectl exec -n services <pod-name> -- env | grep DB_
```

### Search Not Working
```bash
# Check if products table exists
kubectl exec -n storages products-pg-0 -- psql -U postgres -d mydatabase -c "\dt"

# Test full-text search directly
kubectl exec -n storages products-pg-0 -- psql -U postgres -d mydatabase \
  -c "SELECT * FROM products WHERE to_tsvector('english', name) @@ plainto_tsquery('english', 'laptop') LIMIT 5;"
```

## License

Internal Shopiew project - All rights reserved

## Architecture Diagram

```
                    ┌─────────────┐
                    │   Client    │
                    └──────┬──────┘
                           │ HTTP Request
                           ▼
              ┌────────────────────────┐
              │  HTTP Handler Adapter  │ ◄── Input Adapter
              │   (Echo Framework)     │
              └────────────┬───────────┘
                           │ calls
                           ▼
              ┌────────────────────────┐
              │    Input Port          │ ◄── Interface
              │ (SearchServicePort)    │
              └────────────┬───────────┘
                           │ implemented by
                           ▼
              ┌────────────────────────┐
              │   Service Layer        │ ◄── Core Business Logic
              │  (SearchService)       │
              └────────────┬───────────┘
                           │ uses
                           ▼
              ┌────────────────────────┐
              │   Output Port          │ ◄── Interface
              │(ProductRepositoryPort) │
              └────────────┬───────────┘
                           │ implemented by
                           ▼
              ┌────────────────────────┐
              │ Storage Adapter        │ ◄── Output Adapter
              │(PostgresProductRepo)   │
              └────────────┬───────────┘
                           │ queries
                           ▼
                   ┌───────────────┐
                   │  PostgreSQL   │
                   │ (products-db) │
                   └───────────────┘
```

**Key:** The arrows show the direction of dependency. Core (service) never depends on adapters!
