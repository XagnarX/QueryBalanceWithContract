<template>
  <div class="max-w-full mx-auto py-6 sm:px-3 lg:px-4">
    <div class="px-2 py-6 sm:px-0">
      <div class="flex justify-between items-center">
        <div>
          <h1 class="text-3xl font-bold text-gray-900">分组统计汇总</h1>
          <p class="mt-2 text-gray-600">查看所有分组的余额统计和总览</p>
        </div>
        <div class="flex items-center space-x-4">          
          <!-- 设置按钮 -->
          <button @click="showSettings = !showSettings" class="btn-secondary">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path>
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
            </svg>
          </button>
          
          <!-- 查询按钮组 -->
          <div class="flex space-x-2">
            <button @click="showCreateGroupModal = true" class="btn-secondary">
              创建新分组
            </button>
            <button @click="showAddAddressModal = true" class="btn-secondary">
              添加地址
            </button>
            <button @click="loadAllGroupsBalance" :disabled="loading" class="btn-primary">
              <svg v-if="progressInfo.show" class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ progressInfo.show ? '查询中...' : '查询所有余额' }}
            </button>
            <button @click="refreshData" :disabled="loading" class="btn-secondary">
              <svg v-if="loading" class="animate-spin -ml-1 mr-3 h-5 w-5" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              {{ loading ? '加载中...' : '刷新分组' }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 设置面板 -->
    <div v-if="showSettings" class="mb-6">
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">设置</h3>
        </div>
        <div class="card-body">
          <!-- 全局RPC管理 -->
          <div class="mb-6 p-4 bg-blue-50 rounded-lg">
            <div class="flex items-center justify-between mb-3">
              <h4 class="text-md font-medium text-gray-800">全局RPC管理</h4>
              <button 
                @click="showRpcManager = true"
                class="btn-secondary btn-sm"
              >
                添加新RPC
              </button>
            </div>
            <p class="text-sm text-gray-600">管理全局可用的RPC节点列表，添加新的RPC端点。每个分组可以独立选择要使用的RPC节点。</p>
          </div>
          
          <!-- 全局Token管理 -->
          <div class="mb-6 p-4 bg-gray-50 rounded-lg">
            <div class="flex items-center justify-between mb-3">
              <h4 class="text-md font-medium text-gray-800">全局Token管理</h4>
              <button 
                @click="showTokenManager = true"
                class="btn-secondary btn-sm"
              >
                添加新Token
              </button>
            </div>
            <p class="text-sm text-gray-600">管理全局可用的Token列表，添加新的Token合约。每个分组可以独立选择要查询的Token。</p>
          </div>
          

        </div>
      </div>
    </div>



    <!-- 总体统计卡片 -->
    <div class="mb-8">
      <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-5">
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
                  <dt class="text-sm font-medium text-gray-500 truncate">活跃分组</dt>
                  <dd class="text-lg font-medium text-gray-900">{{ totalStats.activeGroups }}</dd>
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
                  <dt class="text-sm font-medium text-gray-500 truncate">总地址数</dt>
                  <dd class="text-lg font-medium text-gray-900">{{ totalStats.totalAddresses }}</dd>
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
                  <dd class="text-lg font-medium text-gray-900">{{ totalStats.totalBNB }} BNB</dd>
                </dl>
              </div>
            </div>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12a2 2 0 01-2-2V6a2 2 0 012-2h6a2 2 0 012 2v4a2 2 0 01-2 2M5 12a2 2 0 00-2 2v4a2 2 0 002 2h6a2 2 0 002-2v-4a2 2 0 00-2-2m8-8a2 2 0 012-2h2a2 2 0 012 2v8a2 2 0 01-2 2h-2a2 2 0 01-2-2V4z"></path>
                  </svg>
                </div>
              </div>
              <div class="ml-5 w-0 flex-1">
                <dl>
                  <dt class="text-sm font-medium text-gray-500 truncate">RPC节点</dt>
                  <dd class="text-lg font-medium text-gray-900">{{ totalStats.totalRpcNodes }}</dd>
                </dl>
              </div>
            </div>
          </div>
        </div>

        <div class="card">
          <div class="card-body">
            <div class="flex items-center">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
                  <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                  </svg>
                </div>
              </div>
              <div class="ml-5 w-0 flex-1">
                <dl>
                  <dt class="text-sm font-medium text-gray-500 truncate">Token合约</dt>
                  <dd class="text-lg font-medium text-gray-900">{{ totalStats.totalTokens }}</dd>
                </dl>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 分组卡片拖拽列表 -->
    <div v-if="groupsData.length > 0" class="space-y-6">
      <!-- 排序提示 -->
      <div class="bg-blue-50 border border-blue-200 rounded-lg p-3">
        <div class="flex items-center space-x-2">
          <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16V4m0 0L3 8m4-4l4 4m6 0v12m0 0l4-4m-4 4l-4-4"></path>
          </svg>
          <span class="text-sm font-medium text-blue-800">拖拽卡片来重新排序分组</span>
          <span v-if="sortingInProgress" class="inline-flex items-center px-2 py-1 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
            保存中...
          </span>
        </div>
      </div>

      <!-- 可拖拽的分组列表 -->
      <draggable 
        v-model="groupsData" 
        @start="onDragStart"
        @end="onDragEnd"
        item-key="group_id"
        :disabled="sortingInProgress"
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-5 gap-6"
        ghost-class="ghost"
        chosen-class="chosen"
      >
        <template #item="{ element: group, index }">
          <div
            :class="[
              'card hover:shadow-lg transition-all duration-200 cursor-move',
              sortingInProgress ? 'opacity-50' : ''
            ]"
          >
        <div class="card-header">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-medium text-gray-900 truncate">
              {{ group.group_name }}
            </h3>
            <div class="flex-shrink-0 space-x-2">
              <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-primary-100 text-primary-800">
                {{ group.address_count }} 地址
              </span>
              <span v-if="!group.address_validation_passed" class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800">
                {{ group.invalid_address_count }} 无效
              </span>
            </div>
          </div>
          
          <!-- 倒计时控制区域 -->
          <div class="mt-3 flex items-center justify-between">
            <!-- 倒计时显示 -->
            <div v-if="groupCountdowns.get(group.group_id)?.enabled" class="flex items-center space-x-2 text-sm">
              <div class="flex items-center space-x-1">
                <svg class="w-4 h-4 text-orange-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
                <span class="font-mono text-orange-600 font-medium">
                  {{ formatCountdownTime(groupCountdowns.get(group.group_id)?.remaining || 0) }}
                </span>
              </div>
              <!-- 倒计时进度条 -->
              <div class="w-16 h-2 bg-gray-200 rounded-full overflow-hidden">
                <div 
                  class="h-full bg-orange-500 transition-all duration-1000 ease-linear"
                  :style="{ width: `${getCountdownProgress(group.group_id)}%` }"
                ></div>
              </div>
            </div>
            
            <!-- 倒计时设置 -->
            <div class="flex items-center space-x-1">
              <!-- 时间输入 -->
              <input 
                type="number" 
                :value="groupCountdowns.get(group.group_id)?.total || 10"
                @change="setGroupCountdownTime(group.group_id, $event.target.value)"
                min="10" 
                max="7200"
                class="w-16 px-1 py-1 text-xs border border-gray-300 rounded focus:ring-1 focus:ring-primary-500 focus:border-primary-500"
                placeholder="10"
              >
              <span class="text-xs text-gray-500">秒</span>
              
              <!-- 开始/停止按钮 -->
              <button 
                @click="toggleGroupCountdown(group.group_id)"
                :class="[
                  'px-2 py-1 text-xs font-medium rounded-md transition-colors',
                  groupCountdowns.get(group.group_id)?.enabled 
                    ? 'bg-red-100 text-red-700 hover:bg-red-200' 
                    : 'bg-green-100 text-green-700 hover:bg-green-200'
                ]"
              >
                {{ groupCountdowns.get(group.group_id)?.enabled ? '停止' : '开始' }}
              </button>
            </div>
          </div>
          
          <!-- 分组RPC选择器 -->
          <div class="mt-3 p-3 bg-blue-50 rounded-lg">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-700">RPC节点</span>
              <button 
                @click="toggleGroupRpcSelector(group.group_id)"
                class="text-xs text-blue-600 hover:text-blue-800"
              >
                {{ isGroupRpcSelectorExpanded(group.group_id) ? '收起' : '选择' }}
              </button>
            </div>
            
            <!-- RPC选择概要 -->
            <div v-if="!isGroupRpcSelectorExpanded(group.group_id)" class="text-xs text-gray-500">
              <span v-if="getGroupSelectedRpc(group.group_id)">
                {{ availableRpcEndpoints.find(rpc => rpc.id === getGroupSelectedRpc(group.group_id))?.name || 'RPC节点' }}
              </span>
              <span v-else>未选择RPC</span>
            </div>
            
            <!-- 展开的RPC选择器 -->
            <div v-if="isGroupRpcSelectorExpanded(group.group_id)" class="space-y-2">
              <div class="space-y-1">
                <label 
                  v-for="rpc in availableRpcEndpoints" 
                  :key="rpc.id"
                  class="flex items-center space-x-2 p-2 border border-gray-200 rounded text-xs cursor-pointer hover:bg-blue-100"
                  :class="getGroupSelectedRpc(group.group_id) === rpc.id ? 'bg-blue-50 border-blue-300' : 'bg-white'"
                >
                  <input 
                    type="radio" 
                    :name="`rpc-${group.group_id}`"
                    :checked="getGroupSelectedRpc(group.group_id) === rpc.id"
                    @change="setGroupSelectedRpc(group.group_id, rpc.id)"
                    class="w-3 h-3 text-blue-600 focus:ring-blue-500"
                  >
                  <div class="flex flex-col flex-1 min-w-0">
                    <span class="text-xs font-medium truncate">{{ rpc.name }}</span>
                    <span class="text-xs text-gray-500 truncate">{{ rpc.url }}</span>
                  </div>
                </label>
              </div>
            </div>
          </div>

          <!-- 分组Token选择器 -->
          <div class="mt-3 p-3 bg-gray-50 rounded-lg">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-medium text-gray-700">查询Token</span>
              <button 
                @click="toggleGroupTokenSelector(group.group_id)"
                class="text-xs text-primary-600 hover:text-primary-800"
              >
                {{ isGroupTokenSelectorExpanded(group.group_id) ? '收起' : '设置' }}
              </button>
            </div>
            
            <!-- Token选择概要 -->
            <div v-if="!isGroupTokenSelectorExpanded(group.group_id)" class="text-xs text-gray-500">
              已选择 {{ getGroupSelectedTokens(group.group_id).length }} 个Token
            </div>
            
            <!-- 展开的Token选择器 -->
            <div v-if="isGroupTokenSelectorExpanded(group.group_id)" class="space-y-2">
              <div class="grid grid-cols-2 gap-1">
                <label 
                  v-for="token in availableTokens" 
                  :key="token.id"
                  class="flex items-center space-x-1 p-1 border border-gray-200 rounded text-xs cursor-pointer"
                  :class="isTokenSelectedForGroup(group.group_id, token.id) ? 'bg-primary-50 border-primary-300' : 'bg-white'"
                >
                  <input 
                    type="checkbox" 
                    :checked="isTokenSelectedForGroup(group.group_id, token.id)"
                    @change="toggleTokenForGroup(group.group_id, token.id)"
                    class="w-3 h-3 rounded border-gray-300 text-primary-600 focus:ring-primary-500"
                  >
                  <div class="flex items-center space-x-1 min-w-0">
                    <div class="w-4 h-4 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center flex-shrink-0">
                      <span class="text-white text-xs font-bold">{{ token.symbol.substring(0, 1) }}</span>
                    </div>
                    <span class="text-xs font-medium truncate">{{ token.symbol }}</span>
                  </div>
                </label>
              </div>
              
              <div class="flex items-center justify-between text-xs">
                <span class="text-gray-500">{{ getGroupSelectedTokens(group.group_id).length }} 个已选</span>
                <div class="space-x-2">
                  <button @click="selectAllTokensForGroup(group.group_id)" class="text-primary-600 hover:text-primary-700">全选</button>
                  <button @click="clearTokensForGroup(group.group_id)" class="text-gray-500 hover:text-gray-700">清空</button>
                </div>
              </div>
            </div>
          </div>
        </div>
        
        <div class="card-body">
          <!-- 地址验证失败状态 -->
          <div v-if="!group.address_validation_passed" class="py-4">
            <div class="text-center mb-3">
              <div class="text-red-400 mb-2">
                <svg class="w-8 h-8 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <p class="text-sm text-red-600 font-medium mb-3">地址验证失败</p>
              <div class="text-sm text-gray-600 mb-3">
                有效地址: {{ group.valid_address_count }}, 无效地址: {{ group.invalid_address_count }}
              </div>
            </div>
            
            <!-- 显示无效地址列表 -->
            <div class="max-h-32 overflow-y-auto">
              <div class="text-xs text-red-600 space-y-1">
                <div v-for="invalidAddr in group.invalid_addresses" :key="invalidAddr.id" class="p-2 bg-red-50 rounded border-l-2 border-red-200">
                  <div class="font-medium">{{ invalidAddr.label || '未命名' }}</div>
                  <div class="font-mono break-all">{{ invalidAddr.address }}</div>
                  <div class="text-red-500">{{ invalidAddr.error }}</div>
                </div>
              </div>
            </div>
          </div>

          <!-- 余额未加载状态（仅在地址验证通过时显示）-->
          <div v-else-if="!group.balance_loaded && !group.balance_loading" class="text-center py-4">
            <div class="text-gray-400 mb-2">
              <svg class="w-8 h-8 mx-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
              </svg>
            </div>
            <p class="text-sm text-gray-500 mb-3">点击查询余额</p>
            <button 
              @click="loadGroupBalance(index)"
              class="btn-primary btn-sm w-full"
            >
              查询余额
            </button>
          </div>

          <!-- 余额加载中状态 -->
          <div v-else-if="group.balance_loading" class="py-4">
            <div class="text-center mb-3">
              <div class="animate-spin w-6 h-6 border-2 border-primary-600 border-t-transparent rounded-full mx-auto mb-2"></div>
              <p class="text-sm text-gray-500">查询 {{ group.address_count }} 个地址...</p>
            </div>
            
            <!-- 分组内进度条 -->
            <div v-if="group.queryProgress" class="mt-3">
              <div class="flex items-center justify-between text-xs text-gray-500 mb-1">
                <span>查询进度</span>
                <span>{{ group.queryProgress.current || 0 }} / {{ group.queryProgress.total || group.address_count }}</span>
              </div>
              <div class="w-full bg-gray-200 rounded-full h-2">
                <div
                  class="bg-primary-600 h-2 rounded-full transition-all duration-300"
                  :style="{ width: `${group.queryProgress.total > 0 ? (group.queryProgress.current / group.queryProgress.total) * 100 : 0}%` }"
                ></div>
              </div>
              <p class="text-xs text-gray-400 mt-1">{{ group.queryProgress.message || '正在查询...' }}</p>
            </div>
          </div>

          <!-- 余额已加载状态 -->
          <div v-else>
            <!-- BNB余额 -->
            <div class="mb-4">
              <div class="flex items-center justify-between">
                <span class="text-sm font-medium text-gray-500">BNB余额</span>
                <div class="text-right">
                  <div class="flex items-center space-x-2">
                    <span class="text-lg font-bold text-primary-600">{{ group.total_bnb || '0' }}</span>
                    <!-- BNB余额变化指示器 -->
                    <div v-if="group.balance_changes && group.balance_changes.bnb_change" class="flex items-center">
                      <div class="flex items-center text-xs px-1 py-0.5 rounded"
                           :class="group.balance_changes.bnb_change.difference > 0 ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'">
                        <svg class="w-3 h-3 mr-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path v-if="group.balance_changes.bnb_change.difference > 0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"></path>
                          <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"></path>
                        </svg>
                        <span>{{ Math.abs(group.balance_changes.bnb_change.difference).toFixed(6) }}</span>
                      </div>
                    </div>
                  </div>
                  <!-- 最后刷新时间 -->
                  <div v-if="group.last_refresh_time" class="text-xs text-gray-400 mt-0.5">
                    {{ group.last_refresh_time }}
                  </div>
                </div>
              </div>
            </div>

            <!-- 主要代币余额 -->
            <div class="space-y-2">
              <div
                v-for="token in getMainTokens(group.token_totals, group.group_id)"
                :key="token.contract_address"
                class="flex items-center justify-between text-sm"
              >
                <span class="text-gray-600">{{ token.symbol }}</span>
                <div class="text-right">
                  <div class="flex items-center space-x-2">
                    <span 
                      class="font-medium"
                      :class="parseFloat(token.balance || 0) > 0 ? 'text-gray-900' : 'text-gray-400'"
                    >
                      {{ formatBalance(token.balance) }}
                    </span>
                    <!-- Token余额变化指示器 -->
                    <div v-if="getTokenChange(group, token)" class="flex items-center">
                      <div class="flex items-center text-xs px-1 py-0.5 rounded"
                           :class="getTokenChange(group, token).difference > 0 ? 'bg-green-100 text-green-600' : 'bg-red-100 text-red-600'">
                        <svg class="w-2.5 h-2.5 mr-0.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path v-if="getTokenChange(group, token).difference > 0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10l7-7m0 0l7 7m-7-7v18"></path>
                          <path v-else stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3"></path>
                        </svg>
                        <span>{{ formatBalance(Math.abs(getTokenChange(group, token).difference).toString()) }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- 查看详情按钮 -->
            <div class="mt-4 pt-3 border-t border-gray-200">
              <button 
                @click="viewGroupDetail(group.group_id)"
                class="text-primary-600 hover:text-primary-500 text-sm font-medium w-full text-left"
              >
                查看详细余额 →
              </button>
            </div>
          </div>
        </div>
          </div>
        </template>
      </draggable>
    </div>

    <!-- 加载状态 -->
    <div v-else-if="loading" class="flex justify-center py-12">
      <div class="text-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
        <p class="mt-4 text-gray-500">正在加载分组数据...</p>
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
      <h3 class="mt-2 text-sm font-medium text-gray-900">暂无分组数据</h3>
      <p class="mt-1 text-sm text-gray-500">
        请先创建分组并添加钱包地址
      </p>
      <div class="mt-6">
        <router-link to="/groups" class="btn-primary">
          创建分组
        </router-link>
      </div>
    </div>

    <!-- 代币总额统计表格 -->
    <div v-if="tokenSummary.length > 0" class="mt-8">
      <div class="card">
        <div class="card-header">
          <h3 class="text-lg font-medium text-gray-900">代币总额统计</h3>
        </div>
        <div class="card-body p-0">
          <div class="overflow-x-auto">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50">
                <tr>
                  <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    代币
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    总额
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    分布分组数
                  </th>
                  <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                    最大持有
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-for="token in tokenSummary" :key="token.symbol">
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
                    {{ formatBalance(token.totalAmount) }}
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm text-gray-500">
                    {{ token.groupCount }} 个分组
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-right text-sm text-gray-500">
                    {{ formatBalance(token.maxAmount) }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Create Group Modal -->
  <div v-if="showCreateGroupModal" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4">
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showCreateGroupModal = false"></div>
      <div class="bg-white rounded-lg p-6 max-w-md w-full relative z-10">
        <h3 class="text-lg font-medium text-gray-900 mb-4">创建新分组</h3>
        <form @submit.prevent="createGroup">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">分组名称</label>
              <input v-model="newGroup.name" class="input mt-1" required />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">描述（可选）</label>
              <textarea v-model="newGroup.description" class="input mt-1" rows="3"></textarea>
            </div>
          </div>
          <div class="mt-6 flex space-x-3">
            <button type="submit" class="btn-primary">创建</button>
            <button type="button" @click="showCreateGroupModal = false" class="btn-secondary">取消</button>
          </div>
        </form>
      </div>
    </div>
  </div>

  <!-- Add Address Modal -->
  <div v-if="showAddAddressModal" class="fixed inset-0 z-50 overflow-y-auto">
    <div class="flex items-center justify-center min-h-screen px-4">
      <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showAddAddressModal = false"></div>
      <div class="bg-white rounded-lg p-6 max-w-md w-full relative z-10">
        <h3 class="text-lg font-medium text-gray-900 mb-4">添加新地址</h3>
        <form @submit.prevent="addAddress">
          <div class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">钱包地址</label>
              <textarea
                v-model="newAddress.address"
                class="input mt-1 min-h-[120px] resize-y"
                placeholder="支持批量添加，多个地址请用逗号分割&#10;0x1234...,0x5678...,0x9abc..."
                required
              ></textarea>
              <p class="text-xs text-gray-500 mt-1">支持批量添加，多个地址请用逗号分割</p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">标签（可选）</label>
              <input v-model="newAddress.label" class="input mt-1" placeholder="给地址添加标签" />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">分组（可选）</label>
              <select v-model="newAddress.group_id" class="input mt-1">
                <option value="">选择分组</option>
                <option v-for="group in walletStore.groups" :key="group.id" :value="group.id">
                  {{ group.name }}
                </option>
              </select>
            </div>
          </div>
          <div class="mt-6 flex space-x-3">
            <button type="submit" class="btn-primary">添加</button>
            <button type="button" @click="showAddAddressModal = false" class="btn-secondary">取消</button>
          </div>
        </form>
      </div>
    </div>
  </div>

  <!-- RPC管理模态框 -->
  <div v-if="showRpcManager" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- 标题 -->
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">RPC节点管理</h3>
          <button 
            @click="showRpcManager = false"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- 添加新RPC -->
        <div class="mb-6 p-4 bg-blue-50 rounded-lg">
          <h4 class="text-md font-medium text-gray-800 mb-3">添加新RPC节点</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">节点名称</label>
              <input 
                v-model="newRpc.name"
                type="text" 
                placeholder="例如: BSC官方节点"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-primary-500 focus:border-primary-500"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">RPC URL</label>
              <input 
                v-model="newRpc.url"
                @input="formatRpcUrl"
                type="text" 
                placeholder="例如: https://bsc-dataseed1.binance.org"
                :class="[
                  'w-full px-3 py-2 border rounded-md focus:ring-primary-500 focus:border-primary-500',
                  newRpc.url && !(newRpc.url.startsWith('http://') || newRpc.url.startsWith('https://'))
                    ? 'border-red-300 bg-red-50' 
                    : 'border-gray-300'
                ]"
              >
              <p v-if="newRpc.url && !(newRpc.url.startsWith('http://') || newRpc.url.startsWith('https://'))" 
                 class="mt-1 text-sm text-red-600">
                RPC URL必须以http://或https://开头
              </p>
              <p v-else class="mt-1 text-sm text-gray-500">
                请输入完整的RPC URL，包括协议前缀
              </p>
            </div>
            <div class="md:col-span-2 flex justify-end">
              <button 
                @click="addNewRpc"
                :disabled="!canAddRpc"
                class="btn-primary"
                :class="!canAddRpc ? 'opacity-50 cursor-not-allowed' : ''"
              >
                添加RPC节点
              </button>
            </div>
          </div>
        </div>

        <!-- RPC列表 -->
        <div>
          <h4 class="text-md font-medium text-gray-800 mb-3">现有RPC节点列表</h4>
          <div class="max-h-64 overflow-y-auto">
            <div class="space-y-2">
              <div 
                v-for="rpc in availableRpcEndpoints" 
                :key="rpc.id"
                class="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50"
              >
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 bg-gradient-to-r from-green-500 to-blue-600 rounded-full flex items-center justify-center">
                    <span class="text-white text-xs font-bold">{{ rpc.name.substring(0, 2) }}</span>
                  </div>
                  <div class="flex-1 min-w-0">
                    <div class="font-medium text-gray-900">{{ rpc.name }}</div>
                    <div class="text-sm text-gray-500 truncate">{{ rpc.url }}</div>
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <span :class="[
                    'px-2 py-1 text-xs rounded-full',
                    rpc.is_active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
                  ]">
                    {{ rpc.is_active ? '活跃' : '禁用' }}
                  </span>
                  <button 
                    @click="toggleRpcStatus(rpc)"
                    class="text-sm text-primary-600 hover:text-primary-800"
                  >
                    {{ rpc.is_active ? '禁用' : '启用' }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="mt-6 flex justify-end space-x-3">
          <button 
            @click="showRpcManager = false"
            class="btn-secondary"
          >
            关闭
          </button>
        </div>
      </div>
    </div>
  </div>

  <!-- Token管理模态框 -->
  <div v-if="showTokenManager" class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50">
    <div class="relative top-20 mx-auto p-5 border w-11/12 max-w-2xl shadow-lg rounded-md bg-white">
      <div class="mt-3">
        <!-- 标题 -->
        <div class="flex items-center justify-between mb-4">
          <h3 class="text-lg font-medium text-gray-900">Token管理</h3>
          <button 
            @click="showTokenManager = false"
            class="text-gray-400 hover:text-gray-600"
          >
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- 添加新Token -->
        <div class="mb-6 p-4 bg-gray-50 rounded-lg">
          <h4 class="text-md font-medium text-gray-800 mb-3">添加新Token</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Token名称</label>
              <input 
                v-model="newToken.name"
                type="text" 
                placeholder="例如: Tether USD"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-primary-500 focus:border-primary-500"
              >
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">Token符号</label>
              <input 
                v-model="newToken.symbol"
                type="text" 
                placeholder="例如: USDT"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-primary-500 focus:border-primary-500"
              >
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700 mb-1">合约地址</label>
              <input 
                v-model="newToken.contractAddress"
                @input="formatContractAddress"
                type="text" 
                placeholder="0xa2ae424d960c26247dd6c32edc70b295c744c43"
                :class="[
                  'w-full px-3 py-2 border rounded-md focus:ring-primary-500 focus:border-primary-500',
                  newToken.contractAddress && !/^0x[a-fA-F0-9]{40}$/.test(newToken.contractAddress) 
                    ? 'border-red-300 bg-red-50' 
                    : 'border-gray-300'
                ]"
              >
              <p v-if="newToken.contractAddress && !/^0x[a-fA-F0-9]{40}$/.test(newToken.contractAddress)" 
                 class="mt-1 text-sm text-red-600">
                合约地址必须以0x开头，且为40位十六进制字符
              </p>
              <p v-else class="mt-1 text-sm text-gray-500">
                请输入以0x开头的合约地址，例如: 0xa2ae424d960c26247dd6c32edc70b295c744c43
              </p>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">精度 (Decimals)</label>
              <input 
                v-model="newToken.decimals"
                type="number" 
                placeholder="18"
                min="0"
                max="18"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:ring-primary-500 focus:border-primary-500"
              >
            </div>
            <div class="flex items-end">
              <button 
                @click="addNewToken"
                :disabled="!canAddToken"
                class="w-full btn-primary"
                :class="!canAddToken ? 'opacity-50 cursor-not-allowed' : ''"
              >
                添加Token
              </button>
            </div>
          </div>
        </div>

        <!-- Token列表 -->
        <div>
          <h4 class="text-md font-medium text-gray-800 mb-3">现有Token列表</h4>
          <div class="max-h-64 overflow-y-auto">
            <div class="space-y-2">
              <div 
                v-for="token in availableTokens" 
                :key="token.id"
                class="flex items-center justify-between p-3 border border-gray-200 rounded-lg hover:bg-gray-50"
              >
                <div class="flex items-center space-x-3">
                  <div class="w-8 h-8 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center">
                    <span class="text-white text-xs font-bold">{{ token.symbol.substring(0, 2) }}</span>
                  </div>
                  <div>
                    <div class="font-medium text-gray-900">{{ token.name }}</div>
                    <div class="text-sm text-gray-500">{{ token.symbol }}</div>
                    <div class="text-xs text-gray-400 font-mono">{{ token.contract_address }}</div>
                  </div>
                </div>
                <div class="flex items-center space-x-2">
                  <span :class="[
                    'px-2 py-1 text-xs rounded-full',
                    token.is_active ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
                  ]">
                    {{ token.is_active ? '活跃' : '禁用' }}
                  </span>
                  <button 
                    @click="toggleTokenStatus(token)"
                    class="text-sm text-primary-600 hover:text-primary-800"
                  >
                    {{ token.is_active ? '禁用' : '启用' }}
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 操作按钮 -->
        <div class="mt-6 flex justify-end space-x-3">
          <button 
            @click="showTokenManager = false"
            class="btn-secondary"
          >
            关闭
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'
import { useChainStore } from '@/stores/chain'
import { groupSettingsAPI } from '@/services/api'
import draggable from 'vuedraggable'

