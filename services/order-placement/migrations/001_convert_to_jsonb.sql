-- Migration: Convert order_items table to JSONB column in orders table
-- This migration drops the order_items table and adds an items JSONB column to orders

-- Step 1: Drop the order_items table (if exists)
DROP TABLE IF EXISTS order_items CASCADE;

-- Step 2: Add items column to orders table as JSONB
ALTER TABLE orders 
ADD COLUMN IF NOT EXISTS items JSONB NOT NULL DEFAULT '[]'::jsonb;

-- Step 3: Add comment for documentation
COMMENT ON COLUMN orders.items IS 'Order items stored as JSONB array';

-- Step 4: Create GIN index on items for better JSON query performance
CREATE INDEX IF NOT EXISTS idx_orders_items_gin ON orders USING GIN (items);

-- Step 5: You can also add specific indexes for common queries
-- Example: Index for searching by product_id within items
CREATE INDEX IF NOT EXISTS idx_orders_items_product_id 
ON orders USING GIN ((items -> 'product_id'));
