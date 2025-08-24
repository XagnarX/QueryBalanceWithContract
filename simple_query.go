package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"QueryBalanceWithContract/bindings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	BSC_MAINNET_RPC  = "https://bsc-dataseed1.binance.org/"
	CONTRACT_ADDRESS = "0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b"

	// 常用代币地址
	USDT_BSC = "0x55d398326f99059fF775485246999027B3197955"
	USDC_BSC = "0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"
	BUSD_BSC = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"
	WBNB_BSC = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法:")
		fmt.Println("  go run simple_query.go <地址> [代币地址...]")
		fmt.Println()
		fmt.Println("示例:")
		fmt.Printf("  go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045\n")
		fmt.Printf("  go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 %s\n", USDT_BSC)
		fmt.Printf("  go run simple_query.go 0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045 %s %s\n", USDT_BSC, USDC_BSC)
		fmt.Println()
		fmt.Println("常用代币地址:")
		fmt.Printf("  USDT: %s\n", USDT_BSC)
		fmt.Printf("  USDC: %s\n", USDC_BSC)
		fmt.Printf("  BUSD: %s\n", BUSD_BSC)
		fmt.Printf("  WBNB: %s\n", WBNB_BSC)
		return
	}

	// 解析地址参数
	targetAddr := common.HexToAddress(os.Args[1])

	// 解析代币地址
	var tokens []common.Address
	for i := 2; i < len(os.Args); i++ {
		tokens = append(tokens, common.HexToAddress(os.Args[i]))
	}

	// 连接BSC主网
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

	fmt.Printf("查询地址: %s\n", targetAddr.Hex())
	fmt.Println("======================")

	if len(tokens) == 0 {
		// 只查询BNB余额
		bnbBalance, err := balanceChecker.GetETHBalance(nil, targetAddr)
		if err != nil {
			log.Fatalf("查询BNB余额失败: %v", err)
		}
		fmt.Printf("BNB余额: %s BNB\n", weiToBNB(bnbBalance))
	} else {
		// 查询BNB和代币余额
		result, err := balanceChecker.GetAddressBalances(nil, targetAddr, tokens)
		if err != nil {
			log.Fatalf("查询余额失败: %v", err)
		}

		fmt.Printf("BNB余额: %s BNB\n", weiToBNB(result.EthBalance))

		for i, balance := range result.TokenBalances {
			tokenName := getTokenName(tokens[i].Hex())
			fmt.Printf("%s余额: %s %s\n", tokenName, tokenBalance(balance, 18), tokenName)
		}
	}
}

// weiToBNB 将wei转换为BNB
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

// getTokenName 根据地址返回代币名称
func getTokenName(address string) string {
	switch address {
	case USDT_BSC:
		return "USDT"
	case USDC_BSC:
		return "USDC"
	case BUSD_BSC:
		return "BUSD"
	case WBNB_BSC:
		return "WBNB"
	default:
		return "TOKEN"
	}
}
