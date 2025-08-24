import { ethers } from 'ethers'

// BalanceChecker合约ABI
const BALANCE_CHECKER_ABI = [
  {
    "inputs": [{"internalType": "address", "name": "addr", "type": "address"}],
    "name": "getETHBalance",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [{"internalType": "address[]", "name": "addresses", "type": "address[]"}],
    "name": "getETHBalances",
    "outputs": [{"internalType": "uint256[]", "name": "", "type": "uint256[]"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "tokenContract", "type": "address"},
      {"internalType": "address", "name": "addr", "type": "address"}
    ],
    "name": "getERC20Balance",
    "outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "tokenContract", "type": "address"},
      {"internalType": "address[]", "name": "addresses", "type": "address[]"}
    ],
    "name": "getERC20Balances",
    "outputs": [{"internalType": "uint256[]", "name": "", "type": "uint256[]"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address[]", "name": "tokenContracts", "type": "address[]"},
      {"internalType": "address", "name": "addr", "type": "address"}
    ],
    "name": "getMultipleERC20Balances",
    "outputs": [{"internalType": "uint256[]", "name": "", "type": "uint256[]"}],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address", "name": "addr", "type": "address"},
      {"internalType": "address[]", "name": "tokenContracts", "type": "address[]"}
    ],
    "name": "getAddressBalances",
    "outputs": [
      {"internalType": "uint256", "name": "ethBalance", "type": "uint256"},
      {"internalType": "uint256[]", "name": "tokenBalances", "type": "uint256[]"}
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {"internalType": "address[]", "name": "addresses", "type": "address[]"},
      {"internalType": "address[]", "name": "tokenContracts", "type": "address[]"}
    ],
    "name": "getMultipleAddressBalances",
    "outputs": [
      {"internalType": "uint256[]", "name": "ethBalances", "type": "uint256[]"},
      {"internalType": "uint256[][]", "name": "tokenBalances", "type": "uint256[][]"}
    ],
    "stateMutability": "view",
    "type": "function"
  }
]

class BlockchainService {
  constructor() {
    this.providers = new Map()
    this.contracts = new Map()
  }

  // 获取或创建Provider
  getProvider(rpcUrl, chainId = null) {
    // 使用 rpcUrl 和 chainId 作为缓存键，确保不同的 chainId 创建不同的 provider
    const cacheKey = chainId ? `${rpcUrl}-${chainId}` : rpcUrl
    
    console.log(`getProvider called with rpcUrl: ${rpcUrl}, chainId: ${chainId}, cacheKey: ${cacheKey}`)
    
    if (!this.providers.has(cacheKey)) {
      // 如果提供了chainId，直接创建带chainId的provider，避免自动调用eth_chainId
      const network = chainId ? { chainId, name: 'unknown' } : undefined
      console.log(`Creating new provider with network:`, network)
      const provider = new ethers.providers.JsonRpcProvider(rpcUrl, network)
      this.providers.set(cacheKey, provider)
    } else {
      console.log(`Using cached provider for key: ${cacheKey}`)
    }
    return this.providers.get(cacheKey)
  }

  // 获取或创建合约实例
  getContract(contractAddress, rpcUrl, chainId = null) {
    // 包含 chainId 在缓存键中，确保不同的 chainId 创建不同的合约实例
    const key = chainId ? `${contractAddress}-${rpcUrl}-${chainId}` : `${contractAddress}-${rpcUrl}`
    if (!this.contracts.has(key)) {
      const provider = this.getProvider(rpcUrl, chainId)
      const contract = new ethers.Contract(contractAddress, BALANCE_CHECKER_ABI, provider)
      this.contracts.set(key, contract)
    }
    return this.contracts.get(key)
  }

  // 测试RPC连接
  async testRPCConnection(rpcUrl, chainId = null) {
    try {
      const provider = this.getProvider(rpcUrl, chainId)
      // 如果提供了chainId，就不需要调用getNetwork来验证了
      if (chainId) {
        // 简单测试连接性，不验证chainId
        await provider.getBlockNumber()
      } else {
        // 没有chainId时才调用getNetwork
        await provider.getNetwork()
      }
      return true
    } catch (error) {
      console.error('RPC连接测试失败:', error)
      return false
    }
  }

  // 格式化余额
  formatBalance(balance, decimals = 18) {
    try {
      const balanceBN = ethers.BigNumber.from(balance)
      const formatted = ethers.utils.formatUnits(balanceBN, decimals)
      return parseFloat(formatted).toFixed(6)
    } catch (error) {
      console.error('格式化余额失败:', error)
      return '0'
    }
  }

