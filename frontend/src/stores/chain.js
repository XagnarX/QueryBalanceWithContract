import { defineStore } from 'pinia'
import api from '@/services/api'

export const useChainStore = defineStore('chain', {
  state: () => ({
    chains: [],
    currentChain: null,
    loading: false,
    error: null
  }),

  getters: {
    activeChains: (state) => state.chains.filter(chain => chain.is_active),
    
    getCurrentChainInfo: (state) => {
      return state.currentChain ? state.chains.find(chain => chain.chain_id === state.currentChain) : null
    },

    getCurrentRPCEndpoints: (state) => {
      const chainInfo = state.chains.find(chain => chain.chain_id === state.currentChain)
      return chainInfo?.rpc_endpoints?.filter(rpc => rpc.is_active) || []
    },

    getCurrentBalanceContracts: (state) => {
      const chainInfo = state.chains.find(chain => chain.chain_id === state.currentChain)
      return chainInfo?.balance_contracts?.filter(contract => contract.is_active) || []
    },

    getCurrentTokens: (state) => {
      const chainInfo = state.chains.find(chain => chain.chain_id === state.currentChain)
      return chainInfo?.tokens?.filter(token => token.is_active) || []
    }
  },

  actions: {
    // 获取所有链信息
    async fetchChains() {
      this.loading = true
      this.error = null
      
      try {
        const response = await api.get('/chains')
        this.chains = response.data.chains
        
        // 如果没有选择链且有可用链，选择第一个活跃的链
        if (!this.currentChain && this.activeChains.length > 0) {
          this.currentChain = this.activeChains[0].chain_id
          this.saveCurrentChain()
        }
        
        return this.chains
      } catch (error) {
        this.error = error.response?.data?.error || '获取链信息失败'
        console.error('获取链信息失败:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    // 获取指定链的详细信息
    async fetchChainInfo(chainId) {
      this.loading = true
      this.error = null

      try {
        const response = await api.get(`/chains/${chainId}`)
        const chainInfo = response.data
        
        // 更新chains数组中对应的链信息
        const index = this.chains.findIndex(chain => chain.chain_id === chainId)
        if (index !== -1) {
          this.chains[index] = chainInfo
        } else {
          this.chains.push(chainInfo)
        }
        
        return chainInfo
      } catch (error) {
        this.error = error.response?.data?.error || '获取链详细信息失败'
        console.error('获取链详细信息失败:', error)
        throw error
      } finally {
        this.loading = false
      }
    },

    // 获取RPC端点
    async fetchRPCEndpoints(chainId) {
      try {
        const response = await api.get(`/chains/${chainId}/rpc`)
        return response.data.rpc_endpoints
      } catch (error) {
        console.error('获取RPC端点失败:', error)
        throw error
      }
    },

    // 获取余额查询合约
    async fetchBalanceContracts(chainId) {
      try {
        const response = await api.get(`/chains/${chainId}/contracts`)
        return response.data.balance_contracts
      } catch (error) {
        console.error('获取余额查询合约失败:', error)
        throw error
      }
    },

    // 获取代币列表
    async fetchTokens(chainId) {
      try {
        const response = await api.get(`/chains/${chainId}/tokens`)
        return response.data.tokens
      } catch (error) {
        console.error('获取代币列表失败:', error)
        throw error
      }
    },

    // 切换当前链
    setCurrentChain(chainId) {
      this.currentChain = chainId
      this.saveCurrentChain()
    },

    // 保存当前选择的链到本地存储
    saveCurrentChain() {
      if (this.currentChain) {
        localStorage.setItem('currentChain', this.currentChain.toString())
      }
    },

    // 从本地存储加载当前选择的链
    loadCurrentChain() {
      const saved = localStorage.getItem('currentChain')
      if (saved) {
        this.currentChain = parseInt(saved)
      }
    },

    // 添加RPC端点
    async addRPCEndpoint(userId, rpcData) {
      try {
        const response = await api.post(`/users/${userId}/admin/rpc`, rpcData)
        
        // 更新本地数据
        const chainInfo = this.chains.find(chain => chain.chain_id === rpcData.chain_id)
        if (chainInfo) {
          if (!chainInfo.rpc_endpoints) {
            chainInfo.rpc_endpoints = []
          }
          chainInfo.rpc_endpoints.push(response.data.endpoint)
        }
        
        return response.data
      } catch (error) {
        console.error('添加RPC端点失败:', error)
        throw error
      }
    },

    // 添加余额查询合约
    async addBalanceContract(userId, contractData) {
      try {
        const response = await api.post(`/users/${userId}/admin/contracts`, contractData)
        
        // 更新本地数据
        const chainInfo = this.chains.find(chain => chain.chain_id === contractData.chain_id)
        if (chainInfo) {
          if (!chainInfo.balance_contracts) {
            chainInfo.balance_contracts = []
          }
          chainInfo.balance_contracts.push(response.data.contract)
        }
        
        return response.data
      } catch (error) {
        console.error('添加余额查询合约失败:', error)
        throw error
      }
    },

    // 添加代币
    async addToken(userId, tokenData) {
      try {
        const response = await api.post(`/users/${userId}/admin/tokens`, tokenData)
        
        // 更新本地数据
        const chainInfo = this.chains.find(chain => chain.chain_id === tokenData.chain_id)
        if (chainInfo) {
          if (!chainInfo.tokens) {
            chainInfo.tokens = []
          }
          chainInfo.tokens.push(response.data.token)
        }
        
        return response.data
      } catch (error) {
        console.error('添加代币失败:', error)
        throw error
      }
    },

    // 清空状态
    clearState() {
      this.chains = []
      this.currentChain = null
      this.loading = false
      this.error = null
      localStorage.removeItem('currentChain')
    }
  }
})
