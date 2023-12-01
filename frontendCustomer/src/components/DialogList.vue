<script setup>
import {Comment, Plus, DeleteFilled} from "@element-plus/icons-vue"
import storage from "good-storage";
import {storageKeyDialog} from "../config/app.js";
import {v4 as uuid} from "uuid";
import {onMounted, reactive} from "vue";
import {ElMessage} from "element-plus";

const emit = defineEmits(['render'])

const storageDialog = storage.get(storageKeyDialog, {
  list: [{title: "默认对话", id: uuid()}],
  active: 0
})

const d = reactive(storageDialog)

onMounted(() => {
  storage.set(storageKeyDialog, d)
  emit('render', d.list[d.active].id)
})

const createDialog = () => {
  let id = uuid()
  d.list.push({
    id,
    title: `新的对话`,
  })
  d.active = d.list.length - 1
  storage.set(storageKeyDialog, d)
  emit('render', id)
}

const delDialog = index => {
  if (d.list.length === 1) {
    ElMessage.warning("最少保留一个对话")
    return
  }
  storage.remove(d.list[index].id)

  if (index <= d.active) {
    d.active !== 0 ? d.active-- : d.active = 0
  }
  d.list.splice(index, 1)
  storage.set(storageKeyDialog, d)
  emit('render', d.list[d.active].id)
}
const clickNode = index => {
  d.active = index
  storage.set(storageKeyDialog, d)
  emit('render', d.list[index].id)
}
const SetTitle = (title = "") => {
  d.list[d.active].title = title.length >= 40 ? title.substring(0, 40) + "..." : title
  storage.set(storageKeyDialog, d)
}
defineExpose({SetTitle})
</script>

<template>
  <div class="dialog-list">
    <div @click="createDialog" class="dropdown-item" style="border:1px solid hsla(0,0%,100%,.2)">
      <div>
        <el-icon>
          <Plus/>
        </el-icon>
        新建对话
      </div>
    </div>
    <br>
    <div class="list">
      <el-tooltip
          effect="dark"
          v-for="(item,i) in d.list" :key="item.id"
          :content="item.title"
          placement="right"
      >
        <div
            class="dropdown-item"
            style="margin-bottom: 10px"
            @click="clickNode(i)"
            :class="{active:i=== d.active}"
            >
          <div class="li-left">
            <el-icon>
              <Comment/>
            </el-icon>
            <div class="title">
              {{ item.title }}
            </div>
          </div>
          <el-icon @click.stop="delDialog(i)">
            <DeleteFilled/>
          </el-icon>
        </div>
      </el-tooltip>

    </div>

  </div>
</template>


<style lang="scss" scoped>

.dialog-list {
  width: 100%;

  .list {
    height: calc(100vh - 246px);
    overflow-y: auto;
    overflow-x: hidden;

    .li-left {
      flex: 1;
      display: flex;

      .title {
        width: 150px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
      }
    }

  }
}

</style>
