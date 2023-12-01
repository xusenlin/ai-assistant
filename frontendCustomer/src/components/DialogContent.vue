<script setup>
import {ref,nextTick} from 'vue'
import {marked} from "marked"
import {baseURL} from "../config/req.js"
import { useUserStore } from "../stores/user.js"
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
import storage from "good-storage";
import promptsData from "./prompts.js"

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
const chats = ref([])
const question = ref("")
const dialogId = ref("")
const answering = ref(false)
const prompts = ref(promptsData)
const promptsVisible = ref(false)

const sendQuestion = async () => {
  question.value = question.value.trim()

  if (question.value.length === 0) {
    ElNotification.warning({
      title: "提示",
      message: "请输入你的问题！"
    })
    return
  }
  answering.value = true

  if(chats.value.length >= 20){
    chats.value.splice(0,2)
  }
  try {
    const response = await fetch(baseURL + "/api/openai/GPT3Dot5Turbo", {
      method: 'post',
      headers: {'Content-Type': 'application/json','Authorization':getUserToken()},
      body: JSON.stringify([...chats.value, {
        role: "user",
        content: question.value
      }]),
    })
    if (!response.ok) {
      const msg = await response.text()
      throw new Error(msg)
    }
    chats.value.push({
      role: "user",
      content: question.value
    }, {
      role: "assistant",
      content: ""
    })
    question.value = ""
    let chatsLen = chats.value.length
    const decoder = new TextDecoder('utf-8');
    const reader = response.body.getReader();
    while (true) {
      const {value, done} = await reader.read();
      if (done) {
        highlightedCode()
        storage.set(dialogId.value,chats.value)
        break
      }
      chats.value[chatsLen-1].content += decoder.decode(value)
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

const markedParse = (t) => {
  return marked.parse(t)
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


const Render = id =>{
  dialogId.value = id
  chats.value =  storage.get(id,[])
  nextTick(()=>{
    setTimeout(()=>{
      highlightedCode(true)
    },300)
  })
}

defineExpose({Render})
</script>

<template>
  <div style="padding: 12px">
    <div>
      <el-card v-for="(n,i) in chats" class="box-card" shadow="never" :key="i">
        <template #header>
          <div class="card-header">
            <span>{{ n.role==="user"?"我":"AI助手" }}</span>
            <el-button link type="warning" @click="chats.splice(i,1)">
              Remove
            </el-button>
          </div>
        </template>
        <div class="dialog-msg"  v-html="markedParse(n.content)"></div>
      </el-card>
      <el-card class="box-card mt" shadow="never">
        <template #header>
          <div class="card-header">
            <span>提问</span>
            <el-button type="success" link  @click="promptsVisible = true">
              Prompts
            </el-button>
          </div>
        </template>
        <div>
          <el-input
              v-model="question"
              :rows="5"
              type="textarea"
              placeholder="使用符号```将你的代码包裹起来，Ai能更好的识别代码。"
          />

        </div>
        <div class="send-group">
          <el-button type="warning" round :disabled="answering" @click="chats.value = []">重置</el-button>
          <el-button type="success" round :loading="answering" :disabled="answering" @click="sendQuestion">发送
          </el-button>
        </div>
      </el-card>
    </div>
    <el-dialog width="80%" v-model="promptsVisible" title="Prompts">
      <el-collapse model-value="">
        <el-collapse-item v-for="n in prompts" :title="n.act">
          <div>
            {{n.prompt}}
          </div>
          <div style="display: flex;justify-content: flex-end">
            <el-button round size="small" type="success" :disabled="answering"  @click="question=n.prompt;promptsVisible=false">使用这个 prompt.</el-button>
          </div>
        </el-collapse-item>
      </el-collapse>
    </el-dialog>
    <el-footer >
      <div style="width: 100%;text-align: center;color: #aa0000;font-size: 12px">
        本项目是一个实验性项目，可能存在不足和错误，仅供学术研究和讨论目的使用，不以任何形式获利，也不得用于任何商业用途，请谨慎使用。
        <div style="color: #999;margin-top: 10px">如果您有任何问题，可以发送邮件到 wumulaozu@gmail.com 或到我的博客 xusenlin.top 进行留言，我会尽快回复您。  😄</div>
      </div>
    </el-footer>
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
