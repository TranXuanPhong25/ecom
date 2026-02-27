# Fulfillment Service - Simplified Version

Platform fulfillment service với **auto-delivery simulation**. Không cần driver app, warehouse scanner, hay delivery partner integration - mọi thứ tự động!

## ⚡ Đơn Giản Hóa

### Thay Vì Phức Tạp:
```
❌ Driver app scan barcode
❌ Warehouse scanner system
❌ Delivery partner webhooks
❌ Manual location updates
```

### Chỉ Cần:
```
✅ Cron job chạy mỗi 30 phút
✅ Tự động tiến package qua các stage
✅ Tự động tạo location hubs ảo
✅ 4 lần update → DELIVERED
```

---

## 🚚 Auto-Delivery Flow

### Khi Seller Marks "Ready to Ship":

```
1. Orders Service calls:
   POST /api/v1/fulfillment/pickup/schedule
   → Tạo package (PENDING_PICKUP)
   → Pickup scheduled: +24h

2. Cron Job #1 (after 30 min):
   PENDING_PICKUP → PICKED_UP
   Location: "Đã lấy hàng từ người bán"
   Event: fulfillment.picked_up

3. Cron Job #2 (after 30 min):
   PICKED_UP → AT_HUB  
   Location: "Kho trung chuyển - Đang phân loại"

4. Cron Job #3 (after 30 min):
   AT_HUB → IN_TRANSIT
   Location: "Kho Hà Nội - Đang vận chuyển" (based on zone)
   Event: fulfillment.in_transit

5. Cron Job #4 (after 30 min):
   IN_TRANSIT → OUT_FOR_DELIVERY
   Location: "Đang giao hàng đến bạn"
   Event: fulfillment.out_for_delivery

6. Cron Job #5 (after 30 min):
   OUT_FOR_DELIVERY → DELIVERED ✅
   Location: "Giao hàng thành công"
   Event: fulfillment.delivered
```

**Total Time:** ~2.5 giờ từ lúc schedule đến delivered (có thể adjust ticker interval)

---

## 📁 Key Files

```
be/services/fulfillment/
├── cmd/server.go                           # Entry + cron job starter
├── internal/
│   └── service/
│       ├── fulfillment_service.go          # Main service (schedule pickup)
│       └── auto_delivery_simulator.go      # 🆕 Cron job logic
```

---

## 🎯 API Usage (Simplified)

### 1. Schedule Pickup (Only Endpoint Orders Service Needs)

```bash
POST /api/v1/fulfillment/pickup/schedule
{
  "order_id": 12345,
  "shop_id": "shop-123",
  "pickup_address": "123 Seller St, Hanoi",
  "delivery_address": "456 Buyer Ave, HCMC",
  "delivery_contact_name": "Nguyen Van A",
  "delivery_contact_phone": "0912345678"
}

Response:
{
  "package_number": "PKG1708677123456789",
  "pickup_scheduled_at": "2026-02-24T09:00:00Z",
  "estimated_delivery": "2026-02-24T12:00:00Z", // ~2.5h
  "message": "Pickup scheduled. Auto-delivery simulation started."
}
```

**Chỉ vậy thôi!** Sau đó cron job tự lo hết.

### 2. Tracking (For Buyer UI)

```bash
GET /api/v1/fulfillment/tracking/PKG1708677123456789

Response:
{
  "package_number": "PKG1708677123456789",
  "order_id": 12345,
  "status": "IN_TRANSIT",
  "current_location": "Kho Hà Nội - Đang vận chuyển",
  "tracking_events": [
    {
      "location": "Đã lấy hàng từ người bán",
      "status": "PICKED_UP",
      "timestamp": "2026-02-24T09:30:00Z"
    },
    {
      "location": "Kho trung chuyển - Đang phân loại",
      "status": "AT_HUB",
      "timestamp": "2026-02-24T10:00:00Z"
    },
    {
      "location": "Kho Hà Nội - Đang vận chuyển",
      "status": "IN_TRANSIT",
      "timestamp": "2026-02-24T10:30:00Z"
    }
  ]
}
```

---

## ⚙️ Configuration

```bash
# .env
AUTO_DELIVERY_INTERVAL=30m  # Cron job interval (default: 30 minutes)
```

Muốn nhanh hơn? Set `10m` hoặc `5m`. Muốn chậm hơn? Set `1h`.

---

## 🚀 Deployment

Giống như trước, chỉ cần:

