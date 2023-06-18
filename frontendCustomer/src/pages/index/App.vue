<template>
<div class="index">
<!--<img src="./img/logo.jpeg" style="width: 30%;margin-top:25%;margin-bottom: 40px;border-radius: 50%" alt="">*-->
<!--  <div class="bing" @click="bing">绑定</div>-->
<!--  <DivideTitle>样本列表</DivideTitle>-->

<!--  <SampleList v-if="!showLoginPopup"/>-->
  <div style="padding-bottom: 10px;">
    <a style="color: #979797;font-size: 13px" href="./userAgreement.html">用户隐私政策</a>
  </div>
  <LoginPopup v-model="showLoginPopup" @success="loginSuccess"/>
</div>
</template>

<script setup>
import { ref } from "vue"
import LoginPopup from "./LoginPopup.vue"

import { isLogin} from "@/utils/common.js"


const showLoginPopup = ref(!isLogin())


const loginSuccess = ()=>{
  showLoginPopup.value = false
}

</script>

<style scoped lang="scss">
  .index{
    display: flex;
    width: 100%;
    height: 100vh;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    .bing{
      margin-top: 20px;
      width: 100px;
      height: 100px;
      color: #fff;
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 16px;
      flex-shrink: 0;
      border-radius: 50%;
      background-color: var(--van-primary-color);
    }
  }
</style>
