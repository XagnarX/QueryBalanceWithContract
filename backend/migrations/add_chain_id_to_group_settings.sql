-- Add chain_id to wallet_group_settings table
-- This allows each group to have different settings for different chains

-- Add chain_id column (default to BSC mainnet: 56)
ALTER TABLE wallet_group_settings
ADD COLUMN chain_id INT NOT NULL DEFAULT 1;

-- Add index on chain_id
CREATE INDEX idx_wallet_group_settings_chain_id ON wallet_group_settings(chain_id);

-- Update unique constraint to include chain_id
-- First drop the old constraint
ALTER TABLE wallet_group_settings
DROP CONSTRAINT IF EXISTS uq_group_settings_user_group;

-- Create new unique constraint with chain_id
-- Now each group can have different settings for different chains
ALTER TABLE wallet_group_settings
ADD CONSTRAINT uq_group_settings_user_group_chain
UNIQUE (user_id, group_id, chain_id);

-- Add comment
COMMENT ON COLUMN wallet_group_settings.chain_id IS 'Blockchain network ID - allows different settings per chain';
