import {RouteRecordRaw} from "vue-router";

export const SensitiveWords: RouteRecordRaw = {
  path: '/sensitive_words',
  component: () => import('@/views/layout/MainLayout.vue'),
  redirect: '/sensitive_words/list',
  meta: {icon: 'Guide', title: "敏感词管理", alwaysShow: true},
  children: [
    {
      path: 'list',
      component: () => import('@/views/sensitiveWord/Index.vue'),
      meta: {icon: 'Guide', title: "敏感词管理", hideMenu: true}
    }
  ]
}
