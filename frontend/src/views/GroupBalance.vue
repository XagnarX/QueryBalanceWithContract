<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex items-center space-x-4">
        <button @click="$router.go(-1)" class="text-gray-500 hover:text-gray-700">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
          </svg>
        </button>
        <h1 class="text-3xl font-bold text-gray-900">
          {{ groupBalance?.group_name }} - 余额详情
        </h1>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
    </div>

    <!-- 余额详情 -->
    <div v-else-if="groupBalance" class="space-y-6">
      <!-- 总览卡片 -->
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">余额总览</h3>
        </div>
        <div class="card-body">
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
            <div class="text-center p-4 bg-primary-50 rounded-lg">
              <div class="text-2xl font-bold text-primary-600">{{ groupBalance.total_bnb }}</div>
              <div class="text-sm text-gray-600">BNB 总额</div>
            </div>
            <div
              v-for="token in groupBalance.token_totals.slice(0, 3)"
              :key="token.symbol"
              class="text-center p-4 bg-gray-50 rounded-lg"
            >
              <div class="text-lg font-bold text-gray-900">{{ formatBalance(token.balance) }}</div>
              <div class="text-sm text-gray-600">{{ token.symbol }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 代币余额详情 -->
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">代币余额</h3>
        </div>
        <div class="card-body">
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <div
              v-for="token in groupBalance.token_totals"
              :key="token.symbol"
              class="flex items-center space-x-3 p-4 border border-gray-200 rounded-lg"
            >
              <div class="flex-shrink-0">
                <div class="w-10 h-10 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
                  <span class="text-white text-xs font-bold">{{ token.symbol.substring(0, 2) }}</span>
                </div>
              </div>
              <div class="flex-1">
                <div class="text-sm font-medium text-gray-900">{{ token.name }}</div>
                <div class="text-lg font-bold text-gray-700">{{ formatBalance(token.balance) }}</div>
                <div class="text-xs text-gray-500">{{ token.symbol }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 地址余额详情 -->
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">地址余额详情</h3>
        </div>
        <div class="card-body">
          <div class="space-y-6">
            <div
              v-for="address in groupBalance.addresses"
              :key="address.address"
              class="border border-gray-200 rounded-lg p-6"
            >
              <div class="flex justify-between items-center mb-4">
                <div>
                  <div class="text-lg font-mono text-gray-900">{{ address.address }}</div>
                  <div class="text-sm text-gray-500">{{ address.label || '无标签' }}</div>
                </div>
                <div class="text-right">
                  <div class="text-lg font-bold text-primary-600">{{ address.bnb_balance }} BNB</div>
                </div>
              </div>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
                <div
                  v-for="token in address.token_balances"
                  :key="token.symbol"
                  class="bg-gray-50 p-3 rounded-lg text-center"
                >
                  <div class="text-sm font-medium text-gray-900">{{ token.symbol }}</div>
                  <div class="text-lg font-bold text-gray-700">{{ formatBalance(token.balance) }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 错误状态 -->
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
          d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z"
        />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">加载失败</h3>
      <p class="mt-1 text-sm text-gray-500">无法加载分组余额数据</p>
      <div class="mt-6">
        <button @click="loadGroupBalance" class="btn-primary">
          重新加载
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'

const route = useRoute()
const authStore = useAuthStore()
const walletStore = useWalletStore()

const groupId = parseInt(route.params.groupId)
const loading = ref(false)
const groupBalance = ref(null)

onMounted(() => {
  loadGroupBalance()
})

const loadGroupBalance = async () => {
  loading.value = true
  try {
    const response = await walletStore.fetchGroupBalance(authStore.userId, groupId)
    groupBalance.value = response
  } catch (error) {
    window.showNotification('error', '加载分组余额失败')
    console.error('Failed to load group balance:', error)
  } finally {
    loading.value = false
  }
}

const formatBalance = (balance) => {
  const num = parseFloat(balance)
  if (num === 0) return '0'
  if (num < 0.000001) return '<0.000001'
  return num.toFixed(6)
}
</script>
