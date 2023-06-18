<template>
  <Popup
      :show="modelValue"
      round
      position="bottom"
      :style="{ height: '40%' }"
  >
    <CellGroup inset style="margin-top: 40px">
      <Field
          v-model="phone"
          name="phone"
          label="用户名"
          placeholder="请填写用户名"
      />
      <Field
          v-model="msg"
          name="msg"
          type="password"
          label="验证码"
          placeholder="请填写密码"
      >
      </Field>
    </CellGroup>
    <div style="margin: 16px;display: flex;flex-direction: column">
      <Button  round block type="primary" @click="onSubmit">
        登录
      </Button>
      <Checkbox icon-size="16px" style="margin-top: 26px;font-size: 13px;" v-model="checked" >
        已阅读并同意《<a style="color: #0077d6" href="./userAgreement.html">用户隐私政策</a>》
      </Checkbox>
    </div>

  </Popup>
</template>

<script setup>
import { ref } from "vue"
import { Button,CellGroup,Field ,showToast,Popup,Checkbox} from "vant"
import { wxUserLogin } from "@/api/inex.js"
import { setToken,setUserInfo } from "@/utils/common.js"

defineProps({
  modelValue:Boolean,
})
const emit  = defineEmits(['success'])


const msg = ref('');
const phone = ref('');

const checked = ref(false)



const onSubmit = async () => {
  if (!checked.value){
    showToast("请阅读并同意《用户隐私政策》")
    return
  }
  if(!msg.value){
    showToast("请填写验证码")
    return
  }
  if(phone.value.length !==11){
    showToast("请正确填写手机号")
    return
  }
  const user = await wxUserLogin({
    mobileCode:msg.value,
    userTel:phone.value,
    vcode:window.$pageParams.code
  })

  let info = user.tokenInfo ||{}
  setUserInfo(info)
  setToken("Bearer "+ info.tokenValue)
  emit('success')
};
</script>

<style scoped lang="scss">

</style>
