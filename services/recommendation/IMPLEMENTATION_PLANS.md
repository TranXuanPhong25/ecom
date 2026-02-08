# Recommendation Service Implementation Plan

## Problem Statement

Implement a hybrid recommendation service that combines collaborative filtering (user behavior) and content-based filtering (product attributes) with event-driven updates via Kafka.

## Current State

- **Backend**: Basic Python/FastAPI skeleton exists at `be/services/recommendation/`
- **Storage**: MongoDB configured in k8s (`recommendation-product-mongo.yaml`)
- **Frontend**: Hardcoded mock data in `product-discovery-grid.tsx`
- **Infrastructure**: Kafka available via Strimzi

## Proposed Architecture

```
┌─────────────┐     ┌─────────────┐     ┌──────────────────┐
│   Frontend  │────▶│   Envoy     │────▶│  Recommendation  │
│  (Next.js)  │     │   Gateway   │     │  Service (FastAPI)│
└─────────────┘     └─────────────┘     └────────┬─────────┘
                                                 │
                    ┌────────────────────────────┼────────────────────────────┐
                    │                            │                            │
                    ▼                            ▼                            ▼
              ┌──────────┐               ┌──────────────┐              ┌──────────┐
              │  MongoDB │               │    Kafka     │              │ Products │
              │  (Cache) │               │   (Events)   │              │   API    │
              └──────────┘               └──────────────┘              └──────────┘
```

## Recommendation Types

| Type              | Algorithm         | Data Source                    | Update Frequency              |
| ----------------- | ----------------- | ------------------------------ | ----------------------------- |
| Daily Discovery   | Hybrid (CF + CB)  | User history + Product attrs   | Daily batch + real-time boost |
| Similar Products  | Content-Based     | Product attributes, categories | On product update             |
| Frequently Bought | Association Rules | Order history                  | Daily batch                   |
| Recently Viewed   | Session-based     | User events                    | Real-time                     |

## Workplan

### Phase 1: Backend Service Foundation

- [ ] **1.1** Set up project structure with proper Python packaging
   - Add dependencies: scikit-learn, motor (async MongoDB), aiokafka, redis
   - Create config management with Pydantic settings
- [ ] **1.2** Implement data models
   - `UserEvent` (view, add_to_cart, purchase, search)
   - `ProductFeatures` (synced from products service)
   - `UserProfile` (computed preferences)
   - `Recommendation` (cached results)
- [ ] **1.3** Set up MongoDB collections and indexes
   - `user_events` - time-series collection
   - `product_features` - product catalog mirror
   - `user_profiles` - computed user preferences
   - `recommendations_cache` - pre-computed recommendations
- [ ] **1.4** Implement Kafka consumers
   - Listen to `user-events` topic (views, purchases, cart actions)
   - Listen to `product-updates` topic (sync product data)

### Phase 2: Recommendation Algorithms

- [ ] **2.1** Content-Based Filtering
   - TF-IDF vectorization for product descriptions
   - Category and attribute similarity using cosine similarity
   - Build product-to-product similarity matrix
- [ ] **2.2** Collaborative Filtering
   - User-item interaction matrix from events
   - Matrix factorization with scikit-learn (TruncatedSVD)
   - Compute user-user and item-item similarities
- [ ] **2.3** Hybrid Scoring
   - Weighted combination of CF and CB scores
   - Business rules (diversity, freshness, popularity boost)
   - A/B testing support for weight tuning

### Phase 3: API Endpoints

- [ ] **3.1** Daily Discovery endpoint
   ```
   GET /api/recommendations/daily-discoveries?user_id={id}&limit=20
   ```
- [ ] **3.2** Similar Products endpoint
   ```
   GET /api/recommendations/similar/{product_id}?limit=10
   ```
- [ ] **3.3** Frequently Bought Together endpoint
   ```
   GET /api/recommendations/bought-together/{product_id}?limit=5
   ```
- [ ] **3.4** Recently Viewed + Related endpoint
   ```
   GET /api/recommendations/recently-viewed?user_id={id}&limit=10
   ```
- [ ] **3.5** Track Events endpoint (for frontend to report actions)
   ```
   POST /api/recommendations/events
   ```

### Phase 4: Kubernetes Deployment

- [ ] **4.1** Create service deployment (`recommendation-svc.yaml`)
- [ ] **4.2** Create Envoy HTTPRoute for API gateway
- [ ] **4.3** Set up Kafka topics for events
- [ ] **4.4** Configure OPA policies for auth