const router = useRouter()
const authStore = useAuthStore()
const walletStore = useWalletStore()

const loading = ref(false)
const groupsData = ref([])
const showSettings = ref(false)
const sortingInProgress = ref(false)

// Modal states
const showCreateGroupModal = ref(false)
const showAddAddressModal = ref(false)
const newGroup = ref({
  name: '',
  description: ''
})
const newAddress = ref({
  address: '',
  label: '',
  group_id: ''
})

// RPC管理相关
const showRpcManager = ref(false)
const newRpc = ref({
  name: '',
  url: ''
})

// Token管理相关
const showTokenManager = ref(false)
const availableTokens = ref([])
const groupTokenSelections = ref(new Map()) // 存储每个分组的Token选择
const groupRpcSelections = ref(new Map()) // 存储每个分组的RPC选择
const newToken = ref({
  name: '',
  symbol: '',
  contractAddress: '',
  decimals: 18
})

// 进度状态
const progressInfo = ref({
  show: false,
  stage: '',
  currentGroup: 0,
  totalGroups: 0,
  groupName: '',
  addressCount: 0,
  message: ''
})


// 自动刷新状态
const autoRefresh = ref({
  enabled: false,
  currentGroup: null,
  nextRoundIn: 0,
  progress: { current: 0, total: 0 },
  history: [],
  timers: {
    groupTimer: null,
    roundTimer: null,
    countdownTimer: null
  }
})

