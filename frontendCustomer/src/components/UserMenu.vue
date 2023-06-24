<script setup>
import {useAppStore} from "../stores/app.js";
import {useUserStore} from "../stores/user.js";

const app = useAppStore()
const user = useUserStore()
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
            <el-button round style="width: 100%"  type="warning" @click="user.loginOut()">
              退出登录
            </el-button>
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
