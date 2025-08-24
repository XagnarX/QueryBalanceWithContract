package main

import (
	"fmt"
	"log"
	"math/big"

	"QueryBalanceWithContract/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// BSC主网RPC端点
	BSC_MAINNET_RPC = "https://bsc-dataseed1.binance.org/"
	// BalanceChecker合约地址
	CONTRACT_ADDRESS = "0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b"

	// 常用BSC代币地址
	USDT_BSC = "0x55d398326f99059fF775485246999027B3197955" // USDT on BSC
	USDC_BSC = "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d" // USDC on BSC
	BUSD_BSC = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56" // BUSD on BSC
	WBNB_BSC = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c" // WBNB on BSC
)

func main() {
	// 连接到BSC主网
	client, err := ethclient.Dial(BSC_MAINNET_RPC)
	if err != nil {
		log.Fatalf("连接BSC主网失败: %v", err)
	}
	defer client.Close()

	// 创建合约实例
	contractAddr := common.HexToAddress(CONTRACT_ADDRESS)
	balanceChecker, err := bindings.NewBalanceChecker(contractAddr, client)
	if err != nil {
		log.Fatalf("创建合约实例失败: %v", err)
	}

	fmt.Println("=== BSC BalanceChecker 示例 ===")
	fmt.Printf("合约地址: %s\n", CONTRACT_ADDRESS)
	fmt.Printf("RPC端点: %s\n", BSC_MAINNET_RPC)
	fmt.Println()

	// 示例地址（可以替换成您想查询的地址）
	testAddresses := []string{
		"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045", // Vitalik地址
		"0x8894E0a0c962CB723c1976a4421c95949bE2D4E3", // 示例地址1
		"0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE", // 示例地址2
	}

	// 1. 查询单个地址的BNB余额
	fmt.Println("1. 查询单个地址的BNB余额")
	targetAddr := common.HexToAddress(testAddresses[0])
	bnbBalance, err := balanceChecker.GetETHBalance(nil, targetAddr)
	if err != nil {
		log.Printf("查询BNB余额失败: %v", err)
	} else {
		fmt.Printf("地址 %s 的BNB余额: %s BNB\n", targetAddr.Hex(), weiToBNB(bnbBalance))
	}
	fmt.Println()

	// 2. 批量查询多个地址的BNB余额
	fmt.Println("2. 批量查询多个地址的BNB余额")
	addresses := make([]common.Address, len(testAddresses))
	for i, addr := range testAddresses {
		addresses[i] = common.HexToAddress(addr)
	}

	bnbBalances, err := balanceChecker.GetETHBalances(nil, addresses)
	if err != nil {
		log.Printf("批量查询BNB余额失败: %v", err)
	} else {
		for i, balance := range bnbBalances {
			fmt.Printf("地址 %s 的BNB余额: %s BNB\n",
				testAddresses[i], weiToBNB(balance))
		}
	}
	fmt.Println()

	// 3. 查询单个地址的USDT余额
	fmt.Println("3. 查询单个地址的USDT余额")
	usdtAddr := common.HexToAddress(USDT_BSC)
	usdtBalance, err := balanceChecker.GetERC20Balance(nil, usdtAddr, targetAddr)
	if err != nil {
		log.Printf("查询USDT余额失败: %v", err)
	} else {
		fmt.Printf("地址 %s 的USDT余额: %s USDT\n",
			targetAddr.Hex(), tokenBalance(usdtBalance, 18))
	}
	fmt.Println()

	// 4. 批量查询多个代币余额
	fmt.Println("4. 批量查询多个代币余额")
	tokens := []common.Address{
		common.HexToAddress(USDT_BSC),
		common.HexToAddress(USDC_BSC),
		common.HexToAddress(BUSD_BSC),
		common.HexToAddress(WBNB_BSC),
	}
	tokenNames := []string{"USDT", "USDC", "BUSD", "WBNB"}

	multiTokenBalances, err := balanceChecker.GetMultipleERC20Balances(nil, tokens, addresses)
	if err != nil {
		log.Printf("批量查询代币余额失败: %v", err)
	} else {
		for i, tokenBalances := range multiTokenBalances {
			fmt.Printf("%s 代币余额:\n", tokenNames[i])
			for j, balance := range tokenBalances {
				fmt.Printf("  地址 %s: %s %s\n",
					testAddresses[j], tokenBalance(balance, 18), tokenNames[i])
			}
		}
	}
	fmt.Println()

	// 5. 查询单个地址的所有余额（BNB + 代币）
	fmt.Println("5. 查询单个地址的所有余额")
	result, err := balanceChecker.GetAddressBalances(nil, targetAddr, tokens)
	if err != nil {
		log.Printf("查询地址余额失败: %v", err)
	} else {
		fmt.Printf("地址 %s:\n", targetAddr.Hex())
		fmt.Printf("  BNB余额: %s BNB\n", weiToBNB(result.EthBalance))
		for i, balance := range result.TokenBalances {
			fmt.Printf("  %s余额: %s %s\n",
				tokenNames[i], tokenBalance(balance, 18), tokenNames[i])
		}
	}
	fmt.Println()

	// 6. 批量查询多个地址的所有余额
	fmt.Println("6. 批量查询多个地址的所有余额")
	multiResult, err := balanceChecker.GetMultipleAddressBalances(nil, addresses, tokens)
	if err != nil {
		log.Printf("批量查询地址余额失败: %v", err)
	} else {
		for i, addr := range testAddresses {
			fmt.Printf("地址 %s:\n", addr)
			fmt.Printf("  BNB余额: %s BNB\n", weiToBNB(multiResult.EthBalances[i]))
			for j, balance := range multiResult.TokenBalances[i] {
				fmt.Printf("  %s余额: %s %s\n",
					tokenNames[j], tokenBalance(balance, 18), tokenNames[j])
			}
			fmt.Println()
		}
	}
}

// weiToBNB 将wei转换为BNB（18位小数）
func weiToBNB(wei *big.Int) string {
	bnb := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return bnb.Text('f', 6)
}

// tokenBalance 将代币余额转换为可读格式
func tokenBalance(balance *big.Int, decimals int) string {
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	tokenAmount := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(divisor))
	return tokenAmount.Text('f', 6)
}