// 分组倒计时相关
const groupCountdowns = ref(new Map()) // 存储每个分组的倒计时状态
const groupTimers = ref(new Map()) // 存储每个分组的定时器

onMounted(async () => {
  // 确保chainStore已经加载了chains数据
  const chainStore = useChainStore()
  if (chainStore.chains.length === 0) {
    console.log('Chain store not loaded, waiting for chains...')
    try {
      await chainStore.fetchChains()
      console.log('Chains loaded successfully')
    } catch (error) {
      console.error('Failed to load chains:', error)
    }
  }

  // Load available RPC endpoints and tokens first (before loading settings)
  console.log('Loading available RPC endpoints and tokens...')
  await loadAvailableRpcEndpoints()
  await loadAvailableTokens()
  console.log('Available resources loaded')

  // Then load settings from database
  console.log('Loading group settings from database...')
  await loadGroupSettingsFromDB()

  // Note: Removed localStorage loading functions since we now use database
  // loadGroupTokenSelections() - REMOVED: Using database instead
  // loadGroupRpcSelections() - REMOVED: Using database instead

  // 清理blockchain service的缓存，确保使用正确的chainId
  import('@/services/blockchain').then(module => {
    module.default.clearCache()
  })
  loadGroupsData()
})

// Watch for chain switching - reload settings when chain changes
watch(() => useChainStore().currentChain, async (newChainId, oldChainId) => {
  if (newChainId && newChainId !== oldChainId) {
    console.log(`=== Chain switched from ${oldChainId} to ${newChainId} ===`)

    // Stop all running countdowns
    groupTimers.value.forEach(timer => {
      if (timer) clearInterval(timer)
    })
    groupTimers.value.clear()

    // Reload chain-specific resources
    console.log('Reloading RPC endpoints and tokens for new chain...')
    await loadAvailableRpcEndpoints()
    await loadAvailableTokens()

    // Reload settings for the new chain
    console.log('Reloading group settings for new chain...')
    await loadGroupSettingsFromDB()

    console.log('✓ Chain switch completed, all configurations reloaded')
  }
})

onBeforeUnmount(() => {
  stopAutoRefresh()

  // 清理分组倒计时定时器
  groupTimers.value.forEach(timer => {
    if (timer) clearInterval(timer)
  })
  groupTimers.value.clear()
})

// 分组倒计时相关方法
const startGroupCountdown = (groupId, seconds) => {
  // 停止现有的倒计时
  stopGroupCountdown(groupId)
  
  // 设置倒计时状态
  groupCountdowns.value.set(groupId, {
    enabled: true,
    remaining: seconds,
    total: seconds
  })
  
  // 启动倒计时定时器
  const timer = setInterval(() => {
    const countdown = groupCountdowns.value.get(groupId)
    if (!countdown) return
    
    countdown.remaining--
    
    if (countdown.remaining <= 0) {
      // 倒计时结束，自动刷新该分组
      stopGroupCountdown(groupId)
      refreshSingleGroup(groupId)
    }
  }, 1000)
  
  groupTimers.value.set(groupId, timer)
}

