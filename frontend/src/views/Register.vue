<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          注册新账号
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          已有账号？
          <router-link to="/login" class="font-medium text-primary-600 hover:text-primary-500">
            立即登录
          </router-link>
        </p>
      </div>
      
      <form class="mt-8 space-y-6" @submit.prevent="handleRegister">
        <div class="space-y-4">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">用户名</label>
            <input
              id="username"
              v-model="form.username"
              name="username"
              type="text"
              required
              class="input mt-1"
              placeholder="请输入用户名"
            />
          </div>
          
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">邮箱</label>
            <input
              id="email"
              v-model="form.email"
              name="email"
              type="email"
              required
              class="input mt-1"
              placeholder="请输入邮箱地址"
            />
          </div>
          
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">密码</label>
            <input
              id="password"
              v-model="form.password"
              name="password"
              type="password"
              required
              class="input mt-1"
              placeholder="请输入密码（至少6位）"
            />
          </div>
          
          <div>
            <label for="confirmPassword" class="block text-sm font-medium text-gray-700">确认密码</label>
            <input
              id="confirmPassword"
              v-model="form.confirmPassword"
              name="confirmPassword"
              type="password"
              required
              class="input mt-1"
              placeholder="请再次输入密码"
            />
          </div>
        </div>

        <div v-if="authStore.error" class="rounded-md bg-red-50 p-4">
          <div class="text-sm text-red-700">
            {{ authStore.error }}
          </div>
        </div>

        <div v-if="validationError" class="rounded-md bg-red-50 p-4">
          <div class="text-sm text-red-700">
            {{ validationError }}
          </div>
        </div>

        <div>
          <button
            type="submit"
            :disabled="authStore.loading || !isFormValid"
            class="btn-primary w-full"
          >
            <span v-if="authStore.loading" class="mr-2">
              <svg class="animate-spin h-4 w-4" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            </span>
            {{ authStore.loading ? '注册中...' : '注册' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

onMounted(() => {
  authStore.clearError()
})

const validationError = computed(() => {
  if (form.password && form.password.length < 6) {
    return '密码长度至少为6位'
  }
  if (form.confirmPassword && form.password !== form.confirmPassword) {
    return '两次输入的密码不一致'
  }
  return null
})

const isFormValid = computed(() => {
  return form.username && 
         form.email && 
         form.password && 
         form.confirmPassword &&
         form.password === form.confirmPassword &&
         form.password.length >= 6 &&
         !validationError.value
})

const handleRegister = async () => {
  if (!isFormValid.value) {
    return
  }

  try {
    await authStore.register({
      username: form.username,
      email: form.email,
      password: form.password
    })
    
    window.showNotification('success', '注册成功！请登录')
    router.push('/login')
  } catch (error) {
    console.error('Registration failed:', error)
  }
}
</script>
