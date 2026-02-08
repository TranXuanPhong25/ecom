# Product Reviews Service

A RESTful API service for managing product reviews in the Shopiew e-commerce platform.

## Architecture

This service follows the standard Go MVC pattern with repository and service layers:

```
product-reviews/
├── server.go           # Main entry point
├── configs/            # Environment configuration
├── controllers/        # HTTP handlers
├── routes/             # Route definitions
├── models/             # GORM models
├── repositories/       # Database access layer
├── services/           # Business logic layer
├── dtos/              # Request/response DTOs
├── validators/        # Input validation
└── utils/             # Helper functions
```

## Features

- ✅ Complete CRUD operations for product reviews
- ✅ Product-based review listing with pagination
- ✅ User-based review listing
- ✅ Ownership validation (users can only edit/delete their own reviews)
- ✅ Input validation (rating 1-5, comment length, etc.)
- ✅ Structured error responses
- 🚧 Product rating statistics (currently mocked - TODO)
- 🚧 JWT authentication integration (TODO)
- 🚧 Database indexes optimization (TODO)

## API Endpoints

### Create Review
```http
POST /api/product-reviews/reviews
Content-Type: application/json

{
  "product_id": 1,
  "user_id": 123,           // TODO: Extract from JWT
  "username": "john_doe",   // TODO: Extract from JWT
  "rating": 5,
  "title": "Great product!",
  "comment": "This product exceeded my expectations. Highly recommended!"
}
```

### Get Single Review
```http
GET /api/product-reviews/reviews/:id
```

### Get Product Reviews (Paginated)
```http
GET /api/product-reviews/products/:productId/reviews?page=1&limit=20
```

### Get Product Statistics
```http
GET /api/product-reviews/products/:productId/reviews/stats
```
**Note:** Currently returns mock data. Real implementation pending.

### Get User Reviews
```http
GET /api/product-reviews/users/:userId/reviews
```

### Update Review
```http
PUT /api/product-reviews/reviews/:id
Content-Type: application/json
X-User-ID: 123              // TODO: Replace with JWT

{
  "rating": 4,
  "title": "Updated title",
  "comment": "Updated comment with at least 20 characters."
}
```

### Delete Review
```http
DELETE /api/product-reviews/reviews/:id
X-User-ID: 123              // TODO: Replace with JWT
```

## Data Models

### Review
```go
type Review struct {
    ID        uint      `json:"id"`
    ProductID uint      `json:"product_id"`
    UserID    uint      `json:"user_id"`
    Username  string    `json:"username"`
    Rating    int       `json:"rating"`        // 1-5
    Title     string    `json:"title"`         // 10-100 chars
    Comment   string    `json:"comment"`       // 20-1000 chars
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Validation Rules

- **Rating**: Required, must be between 1-5
- **Title**: Required, 10-100 characters
- **Comment**: Required, 20-1000 characters
- **ProductID**: Required, positive integer
- **UserID**: Required, positive integer

## Environment Variables

```bash
SERVER_PORT=:8080
DB_HOST=product-reviews-pg-db.storages
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=reviews_db
```

## Kubernetes Deployment

### Deploy Service
```bash
kubectl apply -f /be/k8s/services/product-reviews-svc.yaml
```

### Deploy Database
```bash
kubectl apply -f /be/k8s/storages/product-reviews-pg.yaml
```

### Verify Route
```bash
kubectl get httproute product-reviews-route -o yaml
```

## Build & Deploy

### Local Build
```bash
cd /be/services/product-reviews
go build -o server server.go
./server
```

### Docker Build
```bash
./build.sh
```

### Build All Services
```bash
cd /be/services
./build_and_push_to_dockerhub.sh
```

## Testing

### Test Create Review
```bash
curl -X POST http://localhost:8080/reviews \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "user_id": 123,
    "username": "test_user",
    "rating": 5,
    "title": "Great product!",
    "comment": "This is a test review with enough characters to pass validation."
  }'
```

### Test Get Product Reviews
```bash
curl http://localhost:8080/products/1/reviews?page=1&limit=10
```

### Test Get Statistics (Mock)
```bash
curl http://localhost:8080/products/1/reviews/stats
```

## TODO List

### High Priority
- [ ] **JWT Authentication Integration**
  - Integrate with jwt-service via gRPC
  - Extract UserID and Username from JWT token
  - Remove UserID from request body
  - Add authentication middleware

- [ ] **Real Statistics Calculation**
  - Create ProductRatingStats table
  - Implement aggregation queries
  - Update stats when reviews are created/updated/deleted
  - Replace mock data in GetProductStats

- [ ] **Pagination Total Count**
  - Implement real count query in repository
  - Return accurate total count in ReviewListResponse

### Medium Priority
- [ ] **Database Optimization**
  - Add composite index on (product_id, created_at)
  - Add index on user_id
  - Analyze query performance with EXPLAIN

- [ ] **Validation Enhancements**
  - Verify ProductID exists (call products service)
  - Check if user already reviewed product (one review per user per product)
  - Add rate limiting

### Low Priority
- [ ] **Advanced Features**
  - Review images support (integrate with upload-service)
  - Verified purchase badges
  - Helpful votes system
  - Review replies from sellers
  - Review moderation

## Dependencies

```go
require (
    github.com/labstack/echo/v4 v4.13.3
    github.com/go-playground/validator/v10
    gorm.io/driver/postgres v1.5.11
    gorm.io/gorm v1.25.12
)
```

## Error Responses

All errors follow this format:
```json
{
  "error": "Error Type",
  "message": "Detailed error message"
}
```

### HTTP Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request (validation error)
- `403` - Forbidden (not owner)
- `404` - Not Found
- `500` - Internal Server Error

## Contributing

When implementing TODO items:
1. Remove TODO comments after implementation
2. Update this README
3. Add tests
4. Update API documentation

## License

Internal Shopiew project - All rights reserved