const stopGroupCountdown = (groupId) => {
  const timer = groupTimers.value.get(groupId)
  if (timer) {
    clearInterval(timer)
    groupTimers.value.delete(groupId)
  }
  
  const countdown = groupCountdowns.value.get(groupId)
  if (countdown) {
    countdown.enabled = false
  }
}

const toggleGroupCountdown = (groupId) => {
  const countdown = groupCountdowns.value.get(groupId)
  if (countdown && countdown.enabled) {
    stopGroupCountdown(groupId)
  } else {
    const defaultSeconds = countdown ? countdown.total : 10 // 默认10秒
    startGroupCountdown(groupId, defaultSeconds)
  }
  // Save settings to database
  saveGroupSettings(groupId)
}

const setGroupCountdownTime = (groupId, seconds) => {
  // Convert to integer
  const secondsInt = parseInt(seconds, 10)
  
  const countdown = groupCountdowns.value.get(groupId)
  if (countdown) {
    countdown.total = secondsInt
    if (!countdown.enabled) {
      countdown.remaining = secondsInt
    }
  } else {
    groupCountdowns.value.set(groupId, {
      enabled: false,
      remaining: secondsInt,
      total: secondsInt
    })
  }
  // Trigger reactivity by creating a new Map
  groupCountdowns.value = new Map(groupCountdowns.value)
  // Save settings to database
  saveGroupSettings(groupId)
}

const refreshSingleGroup = async (groupId) => {
  const groupIndex = groupsData.value.findIndex(g => g.group_id === groupId)
  if (groupIndex !== -1) {
    const group = groupsData.value[groupIndex]
    
    console.log(`=== Countdown Refresh for Group ${group.group_name} ===`)
    console.log('Current balance before refresh:', {
      bnb: group.total_bnb,
      tokens: group.token_totals?.length || 0,
      has_previous_data: !!(group.previous_total_bnb !== null || group.previous_token_totals)
    })
    
    // 保存历史数据 - 在重置状态之前保存
    if (group.balance_loaded && (group.total_bnb !== null || group.token_totals)) {
      group.previous_total_bnb = group.total_bnb
      group.previous_token_totals = group.token_totals ? JSON.parse(JSON.stringify(group.token_totals)) : null
      console.log('Saved previous balance data:', {
        previous_bnb: group.previous_total_bnb,
        previous_tokens: group.previous_token_totals?.length || 0
      })
    }
    
    // 重置分组状态
    group.balance_loaded = false
    group.balance_loading = false
    group.queryProgress = null
    
    // 检查该分组是否有选择的Token和RPC
    const selectedTokenIds = getGroupSelectedTokens(groupId)
    const selectedRpcId = getGroupSelectedRpc(groupId)
    const selectedRpc = availableRpcEndpoints.value.find(rpc => rpc.id === selectedRpcId)
    
    console.log(`Selected tokens: ${selectedTokenIds.length}`)
    console.log(`Selected RPC: ${selectedRpc ? selectedRpc.name : 'Default'} (ID: ${selectedRpcId})`)
    
    if (selectedTokenIds.length === 0) {
      // 没有选择Token，只查询BNB余额
      console.log(`Querying BNB only using RPC: ${selectedRpc ? selectedRpc.name : 'Default'}`)
      await loadGroupBNBBalance(groupIndex)
    } else {
      // 有选择Token，查询完整余额
      console.log(`Querying full balance with ${selectedTokenIds.length} tokens using RPC: ${selectedRpc ? selectedRpc.name : 'Default'}`)
      await loadGroupBalance(groupIndex)
    }
    
    console.log(`=== End Countdown Refresh for Group ${group.group_name} ===`)
    
    // 重新启动倒计时
    const countdown = groupCountdowns.value.get(groupId)
    if (countdown) {
      startGroupCountdown(groupId, countdown.total)
    }
  }
}

// 计算余额变化
const calculateBalanceChanges = (group, newBnbBalance, newTokenTotals) => {
  console.log('=== calculateBalanceChanges Debug ===')
  console.log('Group previous BNB:', group.previous_total_bnb)
  console.log('Group previous tokens:', group.previous_token_totals?.length || 0)
  console.log('New BNB:', newBnbBalance)
  console.log('New tokens:', newTokenTotals?.length || 0)
  
  // Check if we have any previous data to compare against
  if (group.previous_total_bnb === null && !group.previous_token_totals) {
    console.log('No previous data, returning null')
    return null
  }
  
  const changes = {
    bnb_change: null,
    token_changes: [],
    has_changes: false,
    timestamp: new Date().toLocaleTimeString()
  }
  
  // 计算BNB变化
  if (group.previous_total_bnb !== null && newBnbBalance !== null) {
    const prevBnb = parseFloat(group.previous_total_bnb || 0)
    const newBnb = parseFloat(newBnbBalance || 0)
    const diff = newBnb - prevBnb
    
    console.log(`BNB comparison: ${prevBnb} -> ${newBnb}, diff: ${diff}`)
    
    if (Math.abs(diff) > 0.000001) { // 忽略极小的误差
      changes.bnb_change = {
        previous: prevBnb,
        current: newBnb,
        difference: diff,
        percentage: prevBnb > 0 ? ((diff / prevBnb) * 100) : 0
      }
      changes.has_changes = true
      console.log('BNB change detected:', changes.bnb_change)
    } else {
      console.log('No significant BNB change')
    }
  }
  
  // 计算Token变化
  if (group.previous_token_totals && newTokenTotals) {
    console.log('Calculating token changes...')
    const tokenChanges = []
    const previousTokenMap = new Map()
    const currentTokenMap = new Map()
    
    // 建立上次的Token余额映射
    group.previous_token_totals.forEach(token => {
      const key = token.contract_address || token.symbol
      previousTokenMap.set(key, parseFloat(token.balance || 0))
    })
    
    // 建立当前的Token余额映射
    newTokenTotals.forEach(token => {
      const key = token.contract_address || token.symbol
      currentTokenMap.set(key, parseFloat(token.balance || 0))
    })
    
    console.log('Previous tokens:', Array.from(previousTokenMap.entries()))
    console.log('Current tokens:', Array.from(currentTokenMap.entries()))
    
    // 检查所有Token的变化
    const allTokenKeys = new Set([...previousTokenMap.keys(), ...currentTokenMap.keys()])
    
    allTokenKeys.forEach(tokenKey => {
      const prevBalance = previousTokenMap.get(tokenKey) || 0
      const currentBalance = currentTokenMap.get(tokenKey) || 0
      const diff = currentBalance - prevBalance
      
      console.log(`Token ${tokenKey}: ${prevBalance} -> ${currentBalance}, diff: ${diff}`)
      
      if (Math.abs(diff) > 0.000001) { // 忽略极小的误差
        const tokenInfo = newTokenTotals.find(t => (t.contract_address || t.symbol) === tokenKey) ||
                         group.previous_token_totals.find(t => (t.contract_address || t.symbol) === tokenKey)
        
        tokenChanges.push({
          symbol: tokenInfo?.symbol || tokenKey,
          contract_address: tokenInfo?.contract_address,
          previous: prevBalance,
          current: currentBalance,
          difference: diff,
          percentage: prevBalance > 0 ? ((diff / prevBalance) * 100) : 0
        })
        changes.has_changes = true
        console.log(`Token change detected for ${tokenInfo?.symbol || tokenKey}:`, diff)
      }
    })
    
    changes.token_changes = tokenChanges
  }
  
  console.log('Final changes result:', changes.has_changes ? changes : null)
  console.log('=== End calculateBalanceChanges Debug ===')
  
  return changes.has_changes ? changes : null
}

const formatCountdownTime = (seconds) => {
  const hours = Math.floor(seconds / 3600)
  const minutes = Math.floor((seconds % 3600) / 60)
  const secs = seconds % 60
  
  if (hours > 0) {
    return `${hours}:${minutes.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`
  }
  return `${minutes}:${secs.toString().padStart(2, '0')}`
}

const getCountdownProgress = (groupId) => {
  const countdown = groupCountdowns.value.get(groupId)
  if (!countdown || countdown.total === 0) return 0
  return ((countdown.total - countdown.remaining) / countdown.total) * 100
}

// 分组Token选择器状态
const groupTokenSelectorExpanded = ref(new Map())
// 分组RPC选择器状态
const groupRpcSelectorExpanded = ref(new Map())
const availableRpcEndpoints = ref([])

// RPC管理相关方法
const canAddRpc = computed(() => {
  return newRpc.value.name && 
         newRpc.value.url && 
         (newRpc.value.url.startsWith('http://') || newRpc.value.url.startsWith('https://'))
})

const formatRpcUrl = () => {
  let url = newRpc.value.url.trim()
  // 如果没有协议前缀，默认添加https://
  if (url && !url.startsWith('http://') && !url.startsWith('https://')) {
    url = 'https://' + url
  }
  newRpc.value.url = url
}

const addNewRpc = async () => {
  if (!canAddRpc.value) return

  try {
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain
    
    const response = await fetch(`/api/users/${authStore.userId}/rpc-endpoints`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({
        chain_id: currentChainId,
        name: newRpc.value.name,
        url: newRpc.value.url
      })
    })

    if (response.ok) {
      window.showNotification('success', 'RPC节点添加成功')
      // 重置表单
      newRpc.value = {
        name: '',
        url: ''
      }
      // 重新加载RPC列表
      await loadAvailableRpcEndpoints()
    } else {
      // 尝试获取详细错误信息
      let errorMessage = 'RPC节点添加失败'
      try {
        const errorData = await response.json()
        if (errorData.error) {
          errorMessage = errorData.error
        }
      } catch (e) {
        errorMessage = `RPC节点添加失败 (HTTP ${response.status})`
      }
      window.showNotification('error', errorMessage)
      return
    }
  } catch (error) {
    console.error('Failed to add RPC endpoint:', error)
    let errorMessage = 'RPC节点添加失败'
    if (error.message.includes('fetch')) {
      errorMessage = '网络连接失败，请检查后端服务是否正常运行'
    }
    window.showNotification('error', errorMessage)
  }
}

