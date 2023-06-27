import {RouteRecordRaw} from "vue-router";

export const DialogRecord: RouteRecordRaw = {
  path: '/dialog_record',
  component: () => import('@/views/layout/MainLayout.vue'),
  redirect: '/dialog_record/list',
  meta: {icon: 'ChatDotSquare', title: "对话记录", alwaysShow: true},
  children: [
    {
      path: 'list',
      component: () => import('@/views/dialogRecord/Index.vue'),
      meta: {icon: 'ChatDotSquare', title: "对话记录列表", hideMenu: true}
    }
  ]
}
