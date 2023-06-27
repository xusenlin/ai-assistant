import {Common} from "./modules/common"
import {Dashboard} from "./modules/dashboard.ts"
import {Users} from "./modules/users.ts"
import {SensitiveWords} from "./modules/sensitiveWords.ts"
import { DialogRecord } from "./modules/dialogRecord.ts"
import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

//这里可以根据权限做动态路由
const menuRoute: RouteRecordRaw[] = [
  Dashboard,
  // System,
  Users,
  SensitiveWords,
  DialogRecord
]


const whiteList: string[] = ['/login','/register']//不需要登录也能查看的路由,最少需要'/login'，要不然会一直重定向到login

export default createRouter({
  history: createWebHashHistory(),
  routes:[...menuRoute, ...Common]
})

export {menuRoute,whiteList}
