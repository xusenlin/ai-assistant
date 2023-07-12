<script setup>
import { Comment,Plus,DeleteFilled } from "@element-plus/icons-vue"
import storage from "good-storage";
import {storageKeyDialog, storageKeyToken} from "../config/app.js";
import {v4 as uuid} from "uuid";
import {ref} from "vue";
import {ElMessage} from "element-plus";

const emit = defineEmits(['dialogNode'])

const storageDialog = storage.get(storageKeyDialog,[{title:"默认对话",id:uuid()}])

const dialog = ref(storageDialog)
const active = ref(0)


const createDialog = ()=>{

  dialog.value.push({
    id:uuid(),
    title:`新的对话`,
  })
  active.value = dialog.value.length-1
  storage.set(storageKeyDialog,dialog.value)
}

const delDialog = index =>{
  if(dialog.value.length===1){
    ElMessage.warning("最少保留一个对话")
    return
  }

  if(index <= active.value){
    active.value--
  }
  dialog.value.splice(index,1)
  storage.set(storageKeyDialog,dialog.value)
}
const clickNode = index =>{
  active.value = index
  emit('dialogNode',dialog.value[index].id)
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
          @click="clickNode(i)"
          :class="{active:i=== active}"
          v-for="(item,i) in dialog" :key="item.id">
        <div class="li-left">
          <el-icon><Comment /></el-icon>
          <div class="title">
            {{ item.title }}
          </div>
        </div>
        <el-icon @click.stop="delDialog(i)"><DeleteFilled /></el-icon>
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