const toggleRpcStatus = async (rpc) => {
  try {
    const response = await fetch(`/api/users/${authStore.userId}/rpc-endpoints/${rpc.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({
        chain_id: rpc.chain_id,
        name: rpc.name,
        url: rpc.url,
        is_active: !rpc.is_active
      })
    })

    if (response.ok) {
      rpc.is_active = !rpc.is_active
      window.showNotification('success', `RPC节点${rpc.is_active ? '启用' : '禁用'}成功`)
      // 重新加载RPC列表以更新选择器
      await loadAvailableRpcEndpoints()
    } else {
      throw new Error('Failed to toggle RPC status')
    }
  } catch (error) {
    console.error('Failed to toggle RPC status:', error)
    window.showNotification('error', 'RPC节点状态切换失败')
  }
}

// Token管理相关方法
const canAddToken = computed(() => {
  return newToken.value.name && 
         newToken.value.symbol && 
         newToken.value.contractAddress && 
         /^0x[a-fA-F0-9]{40}$/.test(newToken.value.contractAddress)
})

const isGroupTokenSelectorExpanded = (groupId) => {
  return groupTokenSelectorExpanded.value.get(groupId) || false
}

const toggleGroupTokenSelector = (groupId) => {
  const current = groupTokenSelectorExpanded.value.get(groupId) || false
  groupTokenSelectorExpanded.value.set(groupId, !current)
}

const getGroupSelectedTokens = (groupId) => {
  const selected = groupTokenSelections.value.get(groupId) || []
  console.log(`getGroupSelectedTokens for group ${groupId}:`, selected)
  return selected
}

const isTokenSelectedForGroup = (groupId, tokenId) => {
  const selectedTokens = groupTokenSelections.value.get(groupId) || []
  return selectedTokens.includes(tokenId)
}

const toggleTokenForGroup = (groupId, tokenId) => {
  const selectedTokens = groupTokenSelections.value.get(groupId) || []
  if (selectedTokens.includes(tokenId)) {
    groupTokenSelections.value.set(groupId, selectedTokens.filter(id => id !== tokenId))
  } else {
    groupTokenSelections.value.set(groupId, [...selectedTokens, tokenId])
  }
  // Trigger reactivity by creating a new Map
  groupTokenSelections.value = new Map(groupTokenSelections.value)
  // Save settings to database
  saveGroupSettings(groupId)
  
  // 如果该分组已经加载过余额，重新查询
  const group = groupsData.value.find(g => g.group_id === groupId)
  if (group && group.balance_loaded) {
    const groupIndex = groupsData.value.findIndex(g => g.group_id === groupId)
    if (groupIndex !== -1) {
      // 重置状态并重新查询
      group.balance_loaded = false
      group.balance_loading = false
      group.queryProgress = null
      
      setTimeout(() => {
        const updatedSelectedTokens = getGroupSelectedTokens(groupId)
        if (updatedSelectedTokens.length === 0) {
          // 没有选择任何Token，只查询BNB
          console.log(`Group ${group.group_name} now has no tokens selected, switching to BNB only`)
          loadGroupBNBBalance(groupIndex)
        } else {
          // 有选择Token，查询完整余额
          console.log(`Group ${group.group_name} has ${updatedSelectedTokens.length} tokens selected, querying full balance`)
          loadGroupBalance(groupIndex)
        }
      }, 100)
    }
  }
}

const selectAllTokensForGroup = (groupId) => {
  const allActiveTokenIds = availableTokens.value.filter(token => token.is_active).map(token => token.id)
  groupTokenSelections.value.set(groupId, allActiveTokenIds)
  // Trigger reactivity by creating a new Map
  groupTokenSelections.value = new Map(groupTokenSelections.value)
  // Save settings to database
  saveGroupSettings(groupId)
}

const clearTokensForGroup = (groupId) => {
  groupTokenSelections.value.set(groupId, [])
  // Trigger reactivity by creating a new Map
  groupTokenSelections.value = new Map(groupTokenSelections.value)
  // Save settings to database
  saveGroupSettings(groupId)
}

const saveGroupTokenSelections = () => {
  const data = Object.fromEntries(groupTokenSelections.value)
  localStorage.setItem('groupTokenSelections', JSON.stringify(data))
}

const loadGroupTokenSelections = () => {
  try {
    const saved = localStorage.getItem('groupTokenSelections')
    if (saved) {
      const data = JSON.parse(saved)
      groupTokenSelections.value = new Map(Object.entries(data))
    }
  } catch (error) {
    console.error('Failed to load group token selections:', error)
  }
}

// RPC选择相关方法
const isGroupRpcSelectorExpanded = (groupId) => {
  return groupRpcSelectorExpanded.value.get(groupId) || false
}

const toggleGroupRpcSelector = (groupId) => {
  const current = groupRpcSelectorExpanded.value.get(groupId) || false
  groupRpcSelectorExpanded.value.set(groupId, !current)
}

const getGroupSelectedRpc = (groupId) => {
  const selected = groupRpcSelections.value.get(groupId) || null
  console.log(`getGroupSelectedRpc for group ${groupId}:`, selected, 'type:', typeof selected)
  return selected
}

const setGroupSelectedRpc = (groupId, rpcEndpointId) => {
  console.log(`setGroupSelectedRpc for group ${groupId}:`, rpcEndpointId, 'type:', typeof rpcEndpointId)
  groupRpcSelections.value.set(groupId, rpcEndpointId)
  // Trigger reactivity by creating a new Map
  groupRpcSelections.value = new Map(groupRpcSelections.value)
  // Save settings to database
  saveGroupSettings(groupId)
  
  console.log('Updated groupRpcSelections:', Object.fromEntries(groupRpcSelections.value))
  
  // 如果该分组已经加载过余额，重新查询
  const group = groupsData.value.find(g => g.group_id === groupId)
  if (group && group.balance_loaded) {
    const groupIndex = groupsData.value.findIndex(g => g.group_id === groupId)
    if (groupIndex !== -1) {
      console.log(`Group ${group.group_name} balance already loaded, triggering refresh with new RPC`)
      // 重置状态并重新查询
      group.balance_loaded = false
      group.balance_loading = false
      group.queryProgress = null
      setTimeout(() => {
        loadGroupBalance(groupIndex)
      }, 100)
    }
  }
}

const saveGroupRpcSelections = () => {
  const data = Object.fromEntries(groupRpcSelections.value)
  localStorage.setItem('groupRpcSelections', JSON.stringify(data))
}

const loadGroupRpcSelections = () => {
  try {
    const saved = localStorage.getItem('groupRpcSelections')
    if (saved) {
      const data = JSON.parse(saved)
      groupRpcSelections.value = new Map(Object.entries(data))
    }
  } catch (error) {
    console.error('Failed to load group RPC selections:', error)
  }
}

const loadAvailableRpcEndpoints = async () => {
  try {
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain
    
    console.log('=== Loading RPC Endpoints ===')
    console.log('Current Chain ID:', currentChainId)
    console.log('Available chains:', chainStore.chains?.length || 0)
    
    if (!currentChainId) {
      console.log('No current chain selected, skipping RPC load')
      return
    }

    // 只加载用户自定义RPC（不再有系统RPC）
    let allRpcEndpoints = []

    try {
      const userRpcResponse = await fetch(`/api/chains/${currentChainId}/rpc-endpoints?user_id=${authStore.userId}`, {
        headers: {
          'Authorization': `Bearer ${authStore.token}`
        }
      })
      if (userRpcResponse.ok) {
        const userData = await userRpcResponse.json()
        allRpcEndpoints = userData.rpc_endpoints || []
        console.log('User RPC endpoints loaded:', allRpcEndpoints.length)
      } else {
        console.warn('Failed to load user RPC endpoints')
      }
    } catch (error) {
      console.warn('Failed to load user RPC endpoints:', error)
    }

    // 筛选活跃的RPC节点
    availableRpcEndpoints.value = allRpcEndpoints.filter(rpc => rpc.is_active)
    console.log('Total active RPC endpoints:', availableRpcEndpoints.value.length)
    console.log('RPC endpoints:', availableRpcEndpoints.value.map(rpc => ({ name: rpc.name, url: rpc.url })))
    
    // 4. 为每个分组初始化默认RPC选择（如果还没有设置的话）
    groupsData.value.forEach(group => {
      if (!groupRpcSelections.value.has(group.group_id) && availableRpcEndpoints.value.length > 0) {
        // 默认选择第一个RPC
        groupRpcSelections.value.set(group.group_id, availableRpcEndpoints.value[0].id)
        console.log(`Set default RPC for group ${group.group_name}: ${availableRpcEndpoints.value[0].name}`)
      }
    })
    // Note: Removed localStorage save since we now use database
    // saveGroupRpcSelections() - REMOVED
    
    console.log('=== End Loading RPC Endpoints ===')
  } catch (error) {
    console.error('Failed to load RPC endpoints:', error)
  }
}

const loadAvailableTokens = async () => {
  try {
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain

    if (!currentChainId) {
      window.showNotification('warning', '请先选择区块链网络')
      return
    }

    const response = await fetch(`/api/chains/${currentChainId}/tokens?user_id=${authStore.userId}`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    if (response.ok) {
      const data = await response.json()
      availableTokens.value = data.tokens || []
      console.log('=== Loaded Available Tokens ===')
      console.log('Total tokens loaded:', availableTokens.value.length)
      availableTokens.value.forEach(token => {
        console.log(`Token: ${token.symbol} (ID: ${token.id}, Contract: ${token.contract_address}, Active: ${token.is_active})`)
      })
      console.log('=== End Available Tokens ===')
      
      // 为每个分组初始化默认Token选择（如果还没有设置的话）
      let needsRefresh = false
      groupsData.value.forEach(group => {
        if (!groupTokenSelections.value.has(group.group_id)) {
          const defaultTokens = availableTokens.value
            .filter(token => token.is_active)
            .map(token => token.id)
          groupTokenSelections.value.set(group.group_id, defaultTokens)
        } else {
          // 检查是否有新的Token需要重新查询余额
          const selectedTokens = getGroupSelectedTokens(group.group_id)
          if (selectedTokens.length > 0 && group.balance_loaded && group.token_totals) {
            // 检查是否有选中的Token不在当前的余额数据中
            const currentTokenAddresses = group.token_totals.map(t => t.contract_address)
            const selectedTokenObjects = availableTokens.value.filter(token => selectedTokens.includes(token.id))
            const missingTokens = selectedTokenObjects.filter(token => 
              !currentTokenAddresses.includes(token.contract_address)
            )
            
            if (missingTokens.length > 0) {
              console.log(`Group ${group.group_name} has ${missingTokens.length} new tokens, need to refresh balance`)
              needsRefresh = true
              // 重置该分组的余额状态，触发重新查询
              group.balance_loaded = false
              group.balance_loading = false
              group.queryProgress = null
            }
          }
        }
      })
      // Note: Removed localStorage save since we now use database
      // saveGroupTokenSelections() - REMOVED

      // 如果有分组需要刷新，延迟后自动刷新
      if (needsRefresh) {
        setTimeout(() => {
          console.log('Auto-refreshing groups with new tokens...')
          groupsData.value.forEach((group, index) => {
            if (!group.balance_loaded && !group.balance_loading && group.address_validation_passed) {
              loadGroupBalance(index)
            }
          })
        }, 1000)
      }
    }
  } catch (error) {
    console.error('Failed to load tokens:', error)
    window.showNotification('error', '加载Token列表失败')
  }
}

const formatContractAddress = () => {
  let address = newToken.value.contractAddress.trim()
  
  // 先检查是否以0x开头，分别处理
  if (address.toLowerCase().startsWith('0x')) {
    // 如果以0x开头，只保留0x和后面的十六进制字符
    const prefix = address.substring(0, 2)
    const hexPart = address.substring(2).replace(/[^0-9a-fA-F]/g, '')
    address = prefix.toLowerCase() + hexPart
  } else {
    // 如果不以0x开头，移除非十六进制字符然后添加前缀
    address = address.replace(/[^0-9a-fA-F]/g, '')
    if (address.length > 0) {
      address = '0x' + address
    }
  }
  
  // 限制长度为42个字符（0x + 40位十六进制）
  if (address.length > 42) {
    address = address.substring(0, 42)
  }
  
  newToken.value.contractAddress = address
}

const addNewToken = async () => {
  if (!canAddToken.value) return

  try {
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain
    
    const response = await fetch(`/api/users/${authStore.userId}/tokens`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({
        chain_id: currentChainId,
        name: newToken.value.name,
        symbol: newToken.value.symbol,
        contract_address: newToken.value.contractAddress,
        decimals: parseInt(newToken.value.decimals)
      })
    })

    if (response.ok) {
      window.showNotification('success', 'Token添加成功')
      // 重置表单
      newToken.value = {
        name: '',
        symbol: '',
        contractAddress: '',
        decimals: 18
      }
      // 重新加载Token列表
      console.log('Before loading tokens after adding new token')
      await loadAvailableTokens()
      console.log('After loading tokens, new token list:', availableTokens.value)
    } else {
      // 尝试获取详细错误信息
      let errorMessage = 'Token添加失败'
      try {
        const errorData = await response.json()
        if (errorData.error) {
          errorMessage = errorData.error
        }
      } catch (e) {
        // 如果无法解析错误响应，使用默认消息
        errorMessage = `Token添加失败 (HTTP ${response.status})`
      }
      window.showNotification('error', errorMessage)
      return
    }
  } catch (error) {
    console.error('Failed to add token:', error)
    let errorMessage = 'Token添加失败'
    if (error.message.includes('fetch')) {
      errorMessage = '网络连接失败，请检查后端服务是否正常运行'
    }
    window.showNotification('error', errorMessage)
  }
}

const toggleTokenStatus = async (token) => {
  try {
    const response = await fetch(`/api/users/${authStore.userId}/tokens/${token.id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({
        chain_id: token.chain_id,
        name: token.name,
        symbol: token.symbol,
        contract_address: token.contract_address,
        decimals: token.decimals,
        is_active: !token.is_active
      })
    })

    if (response.ok) {
      token.is_active = !token.is_active
      window.showNotification('success', `Token${token.is_active ? '启用' : '禁用'}成功`)
      
      // 如果禁用了Token，从选中列表中移除
      if (!token.is_active) {
        selectedTokens.value = selectedTokens.value.filter(id => id !== token.id)
      }
    } else {
      throw new Error('Failed to toggle token status')
    }
  } catch (error) {
    console.error('Failed to toggle token status:', error)
    window.showNotification('error', 'Token状态切换失败')
  }
}

