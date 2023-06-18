<script setup lang="ts">
import { ref } from 'vue'
import { getOptionByKey,setOptionByKey } from "@/api/option.ts"
import {ElMessage, ElMessageBox} from "element-plus";

const openaiUrl = ref("")
const openaiKeys = ref<{val:string}[]>([])

const initData = ()=>{
  getOptionByKey({key:"openai_url"}).then(u=>{
    openaiUrl.value = u.OptionValue
  })
  getOptionByKey({key:"openai_keys"}).then(r=>{
    try {
      if (r.OptionValue == ""){
        return
      }
      let keys = JSON.parse(r.OptionValue) || [] as string[]
      openaiKeys.value = []
      keys.forEach(val=>{
        openaiKeys.value.push({val})
      })

    }catch (e) {
      ElMessage.warning(JSON.stringify(e))
    }
  })
}
initData()

const deleteRow = index=>{
  if(openaiKeys.value.length ===1){
    ElMessage.warning("最少保留一个key")
    return
  }
  let keys = []
  openaiKeys.value.forEach((v,i)=>{
    if(i!==index){
      keys.push(v.val)
    }
  })
  setOptionByKey({key:"openai_keys",val:JSON.stringify(keys)}).then(()=>{
    initData()
  })
}

const addKey = () => {
  ElMessageBox.prompt('添加新的openai key', '提示', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
  }).then(({value}) => {
    if(value==""){
      ElMessage.warning("请输入key")
      return
    }
    let val = JSON.stringify([...openaiKeys.value.map(v=>v.val),value])
    setOptionByKey({key:"openai_keys",val}).then(()=>{
      ElMessage.success("添加key成功")
      initData()
    })
  }).catch(() => {})
}


const editUrl = () => {
  ElMessageBox.prompt('编辑代理地址', '提示', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputValue:openaiUrl.value
  }).then(({value}) => {
    if(value==""){
      ElMessage.warning("请输入url")
      return
    }
    setOptionByKey({key:"openai_url",val:value}).then(()=>{
      ElMessage.success("设置成功")
      initData()
    })
  }).catch(() => {})
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
        <div style="display: flex;align-items: center;justify-content: space-between">
          <p>{{openaiUrl}}</p>
          <el-button  type="warning" @click.prevent="editUrl()">
            编辑
          </el-button>
        </div>
        <el-divider style="margin-top: 50px">Openai Keys</el-divider>

        <el-button  type="success" @click.prevent="addKey()">
          添加
        </el-button>
        <el-table :data="openaiKeys" style="width: 100%">
          <el-table-column prop="val" label="值" />
          <el-table-column prop="address" label="操作">
            <template #default="s">
              <el-button  type="danger" @click.prevent="deleteRow(s.$index)">
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
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
