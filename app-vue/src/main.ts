import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'normalize.css'
import { createPinia } from 'pinia'

const pinia = createPinia()

createApp(App)
  .use(pinia)
  .use(router)
  .mount('#app')