const loadGroupsData = async () => {
  loading.value = true
  try {
    // 只加载分组和地址信息，不查询余额
    await Promise.all([
      walletStore.fetchGroups(authStore.userId),
      walletStore.fetchAddresses(authStore.userId)
    ])
    
    // 初始化分组数据结构（不包含余额信息）
    const groupsWithAddresses = walletStore.groups.filter(group => {
      const groupAddresses = walletStore.getAddressesByGroup(group.id)
      return groupAddresses.length > 0
    })
    
    groupsData.value = groupsWithAddresses.map(group => {
      const addresses = walletStore.getAddressesByGroup(group.id)
      
      // 验证地址格式
      const invalidAddresses = []
      const validAddresses = addresses.filter(addr => {
        const isValid = addr.address && 
                       typeof addr.address === 'string' && 
                       addr.address.length === 42 && 
                       addr.address.startsWith('0x') &&
                       /^0x[a-fA-F0-9]{40}$/.test(addr.address)
        
        if (!isValid) {
          invalidAddresses.push({
            id: addr.id,
            address: addr.address,
            label: addr.label || '',
            error: `无效地址格式 (长度: ${addr.address ? addr.address.length : 0})`
          })
        }
        return isValid
      })
      
      return {
        group_id: group.id,
        group_name: group.name,
        address_count: addresses.length,
        valid_address_count: validAddresses.length,
        invalid_address_count: invalidAddresses.length,
        addresses: validAddresses, // 只保留有效地址用于查询
        invalid_addresses: invalidAddresses, // 保存无效地址用于显示
        // 余额信息暂时为空，等待按需加载
        total_bnb: null,
        token_totals: null,
        balance_loaded: false,
        balance_loading: false,
        queryProgress: null,
        address_validation_passed: invalidAddresses.length === 0,
        // 余额变化追踪
        previous_total_bnb: null,
        previous_token_totals: null,
        balance_changes: null, // 存储余额变化信息
        last_refresh_time: null
      }
    })
    
    // 初始化每个分组的倒计时（默认10秒）
    groupsData.value.forEach(group => {
      if (!groupCountdowns.value.has(group.group_id)) {
        groupCountdowns.value.set(group.group_id, {
          enabled: false,
          remaining: 10, // 10秒
          total: 10
        })
      }
    })

    // Auto-query removed: Users now manually trigger balance queries
    
  } catch (error) {
    window.showNotification('error', '加载分组数据失败')
    console.error('Failed to load groups data:', error)
  } finally {
    loading.value = false
  }
}

// 处理进度更新
const handleProgressUpdate = (progress) => {
  progressInfo.value.show = true
  
  switch (progress.stage) {
    case 'starting':
      progressInfo.value.stage = 'starting'
      progressInfo.value.totalGroups = progress.totalGroups
      progressInfo.value.currentGroup = 0
      progressInfo.value.message = `开始查询 ${progress.totalGroups} 个分组...`
      break
      
    case 'querying_group':
      progressInfo.value.stage = 'querying_group'
      progressInfo.value.currentGroup = progress.currentGroup
      progressInfo.value.totalGroups = progress.totalGroups
      progressInfo.value.groupName = progress.groupName
      progressInfo.value.message = `正在查询分组: ${progress.groupName} (${progress.currentGroup}/${progress.totalGroups})`
      break
      
    case 'querying':
      progressInfo.value.stage = 'querying'
      progressInfo.value.groupName = progress.groupName
      progressInfo.value.addressCount = progress.addressCount
      progressInfo.value.message = `查询 ${progress.groupName} 的 ${progress.addressCount} 个地址...`
      break
      
    case 'completed':
      progressInfo.value.stage = 'completed'
      progressInfo.value.message = `完成查询分组: ${progress.groupName}`
      break
      
    case 'all_completed':
      progressInfo.value.stage = 'all_completed'
      progressInfo.value.message = `全部完成！成功查询 ${progress.successCount} 个分组`
      setTimeout(() => {
        progressInfo.value.show = false
      }, 2000)
      break
  }
}

const refreshData = () => {
  loadGroupsData()
}

// 查询单个分组的BNB余额（仅BNB，不包含Token）
const loadGroupBNBBalance = async (groupIndex) => {
  const group = groupsData.value[groupIndex]
  if (!group || group.balance_loading || group.balance_loaded) {
    return
  }

  // 检查地址验证状态
  if (!group.address_validation_passed) {
    window.showNotification('error', `分组 ${group.group_name} 包含无效地址，无法查询余额`)
    return
  }

  // 检查是否有有效地址
  if (group.valid_address_count === 0) {
    window.showNotification('warning', `分组 ${group.group_name} 没有有效地址`)
    return
  }

  const chainStore = useChainStore()
  const currentChainId = chainStore.currentChain
  
  if (!currentChainId) {
    window.showNotification('warning', '请先选择区块链网络')
    return
  }

  // 设置加载状态
  group.balance_loading = true

  try {
    // 获取该分组选择的RPC ID
    const selectedRpcId = getGroupSelectedRpc(group.group_id)
    
    // 只查询BNB余额，不传递Token ID
    const balanceData = await walletStore.fetchGroupBalance(
      authStore.userId, 
      group.group_id, 
      currentChainId,
      [], // 空的Token ID数组，只查询BNB
      (progress) => {
        // 处理单个分组的进度更新
        handleSingleGroupProgress(group, progress)
      },
      selectedRpcId // 传递选中的RPC ID
    )

    // 更新分组余额数据
    if (balanceData) {
      console.log('=== Processing BNB Balance Data ===')
      console.log('New BNB balance:', balanceData.total_bnb)
      console.log('Has previous data:', !!(group.previous_total_bnb !== null || group.previous_token_totals))
      
      // 计算余额变化
      const changes = calculateBalanceChanges(group, balanceData.total_bnb, balanceData.token_totals || [])
      console.log('Calculated changes:', changes)
      
      // 更新当前余额
      group.total_bnb = balanceData.total_bnb
      group.token_totals = balanceData.token_totals || [] // 空的Token列表
      group.balance_changes = changes
      group.last_refresh_time = new Date().toLocaleTimeString()
      group.balance_loaded = true
      group.balance_loading = false
      
      console.log('Updated group balance changes:', group.balance_changes)
      console.log('=== End Processing BNB Balance Data ===')
    }

  } catch (error) {
    window.showNotification('error', `查询分组 ${group.group_name} BNB余额失败`)
    console.error('Failed to load group BNB balance:', error)
    group.balance_loading = false
  }
}

// 查询单个分组的完整余额（包含选中的Token）
const loadGroupBalance = async (groupIndex) => {
  const group = groupsData.value[groupIndex]
  if (!group || group.balance_loading) {
    return
  }

  // 检查地址验证状态
  if (!group.address_validation_passed) {
    window.showNotification('error', `分组 ${group.group_name} 包含无效地址，无法查询余额`)
    return
  }

  // 检查是否有有效地址
  if (group.valid_address_count === 0) {
    window.showNotification('warning', `分组 ${group.group_name} 没有有效地址`)
    return
  }

  const chainStore = useChainStore()
  const currentChainId = chainStore.currentChain
  
  if (!currentChainId) {
    window.showNotification('warning', '请先选择区块链网络')
    return
  }

  // 设置加载状态
  group.balance_loading = true

  try {
    // 获取该分组选择的Token ID和RPC ID
    const selectedTokenIds = getGroupSelectedTokens(group.group_id)
    const selectedRpcId = getGroupSelectedRpc(group.group_id)
    
    console.log('=== Before fetchGroupBalance ===')
    console.log('Group:', group.group_name, 'ID:', group.group_id)
    console.log('Selected Token IDs:', selectedTokenIds)
    console.log('Available tokens for verification:', availableTokens.value.map(t => ({ id: t.id, symbol: t.symbol })))
    console.log('Selected RPC ID:', selectedRpcId)
    console.log('Current Chain ID:', currentChainId)
    console.log('=== End Before fetchGroupBalance ===')
    
    const balanceData = await walletStore.fetchGroupBalance(
      authStore.userId, 
      group.group_id, 
      currentChainId,
      selectedTokenIds, // 传递该分组选中的Token ID
      (progress) => {
        // 处理单个分组的进度更新
        handleSingleGroupProgress(group, progress)
      },
      selectedRpcId // 传递选中的RPC ID
    )

    // 更新分组余额数据
    if (balanceData) {
      console.log('=== Processing Full Balance Data ===')
      console.log('New BNB balance:', balanceData.total_bnb)
      console.log('New token count:', balanceData.token_totals?.length || 0)
      console.log('Has previous data:', !!(group.previous_total_bnb !== null || group.previous_token_totals))
      
      // 计算余额变化
      const changes = calculateBalanceChanges(group, balanceData.total_bnb, balanceData.token_totals)
      console.log('Calculated changes:', changes)
      
      // 更新当前余额
      group.total_bnb = balanceData.total_bnb
      group.token_totals = balanceData.token_totals
      group.balance_changes = changes
      group.last_refresh_time = new Date().toLocaleTimeString()
      group.balance_loaded = true
      group.balance_loading = false
      
      console.log('Updated group balance changes:', group.balance_changes)
      console.log('=== End Processing Full Balance Data ===')
    }

  } catch (error) {
    window.showNotification('error', `查询分组 ${group.group_name} 余额失败`)
    console.error('Failed to load group balance:', error)
    group.balance_loading = false
  }
}

