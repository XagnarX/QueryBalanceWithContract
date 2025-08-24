<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <!-- 页面标题 -->
    <div class="px-4 py-6 sm:px-0">
      <h1 class="text-3xl font-bold text-gray-900">仪表板</h1>
      <p class="mt-2 text-gray-600">欢迎使用钱包地址管理系统</p>
    </div>

    <!-- 统计卡片 -->
    <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4 mb-8">
      <div class="card">
        <div class="card-body">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-primary-500 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-2m-14 0h2m-2 0h-2m-14 0h2m16 0a2 2 0 002-2V7a2 2 0 00-2-2H9a2 2 0 00-2 2v10a2 2 0 002 2h10z"></path>
                </svg>
              </div>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">分组数量</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.groupCount }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-body">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-success-500 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"></path>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"></path>
                </svg>
              </div>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">钱包地址</dt>
                <dd class="text-lg font-medium text-gray-900">{{ stats.addressCount }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-body">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-warning-500 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                </svg>
              </div>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">BNB总额</dt>
                <dd class="text-lg font-medium text-gray-900">{{ totalBNB }} BNB</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-body">
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <div class="w-8 h-8 bg-danger-500 rounded-md flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
            </div>
            <div class="ml-5 w-0 flex-1">
              <dl>
                <dt class="text-sm font-medium text-gray-500 truncate">代币种类</dt>
                <dd class="text-lg font-medium text-gray-900">{{ walletStore.tokens.length }}</dd>
              </dl>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 快速操作 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg leading-6 font-medium text-gray-900">快速操作</h3>
        </div>
        <div class="card-body">
          <div class="space-y-3">
            <router-link to="/groups" class="btn-primary w-full">
              创建新分组
            </router-link>
            <router-link to="/addresses" class="btn-secondary w-full">
              添加钱包地址
            </router-link>
            <router-link to="/balance" class="btn-secondary w-full">
              查询余额
            </router-link>
            <router-link to="/summary" class="btn-secondary w-full">
              分组统计
            </router-link>
          </div>
        </div>
      </div>

      <div class="card">
        <div class="card-header">
          <h3 class="text-lg leading-6 font-medium text-gray-900">最近分组</h3>
        </div>
        <div class="card-body">
          <div v-if="walletStore.groups.length === 0" class="text-gray-500 text-center py-4">
            暂无分组，<router-link to="/groups" class="text-primary-600 hover:text-primary-500">立即创建</router-link>
          </div>
          <div v-else class="space-y-3">
            <div
              v-for="group in recentGroups"
              :key="group.id"
              class="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
            >
              <div>
                <h4 class="text-sm font-medium text-gray-900">{{ group.name }}</h4>
                <p class="text-xs text-gray-500">{{ group.description || '无描述' }}</p>
              </div>
              <router-link
                :to="`/balance/${group.id}`"
                class="text-primary-600 hover:text-primary-500 text-sm font-medium"
              >
                查看余额
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading" class="flex justify-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-primary-600"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'

const authStore = useAuthStore()
const walletStore = useWalletStore()
const loading = ref(false)

const stats = computed(() => ({
  groupCount: walletStore.groups.length,
  addressCount: walletStore.addresses.length
}))

const totalBNB = computed(() => {
  if (!walletStore.balances.user) return '0.00'
  
  const userBalance = walletStore.balances.user
  if (Array.isArray(userBalance)) {
    return userBalance.reduce((total, group) => {
      const bnb = parseFloat(group.total_bnb || 0)
      return total + bnb
    }, 0).toFixed(6)
  }
  
  return '0.00'
})

const recentGroups = computed(() => {
  return walletStore.groups.slice(0, 3)
})

onMounted(async () => {
  if (!authStore.userId) return
  
  loading.value = true
  try {
    await Promise.all([
      walletStore.fetchGroups(authStore.userId),
      walletStore.fetchAddresses(authStore.userId),
      walletStore.fetchTokens(),
      walletStore.fetchUserBalance(authStore.userId)
    ])
  } catch (error) {
    window.showNotification('error', '加载数据失败')
    console.error('Failed to load dashboard data:', error)
  } finally {
    loading.value = false
  }
})
</script>
