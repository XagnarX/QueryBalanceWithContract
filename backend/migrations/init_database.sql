-- ================================================================
-- Database Initialization Script
-- ================================================================
-- This file contains the complete database schema for the Wallet Manager application
--
-- Tables included:
-- - users (user management)
-- - chains (blockchain networks)
-- - user_rpc_endpoints (user's custom RPC endpoints)
-- - user_tokens (user's custom token definitions)
-- - wallet_groups (wallet organization)
-- - wallet_addresses (individual addresses)
-- - wallet_group_settings (group configuration and preferences)
-- - balance_contracts (balance checking contracts)
--
-- Note: This script is idempotent and can be run multiple times safely.
-- ================================================================

-- ================================================================
-- 1. UTILITY FUNCTIONS
-- ================================================================

-- Create update timestamp trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- ================================================================
-- 2. CORE TABLES CREATION
-- ================================================================

-- Users table - Core user management
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Chains table - Blockchain networks configuration
CREATE TABLE IF NOT EXISTS chains (
    id SERIAL PRIMARY KEY,
    chain_id INTEGER UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    symbol VARCHAR(10) NOT NULL, -- Native token symbol (BNB, ETH, etc.)
    is_testnet BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    block_explorer_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User RPC endpoints table - User's custom RPC endpoints
-- Must be created before wallet_group_settings (referenced by foreign key)
CREATE TABLE IF NOT EXISTS user_rpc_endpoints (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chain_id INTEGER NOT NULL,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Unique constraints
    CONSTRAINT unique_user_rpc_url UNIQUE (user_id, chain_id, url)
);

-- User tokens table - User's custom token definitions
-- Must be created before wallet_group_settings (referenced via selected_token_ids)
CREATE TABLE IF NOT EXISTS user_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chain_id INTEGER NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    name VARCHAR(50) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    decimals INTEGER NOT NULL DEFAULT 18,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Unique constraints
    CONSTRAINT unique_user_token_contract UNIQUE (user_id, chain_id, contract_address)
);

-- Wallet groups table - User's wallet organization
CREATE TABLE IF NOT EXISTS wallet_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    sort_order INTEGER NOT NULL DEFAULT 0, -- For custom ordering
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, name)
);

-- Wallet addresses table - Individual wallet addresses
CREATE TABLE IF NOT EXISTS wallet_addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES wallet_groups(id) ON DELETE SET NULL,
    address VARCHAR(42) NOT NULL,
    label VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Remove old unique constraint if exists and add new one
DO $$
BEGIN
    -- Drop old constraint if it exists
    IF EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'wallet_addresses_user_id_address_key') THEN
        ALTER TABLE wallet_addresses DROP CONSTRAINT wallet_addresses_user_id_address_key;
    END IF;

    -- Add new unique constraint that includes group_id
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'wallet_addresses_user_group_address_idx') THEN
        CREATE UNIQUE INDEX wallet_addresses_user_group_address_idx
        ON wallet_addresses (user_id, COALESCE(group_id, 0), address);
    END IF;
END $$;

