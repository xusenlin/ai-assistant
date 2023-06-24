<script setup lang="ts">
import {ref} from 'vue'
import {getOptionByKey, setOptionByKey} from "@/api/option.ts"
import {ElMessage, ElMessageBox} from "element-plus";
import {getOpenaiKey,add,destroy,ping, OpenaiKey} from "@/api/openaiKey.ts";

const openaiUrl = ref("")
const openaiSysPrompt = ref("")
const openaiKeys = ref<OpenaiKey[]>([])
const initData = () => {
  getOptionByKey({key: "openai_url"}).then(u => {
    openaiUrl.value = u.OptionValue
  })
  getOptionByKey({key: "openai_sys_prompt"}).then(r => {
    openaiSysPrompt.value = r.OptionValue
  })
  getOpenaiKey().then(k => {
    openaiKeys.value = k
  })
}
initData()

const deleteRow = id => {

  if (openaiKeys.value.length === 1) {
    ElMessage.warning("最少保留一个key")
    return
  }

  destroy({id}).then(()=>{
    initData()
  })
}
const testRow = key =>{
  ping({key}).then(pong=>{
    ElMessage.success(pong)
    initData()
  })
}

const addKey = () => {
  ElMessageBox.prompt('添加新的openai key', '提示', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
  }).then(({value}) => {
    if (value == "") {
      ElMessage.warning("请输入key")
      return
    }
    add({Value:value}).then(()=>{
      ElMessage.success("添加key成功")
      initData()
    })
  }).catch(() => {
  })
}


const editUrl = () => {
  ElMessageBox.prompt('编辑代理地址', '提示', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputValue: openaiUrl.value
  }).then(({value}) => {
    if (value == "") {
      ElMessage.warning("请输入url")
      return
    }
    setOptionByKey({key: "openai_url", val: value}).then(() => {
      ElMessage.success("设置成功")
      initData()
    })
  }).catch(() => {
  })
}
const editPrompt = () => {
  ElMessageBox.prompt('编辑系统提示词', '提示', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputType: "textarea",
    inputValue: openaiSysPrompt.value
  }).then(({value}) => {
    if (value == "") {
      ElMessage.warning("请输入提示词")
      return
    }
    setOptionByKey({key: "openai_sys_prompt", val: value}).then(() => {
      ElMessage.success("设置成功")
      initData()
    })
  }).catch(() => {
  })
}

</script>

<template>
  <el-card shadow="never">
    <template #header>
      <div style="display: flex;justify-content: space-between;align-items: center">
        <span>Openai</span>
        <div>

        </div>
      </div>
    </template>
    <div>
      <div>
        <el-divider>请求或代理地址</el-divider>
        <div style="display: flex;align-items: center;justify-content: space-between;padding: 20px">
          <p>{{ openaiUrl }}</p>
          <el-button type="warning" @click.prevent="editUrl()">
            编辑
          </el-button>
        </div>
        <el-divider>Openai SysPrompt</el-divider>
        <div style="display: flex;align-items: center;justify-content: space-between;padding: 20px">
          <div style="max-width: 80%;font-size: 13px;">{{ openaiSysPrompt }}</div>
          <el-button type="warning" @click.prevent="editPrompt()">
            编辑
          </el-button>
        </div>
        <el-divider>Openai Key池</el-divider>
        <div style="padding: 20px">
          <el-button type="success" @click.prevent="addKey()">
            添加
          </el-button>
          <el-table :data="openaiKeys" style="width: 100%">
            <el-table-column width="50" prop="ID" label="ID"/>
            <el-table-column prop="Value" label="值"/>
            <el-table-column prop="Status" label="状态">
              <template #default="s">
                <div style="padding: 10px">
                  <el-badge is-dot :type="s.row.Status==1?'success':'danger'">
                    {{ s.row.Status==1?'正常':'异常' }}
                  </el-badge>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="ExceptionReason" label="异常信息"/>
            <el-table-column prop="address" label="操作">
              <template #default="s">
                <el-button type="danger" @click.prevent="deleteRow(s.row.ID)">
                  删除
                </el-button>
                <el-button type="primary" @click.prevent="testRow(s.row.Value)">
                  测试
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>
  </el-card>
</template>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
