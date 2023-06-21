<script setup>
import {useAppStore} from "../stores/app.js";
import {useUserStore} from "../stores/user.js";

const app = useAppStore()
const user = useUserStore()
</script>

<template>
  <div class="login-info">
    <template v-if="user.info">
      <el-dropdown>
        <div style="position: relative;top: 3px">
          {{ user.userName }}({{user.remainingDialogueCount}})
        </div>
        <template #dropdown>
          <div style="padding: 20px;min-width: 160px">
            <p>
              用户名：{{ user.userName }}
            </p>
            <p>剩余对话次数：{{ user.remainingDialogueCount }}</p>
            <el-button round style="width: 100%"  type="warning" @click="user.loginOut()">
              退出登录
            </el-button>
          </div>

        </template>
      </el-dropdown>
    </template>
    <el-button round v-else @click="app.showLogin = true"  type="primary">
      登录
    </el-button>
  </div>
</template>

<style lang="scss">

</style>
