import axios from 'axios'

// Determine API base URL
// Always use relative path '/api' - Vite proxy will handle it
// The proxy works for both localhost and LAN access
const getApiBaseURL = () => {
  const baseURL = '/api'
  console.log(`📡 API Base URL: ${baseURL} (via Vite proxy)`)
  return baseURL
}

// 创建axios实例
const api = axios.create({
  baseURL: getApiBaseURL(),
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器 - 自动添加认证token
api.interceptors.request.use(
  config => {
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// 响应拦截器 - 处理认证错误
api.interceptors.response.use(
  response => {
    return response
  },
  error => {
    if (error.response?.status === 401) {
      // Token过期或无效，清除本地存储并跳转到登录页
      localStorage.removeItem('auth_token')
      localStorage.removeItem('user_info')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

// 用户相关API
export const userAPI = {
  // 用户注册
  register(userData) {
    return api.post('/users', userData)
  },
  
  // 用户登录
  login(credentials) {
    return api.post('/users/login', credentials)
  },
  
  // 获取用户信息
  getProfile(userId) {
    return api.get(`/users/${userId}`)
  }
}

// 钱包分组相关API
export const groupAPI = {
  // 获取用户所有分组
  getGroups(userId) {
    return api.get(`/users/${userId}/groups`)
  },
  
  // 创建分组
  createGroup(userId, groupData) {
    return api.post(`/users/${userId}/groups`, groupData)
  },
  
  // 删除分组
  deleteGroup(userId, groupId) {
    return api.delete(`/users/${userId}/groups/${groupId}`)
  },
  
  // 获取分组中的地址
  getGroupAddresses(userId, groupId) {
    return api.get(`/users/${userId}/groups/${groupId}/addresses`)
  }
}

// 钱包地址相关API
export const addressAPI = {
  // 获取用户所有地址
  getAddresses(userId) {
    return api.get(`/users/${userId}/addresses`)
  },
  
  // 添加地址
  addAddress(userId, addressData) {
    return api.post(`/users/${userId}/addresses`, addressData)
  },
  
  // 删除地址
  deleteAddress(userId, addressId) {
    return api.delete(`/users/${userId}/addresses/${addressId}`)
  }
}

// 余额查询相关API
export const balanceAPI = {
  // 查询单个地址余额（公共接口）
  getAddressBalance(address, tokens = []) {
    const params = new URLSearchParams({ address })
    if (tokens.length > 0) {
      params.append('tokens', tokens.join(','))
    }
    return api.get(`/balance/address?${params}`)
  },
  
  // 查询用户总余额
  getUserBalance(userId) {
    return api.get(`/users/${userId}/balance`)
  },
  
  // 查询分组余额
  getGroupBalance(userId, groupId) {
    return api.get(`/users/${userId}/groups/${groupId}/balance`)
  },
  
  // 批量查询分组余额
  getMultipleGroupsBalance(userId, groupIds) {
    const params = new URLSearchParams({ group_ids: groupIds.join(',') })
    return api.get(`/users/${userId}/groups/balance?${params}`)
  }
}

// 代币相关API
export const tokenAPI = {
  // 获取支持的代币列表
  getTokens() {
    return api.get('/tokens')
  },

  // 用户添加代币
  addToken(userId, tokenData) {
    return api.post(`/users/${userId}/tokens`, tokenData)
  },

  // 用户更新代币
  updateToken(userId, tokenId, tokenData) {
    return api.put(`/users/${userId}/tokens/${tokenId}`, tokenData)
  }
}

// 分组配置相关API
export const groupSettingsAPI = {
  // 获取所有分组配置 (需要 chain_id)
  getAllSettings(userId, chainId) {
    return api.get(`/users/${userId}/groups/settings?chain_id=${chainId}`)
  },

  // 获取单个分组配置 (需要 chain_id)
  getSettings(userId, groupId, chainId) {
    return api.get(`/users/${userId}/groups/${groupId}/settings?chain_id=${chainId}`)
  },

  // 更新分组配置 (settingsData 必须包含 chain_id)
  updateSettings(userId, groupId, settingsData) {
    return api.put(`/users/${userId}/groups/${groupId}/settings`, settingsData)
  },

  // 删除分组配置 (需要 chain_id)
  deleteSettings(userId, groupId, chainId) {
    return api.delete(`/users/${userId}/groups/${groupId}/settings?chain_id=${chainId}`)
  }
}

export default api
