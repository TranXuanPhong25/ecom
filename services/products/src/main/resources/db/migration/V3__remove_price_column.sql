-- Remove price column from product_variants table (replaced by original_price and sale_price)
ALTER TABLE product_variants DROP COLUMN IF EXISTS price;
