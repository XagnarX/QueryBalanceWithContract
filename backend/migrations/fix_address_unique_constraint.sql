-- ================================================================
-- Fix wallet_addresses unique constraint to allow same address 
-- in different groups
-- ================================================================

-- Step 1: Drop the existing unique constraint
ALTER TABLE wallet_addresses DROP CONSTRAINT IF EXISTS wallet_addresses_user_id_address_key;

-- Step 2: Add new unique constraint that includes group_id
-- This allows same address to exist in different groups for the same user
-- But prevents duplicate addresses within the same group
CREATE UNIQUE INDEX wallet_addresses_user_group_address_idx
  ON wallet_addresses (user_id, COALESCE(group_id, 0), address);

-- Note: Using COALESCE(group_id, 0) to handle NULL group_id values
-- This ensures that addresses without group assignment are also unique per user