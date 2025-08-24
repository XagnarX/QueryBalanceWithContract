-- 003_add_user_custom_resources.sql
-- 添加用户自定义资源表

-- 用户自定义RPC端点表
CREATE TABLE IF NOT EXISTS user_rpc_endpoints (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chain_id INTEGER NOT NULL,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- 唯一约束
    CONSTRAINT unique_user_rpc_url UNIQUE (user_id, chain_id, url)
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_user_rpc_user_id ON user_rpc_endpoints(user_id);
CREATE INDEX IF NOT EXISTS idx_user_rpc_chain_id ON user_rpc_endpoints(chain_id);
CREATE INDEX IF NOT EXISTS idx_user_rpc_active ON user_rpc_endpoints(is_active);

-- 用户自定义代币表
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
    
    -- 唯一约束
    CONSTRAINT unique_user_token_contract UNIQUE (user_id, chain_id, contract_address)
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_user_token_user_id ON user_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_user_token_chain_id ON user_tokens(chain_id);
CREATE INDEX IF NOT EXISTS idx_user_token_active ON user_tokens(is_active);
CREATE INDEX IF NOT EXISTS idx_user_token_contract ON user_tokens(contract_address);

-- 创建更新时间触发器函数（如果不存在）
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 为user_rpc_endpoints表创建触发器
DROP TRIGGER IF EXISTS trigger_user_rpc_endpoints_updated_at ON user_rpc_endpoints;
CREATE TRIGGER trigger_user_rpc_endpoints_updated_at
    BEFORE UPDATE ON user_rpc_endpoints
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- 为user_tokens表创建触发器
DROP TRIGGER IF EXISTS trigger_user_tokens_updated_at ON user_tokens;
CREATE TRIGGER trigger_user_tokens_updated_at
    BEFORE UPDATE ON user_tokens
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();