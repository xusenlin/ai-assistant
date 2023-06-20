import {RouteRecordRaw} from "vue-router";

export const Users: RouteRecordRaw = {
  path: '/users',
  component: () => import('@/views/layout/MainLayout.vue'),
  redirect: '/users/list',
  meta: {icon: 'User', title: "用户管理", alwaysShow: true},
  children: [
    {
      path: 'list',
      component: () => import('@/views/users/Index.vue'),
      meta: {icon: 'User', title: "用户列表", hideMenu: true}
    }
  ]
}
