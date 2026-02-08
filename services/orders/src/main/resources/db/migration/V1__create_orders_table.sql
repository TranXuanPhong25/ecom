-- Drop existing tables
DROP TABLE IF EXISTS order_status_history CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;

-- Orders table
CREATE TABLE IF NOT EXISTS orders
(
    id                    BIGSERIAL PRIMARY KEY,
    order_number          VARCHAR(50) UNIQUE NOT NULL,  -- Mã đơn hàng hiển thị cho user
    user_id               VARCHAR(255) NOT NULL,
    shop_id               VARCHAR(255) NOT NULL,        -- Mỗi order thuộc 1 shop
    
    -- Shipping information
    recipient_name        VARCHAR(255) NOT NULL,
    recipient_phone       VARCHAR(20) NOT NULL,
    delivery_address      TEXT NOT NULL,
    
    -- Order status
    status                VARCHAR(50) NOT NULL DEFAULT 'UNCONFIRMED', 
    -- UNCONFIRMED -> CONFIRMED -> PROCESSING -> SHIPPING -> DELIVERED -> COMPLETED
    -- CANCELLED, RETURNED, REFUNDED
    
    -- Payment information
    payment_method        VARCHAR(50) NOT NULL,  -- COD, CREDIT_CARD, E_WALLET, BANK_TRANSFER
    payment_status        VARCHAR(50) NOT NULL DEFAULT 'UNPAID',  -- UNPAID, PAID, REFUNDED
    paid_at               TIMESTAMPTZ,
    
    -- Pricing
    subtotal              BIGINT NOT NULL DEFAULT 0,  -- Tổng tiền hàng
    shipping_fee          BIGINT NOT NULL DEFAULT 0,  -- Phí vận chuyển
    discount              JSONB DEFAULT '{}',         -- Chi tiết giảm giá (nếu có)
    total_amount          BIGINT NOT NULL DEFAULT 0,  -- Tổng thanh toán   

    -- Shipping
    shipping_method       VARCHAR(50),  -- STANDARD, EXPRESS, SAME_DAY
    shipping_provider     VARCHAR(100), -- GHN, GHTK, VNPost, etc.
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
    
    CONSTRAINT chk_total_amount CHECK (total_amount >= 0),
    CONSTRAINT chk_subtotal CHECK (subtotal >= 0)
);

-- Order items table
CREATE TABLE IF NOT EXISTS order_items
(
    id                BIGSERIAL PRIMARY KEY,
    order_id          BIGINT NOT NULL,
    
    -- Product information (snapshot tại thời điểm đặt hàng)
    product_id        VARCHAR(255) NOT NULL,
    product_name      VARCHAR(500) NOT NULL,
    product_sku       VARCHAR(100),
    image_url         TEXT NOT NULL,
    
    -- Variant information
    variant_id        VARCHAR(255),
    variant_name      VARCHAR(255),  -- Ví dụ: "Màu đỏ, Size M"
    
    -- Pricing
    original_price    BIGINT NOT NULL,  -- Giá gốc
    sale_price        BIGINT NOT NULL,  -- Giá bán
    quantity          INT NOT NULL,

    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    CONSTRAINT chk_quantity CHECK (quantity > 0),
    CONSTRAINT chk_price CHECK (sale_price >= 0)
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

-- Indexes for order_items table
CREATE INDEX idx_order_items_order_id ON order_items (order_id);
CREATE INDEX idx_order_items_product_id ON order_items (product_id);
CREATE INDEX idx_order_items_variant_id ON order_items (variant_id) WHERE variant_id IS NOT NULL;

-- Trigger to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_orders_updated_at
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER trigger_order_items_updated_at
    BEFORE UPDATE ON order_items
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

