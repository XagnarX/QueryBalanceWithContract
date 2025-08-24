-- 创建区块链网络表
CREATE TABLE IF NOT EXISTS chains (
    id SERIAL PRIMARY KEY,
    chain_id INTEGER UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    symbol VARCHAR(10) NOT NULL, -- 原生代币符号，如 BNB, ETH
    is_testnet BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    block_explorer_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 创建RPC端点表
CREATE TABLE IF NOT EXISTS rpc_endpoints (
    id SERIAL PRIMARY KEY,
    chain_id INTEGER NOT NULL REFERENCES chains(chain_id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    priority INTEGER DEFAULT 0, -- 优先级，数字越小优先级越高
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(chain_id, url)
);

-- 创建余额查询合约表
CREATE TABLE IF NOT EXISTS balance_contracts (
    id SERIAL PRIMARY KEY,
    chain_id INTEGER NOT NULL REFERENCES chains(chain_id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    priority INTEGER DEFAULT 0, -- 优先级，数字越小优先级越高
    abi_json TEXT, -- 存储合约ABI
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(chain_id, contract_address)
);

-- 修改tokens表，添加链关联和更多字段
ALTER TABLE tokens DROP CONSTRAINT IF EXISTS tokens_contract_address_key;
ALTER TABLE tokens ADD CONSTRAINT tokens_contract_address_chain_key UNIQUE(contract_address, chain_id);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_chains_chain_id ON chains(chain_id);
CREATE INDEX IF NOT EXISTS idx_rpc_endpoints_chain_id ON rpc_endpoints(chain_id);
CREATE INDEX IF NOT EXISTS idx_rpc_endpoints_priority ON rpc_endpoints(priority);
CREATE INDEX IF NOT EXISTS idx_balance_contracts_chain_id ON balance_contracts(chain_id);
CREATE INDEX IF NOT EXISTS idx_balance_contracts_priority ON balance_contracts(priority);

-- 插入常用区块链网络数据
INSERT INTO chains (chain_id, name, symbol, is_testnet, is_active, block_explorer_url) VALUES
(56, 'BSC Mainnet', 'BNB', FALSE, TRUE, 'https://bscscan.com'),
(97, 'BSC Testnet', 'tBNB', TRUE, TRUE, 'https://testnet.bscscan.com'),
(1, 'Ethereum Mainnet', 'ETH', FALSE, TRUE, 'https://etherscan.io'),
(5, 'Ethereum Goerli', 'ETH', TRUE, FALSE, 'https://goerli.etherscan.io'),
(137, 'Polygon Mainnet', 'MATIC', FALSE, TRUE, 'https://polygonscan.com'),
(80001, 'Polygon Mumbai', 'MATIC', TRUE, FALSE, 'https://mumbai.polygonscan.com')
ON CONFLICT (chain_id) DO NOTHING;

-- 插入BSC RPC端点
INSERT INTO rpc_endpoints (chain_id, name, url, priority) VALUES
(56, 'BSC Official', 'https://bsc-dataseed.binance.org/', 1),
(56, 'BSC Official 2', 'https://bsc-dataseed1.defibit.io/', 2),
(56, 'BSC Official 3', 'https://bsc-dataseed1.ninicoin.io/', 3),
(56, 'NodeReal', 'https://bsc.nodereal.io', 4),
(97, 'BSC Testnet', 'https://data-seed-prebsc-1-s1.binance.org:8545/', 1),
(1, 'Ethereum Mainnet', 'https://mainnet.infura.io/v3/YOUR_INFURA_KEY', 1),
(137, 'Polygon Mainnet', 'https://polygon-rpc.com/', 1)
ON CONFLICT (chain_id, url) DO NOTHING;

-- 插入BalanceChecker合约地址
INSERT INTO balance_contracts (chain_id, name, contract_address, priority) VALUES
(56, 'BalanceChecker BSC', '0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b', 1),
(97, 'BalanceChecker BSC Testnet', '0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b', 1)
ON CONFLICT (chain_id, contract_address) DO NOTHING;

-- 为新表添加更新时间触发器
CREATE TRIGGER update_chains_updated_at BEFORE UPDATE ON chains
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_rpc_endpoints_updated_at BEFORE UPDATE ON rpc_endpoints
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_balance_contracts_updated_at BEFORE UPDATE ON balance_contracts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
