<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <!-- 导航栏 -->
    <nav v-if="authStore.isLoggedIn" class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <h1 class="text-xl font-bold text-gray-900">钱包管理系统</h1>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                to="/dashboard"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
                active-class="border-primary-500 text-gray-900"
              >
                仪表板
              </router-link>
              <router-link
                to="/groups"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
                active-class="border-primary-500 text-gray-900"
              >
                分组管理
              </router-link>
              <router-link
                to="/addresses"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
                active-class="border-primary-500 text-gray-900"
              >
                地址管理
              </router-link>
              <router-link
                to="/balance"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
                active-class="border-primary-500 text-gray-900"
              >
                余额查询
              </router-link>
              <router-link
                to="/summary"
                class="border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 whitespace-nowrap py-2 px-1 border-b-2 font-medium text-sm"
                active-class="border-primary-500 text-gray-900"
              >
                分组统计
              </router-link>
            </div>
          </div>
          <div class="flex items-center">
            <div class="flex items-center space-x-4">
              <!-- 链选择器 -->
              <div class="hidden sm:block">
                <ChainSelector 
                  v-model="selectedChain" 
                  @change="onChainChange"
                  class="w-48"
                />
              </div>
              
              <span class="text-sm text-gray-700">
                欢迎，{{ authStore.user?.username }}
              </span>
              <button
                @click="logout"
                class="text-gray-500 hover:text-gray-700 text-sm font-medium"
              >
                退出登录
              </button>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <!-- 主内容区域 -->
    <main>
      <router-view />
    </main>

    <!-- 全局通知 -->
    <NotificationToast />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'
import { useChainStore } from '@/stores/chain'
import NotificationToast from '@/components/NotificationToast.vue'
import ChainSelector from '@/components/ChainSelector.vue'

const router = useRouter()
const authStore = useAuthStore()
const walletStore = useWalletStore()
const chainStore = useChainStore()

const selectedChain = ref(chainStore.currentChain)

onMounted(async () => {
  // 初始化认证状态
  await authStore.initAuth()
  
  // 初始化链数据
  if (authStore.isLoggedIn) {
    await chainStore.fetchChains()
    selectedChain.value = chainStore.currentChain
  }
})

const onChainChange = (chain) => {
  chainStore.setCurrentChain(chain.chain_id)
  selectedChain.value = chain.chain_id
  window.showNotification('info', `已切换到 ${chain.name}`)
}

const logout = () => {
  authStore.logout()
  walletStore.reset()
  chainStore.clearState()
  router.push('/login')
}
</script>
