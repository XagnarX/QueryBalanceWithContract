<template>
  <div class="relative">
    <label v-if="label" class="block text-sm font-medium text-gray-700 mb-2">
      {{ label }}
    </label>
    
    <div class="relative">
      <button
        @click="showDropdown = !showDropdown"
        class="w-full flex items-center justify-between px-4 py-2 border border-gray-300 rounded-md shadow-sm bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-primary-500 focus:border-primary-500"
        :class="{ 'border-red-300': error }"
      >
        <div v-if="selectedChain" class="flex items-center">
          <div class="flex-shrink-0 h-6 w-6 mr-3">
            <div 
              class="h-6 w-6 rounded-full flex items-center justify-center text-white text-xs font-bold"
              :class="selectedChain.is_testnet ? 'bg-orange-500' : 'bg-green-500'"
            >
              {{ selectedChain.symbol.substring(0, 2) }}
            </div>
          </div>
          <div class="text-left">
            <div class="text-sm font-medium text-gray-900">{{ selectedChain.name }}</div>
            <div class="text-xs text-gray-500">Chain ID: {{ selectedChain.chain_id }}</div>
          </div>
        </div>
        
        <div v-else class="text-gray-500">
          选择区块链网络
        </div>
        
        <svg class="w-5 h-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
        </svg>
      </button>
      
      <!-- 下拉菜单 -->
      <div
        v-if="showDropdown"
        class="absolute z-50 mt-1 w-full bg-white shadow-lg max-h-60 rounded-md py-1 text-base ring-1 ring-black ring-opacity-5 overflow-auto focus:outline-none"
      >
        <div
          v-for="chain in availableChains"
          :key="chain.chain_id"
          @click="selectChain(chain)"
          class="cursor-pointer select-none relative py-2 pl-3 pr-9 hover:bg-gray-100"
          :class="{ 'bg-primary-50 text-primary-900': chain.chain_id === selectedChain?.chain_id }"
        >
          <div class="flex items-center">
            <div class="flex-shrink-0 h-6 w-6 mr-3">
              <div 
                class="h-6 w-6 rounded-full flex items-center justify-center text-white text-xs font-bold"
                :class="chain.is_testnet ? 'bg-orange-500' : 'bg-green-500'"
              >
                {{ chain.symbol.substring(0, 2) }}
              </div>
            </div>
            <div class="flex-1">
              <div class="flex items-center justify-between">
                <div>
                  <div class="text-sm font-medium text-gray-900">{{ chain.name }}</div>
                  <div class="text-xs text-gray-500">
                    Chain ID: {{ chain.chain_id }}
                    <span v-if="chain.is_testnet" class="ml-2 inline-flex items-center px-2 py-0.5 rounded text-xs font-medium bg-orange-100 text-orange-800">
                      测试网
                    </span>
                  </div>
                </div>
                <div class="text-xs text-gray-400">
                  <div v-if="chain.rpc_endpoints?.length">{{ chain.rpc_endpoints.length }} RPC</div>
                  <div v-if="chain.balance_contracts?.length">{{ chain.balance_contracts.length }} 合约</div>
                </div>
              </div>
            </div>
          </div>
          
          <!-- 选中标记 -->
          <span
            v-if="chain.chain_id === selectedChain?.chain_id"
            class="absolute inset-y-0 right-0 flex items-center pr-4 text-primary-600"
          >
            <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
            </svg>
          </span>
        </div>
        
        <!-- 没有可用链的提示 -->
        <div v-if="availableChains.length === 0" class="py-4 text-center text-gray-500 text-sm">
          暂无可用的区块链网络
        </div>
      </div>
    </div>
    
    <!-- 错误提示 -->
    <p v-if="error" class="mt-1 text-sm text-red-600">{{ error }}</p>
    
    <!-- 链信息展示 -->
    <div v-if="showInfo && selectedChain" class="mt-3 p-3 bg-gray-50 rounded-md">
      <div class="space-y-2 text-sm">
        <div class="flex justify-between">
          <span class="text-gray-500">网络名称:</span>
          <span class="font-medium">{{ selectedChain.name }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-500">链ID:</span>
          <span class="font-medium">{{ selectedChain.chain_id }}</span>
        </div>
        <div class="flex justify-between">
          <span class="text-gray-500">原生代币:</span>
          <span class="font-medium">{{ selectedChain.symbol }}</span>
        </div>
        <div v-if="selectedChain.rpc_endpoints?.length" class="flex justify-between">
          <span class="text-gray-500">RPC端点:</span>
          <span class="font-medium">{{ selectedChain.rpc_endpoints.length }} 个</span>
        </div>
        <div v-if="selectedChain.balance_contracts?.length" class="flex justify-between">
          <span class="text-gray-500">查询合约:</span>
          <span class="font-medium">{{ selectedChain.balance_contracts.length }} 个</span>
        </div>
        <div v-if="selectedChain.tokens?.length" class="flex justify-between">
          <span class="text-gray-500">支持代币:</span>
          <span class="font-medium">{{ selectedChain.tokens.length }} 个</span>
        </div>
        <div v-if="selectedChain.block_explorer_url" class="flex justify-between">
          <span class="text-gray-500">浏览器:</span>
          <a 
            :href="selectedChain.block_explorer_url" 
            target="_blank" 
            class="text-primary-600 hover:text-primary-500 font-medium"
          >
            查看 →
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import { useChainStore } from '@/stores/chain'

const props = defineProps({
  modelValue: {
    type: Number,
    default: null
  },
  label: {
    type: String,
    default: ''
  },
  error: {
    type: String,
    default: ''
  },
  showInfo: {
    type: Boolean,
    default: false
  },
  includeTestnet: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue', 'change'])

const chainStore = useChainStore()
const showDropdown = ref(false)

const availableChains = computed(() => {
  let chains = chainStore.activeChains
  if (!props.includeTestnet) {
    chains = chains.filter(chain => !chain.is_testnet)
  }
  return chains
})

const selectedChain = computed(() => {
  return availableChains.value.find(chain => chain.chain_id === props.modelValue)
})

const selectChain = (chain) => {
  emit('update:modelValue', chain.chain_id)
  emit('change', chain)
  showDropdown.value = false
}

const handleClickOutside = (event) => {
  if (!event.target.closest('.relative')) {
    showDropdown.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  
  // 加载链数据
  if (chainStore.chains.length === 0) {
    chainStore.fetchChains()
  }
})

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside)
})

// 监听链数据变化，如果没有选中的链但有可用链，自动选择第一个
watch(() => chainStore.chains, (newChains) => {
  if (!props.modelValue && availableChains.value.length > 0) {
    selectChain(availableChains.value[0])
  }
}, { immediate: true })
</script>
