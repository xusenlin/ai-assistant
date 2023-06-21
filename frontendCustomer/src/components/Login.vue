<script setup>
import { baseURL } from "../config/req.js"
import {useAppStore} from "../stores/app.js";
import {useUserStore} from "../stores/user.js";

import {ref} from "vue";
import {ElNotification} from "element-plus";

const app = useAppStore()
const user = useUserStore()

const form = ref({
  Username:"",
  Password:""
})

const submit = async () => {
  try {
    const response = await fetch(baseURL + "/v1/login", {
      method: 'post',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(form.value),
    })
    if (!response.ok) {
      throw new Error("请求出错")
    }
    const result = await response.json()
    if(!result.Status){
      throw new Error(result.Msg)
    }
    user.updateUserInfo(result.Data)
    app.showLogin = false
    ElNotification.success({
      title: "登录成功",
      message: result.Data.Username + "欢迎你"
    })
  } catch (e) {
    ElNotification.error({
      title: "错误",
      message: e.message
    })
  }
}

</script>

<template>
  <el-dialog
      v-model="app.showLogin"
      title="登录"
      width="360"
      lock-scroll
  >
    <el-form :model="form" label-position="top">
      <el-form-item label="用户名">
        <el-input v-model="form.Username" type="text" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.Password" type="password" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="app.showLogin = false">关闭</el-button>
        <el-button type="primary" @click="submit()">
          登录
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style lang="scss">

</style>
