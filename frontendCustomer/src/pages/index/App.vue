<template>
<div class="index">
  <div class="dialog">
    <div class="block" v-for="(n,i) in chat" :key="i">
      <span>{{ n.role.charAt(0).toUpperCase() + n.role.slice(1) }}:</span>
      <br>
      <div class="dialog-msg" v-if="n.role == 'assistant'" v-html="markedParse(n.content)"></div>
      <div class="dialog-msg" v-else>{{ n.content }}</div>
    </div>
  </div>
  <div class="input">
    <Field
        v-model="question"
        center
        clearable
        placeholder="请输入你的问题"
    >
      <template #button>
        <Button round :disabled="answering" size="small" type="primary" @click="sendQuestion()">发   送</Button>
      </template>
    </Field>
  </div>
  <LoginPopup v-model="showLoginPopup" @success="loginSuccess"/>
</div>
</template>

<script setup>
import { ref } from "vue"
import {marked} from "marked"
import { baseURL } from "@/config/request.js"
import { showToast,Field ,Button } from "vant"
import LoginPopup from "./LoginPopup.vue"
import {getToken, isLogin} from "@/utils/common.js"

const chat = ref([])
const question = ref("")
const answering = ref(false)

const showLoginPopup = ref(!isLogin())



const markedParse = (t) => {
  return marked.parse(t)
}
const sendQuestion = async () => {
  question.value = question.value.trim()
  if (question.value.length === 0) {
    showToast("请输入你的问题！")
    return
  }
  answering.value = true
  // let loading = showLoadingToast({duration:0})
  try {
    const response = await fetch(baseURL + "/api/openai/GPT3Dot5Turbo", {
      method: 'post',
      headers: {
        'Content-Type': 'application/json',
        "Authorization" : getToken()
      },
      body: JSON.stringify([...chat.value, {
        role: "user",
        content: question.value
      }]),
    })
    if (!response.ok) {
      const msg = await response.text()
      throw new Error(msg)
    }
    chat.value.push({
      role: "user",
      content: question.value
    }, {
      role: "assistant",
      content: ""
    })
    question.value = ""

    const decoder = new TextDecoder('utf-8');
    const reader = response.body.getReader();
    // eslint-disable-next-line no-constant-condition
    while (true) {
      const {value, done} = await reader.read();
      if (done) {
        // let msg = document.querySelectorAll('.dialog-msg')
        // let len = chat.value.length
        // if (msg.length !== len || msg.length === 0){
        //   break
        // }
        // let pre = msg[len-1].querySelectorAll('pre')
        // if(pre.length ===0){
        //   break
        // }
        // pre.forEach(el=>{
        //   // hljs.highlightElement(el);
        // })
        break
      }
      chat.value[chat.value.length - 1].content += decoder.decode(value)
    }
    answering.value = false
    // loading.close()
  } catch (e) {
    showToast(e.message)
    answering.value = false
    // setTimeout(()=>{
    //   loading.close()
    // },1000)
  }
}


const loginSuccess = ()=>{
  showLoginPopup.value = false
}

</script>

<style  lang="scss">
  .index{
    display: flex;
    width: 100%;
    height: 100vh;
    background: #fff;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    .dialog{
      width: 100%;
      flex: 1;
      padding: 0 10px;
      overflow-y: auto;
      overflow-x: hidden;
      color: #222;
      .block{
        padding: 18px;
        background: #f4f4f4;
        margin-top: 10px;
        border-radius: 10px;
        font-size: 13px;
        .dialog-msg{
          margin-top: 10px;
          pre{
            background: #eee;
          }
        }
      }
    }
    .input{
      width: 100%;
    }
  }
</style>
