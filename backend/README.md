# 钱包地址管理后端服务

这是一个用于管理钱包地址并查询BSC网络余额的后端服务。支持用户管理、钱包地址分组、批量余额查询等功能。

## 功能特性

- ✅ 用户账号管理（注册、登录）
- ✅ 钱包地址管理（添加、删除、标签）
- ✅ 地址分组管理（创建分组、分组管理）
- ✅ 余额查询（BNB + ERC20代币）
- ✅ 批量余额查询（按分组）
- ✅ 支持主流BSC代币（USDT、USDC、BUSD、WBNB）

## 系统架构

```
├── cmd/
│   └── main.go              # 主程序入口
├── internal/
│   ├── models/              # 数据模型
│   ├── database/            # 数据库连接
│   ├── services/            # 业务逻辑层
│   └── handlers/            # HTTP处理器
├── migrations/              # 数据库迁移文件
├── bindings/                # 智能合约Go绑定
├── docker-compose.yml       # Docker配置
└── Makefile                 # 构建脚本
```

## 快速开始

### 1. 环境要求

- Go 1.22+
- Docker & Docker Compose
- PostgreSQL 15+

### 2. 启动数据库

```bash
# 启动PostgreSQL数据库
make docker-up

# 查看数据库日志
make docker-logs
```

### 3. 安装依赖

```bash
# 安装Go依赖
make install
```

### 4. 启动服务

```bash
# 开发模式启动
make dev

# 或者直接运行
make run
```

服务将在 `http://localhost:8888` 启动

### 5. 验证服务

```bash
# 健康检查
curl http://localhost:8888/health

# 查看服务信息
curl http://localhost:8888/
```

## API 文档

### 用户管理

#### 创建用户
```bash
POST /api/users
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

#### 用户登录
```bash
POST /api/users/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

#### 获取用户信息
```bash
GET /api/users/{user_id}
```

### 钱包分组管理

#### 创建分组
```bash
POST /api/users/{user_id}/groups
Content-Type: application/json

{
  "name": "主要钱包",
  "description": "日常使用的主要钱包地址"
}
```

#### 获取用户分组
```bash
GET /api/users/{user_id}/groups
```

#### 删除分组
```bash
DELETE /api/users/{user_id}/groups/{group_id}
```

### 钱包地址管理

#### 添加钱包地址
```bash
POST /api/users/{user_id}/addresses
Content-Type: application/json

{
  "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
  "label": "主钱包",
  "group_id": 1
}
```

#### 获取用户所有地址
```bash
GET /api/users/{user_id}/addresses
```

#### 获取分组中的地址
```bash
GET /api/users/{user_id}/groups/{group_id}/addresses
```

#### 删除地址
```bash
DELETE /api/users/{user_id}/addresses/{address_id}
```

### 余额查询

#### 查询单个地址余额
```bash
# 查询所有支持的代币余额
GET /api/balance/address?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045

# 查询指定代币余额
GET /api/balance/address?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045&tokens=0x55d398326f99059fF775485246999027B3197955,0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d
```

#### 查询分组余额
```bash
GET /api/users/{user_id}/groups/{group_id}/balance
```

#### 查询用户总余额
```bash
GET /api/users/{user_id}/balance
```

#### 批量查询分组余额
```bash
GET /api/users/{user_id}/groups/balance?group_ids=1,2,3
```

### 代币管理

#### 获取支持的代币列表
```bash
GET /api/tokens
```

## 使用示例

### 完整工作流程示例

```bash
# 1. 创建用户
curl -X POST http://localhost:8888/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "email": "alice@example.com",
    "password": "password123"
  }'

# 2. 用户登录
curl -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "alice",
    "password": "password123"
  }'

# 3. 创建钱包分组
curl -X POST http://localhost:8888/api/users/1/groups \
  -H "Content-Type: application/json" \
  -d '{
    "name": "DeFi钱包",
    "description": "用于DeFi操作的钱包地址"
  }'

# 4. 添加钱包地址
curl -X POST http://localhost:8888/api/users/1/addresses \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    "label": "Vitalik钱包",
    "group_id": 1
  }'

# 5. 查询分组余额
curl http://localhost:8888/api/users/1/groups/1/balance

# 6. 查询用户总余额
curl http://localhost:8888/api/users/1/balance
```

### Python客户端示例

```python
import requests
import json

class WalletManagerClient:
    def __init__(self, base_url="http://localhost:8888"):
        self.base_url = base_url
        self.user_id = None
    
    def create_user(self, username, email, password):
        response = requests.post(f"{self.base_url}/api/users", json={
            "username": username,
            "email": email,
            "password": password
        })
        return response.json()
    
    def login(self, username, password):
        response = requests.post(f"{self.base_url}/api/users/login", json={
            "username": username,
            "password": password
        })
        data = response.json()
        if "user_id" in data:
            self.user_id = data["user_id"]
        return data
    
    def create_group(self, name, description=""):
        response = requests.post(f"{self.base_url}/api/users/{self.user_id}/groups", json={
            "name": name,
            "description": description
        })
        return response.json()
    
    def add_address(self, address, label="", group_id=None):
        data = {"address": address, "label": label}
        if group_id:
            data["group_id"] = group_id
        
        response = requests.post(f"{self.base_url}/api/users/{self.user_id}/addresses", json=data)
        return response.json()
    
    def get_group_balance(self, group_id):
        response = requests.get(f"{self.base_url}/api/users/{self.user_id}/groups/{group_id}/balance")
        return response.json()
    
    def get_total_balance(self):
        response = requests.get(f"{self.base_url}/api/users/{self.user_id}/balance")
        return response.json()

# 使用示例
client = WalletManagerClient()

# 创建用户并登录
client.create_user("alice", "alice@example.com", "password123")
client.login("alice", "password123")

# 创建分组
group = client.create_group("主要钱包", "日常使用的钱包")
group_id = group["id"]

# 添加地址
client.add_address("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", "Vitalik", group_id)

# 查询余额
balance = client.get_group_balance(group_id)
print(json.dumps(balance, indent=2))
```

## 数据库设计

### 表结构

- `users` - 用户表
- `wallet_groups` - 钱包分组表
- `wallet_addresses` - 钱包地址表
- `tokens` - 代币信息表

### 关系图

```
users (1) -----> (n) wallet_groups
users (1) -----> (n) wallet_addresses
wallet_groups (1) -----> (n) wallet_addresses
```

## 部署

### Docker部署

```bash
# 构建镜像
docker build -t wallet-manager .

# 运行服务
docker-compose up -d
```

### 生产环境配置

1. 修改数据库密码
2. 配置环境变量
3. 启用HTTPS
4. 配置反向代理
5. 设置监控和日志

## 开发

### 代码结构

- `models/` - 数据模型定义
- `services/` - 业务逻辑层
- `handlers/` - HTTP处理器
- `database/` - 数据库连接管理

### 添加新功能

1. 在 `models/` 中定义新的数据模型
2. 在 `services/` 中实现业务逻辑
3. 在 `handlers/` 中添加HTTP接口
4. 在 `main.go` 中注册路由

### 运行测试

```bash
make test
```

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查PostgreSQL是否启动
   - 验证数据库配置

2. **合约调用失败**
   - 检查BSC网络连接
   - 验证合约地址

3. **地址格式错误**
   - 确保地址以0x开头
   - 验证地址长度为42字符

### 日志查看

```bash
# 查看应用日志
make docker-logs

# 查看数据库日志
docker-compose logs postgres
```

## 许可证

MIT License
