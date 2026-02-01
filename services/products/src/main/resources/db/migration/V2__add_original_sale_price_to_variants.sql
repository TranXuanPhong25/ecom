-- Add originalPrice and salePrice columns to product_variants table
ALTER TABLE product_variants 
ADD COLUMN IF NOT EXISTS original_price DECIMAL(10, 2),
ADD COLUMN IF NOT EXISTS sale_price DECIMAL(10, 2);

-- Update existing rows: set original_price = price and sale_price = price for existing variants
UPDATE product_variants 
SET original_price = price, sale_price = price 
WHERE original_price IS NULL OR sale_price IS NULL;

-- Make columns NOT NULL after migration
ALTER TABLE product_variants 
ALTER COLUMN original_price SET NOT NULL,
ALTER COLUMN sale_price SET NOT NULL;

-- Add check constraint: sale_price should be less than or equal to original_price
ALTER TABLE product_variants
ADD CONSTRAINT chk_sale_price CHECK (sale_price <= original_price);
