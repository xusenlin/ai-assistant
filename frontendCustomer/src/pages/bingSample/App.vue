<template>
  <div class="bing-sample">
    <img src="./img/logo.jpeg" style="width: 30%;margin-bottom: 60px;border-radius: 50%" alt="">
    <div class="box">
      <CellGroup inset>
        <Field
            v-model="sampleId"
            type="text"
            name="msg"
            required
            center
            label="样本编号"
            placeholder="请输入或者扫描样本编号"
        >
          <template #button>
            <Button @click="scanQRCode" size="mini" round type="primary" icon="scan">扫码</Button>
          </template>
        </Field>
      </CellGroup>
      <div style="padding: 16px;margin-top: 40px;margin-bottom: 40px" >
        <Button round type="primary" block @click="bindSample">确认绑定</Button>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue"
import wx from 'weixin-js-sdk';
import { wechatSignature } from "@/utils/wechat.js"
import { bindOrUnBindSampleUser} from "@/api/inex.js"
import {CellGroup,Button,Field,showToast} from "vant"


wechatSignature()

const sampleId = ref("")
const bindSample = async () => {
  if(!sampleId.value){
    showToast("请正确输入样本编号")
    return
  }
  await bindOrUnBindSampleUser({bind:true,sampleId:sampleId.value})

  showToast("绑定成功")
  // sampleId.value = ""
  window.location.href = "./index.html"
}

const scanQRCode = () => {
  wx.scanQRCode({
    needResult: 1, // 默认为0，扫描结果由微信处理，1则直接返回扫描结果，
    scanType: ["qrCode", "barCode"], // 可以指定扫二维码还是一维码，默认二者都有
    success: function (res) {
      showToast("扫码成功")
      let result = res.resultStr
      let r = result.split(",")
      if(r.length===2){
        sampleId.value = r[1]
      }else {
        sampleId.value = result
      }
    },
    fail: function (res) {
      showToast(JSON.stringify(res))
    },
    cancel: function () {

    }
  });
}
</script>

<style scoped lang="scss">
.bing-sample {
  display: flex;
  width: 100%;
  height: 100vh;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  padding-bottom: 80px;
  .box {
    width: 100%;
  }
}
</style>
