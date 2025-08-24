import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import { useChainStore } from './stores/chain'
import './assets/style.css'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

// 初始化链store
const chainStore = useChainStore()
chainStore.loadCurrentChain()

// 全局通知函数
window.showNotification = (type, message) => {
  const event = new CustomEvent('show-notification', {
    detail: { type, message }
  })
  window.dispatchEvent(event)
}

app.mount('#app')
