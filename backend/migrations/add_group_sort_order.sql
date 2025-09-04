-- Add sort_order field to wallet_groups table
-- This migration adds support for user-defined group ordering

-- Add sort_order column to wallet_groups table
ALTER TABLE wallet_groups ADD COLUMN sort_order INTEGER DEFAULT 0 NOT NULL;

-- Create an index on sort_order for performance  
CREATE INDEX idx_wallet_groups_user_sort ON wallet_groups(user_id, sort_order);

-- Initialize sort_order based on created_at (older groups get lower sort_order)
UPDATE wallet_groups 
SET sort_order = (
    SELECT COUNT(*) 
    FROM wallet_groups wg2 
    WHERE wg2.user_id = wallet_groups.user_id 
    AND wg2.created_at <= wallet_groups.created_at
);