-- 004_remove_system_tables.sql
-- 删除系统表，完全使用用户自定义表

-- 首先删除依赖的外键约束和触发器
DROP TRIGGER IF EXISTS update_rpc_endpoints_updated_at ON rpc_endpoints;
DROP TRIGGER IF EXISTS update_balance_contracts_updated_at ON balance_contracts;

-- 删除索引
DROP INDEX IF EXISTS idx_rpc_endpoints_chain_id;
DROP INDEX IF EXISTS idx_rpc_endpoints_priority;
DROP INDEX IF EXISTS idx_balance_contracts_chain_id;
DROP INDEX IF EXISTS idx_balance_contracts_priority;

-- 删除系统RPC端点表
DROP TABLE IF EXISTS rpc_endpoints CASCADE;

-- 删除余额查询合约表（如果不再需要的话）
-- DROP TABLE IF EXISTS balance_contracts CASCADE;

-- 删除系统token表的链约束（如果tokens表还存在）
ALTER TABLE tokens DROP CONSTRAINT IF EXISTS tokens_contract_address_chain_key;

-- 完全删除系统tokens表
DROP TABLE IF EXISTS tokens CASCADE;

-- 注意：保留chains表和balance_contracts表，因为这些是基础配置
-- 如果你也想删除这些表，取消下面的注释：
-- DROP TABLE IF EXISTS balance_contracts CASCADE;
-- DROP TABLE IF EXISTS chains CASCADE;