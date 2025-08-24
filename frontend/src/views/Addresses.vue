<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex justify-between items-center">
        <h1 class="text-3xl font-bold text-gray-900">地址管理</h1>
        <button @click="showAddModal = true" class="btn-primary">
          添加新地址
        </button>
      </div>
    </div>

    <!-- 地址列表 -->
    <div class="card">
      <div class="card-body p-0">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  地址
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  标签
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  分组
                </th>
                <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  添加时间
                </th>
                <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                  操作
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="address in walletStore.addresses" :key="address.id">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-mono text-gray-900">
                    {{ truncateAddress(address.address) }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">
                    {{ address.label || '-' }}
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span
                    v-if="address.group"
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800"
                  >
                    {{ address.group.name }}
                  </span>
                  <span v-else class="text-sm text-gray-500">未分组</span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ new Date(address.created_at).toLocaleDateString() }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                  <button
                    @click="deleteAddress(address.id)"
                    class="text-red-600 hover:text-red-900"
                  >
                    删除
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- 添加地址模态框 -->
    <div v-if="showAddModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showAddModal = false"></div>
        <div class="bg-white rounded-lg p-6 max-w-md w-full relative">
          <h3 class="text-lg font-medium text-gray-900 mb-4">添加新地址</h3>
          <form @submit.prevent="addAddress">
            <div class="space-y-4">
              <div>
                <label class="block text-sm font-medium text-gray-700">钱包地址</label>
                <textarea
                  v-model="newAddress.address"
                  class="input mt-1 min-h-[120px] resize-y"
                  placeholder="支持批量添加，多个地址请用逗号分割
0x1234...,0x5678...,0x9abc..."
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
              <button type="button" @click="showAddModal = false" class="btn-secondary">取消</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'
import { useWalletStore } from '@/stores/wallet'

const authStore = useAuthStore()
const walletStore = useWalletStore()
const showAddModal = ref(false)

const newAddress = reactive({
  address: '',
  label: '',
  group_id: ''
})

onMounted(async () => {
  await Promise.all([
    walletStore.fetchAddresses(authStore.userId),
    walletStore.fetchGroups(authStore.userId)
  ])
})

const truncateAddress = (address) => {
  return `${address.slice(0, 6)}...${address.slice(-4)}`
}

const addAddress = async () => {
  try {
    // 解析地址，支持批量添加
    const addresses = newAddress.address
      .split(',')
      .map(addr => addr.trim())
      .filter(addr => addr.length > 0)
    
    if (addresses.length === 0) {
      window.showNotification('error', '请输入有效的地址')
      return
    }
    
    // 验证地址格式
    const invalidAddresses = addresses.filter(addr => !addr.match(/^0x[a-fA-F0-9]{40}$/))
    if (invalidAddresses.length > 0) {
      window.showNotification('error', `无效的地址格式: ${invalidAddresses.join(', ')}`)
      return
    }
    
    let successCount = 0
    let failedAddresses = []
    
    // 批量添加地址
    for (const address of addresses) {
      try {
        const addressData = {
          address: address,
          label: newAddress.label || undefined,
          group_id: newAddress.group_id || undefined
        }
        
        await walletStore.addAddress(authStore.userId, addressData)
        successCount++
      } catch (error) {
        failedAddresses.push(address)
      }
    }
    
    // 显示结果
    if (successCount > 0) {
      window.showNotification('success', `成功添加 ${successCount} 个地址`)
    }
    
    if (failedAddresses.length > 0) {
      window.showNotification('error', `${failedAddresses.length} 个地址添加失败`)
    }
    
    // 清空表单并关闭模态框
    showAddModal.value = false
    newAddress.address = ''
    newAddress.label = ''
    newAddress.group_id = ''
    
  } catch (error) {
    window.showNotification('error', '添加失败')
  }
}

const deleteAddress = async (addressId) => {
  if (confirm('确定要删除这个地址吗？')) {
    try {
      await walletStore.deleteAddress(authStore.userId, addressId)
      window.showNotification('success', '地址删除成功')
    } catch (error) {
      window.showNotification('error', '删除失败')
    }
  }
}
</script>
