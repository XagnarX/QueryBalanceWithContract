# 局域网访问配置指南

## 概述

本指南说明如何让局域网内的其他设备访问 Wallet Manager 服务。

## 当前配置

- **本机 IP**: `192.168.71.9`
- **前端端口**: `3000`
- **后端端口**: `8888`

## 快速启动

### 方式 1: 使用启动脚本（推荐）

```bash
# 启动 LAN 模式
./start-lan.sh

# 停止服务
./stop.sh
```

### 方式 2: 手动启动

#### 启动后端
```bash
cd backend
go run cmd/main.go
```

后端会自动监听 `0.0.0.0:8888`，可从局域网访问。

#### 启动前端
```bash
cd frontend
npm run dev
```

前端会监听 `0.0.0.0:3000`，可从局域网访问。

## 访问地址

### 从本机访问
- 前端: http://localhost:3000
- 后端: http://localhost:8888

### 从局域网其他设备访问
- 前端: http://192.168.71.9:3000
- 后端: http://192.168.71.9:8888/api

## 防火墙配置

### macOS
确保防火墙允许端口 3000 和 8888：

```bash
# 查看防火墙状态
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --getglobalstate

# 如果防火墙开启，添加应用程序例外
# (启动服务后，系统会自动提示是否允许)
```

### 如果需要临时关闭防火墙测试
```bash
# 关闭防火墙（不推荐）
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --setglobalstate off

# 重新开启防火墙
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --setglobalstate on
```

## 常见问题

### 1. 无法从其他设备访问

**检查步骤：**

1. 确认本机 IP 地址
```bash
ifconfig | grep "inet " | grep -v 127.0.0.1
```

2. 检查服务是否在正确端口监听
```bash
lsof -i :3000
lsof -i :8888
```

3. 检查防火墙设置
```bash
sudo /usr/libexec/ApplicationFirewall/socketfilterfw --listapps
```

4. 确认设备在同一局域网
   - 手机/平板连接同一个 WiFi
   - 其他电脑连接同一个局域网

### 2. CORS 错误

后端已配置允许所有来源的跨域请求：
```go
Access-Control-Allow-Origin: *
```

如果仍有 CORS 问题，检查前端是否正确配置了 API 代理。

### 3. API 请求失败

前端通过 Vite 代理访问后端 API：
- 前端请求: `/api/*`
- 自动转发到: `http://localhost:8888/api/*`

确保：
1. 后端服务在 8888 端口运行
2. Vite 代理配置正确（vite.config.js）

## 网络配置详情

### 后端 (Go/Gin)
- **监听地址**: `0.0.0.0:8888`
- **CORS**: 允许所有来源 (`*`)
- **支持方法**: GET, POST, PUT, DELETE, OPTIONS

### 前端 (Vue/Vite)
- **监听地址**: `0.0.0.0:3000`
- **开发服务器**: Vite dev server
- **API 代理**: `/api` → `http://localhost:8888`

## 安全建议

1. **仅在可信网络使用**
   - 此配置允许局域网内任何设备访问
   - 不要在公共 WiFi 上使用

2. **生产环境配置**
   - 使用环境变量管理配置
   - 配置具体的 CORS 允许来源
   - 使用 HTTPS/TLS 加密

3. **防火墙**
   - 保持防火墙开启
   - 只开放必要端口

## 切换配置

### 切换到本地模式（仅本机访问）

修改 `frontend/vite.config.js`:
```javascript
server: {
  host: 'localhost',  // 改为 localhost
  port: 3000
}
```

修改 `backend/cmd/main.go`:
```go
host := "127.0.0.1"  // 改为 127.0.0.1
```

### 切换到 LAN 模式（局域网访问）

修改 `frontend/vite.config.js`:
```javascript
server: {
  host: '0.0.0.0',  // 监听所有接口
  port: 3000
}
```

修改 `backend/cmd/main.go`:
```go
host := "0.0.0.0"  // 监听所有接口
```

## 验证配置

### 测试后端可访问性

从其他设备浏览器访问：
```
http://192.168.71.9:8888/health
```

应该返回：
```json
{
  "status": "ok",
  "message": "钱包管理服务运行正常"
}
```

### 测试前端可访问性

从其他设备浏览器访问：
```
http://192.168.71.9:3000
```

应该能看到前端页面。

## 移动设备访问

### iOS/Android 浏览器

直接在浏览器输入：
```
http://192.168.71.9:3000
```

### 注意事项

1. **确保连接同一 WiFi**
2. **某些功能可能需要桌面浏览器模式**
3. **Web3 钱包插件在移动浏览器中可能不可用**

## 故障排除命令

```bash
# 查看占用端口的进程
lsof -i :3000
lsof -i :8888

# 杀死占用端口的进程
lsof -ti:3000 | xargs kill
lsof -ti:8888 | xargs kill

# 查看本机所有 IP 地址
ifconfig

# 测试端口是否可访问（从其他设备）
# 在其他设备上运行
telnet 192.168.71.9 3000
telnet 192.168.71.9 8888
```

## 更新 IP ��址

如果本机 IP 地址变更：

1. 获取新 IP
```bash
ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | head -1
```

2. 更新 `backend/cmd/main.go` 中的日志输出
```go
log.Printf("局域网访问: http://YOUR_NEW_IP:%s", port)
```

3. 重启服务
```bash
./stop.sh
./start-lan.sh
```
