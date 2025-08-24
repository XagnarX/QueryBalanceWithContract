<template>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <div class="px-4 py-6 sm:px-0">
      <div class="flex justify-between items-center">
        <h1 class="text-3xl font-bold text-gray-900">分组管理</h1>
        <button @click="showCreateModal = true" class="btn-primary">
          创建新分组
        </button>
      </div>
    </div>

    <!-- 分组列表 -->
    <div class="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3">
      <div
        v-for="group in walletStore.groups"
        :key="group.id"
        class="card hover:shadow-lg transition-shadow duration-200"
      >
        <div class="card-body">
          <h3 class="text-lg font-medium text-gray-900">{{ group.name }}</h3>
          <p class="text-sm text-gray-500 mt-1">{{ group.description || '无描述' }}</p>
          <p class="text-xs text-gray-400 mt-2">
            创建时间: {{ new Date(group.created_at).toLocaleDateString() }}
          </p>
          <div class="mt-4 flex space-x-2">
            <router-link
              :to="`/balance/${group.id}`"
              class="text-primary-600 hover:text-primary-500 text-sm font-medium"
            >
              查看余额
            </router-link>
            <button
              @click="deleteGroup(group.id)"
              class="text-red-600 hover:text-red-500 text-sm font-medium"
            >
              删除
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 创建分组模态框 -->
    <div v-if="showCreateModal" class="fixed inset-0 z-50 overflow-y-auto">
      <div class="flex items-center justify-center min-h-screen px-4">
        <div class="fixed inset-0 bg-gray-500 bg-opacity-75" @click="showCreateModal = false"></div>
        <div class="bg-white rounded-lg p-6 max-w-md w-full relative">
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
              <button type="button" @click="showCreateModal = false" class="btn-secondary">取消</button>
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
const showCreateModal = ref(false)

const newGroup = reactive({
  name: '',
  description: ''
})

onMounted(async () => {
  await walletStore.fetchGroups(authStore.userId)
})

const createGroup = async () => {
  try {
    await walletStore.createGroup(authStore.userId, newGroup)
    showCreateModal.value = false
    newGroup.name = ''
    newGroup.description = ''
    window.showNotification('success', '分组创建成功')
  } catch (error) {
    window.showNotification('error', '创建失败')
  }
}

const deleteGroup = async (groupId) => {
  if (confirm('确定要删除这个分组吗？')) {
    try {
      await walletStore.deleteGroup(authStore.userId, groupId)
      window.showNotification('success', '分组删除成功')
    } catch (error) {
      window.showNotification('error', '删除失败')
    }
  }
}
</script>
