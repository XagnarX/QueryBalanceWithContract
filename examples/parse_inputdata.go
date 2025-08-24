package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

// 命令常量定义
const (
	V3_SWAP_EXACT_IN            = 0x00
	V3_SWAP_EXACT_OUT           = 0x01
	PERMIT2_TRANSFER_FROM       = 0x02
	PERMIT2_PERMIT_BATCH        = 0x03
	SWEEP                       = 0x04
	TRANSFER                    = 0x05
	PAY_PORTION                 = 0x06
	V2_SWAP_EXACT_IN            = 0x08
	V2_SWAP_EXACT_OUT           = 0x09
	PERMIT2_PERMIT              = 0x0c
	WRAP_ETH                    = 0x0d
	UNWRAP_WETH                 = 0x0e
	PERMIT2_TRANSFER_FROM_BATCH = 0x0f
	BALANCE_CHECK_ERC20         = 0x10
	OWNER_CHECK_721             = 0x11
	OWNER_CHECK_1155            = 0x12
	SWEEP_ERC721                = 0x13
	SWEEP_ERC1155               = 0x14
	SEAPORT_V1_5                = 0x18
	SEAPORT_V1_4                = 0x19
	LOOKS_RARE_V2               = 0x1a
	X2Y2_721                    = 0x1b
	X2Y2_1155                   = 0x1c
	EXECUTE_SUB_PLAN            = 0x20
	APPROVE_ERC20               = 0x21
	STABLE_SWAP_EXACT_IN        = 0x22
	STABLE_SWAP_EXACT_OUT       = 0x23
	PANCAKE_NFT_BNB             = 0x24
	PANCAKE_NFT_WBNB            = 0x25
)

// 命令名称映射
var commandNames = map[int]string{
	V3_SWAP_EXACT_IN:            "V3_SWAP_EXACT_IN",
	V3_SWAP_EXACT_OUT:           "V3_SWAP_EXACT_OUT",
	PERMIT2_TRANSFER_FROM:       "PERMIT2_TRANSFER_FROM",
	PERMIT2_PERMIT_BATCH:        "PERMIT2_PERMIT_BATCH",
	SWEEP:                       "SWEEP",
	TRANSFER:                    "TRANSFER",
	PAY_PORTION:                 "PAY_PORTION",
	V2_SWAP_EXACT_IN:            "V2_SWAP_EXACT_IN",
	V2_SWAP_EXACT_OUT:           "V2_SWAP_EXACT_OUT",
	PERMIT2_PERMIT:              "PERMIT2_PERMIT",
	WRAP_ETH:                    "WRAP_ETH",
	UNWRAP_WETH:                 "UNWRAP_WETH",
	PERMIT2_TRANSFER_FROM_BATCH: "PERMIT2_TRANSFER_FROM_BATCH",
	BALANCE_CHECK_ERC20:         "BALANCE_CHECK_ERC20",
	OWNER_CHECK_721:             "OWNER_CHECK_721",
	OWNER_CHECK_1155:            "OWNER_CHECK_1155",
	SWEEP_ERC721:                "SWEEP_ERC721",
	SWEEP_ERC1155:               "SWEEP_ERC1155",
	SEAPORT_V1_5:                "SEAPORT_V1_5",
	SEAPORT_V1_4:                "SEAPORT_V1_4",
	LOOKS_RARE_V2:               "LOOKS_RARE_V2",
	X2Y2_721:                    "X2Y2_721",
	X2Y2_1155:                   "X2Y2_1155",
	EXECUTE_SUB_PLAN:            "EXECUTE_SUB_PLAN",
	APPROVE_ERC20:               "APPROVE_ERC20",
	STABLE_SWAP_EXACT_IN:        "STABLE_SWAP_EXACT_IN",
	STABLE_SWAP_EXACT_OUT:       "STABLE_SWAP_EXACT_OUT",
	PANCAKE_NFT_BNB:             "PANCAKE_NFT_BNB",
	PANCAKE_NFT_WBNB:            "PANCAKE_NFT_WBNB",
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("使用方法: go run parse_inputdata.go <inputdata文件路径>")
		fmt.Println("或者: go run parse_inputdata.go <inputdata字符串>")
		os.Exit(1)
	}

	input := os.Args[1]
	var inputdata []byte
	var err error

	// 检查是否是文件路径
	if strings.HasSuffix(input, ".txt") || strings.HasSuffix(input, ".hex") || strings.Contains(input, "/") || strings.Contains(input, "\\") {
		// 从文件读取
		inputdata, err = ioutil.ReadFile(input)
		if err != nil {
			log.Fatal("读取文件失败:", err)
		}
		// 移除可能的换行符和空格，然后解析hex
		content := strings.TrimSpace(string(inputdata))
		if strings.HasPrefix(content, "0x") {
			content = content[2:]
		}
		inputdata, err = hex.DecodeString(content)
		if err != nil {
			log.Fatal("解析文件中的hex字符串失败:", err)
		}
	} else {
		// 直接解析字符串
		if strings.HasPrefix(input, "0x") {
			input = input[2:]
		}
		inputdata, err = hex.DecodeString(input)
		if err != nil {
			log.Fatal("解析hex字符串失败:", err)
		}
	}

	fmt.Println("=== Inputdata 解析结果 ===")
	fmt.Printf("原始数据长度: %d 字节\n", len(inputdata))
	fmt.Printf("原始数据: 0x%x\n\n", inputdata)

	parseTransactionData(inputdata)
}