-- Wallet group settings table - Group-specific configurations
CREATE TABLE IF NOT EXISTS wallet_group_settings (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL REFERENCES wallet_groups(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    countdown_enabled BOOLEAN DEFAULT false, -- Whether countdown timer is enabled for this group
    countdown_duration INTEGER DEFAULT 600, -- Countdown duration in seconds (1-7200)
    selected_rpc_id INTEGER REFERENCES user_rpc_endpoints(id) ON DELETE SET NULL, -- User-selected RPC endpoint for this group (NULL = use default)
    selected_token_ids JSONB DEFAULT '[]'::jsonb, -- JSON array of selected token IDs for balance queries
    chain_id INTEGER NOT NULL DEFAULT 1, -- Blockchain network ID - allows different settings per chain
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Check constraints
    CONSTRAINT wallet_group_settings_countdown_duration_check
        CHECK (countdown_duration >= 1 AND countdown_duration <= 7200),

    -- Unique constraint: one settings record per user-group-chain combination
    CONSTRAINT uq_group_settings_user_group_chain UNIQUE (user_id, group_id, chain_id)
);

-- Balance contracts table - Smart contracts for balance checking
CREATE TABLE IF NOT EXISTS balance_contracts (
    id SERIAL PRIMARY KEY,
    chain_id INTEGER NOT NULL REFERENCES chains(chain_id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    priority INTEGER DEFAULT 0, -- Lower number = higher priority
    abi_json TEXT, -- Contract ABI storage
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(chain_id, contract_address)
);

-- ================================================================
-- 3. INDEXES CREATION
-- ================================================================

-- Users table indexes
-- (Primary key index created automatically)

-- Chains table indexes
CREATE INDEX IF NOT EXISTS idx_chains_chain_id ON chains(chain_id);

-- User RPC endpoints table indexes
CREATE INDEX IF NOT EXISTS idx_user_rpc_user_id ON user_rpc_endpoints(user_id);
CREATE INDEX IF NOT EXISTS idx_user_rpc_chain_id ON user_rpc_endpoints(chain_id);
CREATE INDEX IF NOT EXISTS idx_user_rpc_active ON user_rpc_endpoints(is_active);

-- User tokens table indexes
CREATE INDEX IF NOT EXISTS idx_user_token_user_id ON user_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_user_token_chain_id ON user_tokens(chain_id);
CREATE INDEX IF NOT EXISTS idx_user_token_active ON user_tokens(is_active);
CREATE INDEX IF NOT EXISTS idx_user_token_contract ON user_tokens(contract_address);

-- Wallet groups table indexes
CREATE INDEX IF NOT EXISTS idx_wallet_groups_user_id ON wallet_groups(user_id);
CREATE INDEX IF NOT EXISTS idx_wallet_groups_user_sort ON wallet_groups(user_id, sort_order);

-- Wallet addresses table indexes
CREATE INDEX IF NOT EXISTS idx_wallet_addresses_user_id ON wallet_addresses(user_id);
CREATE INDEX IF NOT EXISTS idx_wallet_addresses_group_id ON wallet_addresses(group_id);

-- Wallet group settings table indexes
CREATE INDEX IF NOT EXISTS idx_group_settings_user_id ON wallet_group_settings(user_id);
CREATE INDEX IF NOT EXISTS idx_group_settings_group_id ON wallet_group_settings(group_id);
CREATE INDEX IF NOT EXISTS idx_group_settings_user_group ON wallet_group_settings(user_id, group_id);
CREATE INDEX IF NOT EXISTS idx_wallet_group_settings_chain_id ON wallet_group_settings(chain_id);

-- Balance contracts table indexes
-- (Already created through unique constraint)

-- ================================================================
-- 4. TRIGGERS CREATION
-- ================================================================

-- Users table trigger
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Chains table trigger
DROP TRIGGER IF EXISTS update_chains_updated_at ON chains;
CREATE TRIGGER update_chains_updated_at
    BEFORE UPDATE ON chains
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- User RPC endpoints table trigger
DROP TRIGGER IF EXISTS trigger_user_rpc_endpoints_updated_at ON user_rpc_endpoints;
CREATE TRIGGER trigger_user_rpc_endpoints_updated_at
    BEFORE UPDATE ON user_rpc_endpoints
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- User tokens table trigger
DROP TRIGGER IF EXISTS trigger_user_tokens_updated_at ON user_tokens;
CREATE TRIGGER trigger_user_tokens_updated_at
    BEFORE UPDATE ON user_tokens
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Wallet groups table trigger
DROP TRIGGER IF EXISTS update_wallet_groups_updated_at ON wallet_groups;
CREATE TRIGGER update_wallet_groups_updated_at
    BEFORE UPDATE ON wallet_groups
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Wallet addresses table trigger
DROP TRIGGER IF EXISTS update_wallet_addresses_updated_at ON wallet_addresses;
CREATE TRIGGER update_wallet_addresses_updated_at
    BEFORE UPDATE ON wallet_addresses
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- ================================================================
-- 5. INITIAL DATA INSERTION
-- ================================================================

-- Insert blockchain networks data
INSERT INTO chains (chain_id, name, symbol, is_testnet, is_active, block_explorer_url) VALUES
(56, 'BSC Mainnet', 'BNB', FALSE, TRUE, 'https://bscscan.com'),
(97, 'BSC Testnet', 'tBNB', TRUE, TRUE, 'https://testnet.bscscan.com'),
(1, 'Ethereum Mainnet', 'ETH', FALSE, TRUE, 'https://etherscan.io'),
(5, 'Ethereum Goerli', 'ETH', TRUE, FALSE, 'https://goerli.etherscan.io'),
(137, 'Polygon Mainnet', 'MATIC', FALSE, TRUE, 'https://polygonscan.com'),
(80001, 'Polygon Mumbai', 'MATIC', TRUE, FALSE, 'https://mumbai.polygonscan.com')
ON CONFLICT (chain_id) DO NOTHING;

-- Insert BalanceChecker contract addresses
INSERT INTO balance_contracts (chain_id, name, contract_address, priority) VALUES
(56, 'BalanceChecker BSC', '0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b', 1),
(97, 'BalanceChecker BSC Testnet', '0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b', 1)
ON CONFLICT (chain_id, contract_address) DO NOTHING;

-- ================================================================
-- END OF INITIALIZATION SCRIPT
-- ================================================================

-- Success message
SELECT 'Database initialization completed successfully!' as status;