```bash
# Build & push
./build.sh

# Deploy to K8s
kubectl apply -f ../../k8s/storages/fulfillment-pg.yaml
kubectl apply -f ../../k8s/services/fulfillment-svc.yaml
kubectl apply -f ../../k8s/envoy-gateway/routes/fulfillment.yaml
```

Cron job tự động chạy khi service start.

---

## 🎮 How Orders Service Uses It

```java
// Orders Service
@Service
public class OrderService {
  
  @Transactional
  public void markReadyToShip(Long orderId) {
    Order order = findOrder(orderId);
    order.setStatus(READY_TO_SHIP);
    orderRepository.save(order);
    
    // ✅ ONLY API call needed
    FulfillmentResponse response = fulfillmentClient.schedulePickup(order);
    
    order.setPackageNumber(response.getPackageNumber());
    orderRepository.save(order);
    
    // ✅ Then just listen to Kafka events
    // fulfillment.picked_up → Update order to PICKED_UP
    // fulfillment.delivered → Update order to DELIVERED
  }
}

@KafkaListener(topics = {"fulfillment.picked_up", "fulfillment.delivered"})
public void onFulfillmentEvent(FulfillmentEvent event) {
  Order order = orderRepository.findById(event.getOrderId());
  
  switch(event.getType()) {
    case "picked_up":
      order.setStatus(PICKED_UP);
      break;
    case "delivered":
      order.setStatus(DELIVERED);
      order.setReturnDeadline(Instant.now().plus(15, DAYS));
      break;
  }
  
  orderRepository.save(order);
}
```

**That's it!** Không cần driver app, scanner, webhook gì cả.

---

## 📊 Monitoring

```bash
# Check cron job logs
kubectl logs -n services -l app=fulfillment -f | grep "🚚"

Output:
🚀 Auto-delivery simulator started (runs every 30 minutes)
🚚 Running auto-delivery simulator...
📦 Package PKG123: PENDING_PICKUP → PICKED_UP
📦 Package PKG123: PICKED_UP → AT_HUB
📦 Package PKG123: AT_HUB → IN_TRANSIT (Kho Hà Nội - Đang vận chuyển)
📦 Package PKG123: IN_TRANSIT → OUT_FOR_DELIVERY
✅ Package PKG123: OUT_FOR_DELIVERY → DELIVERED
```

---

## 🎯 Benefits

1. **Zero External Dependencies** - Không cần driver app, delivery partner API
2. **Predictable Timing** - Mỗi stage ~30 phút (adjustable)
3. **Easy Testing** - Có thể set interval = 1 minute cho test
4. **Realistic UI** - Buyer vẫn thấy tracking events như thật
5. **Event-Driven** - Orders Service vẫn nhận Kafka events bình thường

---

## 🔧 Customization

### Adjust Delivery Speed

```go
// cmd/server.go
func startDeliverySimulator(simulator *service.AutoDeliverySimulator) {
  // Fast mode: Every 5 minutes
  ticker := time.NewTicker(5 * time.Minute)
  
  // Or slow mode: Every 2 hours
  ticker := time.NewTicker(2 * time.Hour)
}
```

### Add More Location Hubs

```go
// auto_delivery_simulator.go
func (s *AutoDeliverySimulator) getTransitLocation(pkg *entity.FulfillmentPackage) string {
  locations := []string{
    "Kho Hà Nội - Đang phân loại",
    "Trung tâm Logistics miền Bắc",
    "Đang vận chuyển liên tỉnh",
    "Kho TP.HCM - Đã đến khu vực",
    "Bưu cục quận 1 - Đang giao hàng",
  }
  // Random or sequential
}
```

---

## ✨ Perfect for MVP/Demo

- Buyer experience: Tracking timeline đẹp, professional
- Seller experience: Đơn giản, không phải làm gì thêm
- Developer experience: Zero complexity
- Demo-friendly: Predictable, fast, no flakiness

Done! Fulfillment service giờ đơn giản hơn 10x! 🎉

## Architecture

```
├── cmd/
│   └── server.go                 # Main entry point
├── internal/
│   ├── adapter/
│   │   ├── handler/http/         # HTTP handlers
│   │   ├── storage/postgres/     # Repository implementation
│   │   └── event/                # Kafka event publishers
│   ├── core/
│   │   ├── entity/               # Domain models
│   │   ├── dto/                  # Data transfer objects
│   │   └── port/                 # Interface definitions
│   ├── service/                  # Business logic
│   └── config/                   # Configuration management
├── migrations/                   # Database migrations
└── validators/                   # Request validators
```