// 解析交易数据的主函数
func parseTransactionData(data []byte) {
	if len(data) < 4 {
		fmt.Println("数据太短，无法解析")
		return
	}

	// 解析函数选择器
	selector := data[:4]
	fmt.Printf("函数选择器: 0x%x\n", selector)

	// 根据选择器判断函数类型
	switch fmt.Sprintf("%x", selector) {
	case "3593564c":
		fmt.Println("函数: execute(bytes calldata commands, bytes[] calldata inputs, uint256 deadline)")
		parseExecuteData(data[4:])
	case "24856bc3":
		fmt.Println("函数: execute(bytes calldata commands, bytes[] calldata inputs)")
		parseExecuteDataWithoutDeadline(data[4:])
	default:
		fmt.Printf("未知函数选择器: 0x%x\n", selector)
	}
}

// 解析带deadline的execute函数数据
func parseExecuteData(data []byte) {
	if len(data) < 96 { // 至少需要 3 个 uint256 参数
		fmt.Println("数据太短，无法解析execute参数")
		return
	}

	// 解析参数偏移量
	commandsOffset := new(big.Int).SetBytes(data[:32])
	inputsOffset := new(big.Int).SetBytes(data[32:64])
	deadline := new(big.Int).SetBytes(data[64:96])

	fmt.Printf("\n=== 参数解析 ===\n")
	fmt.Printf("Commands偏移量: %s (0x%x)\n", commandsOffset.String(), commandsOffset)
	fmt.Printf("Inputs偏移量: %s (0x%x)\n", inputsOffset.String(), inputsOffset)
	fmt.Printf("Deadline: %s\n", deadline.String())

	// 解析commands和inputs
	if commandsOffset.Cmp(big.NewInt(0)) > 0 {
		parseCommandsAndInputs(data, commandsOffset, inputsOffset)
	}
}

// 解析不带deadline的execute函数数据
func parseExecuteDataWithoutDeadline(data []byte) {
	if len(data) < 64 { // 至少需要 2 个 uint256 参数
		fmt.Println("数据太短，无法解析execute参数")
		return
	}

	// 解析参数偏移量
	commandsOffset := new(big.Int).SetBytes(data[:32])
	inputsOffset := new(big.Int).SetBytes(data[32:64])

	fmt.Printf("\n=== 参数解析 ===\n")
	fmt.Printf("Commands偏移量: %s (0x%x)\n", commandsOffset.String(), commandsOffset)
	fmt.Printf("Inputs偏移量: %s (0x%x)\n", inputsOffset.String(), inputsOffset)

	// 解析commands和inputs
	if commandsOffset.Cmp(big.NewInt(0)) > 0 {
		parseCommandsAndInputs(data, commandsOffset, inputsOffset)
	}
}