### Phase 5: Frontend Integration

- [ ] **5.1** Create recommendation service client in `fe/apps/main/src/features/recommendations/`
   - API service with React Query hooks
   - Types matching backend responses
- [ ] **5.2** Update `ProductDiscoveryGrid` to fetch from API
- [ ] **5.3** Add Similar Products section to product detail page
- [ ] **5.4** Implement "Frequently Bought Together" on cart page
- [ ] **5.5** Add event tracking (views, cart actions)

### Phase 6: Batch Processing & Optimization

- [ ] **6.1** Implement batch jobs for matrix computation
   - CronJob for nightly similarity matrix rebuild
   - Incremental updates for new products
- [ ] **6.2** Add Redis caching layer for hot recommendations
- [ ] **6.3** Implement fallback strategies (cold start, sparse data)

## File Structure (Backend)

```
be/services/recommendation/
├── src/
│   ├── main.py                 # FastAPI app entry
│   ├── config.py               # Settings & environment
│   ├── routes/
│   │   ├── __init__.py
│   │   ├── daily_discovery.py
│   │   ├── similar_products.py
│   │   ├── bought_together.py
│   │   ├── recently_viewed.py
│   │   └── events.py
│   ├── models/
│   │   ├── __init__.py
│   │   ├── events.py           # UserEvent, EventType
│   │   ├── products.py         # ProductFeatures
│   │   ├── users.py            # UserProfile
│   │   └── recommendations.py  # RecommendationResult
│   ├── services/
│   │   ├── __init__.py
│   │   ├── content_based.py    # TF-IDF, similarity
│   │   ├── collaborative.py    # Matrix factorization
│   │   ├── hybrid.py           # Combined scoring
│   │   └── association.py      # Bought together rules
│   ├── consumers/
│   │   ├── __init__.py
│   │   ├── event_consumer.py   # Kafka user events
│   │   └── product_consumer.py # Kafka product updates
│   ├── repositories/
│   │   ├── __init__.py
│   │   ├── events_repo.py
│   │   ├── products_repo.py
│   │   └── cache_repo.py
│   └── utils/
│       ├── __init__.py
│       └── similarity.py
├── tests/
├── pyproject.toml
├── Dockerfile
└── k8s/
    └── recommendation-svc.yaml
```

## API Response Examples

### Daily Discovery

```json
{
	"recommendations": [
		{
			"product_id": "123",
			"score": 0.95,
			"reason": "Based on your recent purchases",
			"product": {
				/* ProductCardProps */
			}
		}
	],
	"metadata": {
		"algorithm": "hybrid",
		"generated_at": "2026-02-03T10:00:00Z"
	}
}
```

### Similar Products

```json
{
	"similar_products": [
		{
			"product_id": "456",
			"similarity_score": 0.87,
			"shared_attributes": ["category", "brand"],
			"product": {
				/* ProductCardProps */
			}
		}
	]
}
```

## Event Schema (Kafka)

```json
{
	"event_type": "product_view | add_to_cart | purchase | search",
	"user_id": "user-123",
	"product_id": "product-456",
	"session_id": "session-789",
	"timestamp": "2026-02-03T10:00:00Z",
	"metadata": {
		"source": "homepage | search | category",
		"search_query": "optional search term"
	}
}
```

## Notes & Considerations

1. **Cold Start Problem**: New users get popular/trending items; new products get content-based placement
2. **Privacy**: User events anonymized, no PII in recommendations
3. **Performance**: Pre-compute heavy matrices, serve from cache
4. **Fallback**: If personalization fails, serve category-based popular items
5. **A/B Testing**: Support multiple algorithm versions with weighted traffic

## Dependencies to Add

```toml
# pyproject.toml additions
dependencies = [
    "fastapi>=0.128.0",
    "uvicorn[standard]>=0.40.0",
    "motor>=3.3.0",           # Async MongoDB
    "aiokafka>=0.10.0",       # Async Kafka
    "redis>=5.0.0",           # Caching
    "scikit-learn>=1.4.0",    # ML algorithms
    "pandas>=2.3.0",          # Data processing
    "numpy>=1.26.0",          # Numerical ops
    "pydantic-settings>=2.0", # Config management
]
```

## Success Metrics

- API response time < 100ms (p95)
- Recommendation click-through rate > 5%
- Coverage: 90% of users get personalized recommendations
- Diversity: Average recommendation list covers 3+ categories
