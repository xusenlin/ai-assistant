<script setup lang="ts">
import {ref} from 'vue'
import { batchAddUser } from "@/api/user.ts"
import {ElMessageBox} from "element-plus";

const emits = defineEmits(["success"])
const dialogVisible = ref(false)
const form = ref({
  Username: "",
  Password: "",
  Status: 1,
  IsAdmin: false,
  RemainingDialogueCount: 0,
})

const Show = () => {
  dialogVisible.value = true
  form.value = {
    Username: "",
    Password: "",
    Status: 1,
    IsAdmin: false,
    RemainingDialogueCount: 0,
  }
}
const submit = () => {
  batchAddUser(form.value).then(c=>{

    let err = Object.keys(c.error)
    let success = c.success ||[]
    let html = err.length===0 ? `<p>成功导入${ success.length }个用户</p>` :
        `<p>成功导入${ success.length }个用户</p>
        <p>导入失败${ err.length }个用户：</p>
        <div style="display: flex;flex-wrap: wrap;">${ err.map(r=>(`<span>${r}</span>`)) }</div>
        `

    ElMessageBox.alert(
        html,
        '提示',
        {
          dangerouslyUseHTMLString: true,
        }
    )
    dialogVisible.value = false
    emits('success')
  })
}

defineExpose({Show})

</script>

<template>
  <el-dialog
      v-model="dialogVisible"
      title="添加用户"
      width="70%"
      draggable
      lock-scroll
      :close-on-click-modal="false"
  >
    <el-form :model="form" label-width="120px" label-position="top">
      <el-form-item label="用户名">
        <el-input v-model="form.Username" type="textarea" placeholder="用户名是唯一的，多个用户名请用英文逗号隔开,当有多个用户时，大家公用下面的属性" />
      </el-form-item>
      <el-form-item label="状态">
        <el-radio-group v-model="form.Status">
          <el-radio :label="0" >禁用</el-radio>
          <el-radio :label="1" >启用</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="是否是管理员">
        <el-radio-group v-model="form.IsAdmin">
          <el-radio :label="false" >否</el-radio>
          <el-radio :label="true" >是</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="剩余对话次数">
        <el-input-number v-model="form.RemainingDialogueCount"/>
      </el-form-item>
      <el-form-item label="密码">
        <el-alert  title="密码为空时，此用户登录时需要激活(设置密码)才能登录系统" type="warning" show-icon :closable="false" />

        <el-input style="margin-top: 20px" v-model="form.Password" type="password" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">关闭</el-button>
        <el-button type="primary" @click="submit()">
          确认添加
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped>

</style>
