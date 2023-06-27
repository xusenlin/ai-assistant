<script setup>
import {ref,onMounted,watch,computed,nextTick} from 'vue'
import {marked} from "marked"
import {baseURL} from "../config/req.js"
import { useUserStore } from "../stores/user.js"
import { useAppStore } from "../stores/app.js"
import { getUserToken } from "../stores/user.js"
import hljs from "highlight.js/lib/core";
import javascript from "highlight.js/lib/languages/javascript";
import go from "highlight.js/lib/languages/go";
import java from "highlight.js/lib/languages/java";
import python from "highlight.js/lib/languages/python";
import typescript from "highlight.js/lib/languages/typescript";
import php from "highlight.js/lib/languages/php";
import clang from "highlight.js/lib/languages/c";
import rust from "highlight.js/lib/languages/rust";
import objectivec from "highlight.js/lib/languages/objectivec";
import swift from "highlight.js/lib/languages/swift";

import "highlight.js/styles/github.css"
import {ElNotification} from "element-plus"

//为了减少文件大小，只注册会用到的语言，
//也可以全部引入  import hljs from "highlight.js"; 不需要注册
//项目都使用按需引入，最大减少项目体积，后端最好开启gzip压缩
hljs.registerLanguage("javascript",javascript)
hljs.registerLanguage("go",go)
hljs.registerLanguage("java",java)
hljs.registerLanguage("python",python)
hljs.registerLanguage("typescript",typescript)
hljs.registerLanguage("php",php)
hljs.registerLanguage("c",clang)
hljs.registerLanguage("rust",rust)
hljs.registerLanguage("objectivec",objectivec)
hljs.registerLanguage("swift",swift)



const user = useUserStore()
const app = useAppStore()
const chat = computed(() => app.dialog[app.dialogIndex]);
const question = ref("")
const answering = ref(false)

const markedParse = (t) => {
  return marked.parse(t)
}
const sendQuestion = async () => {
  question.value = question.value.trim()
  let dialog = app.dialog[app.dialogIndex]
  let dialogId = dialog.id;
  if (question.value.length === 0) {
    ElNotification.warning({
      title: "提示",
      message: "请输入你的问题！"
    })
    return
  }
  answering.value = true
  try {
    const response = await fetch(baseURL + "/api/openai/GPT3Dot5Turbo", {
      method: 'post',
      headers: {'Content-Type': 'application/json','Authorization':getUserToken()},
      body: JSON.stringify([...dialog.content, {
        role: "user",
        content: question.value
      }]),
    })
    if (!response.ok) {
      const msg = await response.text()
      throw new Error(msg)
    }
    app.pushContent(dialogId,[{
      role: "user",
      content: question.value
    }, {
      role: "assistant",
      content: ""
    }])
    question.value = ""

    const decoder = new TextDecoder('utf-8');
    const reader = response.body.getReader();
    while (true) {
      const {value, done} = await reader.read();
      if (done) {
        highlightedCode()
        break
      }
      app.pushSetAnswer(dialogId,decoder.decode(value))
    }
    answering.value = false
    await user.updateUserInfoByApi()
  } catch (e) {
    ElNotification.error({
      title: "错误",
      message: e.message
    })
    answering.value = false
  }
}


const highlightedCode = (all=false) =>{
  let msg = document.querySelectorAll('.dialog-msg')
  let len = msg.length

  let pre = all ? document.querySelectorAll('pre'):msg[len-1].querySelectorAll('pre')
  if(pre.length ===0){
    return
  }
  pre.forEach(el=>{
    hljs.highlightElement(el);
  })
}

onMounted(()=>{
  highlightedCode(true)
})
watch(()=>app.dialogIndex,()=>{
  nextTick(()=>{//更新成功并等待md解析成功
    setTimeout(()=>{
      highlightedCode(true)
    },300)
  })
},{deep:true})
</script>

<template>
  <div style="padding: 12px">
    <div v-if="app.dialog.length !==0">
      <el-card v-for="(n,i) in chat.content" class="box-card" shadow="never" :key="i">
        <template #header>
          <div class="card-header">
            <span>{{ n.role==="user"?"我":"AI助手" }}</span>
            <!--          <el-button link type="warning" @click="chat.splice(i,1)">-->
            <!--            Remove-->
            <!--          </el-button>-->
          </div>
        </template>
        <div class="dialog-msg" v-if="n.role == 'assistant'" v-html="markedParse(n.content)"></div>
        <div class="dialog-msg" v-else>{{ n.content }}</div>
      </el-card>
      <el-card class="box-card mt" shadow="never">
        <template #header>
          <div class="card-header">
            <span>提问</span>
            <!--          <el-button type="success" link>-->
            <!--            Prompts-->
            <!--          </el-button>-->
          </div>
        </template>
        <div>
          <el-input
              v-model="question"
              :rows="3"
              type="textarea"
              placeholder="提问"
          />

        </div>
        <div class="send-group">
          <el-button type="warning" round :disabled="answering" @click="app.resetDialog(app.dialogIndex)">重置</el-button>
          <el-button type="success" round :loading="answering" :disabled="answering" @click="sendQuestion">发送
          </el-button>
        </div>
      </el-card>
    </div>
    <div v-else style="margin-top: 18%">
      <el-empty description="暂无对话" />
    </div>
  </div>
</template>

<style lang="scss">

.box-card {
    margin-bottom: 20px;
}

.send-group {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}

pre {
    background: #f7f7f7 !important;
    border-radius: 2px;
    padding: 10px;
}

.card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.dialog-msg {
    font-size: 14px;
    white-space: pre-wrap;
    word-wrap: break-word;
}
</style>
