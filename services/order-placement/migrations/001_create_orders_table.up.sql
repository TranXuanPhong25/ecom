-- Migration: Create orders table with JSONB items
-- Drop existing tables
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS outboxes CASCADE;

-- Orders table with JSONB items
CREATE TABLE IF NOT EXISTS orders
(
    id                    BIGSERIAL PRIMARY KEY,
    order_number          VARCHAR(50) UNIQUE NOT NULL,
    user_id               VARCHAR(255) NOT NULL,
    shop_id               VARCHAR(255) NOT NULL,
    
    -- Shipping information
    recipient_name        VARCHAR(255) NOT NULL,
    recipient_phone       VARCHAR(20) NOT NULL,
    delivery_address      TEXT NOT NULL,
    
    -- Order status
    status                VARCHAR(50) NOT NULL DEFAULT 'CREATED',
    
    -- Payment information
    payment_method        VARCHAR(50) NOT NULL,
    payment_status        VARCHAR(50) NOT NULL DEFAULT 'UNPAID',
    paid_at               TIMESTAMPTZ,
    
    -- Pricing
    subtotal              BIGINT NOT NULL DEFAULT 0,
    shipping_fee          BIGINT NOT NULL DEFAULT 0,
    discount              JSONB DEFAULT '{}',
    total_amount          BIGINT NOT NULL DEFAULT 0,

    -- Shipping
    shipping_method       VARCHAR(50),
    shipping_provider     VARCHAR(100),
    tracking_number       VARCHAR(100),
    estimated_delivery    TIMESTAMPTZ,
    actual_delivery       TIMESTAMPTZ,
    
    -- Notes
    customer_note         TEXT,
    seller_note           TEXT,
    cancel_reason         TEXT,
    
    -- Timestamps
    confirmed_at          TIMESTAMPTZ,
    processing_at         TIMESTAMPTZ,
    shipped_at            TIMESTAMPTZ,
    delivered_at          TIMESTAMPTZ,
    completed_at          TIMESTAMPTZ,
    cancelled_at          TIMESTAMPTZ,
    
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- Items stored as JSONB
    items                 JSONB NOT NULL DEFAULT '[]',
    
    CONSTRAINT chk_total_amount CHECK (total_amount >= 0),
    CONSTRAINT chk_subtotal CHECK (subtotal >= 0)
);

-- Indexes for orders table
CREATE INDEX idx_orders_order_number ON orders (order_number);
CREATE INDEX idx_orders_user_id ON orders (user_id);
CREATE INDEX idx_orders_shop_id ON orders (shop_id);
CREATE INDEX idx_orders_status ON orders (status);
CREATE INDEX idx_orders_payment_status ON orders (payment_status);
CREATE INDEX idx_orders_created_at ON orders (created_at DESC);
CREATE INDEX idx_orders_user_status ON orders (user_id, status);
CREATE INDEX idx_orders_shop_status ON orders (shop_id, status);
CREATE INDEX idx_orders_tracking_number ON orders (tracking_number) WHERE tracking_number IS NOT NULL;

-- GIN index for JSONB items
CREATE INDEX idx_orders_items_gin ON orders USING GIN (items);
