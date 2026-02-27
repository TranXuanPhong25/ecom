-- Add fulfillment tracking fields to orders table
ALTER TABLE orders ADD COLUMN IF NOT EXISTS package_number VARCHAR(50);
ALTER TABLE orders ADD COLUMN IF NOT EXISTS pickup_scheduled_at TIMESTAMP;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS pickup_completed_at TIMESTAMP;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS return_deadline TIMESTAMP;

-- Add index for package number lookup
CREATE INDEX IF NOT EXISTS idx_orders_package_number ON orders(package_number);
