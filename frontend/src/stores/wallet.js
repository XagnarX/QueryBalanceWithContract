import { defineStore } from 'pinia'
import api from '@/services/api'
import blockchainService from '@/services/blockchain'
import { useChainStore } from '@/stores/chain'

export const useWalletStore = defineStore('wallet', {
  state: () => ({
    groups: [],
    addresses: [],
    tokens: [],
    balances: {},
    loading: false,
    error: null
  }),

  getters: {
    getGroupById: (state) => (id) => {
      return state.groups.find(group => group.id === id)
    },
    
    getAddressesByGroup: (state) => (groupId) => {
      return state.addresses.filter(addr => addr.group_id === groupId)
    },
    
    getUngroupedAddresses: (state) => {
      return state.addresses.filter(addr => !addr.group_id)
    }
  },

  actions: {
    // 获取用户分组
    async fetchGroups(userId) {
      this.loading = true
      try {
        const response = await api.get(`/users/${userId}/groups`)
        this.groups = response.data
      } catch (error) {
        this.error = error.response?.data?.error || '获取分组失败'
        throw error
      } finally {
        this.loading = false
      }
    },

    // 创建分组
    async createGroup(userId, groupData) {
      try {
        const response = await api.post(`/users/${userId}/groups`, groupData)
        this.groups.push(response.data)
        return response.data
      } catch (error) {
        this.error = error.response?.data?.error || '创建分组失败'
        throw error
      }
    },

    // 删除分组
    async deleteGroup(userId, groupId) {
      try {
        await api.delete(`/users/${userId}/groups/${groupId}`)
        this.groups = this.groups.filter(group => group.id !== groupId)
        // 将该分组下的地址设为未分组
        this.addresses.forEach(addr => {
          if (addr.group_id === groupId) {
            addr.group_id = null
          }
        })
      } catch (error) {
        this.error = error.response?.data?.error || '删除分组失败'
        throw error
      }
    },

    // 获取用户地址
    async fetchAddresses(userId) {
      this.loading = true
      try {
        const response = await api.get(`/users/${userId}/addresses`)
        this.addresses = response.data
      } catch (error) {
        this.error = error.response?.data?.error || '获取地址失败'
        throw error
      } finally {
        this.loading = false
      }
    },

    // 添加地址
    async addAddress(userId, addressData) {
      try {
        const response = await api.post(`/users/${userId}/addresses`, addressData)
        this.addresses.push(response.data)
        return response.data
      } catch (error) {
        this.error = error.response?.data?.error || '添加地址失败'
        throw error
      }
    },

    // 删除地址
    async deleteAddress(userId, addressId) {
      try {
        await api.delete(`/users/${userId}/addresses/${addressId}`)
        this.addresses = this.addresses.filter(addr => addr.id !== addressId)
      } catch (error) {
        this.error = error.response?.data?.error || '删除地址失败'
        throw error
      }
    },

    // 查询单个地址余额
    async fetchAddressBalance(address, chainId) {
      try {
        const chainStore = useChainStore()
        const chainInfo = chainStore.chains.find(chain => chain.chain_id === chainId)
        if (!chainInfo) throw new Error('不支持的区块链网络')

        const rpcEndpoints = chainInfo.rpc_endpoints?.filter(rpc => rpc.is_active) || []
        const balanceContracts = chainInfo.balance_contracts?.filter(contract => contract.is_active) || []
        const tokens = chainInfo.tokens?.filter(token => token.is_active) || []

        if (rpcEndpoints.length === 0) throw new Error('没有可用的RPC端点')
        if (balanceContracts.length === 0) throw new Error('没有可用的余额查询合约')

        const rpcUrl = rpcEndpoints[0].url
        const contractAddress = balanceContracts[0].contract_address
        const tokenContracts = tokens.map(token => token.contract_address)

        const result = await blockchainService.getAddressBalances(
          address, 
          tokenContracts, 
          contractAddress, 
          rpcUrl,
          18, // 原生代币精度
          chainId // 传递chainId，避免重复请求eth_chainId
        )

        // 格式化结果
        const tokenBalances = result.tokenBalances.map((tokenBalance, index) => ({
          ...tokens[index],
          balance: tokenBalance.balance
        }))

        return {
          address,
          nativeBalance: result.nativeBalance,
          nativeSymbol: chainInfo.symbol,
          tokenBalances
        }
      } catch (error) {
        this.error = error.message || '查询地址余额失败'
        throw error
      }
    },

    // 查询分组余额（优化版本，支持大批量地址和进度回调）
    async fetchGroupBalance(userId, groupId, chainId, selectedTokenIds = null, onProgress = null, selectedRpcId = null) {
      try {
        const group = this.getGroupById(groupId)
        if (!group) throw new Error('分组不存在')

        const groupAddresses = this.getAddressesByGroup(groupId)
        if (groupAddresses.length === 0) {
          return {
            group_id: groupId,
            group_name: group.name,
            total_bnb: '0',
            tokenTotals: [],
            addressBalances: []
          }
        }

        const chainStore = useChainStore()
        const chainInfo = chainStore.chains.find(chain => chain.chain_id === chainId)
        if (!chainInfo) throw new Error('不支持的区块链网络')

        // Get user RPCs and tokens from API instead of chainInfo
        let rpcEndpoints = []
        let tokens = []
        
        // Get auth info from the calling parameters (userId is already available)
        const authStore = await import('@/stores/auth').then(m => m.useAuthStore())
        const authToken = authStore.token
        const authUserId = userId

        // Load user RPCs
        try {
          const response = await fetch(`/api/chains/${chainId}/rpc-endpoints?user_id=${authUserId}`, {
            headers: {
              'Authorization': `Bearer ${authToken}`
            }
          })
          if (response.ok) {
            const data = await response.json()
            rpcEndpoints = data.rpc_endpoints?.filter(rpc => rpc.is_active) || []
          }
        } catch (error) {
          console.warn('Failed to load user RPCs:', error)
        }

        // Load user tokens if selectedTokenIds provided
        if (selectedTokenIds !== null && selectedTokenIds.length > 0) {
          try {
            const response = await fetch(`/api/chains/${chainId}/tokens?user_id=${authUserId}`, {
              headers: {
                'Authorization': `Bearer ${authToken}`
              }
            })
            if (response.ok) {
              const data = await response.json()
              tokens = data.tokens?.filter(token => token.is_active && selectedTokenIds.includes(token.id)) || []
            }
          } catch (error) {
            console.warn('Failed to load user tokens:', error)
          }
        }

        // Still get balance contracts from chainInfo (these remain system-wide)
        const balanceContracts = chainInfo.balance_contracts?.filter(contract => contract.is_active) || []
        console.log('=== Token Selection Debug ===')
        console.log('selectedTokenIds:', selectedTokenIds)
        console.log('loaded tokens count:', tokens.length)
        console.log('Final tokens:', tokens.map(t => ({ id: t.id, symbol: t.symbol, contract: t.contract_address })))
        
        // 验证所有Token合约地址格式
        tokens.forEach(token => {
          if (!token.contract_address || typeof token.contract_address !== 'string' || 
              token.contract_address.length !== 42 || !token.contract_address.startsWith('0x')) {
            console.error('Invalid token contract address:', token.symbol, token.contract_address)
            console.error('Token object full data:', JSON.stringify(token, null, 2))
            throw new Error(`Invalid token contract address for ${token.symbol}: ${token.contract_address}`)
          } else {
            console.log('Valid token:', token.symbol, token.contract_address)
          }
        })
        console.log('=== End Token Selection Debug ===')

        if (rpcEndpoints.length === 0) throw new Error('没有可用的RPC端点')
        if (balanceContracts.length === 0) throw new Error('没有可用的余额查询合约')

        // 选择RPC端点
        console.log('=== RPC Selection Debug ===')
        console.log('selectedRpcId:', selectedRpcId)
        console.log('Available RPCs:', rpcEndpoints.map(rpc => ({ id: rpc.id, name: rpc.name, url: rpc.url })))
        
        let selectedRpcUrl
        if (selectedRpcId) {
          // 使用指定的RPC
          const selectedRpc = rpcEndpoints.find(rpc => rpc.id === selectedRpcId)
          console.log('Found selected RPC:', selectedRpc)
          if (selectedRpc) {
            selectedRpcUrl = selectedRpc.url
            console.log('Using selected RPC:', selectedRpc.name, selectedRpc.url)
          } else {
            console.warn('Selected RPC not found, using first available')
            console.warn('Searching for ID:', selectedRpcId, 'in', rpcEndpoints.map(rpc => rpc.id))
            selectedRpcUrl = rpcEndpoints[0].url
            console.log('Fallback to first RPC:', rpcEndpoints[0].name, rpcEndpoints[0].url)
          }
        } else {
          // 使用第一个RPC
          console.log('No RPC selected, using first available')
          selectedRpcUrl = rpcEndpoints[0].url
          console.log('Using first RPC:', rpcEndpoints[0].name, rpcEndpoints[0].url)
        }
        console.log('Final selected RPC URL:', selectedRpcUrl)
        console.log('=== End RPC Selection Debug ===')

        const rpcUrl = selectedRpcUrl
        const contractAddress = balanceContracts[0].contract_address
        const addresses = groupAddresses.map(addr => addr.address)

        // 报告进度：开始查询
        if (onProgress) {
          onProgress({
            stage: 'querying',
            groupName: group.name,
            addressCount: addresses.length,
            current: 0,
            total: addresses.length
          })
        }

        const result = await blockchainService.getGroupBalances(
          addresses,
          tokens,
          contractAddress,
          rpcUrl,
          chainInfo.symbol,
          200, // 批次大小
          onProgress, // 传递进度回调
          chainId // 传递chainId，避免重复请求eth_chainId
        )

        // 报告进度：完成查询
        if (onProgress) {
          onProgress({
            stage: 'completed',
            groupName: group.name,
            addressCount: addresses.length,
            current: addresses.length,
            total: addresses.length
          })
        }

        const balanceData = {
          group_id: groupId,
          group_name: group.name,
          total_bnb: result.totalNative,
          addresses: result.addressBalances.map(addrBalance => ({
            address: addrBalance.address,
            label: groupAddresses.find(addr => addr.address === addrBalance.address)?.label || '',
            bnb_balance: addrBalance.nativeBalance,
            token_balances: addrBalance.tokenBalances
          })),
          token_totals: result.tokenTotals
        }

        this.balances[`group_${groupId}`] = balanceData
        return balanceData
      } catch (error) {
        this.error = error.message || '查询分组余额失败'
        throw error
      }
    },

    // 批量查询所有分组余额（优化版本）
    async fetchAllGroupsBalance(userId, chainId, selectedTokenIds = null, onProgress = null) {
      try {
        // 确保已加载分组和地址数据
        if (this.groups.length === 0) {
          await this.fetchGroups(userId)
        }
        if (this.addresses.length === 0) {
          await this.fetchAddresses(userId)
        }

        const groupsWithAddresses = this.groups.filter(group => {
          const groupAddresses = this.getAddressesByGroup(group.id)
          return groupAddresses.length > 0
        })

        if (groupsWithAddresses.length === 0) {
          return []
        }

        // 报告进度：开始批量查询
        if (onProgress) {
          onProgress({
            stage: 'starting',
            totalGroups: groupsWithAddresses.length,
            currentGroup: 0
          })
        }

        const results = []
        
        // 按组依次查询
        for (let i = 0; i < groupsWithAddresses.length; i++) {
          const group = groupsWithAddresses[i]
          
          try {
            // 报告进度：当前查询的分组
            if (onProgress) {
              onProgress({
                stage: 'querying_group',
                totalGroups: groupsWithAddresses.length,
                currentGroup: i + 1,
                groupName: group.name
              })
            }

            const balanceData = await this.fetchGroupBalance(
              userId, 
              group.id, 
              chainId,
              selectedTokenIds,
              onProgress
            )
            
            if (balanceData && balanceData.addresses && balanceData.addresses.length > 0) {
              results.push(balanceData)
            }
          } catch (error) {
            console.error(`查询分组 ${group.name} 余额失败:`, error)
            // 继续查询其他分组，不中断整个流程
          }
        }

        // 报告进度：全部完成
        if (onProgress) {
          onProgress({
            stage: 'all_completed',
            totalGroups: groupsWithAddresses.length,
            currentGroup: groupsWithAddresses.length,
            successCount: results.length
          })
        }

        return results
      } catch (error) {
        this.error = error.message || '批量查询分组余额失败'
        throw error
      }
    },

    // 清除错误
    clearError() {
      this.error = null
    },

    // 重置状态
    reset() {
      this.groups = []
      this.addresses = []
      this.tokens = []
      this.balances = {}
      this.error = null
    }
  }
})