## Package Status Flow

```
PENDING_PICKUP → PICKED_UP → AT_HUB → IN_TRANSIT → OUT_FOR_DELIVERY → DELIVERED
                                                                      ↓
                                                              DELIVERY_FAILED
                                                                      ↓
                                                          (retry max 3 times)
                                                                      ↓
                                                            RETURNED_TO_SELLER
```

## API Endpoints

### Schedule Pickup
```bash
POST /api/v1/fulfillment/pickup/schedule
Content-Type: application/json

{
  "order_id": 12345,
  "shop_id": "shop-123",
  "pickup_address": "123 Nguyen Trai, Hanoi",
  "pickup_contact_name": "Seller Name",
  "pickup_contact_phone": "0912345678",
  "delivery_address": "456 Le Loi, HCMC",
  "delivery_contact_name": "Buyer Name",
  "delivery_contact_phone": "0987654321",
  "weight_grams": 500,
  "dimensions": {"length": 20, "width": 15, "height": 10},
  "special_instructions": "Fragile - handle with care"
}

Response:
{
  "package_number": "PKG1708677123456789",
  "pickup_scheduled_at": "2026-02-24T09:00:00Z",
  "estimated_delivery": "2026-02-27T17:00:00Z",
  "message": "Pickup scheduled successfully. Driver will arrive at scheduled time."
}
```

### Mark Picked Up
```bash
POST /api/v1/fulfillment/pickup/confirm
Content-Type: application/json

{
  "package_number": "PKG1708677123456789",
  "pickup_by": "Driver A",
  "notes": "Package condition good"
}
```

### Update Location
```bash
POST /api/v1/fulfillment/location/update
Content-Type: application/json

{
  "package_number": "PKG1708677123456789",
  "location": "Fulfillment Hub North - Hanoi",
  "scanned_at": "2026-02-24T10:30:00Z"
}
```

### Update Delivery Status
```bash
POST /api/v1/fulfillment/delivery/status
Content-Type: application/json

{
  "package_number": "PKG1708677123456789",
  "status": "DELIVERED",
  "delivery_signature_url": "https://cdn.shopiew.vn/signatures/abc123.jpg",
  "attempted_at": "2026-02-27T16:45:00Z"
}
```

### Get Package Tracking
```bash
GET /api/v1/fulfillment/tracking/{packageNumber}

Response:
{
  "package_number": "PKG1708677123456789",
  "order_id": 12345,
  "status": "IN_TRANSIT",
  "current_location": "Fulfillment Hub South - HCMC",
  "last_scan_at": "2026-02-26T14:20:00Z",
  "estimated_delivery": "2026-02-27T17:00:00Z",
  "delivery_attempts": 0,
  "tracking_events": [
    {
      "location": "Pickup Scheduled",
      "status": "PENDING_PICKUP",
      "timestamp": "2026-02-24T09:00:00Z"
    },
    {
      "location": "Picked Up from Seller",
      "status": "PICKED_UP",
      "timestamp": "2026-02-24T10:15:00Z"
    },
    {
      "location": "Fulfillment Hub North - Hanoi",
      "status": "AT_HUB",
      "timestamp": "2026-02-24T12:00:00Z"
    },
    {
      "location": "Fulfillment Hub South - HCMC",
      "status": "IN_TRANSIT",
      "timestamp": "2026-02-26T14:20:00Z"
    }
  ],
  "created_at": "2026-02-23T15:30:00Z",
  "updated_at": "2026-02-26T14:20:00Z"
}
```

### List Packages
```bash
GET /api/v1/fulfillment/packages?shop_id=shop-123&status=IN_TRANSIT&page=1&page_size=20

Response:
{
  "content": [
    {
      "id": 1,
      "package_number": "PKG1708677123456789",
      "order_id": 12345,
      "shop_id": "shop-123",
      "status": "IN_TRANSIT",
      "pickup_scheduled_at": "2026-02-24T09:00:00Z",
      "estimated_delivery": "2026-02-27T17:00:00Z",
      "delivery_zone": "ZONE_SOUTH",
      "created_at": "2026-02-23T15:30:00Z"
    }
  ],
  "total_elements": 1,
  "total_pages": 1,
  "page_number": 1,
  "page_size": 20
}
```

## Kafka Events

Service publishes following events:

- `fulfillment.pickup_scheduled`: Khi pickup được schedule
- `fulfillment.picked_up`: Khi package được pickup
- `fulfillment.in_transit`: Khi package đang transit
- `fulfillment.out_for_delivery`: Khi package out for delivery
- `fulfillment.delivered`: Khi delivery thành công
- `fulfillment.delivery_failed`: Khi delivery thất bại

