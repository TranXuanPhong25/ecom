CREATE TABLE IF NOT EXISTS fulfillment_packages (
    id BIGSERIAL PRIMARY KEY,
    package_number VARCHAR(50) UNIQUE NOT NULL,
    order_id BIGINT NOT NULL,
    
    -- Seller/pickup info
    shop_id VARCHAR(50) NOT NULL,
    pickup_address TEXT NOT NULL,
    pickup_contact_name VARCHAR(100),
    pickup_contact_phone VARCHAR(20),
    pickup_scheduled_at TIMESTAMP,
    pickup_completed_at TIMESTAMP,
    
    -- Package status
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING_PICKUP',
    
    -- Transit tracking
    current_hub_location VARCHAR(100),
    last_scan_at TIMESTAMP,
    estimated_delivery TIMESTAMP,
    
    -- Delivery info
    delivery_address TEXT NOT NULL,
    delivery_contact_name VARCHAR(100) NOT NULL,
    delivery_contact_phone VARCHAR(20) NOT NULL,
    delivery_zone VARCHAR(50),
    delivery_partner VARCHAR(100),
    tracking_number VARCHAR(100),
    
    -- Delivery attempts
    delivery_attempts INT DEFAULT 0,
    last_delivery_attempt_at TIMESTAMP,
    delivery_failure_reason TEXT,
    delivered_at TIMESTAMP,
    delivery_signature_url TEXT,
    
    -- Package details
    weight_grams INT,
    dimensions JSONB,
    special_instructions TEXT,
    
    -- Metadata
    metadata JSONB,
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fulfillment_order ON fulfillment_packages(order_id);
CREATE INDEX idx_fulfillment_status ON fulfillment_packages(status);
CREATE INDEX idx_fulfillment_shop ON fulfillment_packages(shop_id);
CREATE INDEX idx_fulfillment_delivery_zone ON fulfillment_packages(delivery_zone);
CREATE INDEX idx_fulfillment_package_number ON fulfillment_packages(package_number);

-- Package status enum values:
-- PENDING_PICKUP: Waiting for pickup from seller
-- PICKED_UP: Collected from seller
-- AT_HUB: Arrived at fulfillment hub
-- IN_TRANSIT: Moving between hubs
-- OUT_FOR_DELIVERY: Out for last-mile delivery
-- DELIVERED: Successfully delivered
-- DELIVERY_FAILED: Failed delivery
-- RETURNED_TO_SELLER: Returned to seller after failed delivery
