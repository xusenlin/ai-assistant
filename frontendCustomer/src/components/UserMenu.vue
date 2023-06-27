<script setup>
import {useAppStore} from "../stores/app.js";
import {useUserStore} from "../stores/user.js";
import {ElMessageBox, ElMessage} from "element-plus"
import { updatePwd } from "../api/user.js"
import {MoreFilled, ChatDotRound, Avatar, Histogram,Edit, CircleCloseFilled} from "@element-plus/icons-vue"

const app = useAppStore()
const user = useUserStore()


const editPwd = () => {
  ElMessageBox.prompt('请输入新密码，必须大于等于6位', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '关闭',
  }).then(({value}) => {
    if (value.length < 6) {
      ElMessage.warning("输入的新密码少于6位")
      return
    }
    updatePwd(value).then(()=>{
      ElMessage.success("修改成功，请重新登录")
      setTimeout(() => {
        user.loginOut()
      }, 1000)
    })
  }).catch(() => {})
}

</script>

<template>
  <div class="login-info">
    <template v-if="user.info">
      <div class="dropdown-item">
        <div>
          <el-icon><Histogram /></el-icon>
          总共消耗Token
        </div>
        <div>{{ user.tokenConsumed }}</div>
      </div>
      <div class="dropdown-item">
        <div>
          <el-icon>
            <ChatDotRound/>
          </el-icon>
          剩余对话次数
        </div>
        <div>{{ user.remainingDialogueCount }}</div>
      </div>

      <el-dropdown trigger="click" popper-class="dropdown" style="width: 100%">
        <div class="dropdown-item">
          <div>
            <el-icon>
              <Avatar/>
            </el-icon>
            {{ user.userName }}
          </div>
          <el-icon>
            <MoreFilled/>
          </el-icon>
        </div>

        <template #dropdown>
          <div class="dro">
            <div class="dropdown-item" @click="editPwd()">
              <div>
                <el-icon>
                  <Edit/>
                </el-icon>
                修改密码
              </div>

            </div>
            <div class="dropdown-item" @click="user.loginOut()">
              <div>
                <el-icon>
                  <CircleCloseFilled/>
                </el-icon>
                退出登录
              </div>

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

<style lang="scss">
.dropdown {
  background: #000 !important;
  width: 250px;
  border: none !important;

  .dro {
    color: #fff;
    width: 100%;
  }

  .el-popper__arrow {
    display: none;
  }
}
@media only screen and (max-width: 720px) {
  .dropdown{
    width: 97%;
  }
}

</style>
<style lang="scss" scoped>

.login-info {
  width: 100%;
}

</style>