Event payload example:
```json
{
  "event_type": "fulfillment.picked_up",
  "package_number": "PKG1708677123456789",
  "order_id": 12345,
  "shop_id": "shop-123",
  "picked_up_at": "2026-02-24T10:15:00Z"
}
```

## Environment Variables

```bash
SERVER_PORT=8080

# Database
DB_HOST=fulfillment-pg-svc.services.svc.cluster.local
DB_PORT=5432
DB_USER=fulfillment_user
DB_PASSWORD=changeme
DB_NAME=fulfillment_db

# Kafka
KAFKA_BROKERS=kafka-kafka-bootstrap.kafka.svc.cluster.local:9092

# Business Logic
DEFAULT_PICKUP_WINDOW=24       # Hours ahead to schedule pickup
MAX_DELIVERY_ATTEMPTS=3        # Max retry attempts
ESTIMATED_DELIVERY_DAYS=3      # Default delivery estimation
```

## Database Setup

1. Apply storage deployment:
```bash
kubectl apply -f be/k8s/storages/fulfillment-pg.yaml
```

2. Run migrations:
```bash
# Connect to pod
kubectl exec -it -n services fulfillment-pg-deployment-xxx -- psql -U fulfillment_user -d fulfillment_db

# Run migration
\i /path/to/migrations/001_create_fulfillment_packages_table.up.sql
```

Or use a migration tool like [golang-migrate](https://github.com/golang-migrate/migrate):
```bash
migrate -path ./migrations -database "postgresql://fulfillment_user:changeme@localhost:5432/fulfillment_db?sslmode=disable" up
```

## Deployment

### Local Development
```bash
# Copy env file
cp .env.example .env

# Install dependencies
go mod download

# Run service
go run cmd/server.go
```

### Docker
```bash
# Build image
./build.sh

# Or manually
docker build -t rengumin/fulfillment:1.0 .
docker push rengumin/fulfillment:1.0
```

### Kubernetes
```bash
# Apply storage first
kubectl apply -f be/k8s/storages/fulfillment-pg.yaml

# Wait for database to be ready
kubectl wait --for=condition=ready pod -l app=fulfillment-pg -n services --timeout=120s

# Apply service deployment
kubectl apply -f be/k8s/services/fulfillment-svc.yaml

# Apply gateway route
kubectl apply -f be/k8s/envoy-gateway/routes/fulfillment.yaml

# Check status
kubectl get pods -n services -l app=fulfillment
kubectl logs -n services -l app=fulfillment -f
```

## Business Logic

### Pickup Scheduling
- Khi seller marks order as "Ready to Ship"
- Service tự động calculate pickup time (default: next day, 9-11 AM)
- Generate unique package number
- Calculate estimated delivery (pickup time + 3 days)
- Determine delivery zone from address

### Delivery Retry Logic
- Max 3 delivery attempts
- Attempt 1 fails → Auto-schedule retry next day
- Attempt 2 fails → Contact buyer to confirm address
- Attempt 3 fails → Status = RETURNED_TO_SELLER

### Zone Determination
Current implementation: simple first-character logic
- A-M → ZONE_NORTH
- N-Z → ZONE_SOUTH

Production: Use geocoding API or address parsing service

## Testing

```bash
# Unit tests
go test ./...

# Integration tests
go test -tags=integration ./...

# Test API locally
curl http://localhost:8080/health

# Test pickup scheduling
curl -X POST http://localhost:8080/api/v1/fulfillment/pickup/schedule \
  -H "Content-Type: application/json" \
  -d @test-data/schedule-pickup.json
```

## Monitoring

Health check endpoint:
```bash
GET /health

Response: {"status": "healthy"}
```

Metrics to monitor:
- Pickup success rate
- Average transit time
- Delivery success rate (first attempt)
- Failed deliveries count

## Future Enhancements

- [ ] Real delivery partner integrations (GHN, GHTK, Viettel Post)
- [ ] GPS tracking for delivery vehicles
- [ ] Automated routing and driver assignment
- [ ] Return package handling
- [ ] Batch pickup scheduling optimization
- [ ] SLA violation alerts
- [ ] Delivery time predictions with ML

## Support

For issues or questions:
- Check logs: `kubectl logs -n services -l app=fulfillment`
- Check database: Connect to fulfillment-pg-svc
- Check Kafka topics: `fulfillment.*`
