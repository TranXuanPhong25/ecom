# Voucher Service

A microservice for managing discount vouchers/coupons in the Shopiew e-commerce platform. Built with **Go** using **Hexagonal Architecture** (Clean Architecture).

## Architecture

This service demonstrates **Hexagonal Architecture** (also known as Ports and Adapters pattern):

```
┌─────────────────────────────────────────────────────────┐
│                    HTTP Handler (IN)                     │
│               (adapter/handler/*.go)                     │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              VoucherServicePort (IN)                     │
│           (core/port/in/*.go - Interface)                │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│               Voucher Service (Core)                     │
│              (service/voucher_service.go)                │
│                                                           │
│   Business Logic:                                        │
│   - Validate voucher codes                               │
│   - Check discount value ranges                          │
│   - Verify expiration dates                              │
│   - Enforce business rules                               │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│          VoucherRepositoryPort (OUT)                     │
│          (core/port/out/*.go - Interface)                │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│          PostgreSQL Repository (OUT)                     │
│       (adapter/storage/*.go - Implementation)            │
└─────────────────────────────────────────────────────────┘
```

### Benefits of Hexagonal Architecture

1. **Testability**: Easy to test business logic by mocking repository interfaces
2. **Flexibility**: Can swap PostgreSQL with another database (e.g., MongoDB) by just changing the adapter
3. **Independence**: Core domain has zero external dependencies
4. **Clear Boundaries**: Ports (interfaces) define contracts between layers
5. **Maintainability**: Each layer has a single responsibility

### Project Structure

```
voucher-service/
├── cmd/
│   └── server.go                    # Application entry point (DI wiring)
├── internal/
│   ├── adapter/                     # Adapters (IN and OUT)
│   │   ├── handler/                 # HTTP handlers (IN adapter)
│   │   │   ├── voucher_handler.go
│   │   │   └── routes.go
│   │   └── storage/                 # Database repository (OUT adapter)
│   │       └── postgres_voucher_repository.go
│   ├── core/                        # Domain layer (pure business logic)
│   │   ├── entity/                  # Domain models
│   │   │   └── voucher.go
│   │   ├── dto/                     # Data transfer objects
│   │   │   └── voucher_dto.go
│   │   └── port/                    # Ports (interfaces)
│   │       ├── in/                  # Input ports (what service provides)
│   │       │   └── voucher_service_port.go
│   │       └── out/                 # Output ports (what service needs)
│   │           └── voucher_repository_port.go
│   ├── service/                     # Business logic implementation
│   │   └── voucher_service.go
│   └── config/                      # Configuration
│       ├── config.go
│       └── database.go
└── validators/                      # Input validation
    └── voucher_validator.go
```

## Features

- **CRUD Operations**: Create, Read, Update, Delete vouchers
- **Discount Types**: 
  - Percentage discount (e.g., 10% off)
  - Fixed amount discount (e.g., $5 off)
- **Validation**:
  - Unique voucher codes
  - Discount value validation (0-100% for percentage, >0 for fixed)
  - Future expiration dates
  - Alphanumeric codes with hyphens/underscores
- **Pagination**: List vouchers with page/limit
- **Filtering**: Filter by active status
- **Health Check**: Kubernetes-ready health endpoint

## Voucher Entity

```go
type Voucher struct {
    ID              uint         // Primary key
    Code            string       // Unique voucher code (e.g., "SUMMER2024")
    DiscountType    DiscountType // "PERCENTAGE" or "FIXED_AMOUNT"
    DiscountValue   float64      // 20.0 for 20% or $20
    MinOrderValue   float64      // Minimum order value to apply
    MaxUsage        int          // 0 = unlimited (for future)
    UsedCount       int          // Tracking usage (for future)
    ExpiresAt       *time.Time   // Optional expiration
    IsActive        bool         // Active status
    Description     string       // Description (max 500 chars)
    CreatedAt       time.Time
    UpdatedAt       time.Time
}
```

## API Endpoints

### Create Voucher
```http
POST /api/vouchers
Content-Type: application/json

{
  "code": "SUMMER2024",
  "discount_type": "PERCENTAGE",
  "discount_value": 20,
  "min_order_value": 50,
  "expires_at": "2024-12-31T23:59:59Z",
  "description": "Summer sale - 20% off on orders over $50"
}
```

**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "code": "SUMMER2024",
    "discount_type": "PERCENTAGE",
    "discount_value": 20,
    "min_order_value": 50,
    "max_usage": 0,
    "used_count": 0,
    "expires_at": "2024-12-31T23:59:59Z",
    "is_active": true,
    "description": "Summer sale - 20% off on orders over $50",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

### Get Voucher by Code
```http
GET /api/vouchers/code/SUMMER2024
```

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "code": "SUMMER2024",
    ...
  }
}
```

### Get Voucher by ID
```http
GET /api/vouchers/1
```

### List Vouchers
```http
GET /api/vouchers?page=1&limit=20&active=true
```

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 20, max: 100)
- `active` (optional): Filter by active status (true/false)

**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "vouchers": [...],
    "total": 15,
    "page": 1,
    "limit": 20
  }
}
```

### Update Voucher
```http
PUT /api/vouchers/1
Content-Type: application/json

{
  "discount_value": 25,
  "is_active": false,
  "description": "Updated description"
}
```

**Note**: Voucher code is immutable and cannot be updated.

**Updatable Fields:**
- `discount_value`
- `min_order_value`
- `max_usage`
- `expires_at`
- `is_active`
- `description`