// 解析commands和inputs
func parseCommandsAndInputs(data []byte, commandsOffset, inputsOffset *big.Int) {
	fmt.Printf("\n=== Commands和Inputs解析 ===\n")

	// 解析commands
	cmdOffset := int(commandsOffset.Int64())
	if cmdOffset >= len(data) {
		fmt.Println("Commands偏移量超出数据范围")
		return
	}

	// 读取commands长度
	if cmdOffset+32 > len(data) {
		fmt.Println("无法读取commands长度")
		return
	}
	commandsLength := new(big.Int).SetBytes(data[cmdOffset : cmdOffset+32])
	fmt.Printf("Commands长度: %s\n", commandsLength.String())

	// 读取commands数据
	cmdDataStart := cmdOffset + 32
	cmdDataEnd := cmdDataStart + int(commandsLength.Int64())
	if cmdDataEnd > len(data) {
		fmt.Println("Commands数据超出范围")
		return
	}
	commands := data[cmdDataStart:cmdDataEnd]
	fmt.Printf("Commands数据: 0x%x\n", commands)

	// 解析每个命令
	fmt.Printf("\n=== 命令详情 ===\n")
	for i, cmd := range commands {
		cmdName := commandNames[int(cmd)]
		if cmdName == "" {
			cmdName = fmt.Sprintf("未知命令(0x%02x)", cmd)
		}
		fmt.Printf("命令 %d: 0x%02x (%s)\n", i+1, cmd, cmdName)
	}

	// 解析inputs
	inputOffset := int(inputsOffset.Int64())
	if inputOffset >= len(data) {
		fmt.Println("Inputs偏移量超出数据范围")
		return
	}

	// 读取inputs数组长度
	if inputOffset+32 > len(data) {
		fmt.Println("无法读取inputs数组长度")
		return
	}
	inputsArrayLength := new(big.Int).SetBytes(data[inputOffset : inputOffset+32])
	fmt.Printf("\nInputs数组长度: %s\n", inputsArrayLength.String())

	// 解析每个input
	parseInputsArray(data, inputOffset+32, int(inputsArrayLength.Int64()), commands)
}

// 解析inputs数组
func parseInputsArray(data []byte, startOffset int, arrayLength int, commands []byte) {
	fmt.Printf("\n=== Inputs详情 ===\n")

	currentOffset := startOffset
	for i := 0; i < arrayLength; i++ {
		if currentOffset+32 > len(data) {
			fmt.Printf("Input %d: 偏移量超出范围\n", i+1)
			break
		}

		// 读取input的偏移量
		inputOffset := new(big.Int).SetBytes(data[currentOffset : currentOffset+32])
		fmt.Printf("\nInput %d 偏移量: %s (0x%x)\n", i+1, inputOffset.String(), inputOffset)

		// 读取input数据 - 注意：这里的偏移量是相对于原始数据的起始位置
		// 但是需要加上函数选择器的4字节偏移
		inputStart := int(inputOffset.Int64()) + 4
		if inputStart >= len(data) {
			fmt.Printf("Input %d: 数据偏移量超出范围\n", i+1)
			currentOffset += 32
			continue
		}

		// 读取input长度
		if inputStart+32 > len(data) {
			fmt.Printf("Input %d: 无法读取长度\n", i+1)
			currentOffset += 32
			continue
		}

		inputLength := new(big.Int).SetBytes(data[inputStart : inputStart+32])
		fmt.Printf("Input %d 长度: %s\n", i+1, inputLength.String())

		// 读取input内容
		inputDataStart := inputStart + 32
		inputDataEnd := inputStart + 32 + int(inputLength.Int64())
		if inputDataEnd > len(data) {
			fmt.Printf("Input %d: 数据超出范围\n", i+1)
			currentOffset += 32
			continue
		}

		inputContent := data[inputDataStart:inputDataEnd]
		fmt.Printf("Input %d 内容: 0x%x\n", i+1, inputContent)

		// 尝试解析常见的input格式
		parseCommonInputFormat(inputContent, i+1)

		// 特殊处理：如果是第一个input且长度很小，可能是commands数据
		if i == 0 && len(inputContent) <= 10 {
			fmt.Printf("Input %d 特殊解析: 这是commands数据\n", i+1)
		}

		// 根据命令类型解析具体的交易参数
		if i < len(commands) {
			parseCommandSpecificData(int(commands[i]), inputContent, i+1)
		}

		currentOffset += 32
	}
}

