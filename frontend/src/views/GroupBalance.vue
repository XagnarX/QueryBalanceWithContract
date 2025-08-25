<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex items-center space-x-4 mb-6">
        <button @click="$router.go(-1)" class="text-gray-500 hover:text-gray-700">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
        </button>
        <h1 class="text-3xl font-bold text-gray-900">
          {{ groupInfo?.name || '分组' }} - 余额详情
        </h1>
      </div>

      <!-- 设置区域 -->
      <div class="card mb-6">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">查询设置</h3>
        </div>
        <div class="card-body">
          <!-- RPC节点选择 - 平铺单选 -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-3">RPC节点 (单选)</label>
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
              <label 
                v-for="rpc in availableRpcs" 
                :key="rpc.id"
                class="flex items-center p-3 border border-gray-200 rounded-lg cursor-pointer hover:bg-gray-50 transition-colors"
                :class="selectedRpcId === rpc.id ? 'border-primary-500 bg-primary-50' : ''"
              >
                <input 
                  type="radio" 
                  :value="rpc.id"
                  v-model="selectedRpcId"
                  @change="handleRpcChange"
                  class="w-4 h-4 text-primary-600 focus:ring-primary-500 border-gray-300"
                >
                <div class="ml-3 flex-1 min-w-0">
                  <div class="text-sm font-medium text-gray-900">{{ rpc.name }}</div>
                  <div class="text-xs text-gray-500 truncate">{{ rpc.url }}</div>
                </div>
              </label>
            </div>
          </div>

          <!-- Token选择 - 平铺多选 -->
          <div class="mb-6">
            <label class="block text-sm font-medium text-gray-700 mb-3">Token选择 (多选)</label>
            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
              <label 
                v-for="token in availableTokens" 
                :key="token.id"
                class="flex items-center p-3 border border-gray-200 rounded-lg cursor-pointer hover:bg-gray-50 transition-colors"
                :class="selectedTokenIds.includes(token.id) ? 'border-primary-500 bg-primary-50' : ''"
              >
                <input 
                  type="checkbox" 
                  :value="token.id"
                  v-model="selectedTokenIds"
                  @change="handleTokenChange"
                  class="w-4 h-4 text-primary-600 focus:ring-primary-500 border-gray-300 rounded"
                >
                <div class="ml-3 flex items-center space-x-2 flex-1 min-w-0">
                  <div class="w-6 h-6 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center flex-shrink-0">
                    <span class="text-white text-xs font-bold">{{ token.symbol.substring(0, 1) }}</span>
                  </div>
                  <div class="min-w-0">
                    <div class="text-sm font-medium text-gray-900 truncate">{{ token.symbol }}</div>
                    <div class="text-xs text-gray-500 truncate">{{ token.name }}</div>
                  </div>
                </div>
              </label>
            </div>
            <p class="text-xs text-gray-500 mt-2">已选择 {{ selectedTokenIds.length }} 个Token</p>
          </div>

          <!-- 查询按钮 -->
          <div class="flex items-center justify-between">
            <button 
              @click="queryBalance" 
              :disabled="loading || !selectedRpcId"
              class="btn-primary"
              :class="{ 'opacity-50 cursor-not-allowed': loading || !selectedRpcId }"
            >
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 714 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ loading ? '查询中...' : '查询余额' }}
            </button>

            <!-- 下载按钮组 -->
            <div v-if="groupBalance && groupBalance.addresses" class="relative">
              <div class="flex space-x-2">
                <!-- JSON下载按钮 -->
                <button 
                  @click="downloadBalanceData" 
                  class="btn-secondary"
                  title="下载 JSON 格式"
                >
                  <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                  JSON
                </button>
                
                <!-- CSV下载按钮 -->
                <button 
                  @click="downloadBalanceCSV" 
                  class="btn-secondary"
                  title="下载 CSV 格式 (Excel 兼容)"
                >
                  <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                  CSV
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
    </div>

    <!-- 余额表格 -->
    <div v-else-if="groupBalance && groupBalance.addresses" class="space-y-6">
      <!-- 总览卡片 -->
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">余额总览</h3>
        </div>
        <div class="card-body">
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="text-center p-4 bg-primary-50 rounded-lg">
              <div class="text-2xl font-bold text-primary-600">{{ groupBalance.total_bnb || '0' }}</div>
              <div class="text-sm text-gray-600">BNB 总额</div>
            </div>
            <div class="text-center p-4 bg-gray-50 rounded-lg">
              <div class="text-lg font-bold text-gray-900">{{ groupBalance.addresses?.length || 0 }}</div>
              <div class="text-sm text-gray-600">地址数量</div>
            </div>
            <div
              v-for="token in (groupBalance.token_totals || []).slice(0, 2)"
              :key="token.symbol"
              class="text-center p-4 bg-gray-50 rounded-lg"
            >
              <div class="text-lg font-bold text-gray-900">{{ formatBalance(token.balance) }}</div>
              <div class="text-sm text-gray-600">{{ token.symbol }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 余额表格 -->
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">钱包地址余额详情</h3>
        </div>
        <div class="card-body p-0">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    钱包地址
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    BNB余额
                  </th>
                  <!-- 动态Token列 -->
                  <th
                    v-for="token in selectedTokensInfo"
                    :key="token.symbol"
                    class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
                  >
                    {{ token.symbol }}
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr
                  v-for="address in groupBalance.addresses"
                  :key="address.address"
                  class="hover:bg-gray-50"
                >
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div>
                      <div class="text-sm font-mono text-gray-900">
                        {{ truncateAddress(address.address) }}
                      </div>
                      <div class="text-sm text-gray-500">{{ address.label || '无标签' }}</div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium text-gray-900">
                    {{ formatBalance(address.bnb_balance) }}
                  </td>
                  <!-- 动态Token余额列 -->
                  <td
                    v-for="token in selectedTokensInfo"
                    :key="token.symbol"
                    class="px-6 py-4 whitespace-nowrap text-right text-sm text-gray-500"
                  >
                    {{ getTokenBalanceForAddress(address, token.contract_address) }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Token总额统计 -->
      <div v-if="groupBalance.token_totals && groupBalance.token_totals.length > 0" class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">Token总额统计</h3>
        </div>
        <div class="card-body p-0">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Token
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    总余额
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    持有地址数
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="token in groupBalance.token_totals" :key="token.symbol">
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="flex-shrink-0 h-8 w-8">
                        <div class="h-8 w-8 rounded-full bg-gradient-to-r from-blue-500 to-purple-600 flex items-center justify-center">
                          <span class="text-white text-xs font-bold">{{ token.symbol.substring(0, 2) }}</span>
                        </div>
                      </div>
                      <div class="ml-4">
                        <div class="text-sm font-medium text-gray-900">{{ token.name }}</div>
                        <div class="text-sm text-gray-500">{{ token.symbol }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium text-gray-900">
                    {{ formatBalance(token.balance) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm text-gray-500">
                    {{ getTokenHolderCount(token.contract_address) }} 个地址
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-else class="text-center py-12">
      <svg
        class="mx-auto h-12 w-12 text-gray-400"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10"
        />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">暂无余额数据</h3>
      <p class="mt-1 text-sm text-gray-500">
        请选择RPC节点并点击查询余额
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'
import { useChainStore } from '@/stores/chain'
import api from '@/services/api'

const route = useRoute()
const authStore = useAuthStore()
const walletStore = useWalletStore()
const chainStore = useChainStore()

const groupId = parseInt(route.params.groupId)
const loading = ref(false)
const groupBalance = ref(null)
const groupInfo = ref(null)

// RPC和Token选择
const availableRpcs = ref([])
const availableTokens = ref([])
const selectedRpcId = ref('')
const selectedTokenIds = ref([])

onMounted(async () => {
  try {
    // 确保chainStore已经加载了chains数据
    if (chainStore.chains.length === 0) {
      await chainStore.fetchChains()
    }
    
    // 加载分组信息
    await loadGroupInfo()
    // 加载可用的RPC和Token
    await loadAvailableRpcs()
    await loadAvailableTokens()
    
    // 默认查询BNB余额（不选择任何Token）
    if (selectedRpcId.value && groupInfo.value) {
      await queryBalance()
    }
  } catch (error) {
    console.error('Initialization failed:', error)
    window.showNotification('error', '初始化失败')
  }
})

const loadGroupInfo = async () => {
  try {
    // 先确保加载了分组数据
    if (walletStore.groups.length === 0) {
      await walletStore.fetchGroups(authStore.userId)
    }
    
    groupInfo.value = walletStore.groups.find(group => group.id === groupId)
    
    if (!groupInfo.value) {
      window.showNotification('error', '分组不存在')
      return
    }
  } catch (error) {
    console.error('Failed to load group info:', error)
    window.showNotification('error', '加载分组信息失败')
  }
}

const loadAvailableRpcs = async () => {
  try {
    const currentChainId = chainStore.currentChain
    if (!currentChainId) {
      window.showNotification('warning', '请先选择区块链网络')
      return
    }

    const response = await api.get(`/chains/${currentChainId}/rpc-endpoints?user_id=${authStore.userId}`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (response.data && response.data.rpc_endpoints) {
      availableRpcs.value = response.data.rpc_endpoints.filter(rpc => rpc.is_active)
      
      // 默认选择第一个RPC
      if (availableRpcs.value.length > 0) {
        selectedRpcId.value = availableRpcs.value[0].id
      }
    }
  } catch (error) {
    console.error('Failed to load RPC endpoints:', error)
    window.showNotification('error', '加载RPC节点失败')
  }
}

const loadAvailableTokens = async () => {
  try {
    const currentChainId = chainStore.currentChain
    if (!currentChainId) return

    const response = await api.get(`/chains/${currentChainId}/tokens?user_id=${authStore.userId}`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (response.data && response.data.tokens) {
      availableTokens.value = response.data.tokens.filter(token => token.is_active)
    }
  } catch (error) {
    console.error('Failed to load tokens:', error)
    window.showNotification('error', '加载Token列表失败')
  }
}

const handleRpcChange = () => {
  // RPC变更时，如果已经有余额数据，重新查询
  if (groupBalance.value) {
    queryBalance()
  }
}

const handleTokenChange = () => {
  // Token选择变更时，如果已经有余额数据，重新查询
  if (groupBalance.value) {
    queryBalance()
  }
}

const queryBalance = async () => {
  if (!selectedRpcId.value) {
    window.showNotification('warning', '请选择RPC节点')
    return
  }

  const currentChainId = chainStore.currentChain
  if (!currentChainId) {
    window.showNotification('warning', '请先选择区块链网络')
    return
  }

  loading.value = true
  try {
    const response = await walletStore.fetchGroupBalance(
      authStore.userId,
      groupId,
      currentChainId,
      selectedTokenIds.value, // 选中的Token ID列表
      null, // 进度回调暂时为null
      selectedRpcId.value // 选中的RPC ID
    )

    groupBalance.value = response
    window.showNotification('success', '余额查询成功')
  } catch (error) {
    console.error('Failed to query balance:', error)
    window.showNotification('error', '余额查询失败')
  } finally {
    loading.value = false
  }
}

// 计算选中的Token信息（用于表格列标题）
const selectedTokensInfo = computed(() => {
  return availableTokens.value.filter(token => selectedTokenIds.value.includes(token.id))
})

// 获取指定地址的Token余额
const getTokenBalanceForAddress = (address, contractAddress) => {
  const tokenBalance = address.token_balances?.find(token => token.contract_address === contractAddress)
  return tokenBalance ? formatBalance(tokenBalance.balance) : '0'
}

// 计算某个Token的持有地址数量
const getTokenHolderCount = (contractAddress) => {
  if (!groupBalance.value || !groupBalance.value.addresses) return 0
  
  return groupBalance.value.addresses.filter(address => 
    address.token_balances?.some(token => 
      token.contract_address === contractAddress && parseFloat(token.balance || 0) > 0
    )
  ).length
}

const truncateAddress = (address) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const formatBalance = (balance) => {
  const num = parseFloat(balance)
  if (num === 0) return '0'
  if (num < 0.000001) return '<0.000001'
  if (num >= 1000000) return (num / 1000000).toFixed(2) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(2) + 'K'
  return num.toFixed(6)
}

// Download balance data functionality
const downloadBalanceData = () => {
  if (!groupBalance.value || !groupBalance.value.addresses) {
    window.showNotification('warning', '没有可下载的余额数据')
    return
  }

  try {
    // Prepare data for export
    const exportData = {
      group_info: {
        group_id: groupId,
        group_name: groupInfo.value?.name || `Group ${groupId}`,
        export_time: new Date().toLocaleString('zh-CN'),
        total_addresses: groupBalance.value.addresses.length,
        selected_rpc: availableRpcs.value.find(rpc => rpc.id === selectedRpcId.value)?.name || 'Unknown RPC',
        selected_tokens: selectedTokensInfo.value.map(token => ({
          symbol: token.symbol,
          name: token.name,
          contract_address: token.contract_address
        }))
      },
      summary: {
        total_bnb: groupBalance.value.total_bnb || '0',
        token_totals: groupBalance.value.token_totals || []
      },
      address_details: groupBalance.value.addresses.map(address => {
        const addressData = {
          address: address.address,
          label: address.label || '',
          bnb_balance: address.bnb_balance || '0'
        }
        
        // Add token balances
        selectedTokensInfo.value.forEach(token => {
          const tokenBalance = address.token_balances?.find(tb => 
            tb.contract_address === token.contract_address
          )
          addressData[`${token.symbol}_balance`] = tokenBalance ? tokenBalance.balance : '0'
        })
        
        return addressData
      })
    }

    // Generate filename with timestamp
    const timestamp = new Date().toISOString().slice(0, 19).replace(/[:-]/g, '')
    const filename = `balance_${groupInfo.value?.name || `group_${groupId}`}_${timestamp}.json`

    // Create and download file
    const blob = new Blob([JSON.stringify(exportData, null, 2)], {
      type: 'application/json;charset=utf-8'
    })
    
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = filename
    document.body.appendChild(link)
    link.click()
    
    // Cleanup
    document.body.removeChild(link)
    URL.revokeObjectURL(url)
    
    window.showNotification('success', `余额详情已下载: ${filename}`)
  } catch (error) {
    console.error('Download failed:', error)
    window.showNotification('error', '下载失败')
  }
}

// Download as CSV format
const downloadBalanceCSV = () => {
  if (!groupBalance.value || !groupBalance.value.addresses) {
    window.showNotification('warning', '没有可下载的余额数据')
    return
  }

  try {
    // Prepare CSV headers
    const headers = ['钱包地址', '标签', 'BNB余额']
    selectedTokensInfo.value.forEach(token => {
      headers.push(`${token.symbol}余额`)
    })

    // Prepare CSV rows
    const rows = [headers.join(',')]
    
    groupBalance.value.addresses.forEach(address => {
      const row = [
        `"${address.address}"`,
        `"${address.label || ''}"`,
        address.bnb_balance || '0'
      ]
      
      // Add token balances
      selectedTokensInfo.value.forEach(token => {
        const tokenBalance = address.token_balances?.find(tb => 
          tb.contract_address === token.contract_address
        )
        row.push(tokenBalance ? tokenBalance.balance : '0')
      })
      
      rows.push(row.join(','))
    })

    // Generate filename
    const timestamp = new Date().toISOString().slice(0, 19).replace(/[:-]/g, '')
    const filename = `balance_${groupInfo.value?.name || `group_${groupId}`}_${timestamp}.csv`

    // Create and download file
    const csvContent = '\uFEFF' + rows.join('\n') // Add BOM for Excel compatibility
    const blob = new Blob([csvContent], {
      type: 'text/csv;charset=utf-8'
    })
    
    const url = URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = filename
    document.body.appendChild(link)
    link.click()
    
    // Cleanup
    document.body.removeChild(link)
    URL.revokeObjectURL(url)
    
    window.showNotification('success', `余额详情已下载: ${filename}`)
  } catch (error) {
    console.error('CSV download failed:', error)
    window.showNotification('error', 'CSV下载失败')
  }
}
</script>
