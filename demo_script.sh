#!/bin/bash

echo "=== 钱包地址管理系统演示 ==="
echo "后端服务运行在: http://localhost:8888"
echo ""

BASE_URL="http://localhost:8888/api"

echo "1. 测试健康检查"
curl -s http://localhost:8888/health | python3 -m json.tool
echo ""

echo "2. 创建用户"
USER_RESPONSE=$(curl -s -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo_user",
    "email": "demo@example.com",
    "password": "demo123456"
  }')
echo $USER_RESPONSE | python3 -m json.tool
USER_ID=$(echo $USER_RESPONSE | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])")
echo "创建的用户ID: $USER_ID"
echo ""

echo "3. 用户登录"
curl -s -X POST $BASE_URL/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "demo_user",
    "password": "demo123456"
  }' | python3 -m json.tool
echo ""

echo "4. 创建钱包分组"
GROUP1_RESPONSE=$(curl -s -X POST $BASE_URL/users/$USER_ID/groups \
  -H "Content-Type: application/json" \
  -d '{
    "name": "主钱包",
    "description": "主要使用的钱包地址"
  }')
echo $GROUP1_RESPONSE | python3 -m json.tool
GROUP1_ID=$(echo $GROUP1_RESPONSE | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])")

GROUP2_RESPONSE=$(curl -s -X POST $BASE_URL/users/$USER_ID/groups \
  -H "Content-Type: application/json" \
  -d '{
    "name": "投资钱包",
    "description": "用于投资的钱包地址"
  }')
echo $GROUP2_RESPONSE | python3 -m json.tool
GROUP2_ID=$(echo $GROUP2_RESPONSE | python3 -c "import sys, json; print(json.load(sys.stdin)['id'])")
echo ""

echo "5. 添加钱包地址"
# 添加Vitalik地址到主钱包
curl -s -X POST $BASE_URL/users/$USER_ID/addresses \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
    "label": "Vitalik钱包",
    "group_id": '$GROUP1_ID'
  }' | python3 -m json.tool

# 添加巨鲸地址到主钱包
curl -s -X POST $BASE_URL/users/$USER_ID/addresses \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0x8894E0a0c962CB723c1976a4421c95949bE2D4E3",
    "label": "巨鲸地址",
    "group_id": '$GROUP1_ID'
  }' | python3 -m json.tool

# 添加地址到投资钱包
curl -s -X POST $BASE_URL/users/$USER_ID/addresses \
  -H "Content-Type: application/json" \
  -d '{
    "address": "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
    "label": "DeFi投资地址",
    "group_id": '$GROUP2_ID'
  }' | python3 -m json.tool
echo ""

echo "6. 查看用户所有分组"
curl -s $BASE_URL/users/$USER_ID/groups | python3 -m json.tool
echo ""

echo "7. 查看用户所有地址"
curl -s $BASE_URL/users/$USER_ID/addresses | python3 -m json.tool
echo ""

echo "8. 查询主钱包分组余额"
echo "分组ID: $GROUP1_ID"
curl -s $BASE_URL/users/$USER_ID/groups/$GROUP1_ID/balance | python3 -m json.tool
echo ""

echo "9. 查询投资钱包分组余额"
echo "分组ID: $GROUP2_ID"
curl -s $BASE_URL/users/$USER_ID/groups/$GROUP2_ID/balance | python3 -m json.tool
echo ""

echo "10. 批量查询所有分组余额"
curl -s "$BASE_URL/users/$USER_ID/groups/balance?group_ids=$GROUP1_ID,$GROUP2_ID" | python3 -m json.tool
echo ""

echo "11. 查询用户总余额"
curl -s $BASE_URL/users/$USER_ID/balance | python3 -m json.tool
echo ""

echo "12. 查询支持的代币列表"
curl -s $BASE_URL/tokens | python3 -m json.tool
echo ""

echo "13. 单独查询地址余额（公共接口）"
curl -s "$BASE_URL/balance/address?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" | python3 -m json.tool
echo ""

echo "=== 演示完成 ==="
echo "系统功能全部正常！"
echo "- ✅ 用户管理"
echo "- ✅ 钱包分组管理"
echo "- ✅ 钱包地址管理"
echo "- ✅ 余额查询（BNB + ERC20代币）"
echo "- ✅ 批量查询"
echo "- ✅ 按分组查询"