// 解析常见的input格式
func parseCommonInputFormat(inputData []byte, inputIndex int) {
	if len(inputData) < 32 {
		return
	}

	fmt.Printf("Input %d 解析:\n", inputIndex)

	// 尝试解析为地址
	if len(inputData) >= 32 {
		// 前32字节可能是地址
		addr := inputData[:32]
		// 检查是否是有效的地址格式（20字节，前面补0）
		if len(addr) == 32 && isAddressFormat(addr) {
			address := addr[12:] // 取后20字节
			fmt.Printf("  地址: 0x%x\n", address)
		}
	}

	// 尝试解析为数量
	if len(inputData) >= 32 {
		amount := new(big.Int).SetBytes(inputData[:32])
		if amount.Cmp(big.NewInt(0)) > 0 {
			fmt.Printf("  数量: %s (0x%x)\n", amount.String(), amount)
		}
	}

	// 如果数据很长，可能是路径或其他复杂数据
	if len(inputData) > 64 {
		fmt.Printf("  复杂数据长度: %d 字节\n", len(inputData))
	}
}

// 检查是否是地址格式（32字节，前12字节为0）
func isAddressFormat(data []byte) bool {
	if len(data) != 32 {
		return false
	}
	// 检查前12字节是否为0
	for i := 0; i < 12; i++ {
		if data[i] != 0 {
			return false
		}
	}
	return true
}

// 根据命令类型解析具体的交易参数
func parseCommandSpecificData(commandType int, inputData []byte, inputIndex int) {
	fmt.Printf("Input %d 命令类型解析 (%s):\n", inputIndex, commandNames[commandType])

	switch commandType {
	case V3_SWAP_EXACT_IN:
		parseV3SwapExactIn(inputData, inputIndex)
	case V2_SWAP_EXACT_IN:
		parseV2SwapExactIn(inputData, inputIndex)
	case PERMIT2_PERMIT:
		parsePermit2Permit(inputData, inputIndex)
	default:
		fmt.Printf("  暂不支持该命令类型的详细解析\n")
	}
}

// 解析V3_SWAP_EXACT_IN命令
func parseV3SwapExactIn(inputData []byte, inputIndex int) {
	if len(inputData) < 160 { // 至少需要5个32字节参数
		fmt.Printf("  V3_SWAP_EXACT_IN数据长度不足\n")
		return
	}

	// 解析参数: recipient, amountIn, amountOutMin, path, payerIsUser
	recipient := inputData[12:32] // 跳过前12字节的0
	amountIn := new(big.Int).SetBytes(inputData[32:64])
	amountOutMin := new(big.Int).SetBytes(inputData[64:96])
	payerIsUser := new(big.Int).SetBytes(inputData[128:160])

	fmt.Printf("  接收地址: 0x%x\n", recipient)
	fmt.Printf("  输入数量: %s (0x%x)\n", amountIn.String(), amountIn)
	fmt.Printf("  最小输出: %s (0x%x)\n", amountOutMin.String(), amountOutMin)
	fmt.Printf("  用户支付: %t\n", payerIsUser.Cmp(big.NewInt(0)) > 0)

	// 尝试解析路径（如果数据足够长）
	if len(inputData) > 160 {
		fmt.Printf("  路径数据: 0x%x\n", inputData[160:])
	}
}

// 解析V2_SWAP_EXACT_IN命令
func parseV2SwapExactIn(inputData []byte, inputIndex int) {
	if len(inputData) < 160 { // 至少需要5个32字节参数
		fmt.Printf("  V2_SWAP_EXACT_IN数据长度不足\n")
		return
	}

	// 解析参数: recipient, amountIn, amountOutMin, path, payerIsUser
	recipient := inputData[12:32] // 跳过前12字节的0
	amountIn := new(big.Int).SetBytes(inputData[32:64])
	amountOutMin := new(big.Int).SetBytes(inputData[64:96])
	payerIsUser := new(big.Int).SetBytes(inputData[128:160])

	fmt.Printf("  接收地址: 0x%x\n", recipient)
	fmt.Printf("  输入数量: %s (0x%x)\n", amountIn.String(), amountIn)
	fmt.Printf("  最小输出: %s (0x%x)\n", amountOutMin.String(), amountOutMin)
	fmt.Printf("  用户支付: %t\n", payerIsUser.Cmp(big.NewInt(0)) > 0)

	// 尝试解析路径（如果数据足够长）
	if len(inputData) > 160 {
		fmt.Printf("  路径数据: 0x%x\n", inputData[160:])
	}
}

// 解析PERMIT2_PERMIT命令
func parsePermit2Permit(inputData []byte, inputIndex int) {
	fmt.Printf("  PERMIT2_PERMIT数据: 0x%x\n", inputData)
	fmt.Printf("  数据长度: %d 字节\n", len(inputData))
}
