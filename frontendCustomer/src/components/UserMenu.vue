<script setup>
import {useAppStore} from "../stores/app.js";
import {getUserToken, useUserStore} from "../stores/user.js";
import { ElMessageBox,ElMessage } from "element-plus"
import {baseURL} from "../config/req.js";

const app = useAppStore()
const user = useUserStore()


const editPwd = ()=>{
  ElMessageBox.prompt('请输入新密码，必须大于等于6位', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '关闭',
  }).then(({value}) => {
    if (value.length < 6) {
      ElMessage.warning("输入的新密码少于6位")
      return
    }
    fetch(baseURL + "/api/user/updatePwd", {
      method: 'post',
      headers: {'Content-Type': 'application/json','Authorization':getUserToken()},
      body: JSON.stringify({password:value}),
    }).then(r=>r.json()).then(response=>{
      console.log(response)
      if (!response.Status) {
        ElMessage.warning(response.Msg)
        return
      }
      ElMessage.success("修改成功，请重新登录")
      setTimeout(()=>{
        user.loginOut()
      },1000)
    })

  }).catch(() => {})
}

</script>

<template>
  <div class="login-info">
    <template v-if="user.info">
      <el-dropdown style="width: 100%">
        <div class="item">
          {{ user.userName }}
          <div style="font-size: 12px">
            剩余对话次数：{{user.remainingDialogueCount}}
          </div>
        </div>
        <template #dropdown>
          <div style="padding: 20px;width: 240px">
            <p>
              用户名：{{ user.userName }}
            </p>
            <br>
            <p>剩余对话次数：{{ user.remainingDialogueCount }}</p>
            <br>
            <div>
              <el-button round style="width: 100%"  type="primary" @click="editPwd()">
                修改密码
              </el-button>
            </div>
            <br>
            <div>
              <el-button round style="width: 100%"  type="warning" @click="user.loginOut()">
                退出登录
              </el-button>
            </div>
          </div>

        </template>
      </el-dropdown>
    </template>
    <el-button round v-else @click="app.showLogin = true" style="width: 100%" type="primary">
      登录
    </el-button>
  </div>
</template>

<style lang="scss" scoped>
  .item{
    color: #fff;
    cursor: pointer;
    height: 44px;
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-radius: 4px;
    padding: 0 8px;
    &:hover{
      background: #343541;
    }
  }
</style>