  // 查询单个地址的原生代币余额
  async getNativeBalance(address, rpcUrl, chainId = null) {
    try {
      const provider = this.getProvider(rpcUrl, chainId)
      const balance = await provider.getBalance(address)
      return this.formatBalance(balance)
    } catch (error) {
      console.error('查询原生代币余额失败:', error)
      throw error
    }
  }

  // 查询单个地址的ERC20代币余额
  async getTokenBalance(address, tokenContract, rpcUrl, decimals = 18, chainId = null) {
    try {
      const provider = this.getProvider(rpcUrl, chainId)
      const tokenABI = [
        "function balanceOf(address owner) view returns (uint256)"
      ]
      const contract = new ethers.Contract(tokenContract, tokenABI, provider)
      const balance = await contract.balanceOf(address)
      return this.formatBalance(balance, decimals)
    } catch (error) {
      console.error('查询代币余额失败:', error)
      throw error
    }
  }

  // 使用BalanceChecker合约查询单个地址余额
  async getAddressBalances(address, tokenContracts, contractAddress, rpcUrl, nativeDecimals = 18, chainId = null) {
    try {
      const contract = this.getContract(contractAddress, rpcUrl, chainId)
      const result = await contract.getAddressBalances(address, tokenContracts)
      
      return {
        nativeBalance: this.formatBalance(result.ethBalance, nativeDecimals),
        tokenBalances: result.tokenBalances.map((balance, index) => ({
          contractAddress: tokenContracts[index],
          balance: this.formatBalance(balance)
        }))
      }
    } catch (error) {
      console.error('查询地址余额失败:', error)
      throw error
    }
  }

  // 使用BalanceChecker合约批量查询多个地址余额
  async getMultipleAddressBalances(addresses, tokenContracts, contractAddress, rpcUrl, nativeDecimals = 18, chainId = null) {
    try {
      // 验证地址格式（只检查长度和0x前缀）
      const validAddresses = addresses.map(addr => {
        if (!addr || typeof addr !== 'string' || addr.length !== 42 || !addr.startsWith('0x')) {
          throw new Error(`Invalid address format: ${addr}`)
        }
        return addr
      })
      
      // 验证代币合约地址格式（只检查长度和0x前缀）
      const validTokenContracts = tokenContracts.map(addr => {
        if (!addr || typeof addr !== 'string' || addr.length !== 42 || !addr.startsWith('0x')) {
          throw new Error(`Invalid token contract address: ${addr}`)
        }
        return addr
      })
      
      const provider = this.getProvider(rpcUrl, chainId)
      
      // 使用简化的ABI，只包含我们需要的方法，避免方法选择冲突
      const specificABI = [
        {
          "inputs": [
            {"internalType": "address[]", "name": "targets", "type": "address[]"},
            {"internalType": "address[]", "name": "tokens", "type": "address[]"}
          ],
          "name": "getMultipleAddressBalances",
          "outputs": [
            {"internalType": "uint256[]", "name": "ethBalances", "type": "uint256[]"},
            {"internalType": "uint256[][]", "name": "tokenBalances", "type": "uint256[][]"}
          ],
          "stateMutability": "view",
          "type": "function"
        }
      ]
      
      const contract = new ethers.Contract(contractAddress, specificABI, provider)
      
      console.log('Calling getMultipleAddressBalances with:')
      console.log('Addresses:', validAddresses.length, 'items:', validAddresses)
      console.log('Token contracts:', validTokenContracts.length, 'items:', validTokenContracts)
      console.log('Contract address:', contractAddress)
      console.log('RPC URL:', rpcUrl)
      console.log('Chain ID:', chainId)
      
      // 调用正确的合约方法：getMultipleAddressBalances(address[] targets, address[] tokens)
      // 参数顺序：第一个是地址数组，第二个是代币合约数组
      
      // 临时测试：先尝试一个简单的调用来确认合约是否正常
      try {
        console.log('Testing contract call with first address only...')
        const testResult = await contract.getMultipleAddressBalances([validAddresses[0]], validTokenContracts)
        console.log('Test call successful:', testResult)
      } catch (testError) {
        console.error('Test call failed:', testError)
      }
      
      const result = await contract.getMultipleAddressBalances(validAddresses, validTokenContracts)
      
      console.log('Contract call successful, result:', result)
      
      return {
        addresses: validAddresses.map((address, index) => ({
          address,
          nativeBalance: this.formatBalance(result.ethBalances[index], nativeDecimals),
          tokenBalances: result.tokenBalances[index].map((balance, tokenIndex) => ({
            contractAddress: validTokenContracts[tokenIndex],
            balance: this.formatBalance(balance)
          }))
        }))
      }
    } catch (error) {
      console.error('批量查询地址余额失败:', error)
      console.error('Error message:', error.message)
      console.error('Error code:', error.code)
      console.error('Error data:', error.data)
      console.error('Error reason:', error.reason)
      console.error('Error transaction:', error.transaction)
      
      // 如果是revert错误，尝试解析revert reason
      if (error.data) {
        console.error('Raw error data:', error.data)
      }
      
      throw error
    }
  }

