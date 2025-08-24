# BSC BalanceChecker 余额查询工具

这是一个用于批量查询BSC (Binance Smart Chain) 主网上地址余额的Go工具。

## 合约信息

- **合约地址**: `0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b`
- **网络**: BSC主网
- **RPC端点**: `https://bsc-dataseed1.binance.org/`

## 功能特性

- 查询单个地址的BNB余额
- 批量查询多个地址的BNB余额
- 查询ERC20代币余额
- 批量查询多个代币、多个地址的余额
- 一次性查询地址的所有余额（BNB + 代币）

## 支持的代币

| 代币 | 合约地址 |
|------|----------|
| USDT | `0x55d398326f99059fF775485246999027B3197955` |
| USDC | `0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d` |
| BUSD | `0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56` |
| WBNB | `0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c` |

## 安装和使用

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 运行完整示例

```bash
go run main.go
```

这将运行一个完整的示例，展示所有功能：
- 查询Vitalik地址等示例地址的BNB余额
- 批量查询多个地址的余额
- 查询各种代币余额

### 3. 运行简单查询

```bash
# 只查询BNB余额
go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045

# 查询BNB + USDT余额
go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 0x55d398326f99059fF775485246999027B3197955

# 查询BNB + 多个代币余额
go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 0x55d398326f99059fF775485246999027B3197955 0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d
```

### 4. 使用帮助

```bash
go run simple_query.go
```

## 代码结构

```
├── contracts/              # Solidity合约源码
│   └── BalanceChecker.sol  # 余额查询合约
├── bindings/               # Go合约绑定文件
│   └── BalanceChecker.go   # 自动生成的绑定
├── abi/                    # ABI文件
│   └── BalanceChecker.abi  # 合约ABI
├── main.go                 # 完整功能示例
├── simple_query.go         # 简单查询工具
└── go.mod                  # Go模块文件
```

## API说明

### BalanceChecker合约方法

1. `getETHBalance(address target)` - 查询单个地址的BNB余额
2. `getETHBalances(address[] targets)` - 批量查询BNB余额
3. `getERC20Balance(address token, address target)` - 查询ERC20代币余额
4. `getERC20Balances(address token, address[] targets)` - 批量查询代币余额
5. `getAddressBalances(address target, address[] tokens)` - 查询地址的所有余额
6. `getMultipleAddressBalances(address[] targets, address[] tokens)` - 批量查询所有余额

### Go绑定使用示例

```go
package main

import (
    "fmt"
    "log"
    
    "QueryBalanceWithContract/bindings"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // 连接BSC主网
    client, err := ethclient.Dial("https://bsc-dataseed1.binance.org/")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // 创建合约实例
    contractAddr := common.HexToAddress("0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b")
    contract, err := bindings.NewBalanceChecker(contractAddr, client)
    if err != nil {
        log.Fatal(err)
    }

    // 查询BNB余额
    addr := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
    balance, err := contract.GetETHBalance(nil, addr)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("BNB余额: %s wei\n", balance.String())
}
```

## 注意事项

1. 所有余额返回值都是以最小单位计算的（wei对于BNB，代币的最小单位对于ERC20）
2. 脚本中已经包含了转换函数来显示人类可读的格式
3. 合约调用是只读的，不会产生gas费用
4. 请确保网络连接正常，因为需要访问BSC主网

## 依赖

- Go 1.22+
- github.com/ethereum/go-ethereum v1.13.10+

## 许可证

MIT License