// 批量查询所有分组BNB余额（自动查询时使用）
const loadAllGroupsBNBBalance = async () => {
  const chainStore = useChainStore()
  const currentChainId = chainStore.currentChain
  
  if (!currentChainId) {
    window.showNotification('warning', '请先选择区块链网络')
    return
  }

  progressInfo.value.show = true
  progressInfo.value.stage = 'starting'
  progressInfo.value.totalGroups = groupsData.value.length
  progressInfo.value.currentGroup = 0
  progressInfo.value.message = `开始查询 ${groupsData.value.length} 个分组的BNB余额...`

  for (let i = 0; i < groupsData.value.length; i++) {
    const group = groupsData.value[i]
    
    if (group.balance_loaded) {
      continue // 跳过已加载的分组
    }

    // 跳过地址验证失败的分组
    if (!group.address_validation_passed) {
      console.warn(`跳过分组 ${group.group_name}: 包含无效地址`)
      continue
    }

    progressInfo.value.currentGroup = i + 1
    progressInfo.value.groupName = group.group_name
    progressInfo.value.message = `正在查询分组: ${group.group_name} 的BNB余额 (${i + 1}/${groupsData.value.length})`

    await loadGroupBNBBalance(i)
  }

  progressInfo.value.stage = 'all_completed'
  progressInfo.value.message = `BNB余额查询完成！`
  setTimeout(() => {
    progressInfo.value.show = false
  }, 2000)
}

// 批量查询所有分组余额（手动点击"查询所有余额"时使用）
const loadAllGroupsBalance = async () => {
  const chainStore = useChainStore()
  const currentChainId = chainStore.currentChain
  
  if (!currentChainId) {
    window.showNotification('warning', '请先选择区块链网络')
    return
  }

  progressInfo.value.show = true
  progressInfo.value.stage = 'starting'
  progressInfo.value.totalGroups = groupsData.value.length
  progressInfo.value.currentGroup = 0
  progressInfo.value.message = `开始查询 ${groupsData.value.length} 个分组的完整余额...`

  for (let i = 0; i < groupsData.value.length; i++) {
    const group = groupsData.value[i]
    
    // 跳过地址验证失败的分组
    if (!group.address_validation_passed) {
      console.warn(`跳过分组 ${group.group_name}: 包含无效地址`)
      continue
    }

    progressInfo.value.currentGroup = i + 1
    progressInfo.value.groupName = group.group_name
    progressInfo.value.message = `正在查询分组: ${group.group_name} (${i + 1}/${groupsData.value.length})`

    // 重置状态，重新查询完整余额
    group.balance_loaded = false
    group.balance_loading = false
    group.queryProgress = null
    
    await loadGroupBalance(i)
  }

  progressInfo.value.stage = 'all_completed'
  progressInfo.value.message = `全部完成！`
  setTimeout(() => {
    progressInfo.value.show = false
  }, 2000)
}

// 处理单个分组查询进度
const handleSingleGroupProgress = (group, progress) => {
  // 更新分组内部的进度信息
  if (!group.queryProgress) {
    group.queryProgress = {
      current: 0,
      total: group.address_count,
      message: ''
    }
  }

  switch (progress.stage) {
    case 'querying':
      group.queryProgress.current = 0
      group.queryProgress.total = progress.addressCount || group.address_count
      group.queryProgress.message = `开始查询 ${progress.addressCount} 个地址...`
      break
    case 'batch_processing':
      group.queryProgress.current = progress.processedAddresses || 0
      group.queryProgress.total = progress.totalAddresses || group.address_count
      group.queryProgress.message = `正在处理第 ${progress.currentBatch}/${progress.totalBatches} 批 (${progress.batchSize} 个地址)...`
      break
    case 'batch_completed':
      group.queryProgress.current = progress.processedAddresses || 0
      group.queryProgress.total = progress.totalAddresses || group.address_count
      group.queryProgress.message = `已完成 ${progress.currentBatch}/${progress.totalBatches} 批查询`
      break
    case 'completed':
      group.queryProgress.current = group.queryProgress.total
      group.queryProgress.message = `查询完成`
      // 清理进度信息
      setTimeout(() => {
        group.queryProgress = null
      }, 1000)
      break
  }
}


const toggleAutoRefresh = () => {
  if (autoRefresh.value.enabled) {
    stopAutoRefresh()
  } else {
    startAutoRefresh()
  }
}

const startAutoRefresh = () => {
  if (groupsData.value.length === 0) {
    window.showNotification('warning', '没有分组数据，无法启动自动刷新')
    return
  }

  autoRefresh.value.enabled = true
  autoRefresh.value.progress.total = groupsData.value.length
  autoRefresh.value.progress.current = 0
  autoRefresh.value.history = []
  
  window.showNotification('success', '自动刷新已启动')
  startRefreshCycle()
}

const stopAutoRefresh = () => {
  autoRefresh.value.enabled = false
  autoRefresh.value.currentGroup = null
  autoRefresh.value.nextRoundIn = 0
  autoRefresh.value.progress = { current: 0, total: 0 }
  
  // 清除所有定时器
  Object.values(autoRefresh.value.timers).forEach(timer => {
    if (timer) clearTimeout(timer)
  })
  autoRefresh.value.timers = {
    groupTimer: null,
    roundTimer: null,
    countdownTimer: null
  }
  
  if (autoRefresh.value.enabled === false) {
    window.showNotification('info', '自动刷新已停止')
  }
}

const startRefreshCycle = () => {
  if (!autoRefresh.value.enabled) return
  
  autoRefresh.value.progress.current = 0
  refreshNextGroup(0)
}

const refreshNextGroup = async (groupIndex) => {
  if (!autoRefresh.value.enabled || groupIndex >= groupsData.value.length) {
    // 完成一轮刷新，等待下一轮
    startRoundCountdown()
    return
  }

  const group = groupsData.value[groupIndex]
  autoRefresh.value.currentGroup = group.group_name
  autoRefresh.value.progress.current = groupIndex + 1

  try {
    // 获取当前选择的链ID
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain
    
    if (!currentChainId) {
      throw new Error('未选择区块链网络')
    }

    // 刷新单个分组
    const refreshedData = await walletStore.fetchGroupBalance(authStore.userId, group.group_id, currentChainId)
    
    // 更新分组数据
    const index = groupsData.value.findIndex(g => g.group_id === group.group_id)
    if (index !== -1 && refreshedData) {
      groupsData.value[index] = refreshedData
    }

    // 记录成功
    autoRefresh.value.history.push({
      groupName: group.group_name,
      success: true,
      time: new Date().toLocaleTimeString()
    })

  } catch (error) {
    // 记录失败
    autoRefresh.value.history.push({
      groupName: group.group_name,
      success: false,
      time: new Date().toLocaleTimeString()
    })
  }

  // 等待组别间隔后刷新下一个分组
  autoRefresh.value.timers.groupTimer = setTimeout(() => {
    refreshNextGroup(groupIndex + 1)
  }, settings.value.groupInterval * 1000)
}

const startRoundCountdown = () => {
  if (!autoRefresh.value.enabled) return
  
  autoRefresh.value.currentGroup = null
  autoRefresh.value.nextRoundIn = settings.value.roundInterval

  const countdown = () => {
    if (!autoRefresh.value.enabled) return
    
    autoRefresh.value.nextRoundIn--
    
    if (autoRefresh.value.nextRoundIn <= 0) {
      startRefreshCycle()
    } else {
      autoRefresh.value.timers.countdownTimer = setTimeout(countdown, 1000)
    }
  }

  autoRefresh.value.timers.countdownTimer = setTimeout(countdown, 1000)
}

const totalStats = computed(() => {
  const stats = {
    activeGroups: groupsData.value.length,
    totalAddresses: 0,
    totalBNB: 0,
    totalRpcNodes: availableRpcEndpoints.value.length,
    totalTokens: availableTokens.value.filter(token => token.is_active).length
  }

  groupsData.value.forEach(group => {
    stats.totalAddresses += group.valid_address_count || 0
    stats.totalBNB += parseFloat(group.total_bnb || 0)
  })

  return {
    activeGroups: stats.activeGroups,
    totalAddresses: stats.totalAddresses,
    totalBNB: stats.totalBNB.toFixed(6),
    totalRpcNodes: stats.totalRpcNodes,
    totalTokens: stats.totalTokens
  }
})

const tokenSummary = computed(() => {
  const tokenMap = new Map()

  groupsData.value.forEach(group => {
    group.token_totals?.forEach(token => {
      const key = token.symbol
      const amount = parseFloat(token.balance || 0)
      
      if (!tokenMap.has(key)) {
        tokenMap.set(key, {
          symbol: token.symbol,
          name: token.name,
          totalAmount: 0,
          maxAmount: 0,
          groupCount: 0,
          groups: new Set()
        })
      }
      
      const tokenData = tokenMap.get(key)
      tokenData.totalAmount += amount
      tokenData.maxAmount = Math.max(tokenData.maxAmount, amount)
      if (!tokenData.groups.has(group.group_id)) {
        tokenData.groups.add(group.group_id)
        tokenData.groupCount++
      }
    })
  })

  return Array.from(tokenMap.values())
    .sort((a, b) => b.totalAmount - a.totalAmount) // 按总余额排序，余额大的在前
})

const getMainTokens = (tokens, groupId = null) => {
  if (!tokens) tokens = []
  
  // 如果提供了groupId，确保显示所有选中的Token，即使没有余额数据
  if (groupId) {
    const selectedTokenIds = getGroupSelectedTokens(groupId)
    const selectedTokenObjects = availableTokens.value.filter(token => 
      selectedTokenIds.includes(token.id) && token.is_active
    )
    
    // 创建一个完整的Token列表，包含余额数据或默认值
    const completeTokens = selectedTokenObjects.map(selectedToken => {
      // 查找是否有对应的余额数据
      const balanceToken = tokens.find(t => 
        t.contract_address === selectedToken.contract_address
      )
      
      return balanceToken || {
        symbol: selectedToken.symbol,
        name: selectedToken.name,
        contract_address: selectedToken.contract_address,
        balance: '0', // 默认余额为0
        decimals: selectedToken.decimals
      }
    })
    
    // 优先显示有余额的代币，然后显示余额为0的代币
    const tokensWithBalance = completeTokens.filter(token => parseFloat(token.balance || 0) > 0)
    const tokensWithoutBalance = completeTokens.filter(token => parseFloat(token.balance || 0) === 0)
    
    // 最多显示5个代币：优先有余额的，然后是余额为0的
    const result = [...tokensWithBalance, ...tokensWithoutBalance].slice(0, 5)
    return result
  }
  
  // 原有逻辑，用于其他地方
  const tokensWithBalance = tokens.filter(token => parseFloat(token.balance || 0) > 0)
  const tokensWithoutBalance = tokens.filter(token => parseFloat(token.balance || 0) === 0)
  
  // 最多显示5个代币：优先有余额的，然后是余额为0的
  const result = [...tokensWithBalance, ...tokensWithoutBalance].slice(0, 5)
  return result
}

