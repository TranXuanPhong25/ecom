-- Outbox table theo Debezium Outbox Event Router pattern
CREATE TABLE outboxes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Debezium routing fields
    aggregate_type VARCHAR(255) NOT NULL,      -- Dùng làm topic routing
    aggregate_id VARCHAR(255) NOT NULL,        -- Dùng làm message key
    type VARCHAR(255) NOT NULL,               -- Event type
    
    -- Event data
    payload JSONB NOT NULL,                   -- Event payload
    
    -- Metadata
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    
    -- Optional: để track và debug
    tracing_span_id VARCHAR(255)
    
    -- -- Debezium sẽ tự động detect các column này
    -- CONSTRAINT outboxes_pkey PRIMARY KEY (id)
);

-- Index để query nhanh (optional, chỉ cần nếu query manual)
CREATE INDEX idx_outboxes_timestamp ON outboxes(timestamp);
CREATE INDEX idx_outboxes_aggregate_type ON outboxes(aggregate_type);

-- Partition theo timestamp để dễ cleanup (optional)
-- CREATE TABLE outbox (
--     ...
-- ) PARTITION BY RANGE (timestamp);