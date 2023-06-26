import { createApp } from 'vue'
import {createPinia} from "pinia"
import './assets/style.scss'
import "element-plus/theme-chalk/el-notification.css"
import "element-plus/theme-chalk/el-loading.css"
import "element-plus/theme-chalk/el-message.css"
import "element-plus/theme-chalk/el-message-box.css"

import App from './App.vue'
const Pinia = createPinia()

createApp(App).use(Pinia).mount('#app')