  // 分批查询多个地址余额（单批最多200个地址，支持进度回调）
  async getBatchedAddressBalances(addresses, tokenContracts, contractAddress, rpcUrl, batchSize = 200, nativeDecimals = 18, onProgress = null, chainId = null) {
    try {
      if (!addresses || addresses.length === 0) {
        return { addresses: [] }
      }

      const batches = []
      for (let i = 0; i < addresses.length; i += batchSize) {
        batches.push(addresses.slice(i, i + batchSize))
      }

      const allResults = []
      let processedAddresses = 0

      for (let i = 0; i < batches.length; i++) {
        const batch = batches[i]
        
        // 报告进度
        if (onProgress) {
          onProgress({
            stage: 'batch_processing',
            currentBatch: i + 1,
            totalBatches: batches.length,
            processedAddresses,
            totalAddresses: addresses.length,
            batchSize: batch.length
          })
        }

        const result = await this.getMultipleAddressBalances(batch, tokenContracts, contractAddress, rpcUrl, nativeDecimals, chainId)
        allResults.push(...result.addresses)
        processedAddresses += batch.length

        // 报告批次完成进度
        if (onProgress) {
          onProgress({
            stage: 'batch_completed',
            currentBatch: i + 1,
            totalBatches: batches.length,
            processedAddresses,
            totalAddresses: addresses.length
          })
        }
      }

      return { addresses: allResults }
    } catch (error) {
      console.error('分批查询地址余额失败:', error)
      throw error
    }
  }

  // 查询分组余额（汇总计算，支持大批量地址）
  async getGroupBalances(addresses, tokens, contractAddress, rpcUrl, nativeSymbol = 'BNB', batchSize = 200, onProgress = null, chainId = null) {
    try {
      if (!addresses || addresses.length === 0) {
        return {
          totalNative: '0',
          tokenTotals: [],
          addressBalances: []
        }
      }

      const tokenContracts = tokens.map(token => token.contract_address)
      
      // 使用分批查询，传递进度回调和chainId
      const result = await this.getBatchedAddressBalances(addresses, tokenContracts, contractAddress, rpcUrl, batchSize, 18, onProgress, chainId)
      
      // 计算总余额
      let totalNative = 0
      const tokenTotals = new Map()

      // 初始化代币总额
      tokens.forEach(token => {
        tokenTotals.set(token.contract_address, {
          symbol: token.symbol,
          name: token.name,
          contract_address: token.contract_address,
          decimals: token.decimals,
          total: 0
        })
      })

      // 汇总计算
      const addressBalances = result.addresses.map(addrResult => {
        const nativeBalance = parseFloat(addrResult.nativeBalance)
        totalNative += nativeBalance

        const tokenBalances = addrResult.tokenBalances.map((tokenBalance, index) => {
          const token = tokens[index]
          const balance = parseFloat(tokenBalance.balance)
          
          if (tokenTotals.has(token.contract_address)) {
            const tokenTotal = tokenTotals.get(token.contract_address)
            tokenTotal.total += balance
          }

          return {
            symbol: token.symbol,
            name: token.name,
            contract_address: token.contract_address,
            balance: tokenBalance.balance,
            decimals: token.decimals
          }
        })

        return {
          address: addrResult.address,
          nativeBalance: addrResult.nativeBalance,
          tokenBalances
        }
      })

      return {
        totalNative: totalNative.toFixed(6),
        tokenTotals: Array.from(tokenTotals.values()).map(token => ({
          ...token,
          balance: token.total.toFixed(6)
        })),
        addressBalances
      }
    } catch (error) {
      console.error('查询分组余额失败:', error)
      throw error
    }
  }

  // 清理缓存
  clearCache() {
    this.providers.clear()
    this.contracts.clear()
  }
}

export default new BlockchainService()
