import { defineStore } from 'pinia'
import { userAPI } from '@/services/api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: localStorage.getItem('auth_token'),
    isAuthenticated: false,
    loading: false,
    error: null
  }),

  getters: {
    isLoggedIn: (state) => !!state.token && !!state.user,
    userId: (state) => state.user?.id
  },

  actions: {
    // 初始化认证状态
    async initAuth() {
      if (this.token) {
        const userInfo = localStorage.getItem('user_info')
        if (userInfo) {
          this.user = JSON.parse(userInfo)
          this.isAuthenticated = true
        }
      }
    },

    // 用户登录
    async login(credentials) {
      this.loading = true
      this.error = null
      
      try {
        const response = await userAPI.login(credentials)
        const { token, user_id, username } = response.data
        
        this.token = token
        this.user = { id: user_id, username }
        this.isAuthenticated = true
        
        // 保存到本地存储
        localStorage.setItem('auth_token', token)
        localStorage.setItem('user_info', JSON.stringify(this.user))
        
        return response.data
      } catch (error) {
        this.error = error.response?.data?.error || '登录失败'
        throw error
      } finally {
        this.loading = false
      }
    },

    // 用户注册
    async register(userData) {
      this.loading = true
      this.error = null
      
      try {
        const response = await userAPI.register(userData)
        return response.data
      } catch (error) {
        this.error = error.response?.data?.error || '注册失败'
        throw error
      } finally {
        this.loading = false
      }
    },

    // 用户登出
    logout() {
      this.user = null
      this.token = null
      this.isAuthenticated = false
      
      // 清除本地存储
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user_info')
    },

    // 清除错误
    clearError() {
      this.error = null
    }
  }
})
