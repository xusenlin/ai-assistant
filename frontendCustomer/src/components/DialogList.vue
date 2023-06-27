<script setup>
import {useAppStore} from "../stores/app.js";
import { Comment,Plus,DeleteFilled } from "@element-plus/icons-vue"

const app = useAppStore()

const createDialog = ()=>{
  app.addDialog()
}
</script>

<template>
  <div class="dialog-list">
    <div @click="createDialog" class="dropdown-item" style="border:1px solid hsla(0,0%,100%,.2)">
      <div>
        <el-icon><Plus /></el-icon>
        新建对话
      </div>
    </div>
    <br>
    <div class="list">
      <div
          class="dropdown-item"
          style="margin-bottom: 10px"
          @click="app.dialogIndex=i"
          :class="{active:i===app.dialogIndex}"
          v-for="(item,i) in app.dialog" :key="item.id">
        <div class="li-left">
          <el-icon><Comment /></el-icon>
          <div class="title">
            {{ item.title }}
          </div>
        </div>
        <el-icon @click.stop="app.delDialog(i)"><DeleteFilled /></el-icon>
      </div>
    </div>

  </div>
</template>


<style lang="scss" scoped>

.dialog-list {
  width: 100%;
  .list{
    height: calc(100vh - 246px);
    overflow-y: auto;
    overflow-x: hidden;
    .li-left{
      flex: 1;
      display: flex;
      .title{
        width: 150px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
      }
    }

  }
}

</style>
