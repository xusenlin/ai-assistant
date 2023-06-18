import {RouteRecordRaw} from "vue-router";

export const Common: RouteRecordRaw[] = [
  {
    path: '/login',
    component: () => import('@/views/common/Login.vue'),
    meta: {title: '登录', hideTabs: true}
  },
  {
    path: '/register',
    component: () => import('@/views/common/Register.vue'),
    meta: {title: '注册', hideTabs: true}
  },
  {
    // 找不到路由重定向到404页面
    path: "/:pathMatch(.*)",
    component: () => import('@/views/common/404.vue'),
  },
]
