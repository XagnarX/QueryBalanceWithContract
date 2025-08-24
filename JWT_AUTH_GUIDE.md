# JWT认证系统使用指南

## 🔐 认证系统概述

我们已经成功为钱包地址管理系统添加了JWT Token认证机制，确保用户数据的安全性和隔离性。

## ✅ 已实现功能

### 1. JWT Token生成与验证
- ✅ 用户登录时自动生成JWT Token
- ✅ Token包含用户ID、用户名等信息
- ✅ Token有效期为24小时
- ✅ 使用HMAC-SHA256算法签名

### 2. 认证中间件
- ✅ 自动验证请求中的Authorization Header
- ✅ 支持Bearer Token格式
- ✅ 防止跨用户数据访问
- ✅ 详细的错误提示

### 3. 权限控制
- ✅ 用户只能访问自己的数据
- ✅ 公共接口无需认证
- ✅ 所有用户相关操作都需要认证

## 📋 API接口分类

### 公共接口（无需认证）
```bash
POST /api/users              # 创建用户
POST /api/users/login        # 用户登录
GET  /api/tokens             # 获取代币列表  
GET  /api/balance/address    # 查询地址余额
GET  /health                 # 健康检查
```

### 受保护接口（需要认证）
```bash
GET    /api/users/{user_id}                              # 获取用户信息
POST   /api/users/{user_id}/groups                       # 创建分组
GET    /api/users/{user_id}/groups                       # 获取分组列表
DELETE /api/users/{user_id}/groups/{group_id}            # 删除分组
POST   /api/users/{user_id}/addresses                    # 添加地址
GET    /api/users/{user_id}/addresses                    # 获取地址列表
DELETE /api/users/{user_id}/addresses/{address_id}       # 删除地址
GET    /api/users/{user_id}/balance                      # 查询用户总余额
GET    /api/users/{user_id}/groups/{group_id}/balance    # 查询分组余额
GET    /api/users/{user_id}/groups/balance               # 批量查询分组余额
```

## 🔧 使用方法

### 1. 用户登录获取Token

```bash
curl -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "your_username",
    "password": "your_password"
  }'
```

**响应示例：**
```json
{
  "message": "登录成功",
  "user_id": 2,
  "username": "your_username",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "token_type": "Bearer",
  "expires_in": 86400
}
```

### 2. 使用Token访问受保护的API

```bash
curl -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  http://localhost:8888/api/users/{user_id}/groups
```

### 3. 完整示例

```bash
# 1. 登录获取token
TOKEN=$(curl -s -X POST http://localhost:8888/api/users/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"password123"}' | \
  python3 -c "import sys, json; print(json.load(sys.stdin)['token'])")

# 2. 使用token创建分组
curl -X POST http://localhost:8888/api/users/1/groups \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name":"我的钱包","description":"主要钱包地址"}'

# 3. 使用token添加地址
curl -X POST http://localhost:8888/api/users/1/addresses \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "address":"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    "label":"Vitalik钱包",
    "group_id":1
  }'

# 4. 使用token查询余额
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8888/api/users/1/groups/1/balance
```

## 🛡️ 安全特性

### 1. Token安全
- **过期时间**: Token有效期24小时，降低泄露风险
- **签名验证**: 使用HMAC-SHA256确保Token完整性
- **用户绑定**: Token包含用户ID，防止令牌滥用

### 2. 权限隔离
- **用户隔离**: 用户只能访问自己的数据
- **路径验证**: 验证URL中的user_id与Token中的用户ID匹配
- **详细错误**: 提供清晰的认证错误信息

### 3. 错误处理
```json
// 缺少Token
{"error": "缺少认证token"}

// Token格式错误
{"error": "无效的token格式，应为: Bearer <token>"}

// Token过期或无效
{"error": "无效的token: token已过期"}

// 跨用户访问
{"error": "无权限访问其他用户的数据"}
```

## 🔄 Token刷新

当前实现包含Token刷新功能（在Token过期前1小时内可刷新）：

```bash
curl -X POST http://localhost:8888/api/auth/refresh \
  -H "Authorization: Bearer YOUR_CURRENT_TOKEN"
```

## 💡 最佳实践

### 1. 客户端使用
```javascript
// JavaScript示例
class WalletAPIClient {
  constructor(baseURL) {
    this.baseURL = baseURL;
    this.token = localStorage.getItem('auth_token');
  }

  async login(username, password) {
    const response = await fetch(`${this.baseURL}/api/users/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    
    const data = await response.json();
    if (data.token) {
      this.token = data.token;
      localStorage.setItem('auth_token', this.token);
    }
    return data;
  }

  async apiCall(endpoint, options = {}) {
    const headers = {
      'Content-Type': 'application/json',
      ...options.headers
    };
    
    if (this.token) {
      headers.Authorization = `Bearer ${this.token}`;
    }

    return fetch(`${this.baseURL}${endpoint}`, {
      ...options,
      headers
    });
  }
}
```

### 2. Python客户端示例
```python
import requests

class WalletClient:
    def __init__(self, base_url="http://localhost:8888"):
        self.base_url = base_url
        self.token = None
    
    def login(self, username, password):
        response = requests.post(f"{self.base_url}/api/users/login", json={
            "username": username,
            "password": password
        })
        data = response.json()
        if "token" in data:
            self.token = data["token"]
        return data
    
    def _headers(self):
        headers = {"Content-Type": "application/json"}
        if self.token:
            headers["Authorization"] = f"Bearer {self.token}"
        return headers
    
    def get_user_groups(self, user_id):
        response = requests.get(
            f"{self.base_url}/api/users/{user_id}/groups",
            headers=self._headers()
        )
        return response.json()
```

## 🚀 生产环境配置

### 1. 环境变量配置
```bash
# 设置JWT密钥（生产环境必须更改）
export JWT_SECRET_KEY="your-very-secure-secret-key-here"

# 设置Token过期时间（秒）
export JWT_EXPIRES_IN=86400
```

### 2. 安全建议
- **更换密钥**: 生产环境必须使用强随机密钥
- **HTTPS**: 生产环境必须使用HTTPS传输
- **日志记录**: 记录认证失败尝试
- **频率限制**: 添加登录频率限制

## 📊 测试结果

我们的测试验证了以下功能：

✅ **登录认证**: 成功返回JWT Token  
✅ **Token验证**: 正确验证Bearer Token格式  
✅ **权限控制**: 阻止跨用户数据访问  
✅ **错误处理**: 清晰的错误消息  
✅ **API保护**: 受保护的API需要认证  
✅ **公共接口**: 公共API无需认证正常工作  

## 🎯 总结

JWT认证系统已经完全集成到钱包地址管理系统中，提供了：

- **安全性**: 保护用户数据不被未授权访问
- **隔离性**: 确保用户只能访问自己的数据
- **易用性**: 标准的JWT Token使用方式
- **扩展性**: 易于集成到前端应用和API客户端

系统现在可以安全地支持多用户环境，每个用户都能独立管理自己的钱包地址和分组！
