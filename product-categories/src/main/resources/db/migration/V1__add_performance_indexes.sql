-- Performance optimization indexes for product_category_closure table
-- These indexes improve query performance for hierarchical queries

-- Index for finding direct children (depth = 1)
CREATE INDEX IF NOT EXISTS idx_closure_ancestor_depth 
ON product_category_closure(ancestor_id, depth);

-- Index for finding parent categories
CREATE INDEX IF NOT EXISTS idx_closure_descendant_depth 
ON product_category_closure(descendant_id, depth);

-- Index for identifying root categories (categories with no parents)
CREATE INDEX IF NOT EXISTS idx_closure_descendant_depth_nonzero 
ON product_category_closure(descendant_id) 
WHERE depth > 0;

-- Composite index for efficient relationship lookups
CREATE INDEX IF NOT EXISTS idx_closure_ancestor_descendant 
ON product_category_closure(ancestor_id, descendant_id);
