-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建钱包分组表
CREATE TABLE IF NOT EXISTS wallet_groups (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, name)
);

-- 创建钱包地址表
CREATE TABLE IF NOT EXISTS wallet_addresses (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id INTEGER REFERENCES wallet_groups(id) ON DELETE SET NULL,
    address VARCHAR(42) NOT NULL,
    label VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, address)
);

-- 创建代币表
CREATE TABLE IF NOT EXISTS tokens (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(10) NOT NULL,
    name VARCHAR(50) NOT NULL,
    contract_address VARCHAR(42) UNIQUE NOT NULL,
    decimals INTEGER NOT NULL DEFAULT 18,
    chain_id INTEGER NOT NULL DEFAULT 56, -- BSC mainnet
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_wallet_addresses_user_id ON wallet_addresses(user_id);
CREATE INDEX IF NOT EXISTS idx_wallet_addresses_group_id ON wallet_addresses(group_id);
CREATE INDEX IF NOT EXISTS idx_wallet_groups_user_id ON wallet_groups(user_id);
CREATE INDEX IF NOT EXISTS idx_tokens_contract_address ON tokens(contract_address);
CREATE INDEX IF NOT EXISTS idx_tokens_chain_id ON tokens(chain_id);

-- 插入常用BSC代币数据
INSERT INTO tokens (symbol, name, contract_address, decimals, chain_id) VALUES
('USDT', 'Tether USD', '0x55d398326f99059fF775485246999027B3197955', 18, 56),
('USDC', 'USD Coin', '0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d', 18, 56),
('BUSD', 'Binance USD', '0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56', 18, 56),
('WBNB', 'Wrapped BNB', '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c', 18, 56)
ON CONFLICT (contract_address) DO NOTHING;

-- 创建更新时间触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 为各表添加更新时间触发器
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_wallet_groups_updated_at BEFORE UPDATE ON wallet_groups
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_wallet_addresses_updated_at BEFORE UPDATE ON wallet_addresses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