// 获取Token的变化信息
const getTokenChange = (group, token) => {
  if (!group.balance_changes || !group.balance_changes.token_changes) {
    return null
  }
  
  return group.balance_changes.token_changes.find(change => 
    change.contract_address === token.contract_address || 
    change.symbol === token.symbol
  ) || null
}

const formatBalance = (balance) => {
  const num = parseFloat(balance)
  if (num === 0) return '0'
  if (num < 0.000001) return '<0.000001'
  if (num >= 1000000) return (num / 1000000).toFixed(2) + 'M'
  if (num >= 1000) return (num / 1000).toFixed(2) + 'K'
  return num.toFixed(6)
}

const viewGroupDetail = (groupId) => {
  router.push(`/balance/${groupId}`)
}

// Drag and Drop handlers
const onDragStart = (evt) => {
  console.log('Drag started:', evt)
}

const onDragEnd = async (evt) => {
  console.log('Drag ended:', evt)
  
  if (evt.oldIndex === evt.newIndex) {
    // No position change, no need to update
    return
  }

  sortingInProgress.value = true

  try {
    // Create the reorder request payload
    const groupOrders = groupsData.value.map((group, index) => ({
      group_id: group.group_id,
      sort_order: index + 1 // 1-based ordering
    }))

    console.log('Reordering groups:', groupOrders)

    // Call API to update sort order
    const response = await fetch(`/api/users/${authStore.userId}/groups/reorder`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({
        group_orders: groupOrders
      })
    })

    if (response.ok) {
      window.showNotification('success', '分组排序已保存')
      console.log('Group order updated successfully')
    } else {
      throw new Error('Failed to update group order')
    }
  } catch (error) {
    console.error('Failed to update group order:', error)
    window.showNotification('error', '排序保存失败')
    
    // Revert the order on error
    await loadGroupsData()
  } finally {
    sortingInProgress.value = false
  }
}

// Create group function
const createGroup = async () => {
  try {
    await walletStore.createGroup(authStore.userId, newGroup.value)
    showCreateGroupModal.value = false
    newGroup.value = { name: '', description: '' }
    window.showNotification('success', '分组创建成功')
    await refreshData()
  } catch (error) {
    window.showNotification('error', '创建失败')
  }
}

// Add address function
const addAddress = async () => {
  try {
    // Parse addresses, support batch add
    const addresses = newAddress.value.address
      .split(',')
      .map(addr => addr.trim())
      .filter(addr => addr.length > 0)

    if (addresses.length === 0) {
      window.showNotification('error', '请输入有效的地址')
      return
    }

    // Validate address format
    const invalidAddresses = addresses.filter(addr => !addr.match(/^0x[a-fA-F0-9]{40}$/))
    if (invalidAddresses.length > 0) {
      window.showNotification('error', `无效的地址格式: ${invalidAddresses.join(', ')}`)
      return
    }

    let successCount = 0
    let failedAddresses = []

    // Batch add addresses
    for (const address of addresses) {
      try {
        const addressData = {
          address: address,
          label: newAddress.value.label || undefined,
          group_id: newAddress.value.group_id || undefined
        }

        await walletStore.addAddress(authStore.userId, addressData)
        successCount++
      } catch (error) {
        failedAddresses.push(address)
      }
    }

    // Show result
    if (successCount > 0) {
      window.showNotification('success', `成功添加 ${successCount} 个地址`)
    }

    if (failedAddresses.length > 0) {
      window.showNotification('error', `${failedAddresses.length} 个地址添加失败`)
    }

    // Clear form and close modal
    showAddAddressModal.value = false
    newAddress.value = { address: '', label: '', group_id: '' }
    await refreshData()

  } catch (error) {
    window.showNotification('error', '添加失败')
  }
}

// Helper function to truncate address
const truncateAddress = (address) => {
  if (!address) return ''
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

// ====== Group Settings Persistence Functions ======

// Load settings from database
const loadGroupSettingsFromDB = async () => {
  try {
    const chainStore = useChainStore()
    const currentChainId = chainStore.currentChain

    if (!currentChainId) {
      console.warn('Cannot load group settings: no chain selected')
      return
    }

    const response = await groupSettingsAPI.getAllSettings(authStore.userId, currentChainId)
    const settingsList = response.data

    console.log('=== Loading Group Settings from DB ===')
    console.log('Chain ID:', currentChainId)
    console.log('Raw response:', response)
    console.log('Settings list:', settingsList)
    console.log('Settings count:', settingsList.length)

    // Note: Don't clear existing selections - preserve any defaults set by loadAvailableRpcEndpoints

    // Apply loaded settings to each group
    settingsList.forEach(settings => {
      const groupId = settings.group_id

      console.log(`\n=== Processing settings for group ${groupId} ===`)
      console.log('Full settings object:', settings)
      console.log('selected_rpc_id value:', settings.selected_rpc_id, 'type:', typeof settings.selected_rpc_id)
      console.log('selected_token_ids value:', settings.selected_token_ids, 'type:', typeof settings.selected_token_ids, 'isArray:', Array.isArray(settings.selected_token_ids))

      // Apply countdown settings (always load duration, even if disabled)
      groupCountdowns.value.set(groupId, {
        enabled: settings.countdown_enabled || false,
        total: settings.countdown_duration || 600,
        remaining: settings.countdown_duration || 600
      })

      // Apply RPC selection
      console.log(`\nChecking RPC selection for group ${groupId}:`)
      console.log('  - selected_rpc_id:', settings.selected_rpc_id)
      console.log('  - Condition check (settings.selected_rpc_id):', !!settings.selected_rpc_id)

      if (settings.selected_rpc_id) {
        console.log(`  ✓ Setting RPC ${settings.selected_rpc_id} for group ${groupId}`)
        groupRpcSelections.value.set(groupId, settings.selected_rpc_id)
        console.log(`  ✓ RPC set successfully. Verifying: ${groupRpcSelections.value.get(groupId)}`)
      } else {
        console.log(`  ✗ No RPC set for group ${groupId} (selected_rpc_id is falsy)`)
      }

      // Apply token selection
      console.log(`\nChecking Token selection for group ${groupId}:`)
      console.log('  - selected_token_ids:', settings.selected_token_ids)
      console.log('  - Is array:', Array.isArray(settings.selected_token_ids))
      console.log('  - Length:', settings.selected_token_ids?.length)

      if (settings.selected_token_ids && settings.selected_token_ids.length > 0) {
        console.log(`  ✓ Setting ${settings.selected_token_ids.length} tokens for group ${groupId}:`, settings.selected_token_ids)
        groupTokenSelections.value.set(groupId, settings.selected_token_ids)
        console.log(`  ✓ Tokens set successfully. Verifying:`, groupTokenSelections.value.get(groupId))
      } else {
        console.log(`  ✗ No tokens to set for group ${groupId}`)
      }
    })

    // Trigger reactivity for all Maps
    console.log('\n=== Triggering Map Reactivity ===')
    groupCountdowns.value = new Map(groupCountdowns.value)
    groupRpcSelections.value = new Map(groupRpcSelections.value)
    groupTokenSelections.value = new Map(groupTokenSelections.value)

    console.log('\n=== Final Verification ===')
    console.log('✓ Settings applied to groups')
    console.log('✓ RPC selections Map size:', groupRpcSelections.value.size)
    console.log('✓ RPC selections content:', Object.fromEntries(groupRpcSelections.value))
    console.log('✓ Token selections Map size:', groupTokenSelections.value.size)
    console.log('✓ Token selections content:', Object.fromEntries(groupTokenSelections.value))

    // Start countdowns for groups that have countdown enabled
    console.log('\n=== Starting Enabled Countdowns ===')
    groupCountdowns.value.forEach((countdown, groupId) => {
      if (countdown.enabled) {
        console.log(`✓ Starting countdown for group ${groupId}: ${countdown.remaining}s`)
        startGroupCountdown(groupId, countdown.remaining)
      }
    })
  } catch (error) {
    console.error('Failed to load group settings:', error)
  }
}

// Save settings to database (debounced)
let saveSettingsTimers = new Map()

const saveGroupSettings = async (groupId) => {
  // Clear existing timer for this group
  if (saveSettingsTimers.has(groupId)) {
    clearTimeout(saveSettingsTimers.get(groupId))
  }

  // Debounce: save after 1 second of no changes
  const timer = setTimeout(async () => {
    try {
      const chainStore = useChainStore()
      const currentChainId = chainStore.currentChain

      if (!currentChainId) {
        console.warn(`Cannot save settings for group ${groupId}: no chain selected`)
        return
      }

      const countdown = groupCountdowns.value.get(groupId)
      const selectedRpcId = groupRpcSelections.value.get(groupId)
      const selectedTokenIds = groupTokenSelections.value.get(groupId) || []

      console.log(`=== Saving settings for group ${groupId} ===`)
      console.log('Chain ID:', currentChainId)
      console.log('countdown:', countdown)
      console.log('selectedRpcId:', selectedRpcId, 'type:', typeof selectedRpcId)
      console.log('selectedTokenIds:', selectedTokenIds)

      const settingsData = {
        chain_id: currentChainId, // Add chain_id to request body
        countdown_enabled: countdown?.enabled || false,
        countdown_duration: parseInt(countdown?.total, 10) || 600,
        selected_rpc_id: selectedRpcId !== undefined ? selectedRpcId : null,
        selected_token_ids: selectedTokenIds
      }

      console.log('settingsData to save:', settingsData)

      await groupSettingsAPI.updateSettings(authStore.userId, groupId, settingsData)
      console.log(`Settings saved successfully for group ${groupId}`)
    } catch (error) {
      console.error(`Failed to save settings for group ${groupId}:`, error)
    } finally {
      saveSettingsTimers.delete(groupId)
    }
  }, 1000)

  saveSettingsTimers.set(groupId, timer)
}

// Auto-save is now handled in individual setter functions instead of global watch

</script>

<style scoped>
/* Drag and drop styles */
.ghost {
  opacity: 0.5;
}

.chosen {
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
}

.sortable-fallback {
  display: none;
}

.sortable-ghost {
  opacity: 0.4;
}

.sortable-chosen {
  transform: scale(1.05);
  transition: all 0.2s ease;
}
</style>
