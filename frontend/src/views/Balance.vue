<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <h1 class="text-3xl font-bold text-gray-900">余额查询</h1>
      <p class="mt-2 text-gray-600">查看您的钱包地址余额汇总</p>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
    </div>

    <!-- 余额展示 -->
    <div v-else-if="balanceData.length > 0" class="space-y-6">
      <div v-for="group in balanceData" :key="group.group_id" class="card">
        <div class="card-header">
          <div class="flex justify-between items-center">
            <h3 class="text-lg font-medium text-gray-900">{{ group.group_name }}</h3>
            <div class="text-sm text-gray-500">
              总BNB: {{ group.total_bnb }} BNB
            </div>
          </div>
        </div>
        <div class="card-body">
          <!-- 代币总额 -->
          <div class="mb-6">
            <h4 class="text-md font-medium text-gray-700 mb-3">代币总额</h4>
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
              <div
                v-for="token in group.token_totals"
                :key="token.symbol"
                class="bg-gray-50 p-3 rounded-lg"
              >
                <div class="text-sm font-medium text-gray-900">{{ token.symbol }}</div>
                <div class="text-lg font-bold text-primary-600">{{ formatBalance(token.balance) }}</div>
              </div>
            </div>
          </div>

          <!-- 地址详情 -->
          <div>
            <h4 class="text-md font-medium text-gray-700 mb-3">地址详情</h4>
            <div class="space-y-4">
              <div
                v-for="address in group.addresses"
                :key="address.address"
                class="border border-gray-200 rounded-lg p-4"
              >
                <div class="flex justify-between items-center mb-2">
                  <div>
                    <div class="text-sm font-mono text-gray-900">
                      {{ truncateAddress(address.address) }}
                    </div>
                    <div class="text-xs text-gray-500">{{ address.label || '无标签' }}</div>
                  </div>
                  <div class="text-sm font-medium text-gray-900">
                    {{ address.bnb_balance }} BNB
                  </div>
                </div>
                
                <div class="grid grid-cols-2 sm:grid-cols-4 gap-2 mt-3">
                  <div
                    v-for="token in address.token_balances"
                    :key="token.symbol"
                    class="text-center"
                  >
                    <div class="text-xs text-gray-500">{{ token.symbol }}</div>
                    <div class="text-sm font-medium">{{ formatBalance(token.balance) }}</div>
                  </div>
                </div>
              </div>
            </div>
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
        请先添加钱包地址
      </p>
      <div class="mt-6">
        <router-link to="/addresses" class="btn-primary">
          添加地址
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'
import { useChainStore } from '@/stores/chain'

const authStore = useAuthStore()
const walletStore = useWalletStore()
const chainStore = useChainStore()
const loading = ref(false)
const balanceData = ref([])

onMounted(async () => {
  loading.value = true
  try {
    // 获取当前选择的链ID
    const currentChainId = chainStore.currentChain
    
    if (!currentChainId) {
      window.showNotification('warning', '请先选择区块链网络')
      return
    }
    
    // 使用优化的批量查询方法
    const results = await walletStore.fetchAllGroupsBalance(
      authStore.userId, 
      currentChainId,
      (progress) => {
        // 可以在这里添加进度处理逻辑
        console.log('Balance loading progress:', progress)
      }
    )
    
    balanceData.value = results
    
  } catch (error) {
    window.showNotification('error', '加载余额数据失败')
    console.error('Failed to load balance data:', error)
  } finally {
    loading.value = false
  }
})

const truncateAddress = (address) => {
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const formatBalance = (balance) => {
  const num = parseFloat(balance)
  if (num === 0) return '0'
  if (num < 0.000001) return '<0.000001'
  return num.toFixed(6)
}
</script>