### Delete Voucher
```http
DELETE /api/vouchers/1
```

**Response (200 OK):**
```json
{
  "success": true,
  "message": "Voucher deleted successfully"
}
```

### Health Check
```http
GET /health
```

**Response (200 OK):**
```json
{
  "status": "healthy"
}
```

## Validation Rules

### Voucher Code
- Required
- Length: 4-50 characters
- Format: Alphanumeric + hyphen/underscore only
- Automatically converted to uppercase
- Must be unique

### Discount Value
- **For PERCENTAGE**: 0 < value ≤ 100
- **For FIXED_AMOUNT**: value > 0
- Required

### Min Order Value
- Optional (default: 0)
- If provided: value ≥ 0

### Expiration Date
- Optional
- If provided: must be a future date

### Business Rules
- Cannot create voucher with past expiration
- Cannot update voucher code (immutable)
- Cannot set negative discount values
- Duplicate codes are rejected

## Environment Variables

```bash
SERVER_PORT=8080
DB_HOST=voucher-pg-db.storages
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=voucherdb
```

## Building

### Local Build
```bash
# Build binary
./build.sh

# Run locally (requires PostgreSQL)
./voucher-service
```

### Docker Build
```bash
# Build Docker image
docker build -t rengumin/voucher-svc:1.0 .

# Run with Docker
docker run -p 8080:8080 \
  -e DB_HOST=postgres-host \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  rengumin/voucher-svc:1.0
```

### Build All Services
```bash
# From be/services directory
./build_and_push_to_dockerhub.sh
```

## Deployment

### Prerequisites
- Kubernetes cluster
- `services` and `storages` namespaces
- `app-secrets` secret with POSTGRES_USER and POSTGRES_PASSWORD

### Deploy Database
```bash
kubectl apply -f be/k8s/storages/voucher-pg.yaml
```

### Deploy Service
```bash
kubectl apply -f be/k8s/services/voucher-svc.yaml
```

### Deploy HTTPRoute
```bash
kubectl apply -f be/k8s/envoy-gateway/routes/voucher.yaml
```

### Verify Deployment
```bash
# Check pods
kubectl get pods -n services -l app=voucher-svc
kubectl get pods -n storages -l app=voucher-pg-db

# Check logs
kubectl logs -n services -l app=voucher-svc -f

# Test health check
kubectl port-forward -n services svc/voucher-svc 8080:8080
curl http://localhost:8080/health
```

## Testing

### Create a Test Voucher
```bash
curl -X POST http://localhost:8000/api/vouchers \
  -H "Content-Type: application/json" \
  -d '{
    "code": "NEWYEAR2024",
    "discount_type": "PERCENTAGE",
    "discount_value": 15,
    "min_order_value": 30,
    "description": "New Year Sale"
  }'
```

### Get Voucher by Code
```bash
curl http://localhost:8000/api/vouchers/code/NEWYEAR2024
```

### List All Vouchers
```bash
curl http://localhost:8000/api/vouchers?page=1&limit=10
```

### Update Voucher
```bash
curl -X PUT http://localhost:8000/api/vouchers/1 \
  -H "Content-Type: application/json" \
  -d '{
    "discount_value": 20,
    "description": "Updated: New Year Sale - 20% off"
  }'
```

### Delete Voucher
```bash
curl -X DELETE http://localhost:8000/api/vouchers/1
```

## Error Responses

### Validation Error (400 Bad Request)
```json
{
  "success": false,
  "error": "percentage discount must be between 0 and 100"
}
```

### Not Found (404 Not Found)
```json
{
  "success": false,
  "error": "voucher not found"
}
```

### Duplicate Code (400 Bad Request)
```json
{
  "success": false,
  "error": "voucher code already exists"
}
```

## Future Enhancements

These features are prepared but not yet implemented:

- **Usage Tracking**: Enforce max_usage limits and track used_count
- **User-Specific Vouchers**: Assign vouchers to specific users
- **Multi-Use vs Single-Use**: Different voucher types
- **Categories/Tags**: Organize vouchers by category
- **Apply to Cart**: Integration with cart service for checkout
- **Analytics**: Usage statistics and reports
- **Batch Operations**: Create multiple vouchers at once
- **Auto-Expiration**: Background job to deactivate expired vouchers

## Dependencies

```
github.com/labstack/echo/v4              # HTTP framework
gorm.io/gorm                             # ORM
gorm.io/driver/postgres                   # PostgreSQL driver
github.com/go-playground/validator/v10    # Validation
```

## Database Schema

```sql
CREATE TABLE vouchers (
    id              SERIAL PRIMARY KEY,
    code            VARCHAR(50) UNIQUE NOT NULL,
    discount_type   VARCHAR(20) NOT NULL,
    discount_value  FLOAT NOT NULL,
    min_order_value FLOAT DEFAULT 0,
    max_usage       INTEGER DEFAULT 0,
    used_count      INTEGER DEFAULT 0,
    expires_at      TIMESTAMP,
    is_active       BOOLEAN DEFAULT true,
    description     VARCHAR(500),
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL
);

CREATE INDEX idx_vouchers_code ON vouchers(code);
CREATE INDEX idx_vouchers_is_active ON vouchers(is_active);
```

## Performance Considerations

- **Unique Index** on voucher code for fast lookups
- **Pagination** to limit result sets
- **Connection Pooling** via GORM
- **Context Timeouts** for database queries

## License

Part of the Shopiew e-commerce platform.
