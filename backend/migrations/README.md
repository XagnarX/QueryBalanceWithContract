# Database Migrations

## 概述

此目录包含数据库初始化和迁移脚本。

## 文件说明

### `init_database.sql`
- **主初始化脚本** - 包含创建整个数据库结构所需的完整SQL
- 合并了所有历史迁移文件的最终状态
- 可独立运行，无需其他依赖
- 支持重复执行（幂等性）

### `backup/` 目录
包含原始的分步迁移文件（已合并）：
- `001_init.sql` - 初始基础表结构
- `002_add_multi_chain_support.sql` - 多链支持功能
- `003_add_user_custom_resources.sql` - 用户自定义资源
- `004_remove_system_tables.sql` - 清理系统表

## 使用方法

### 新数据库初始化
对于全新的数据库，直接运行：
```bash
psql -d your_database -f init_database.sql
```

### 数据库结构

#### 核心表
- `users` - 用户管理
- `chains` - 区块链网络配置
- `wallet_groups` - 钱包分组
- `wallet_addresses` - 钱包地址
- `balance_contracts` - 余额查询合约
- `user_rpc_endpoints` - 用户自定义RPC端点
- `user_tokens` - 用户自定义代币

#### 特性
- ✅ 自动时间戳更新
- ✅ 完整的外键约束
- ✅ 性能优化索引
- ✅ 多链架构支持
- ✅ 用户数据隔离

## 架构说明

### 用户自定义资源
系统采用用户自定义资源模式：
- 每个用户可以添加自己的RPC端点
- 每个用户可以定义自己的Token合约
- 系统只提供基础的链配置和合约地址

### 数据隔离
- 所有用户数据通过 `user_id` 外键进行隔离
- 用户删除时级联删除相关数据
- 支持多用户并发使用

## 注意事项

1. **数据库权限**：确保数据库用户具有创建表、索引和触发器的权限
2. **PostgreSQL版本**：建议使用PostgreSQL 12+
3. **备份**：在生产环境执行前请备份现有数据
4. **测试**：建议先在测试环境验证脚本

## 开发指南

如需添加新的数据库更改：
1. 直接修改 `init_database.sql`
2. 确保更改是向后兼容的
3. 更新此README文档
4. 在测试环境验证更改

## 故障排除

### 常见问题
1. **权限不足**：确保数据库用户有足够权限
2. **表已存在**：脚本使用 `IF NOT EXISTS`，可安全重复运行
3. **外键约束**：检查引用表是否已正确创建

### 重建数据库
如需完全重建：
```sql
-- 删除所有表（谨慎使用！）
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO your_user;

-- 然后重新运行初始化脚本
\i init_database.sql
```