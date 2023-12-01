<script setup>
import Login from "./components/Login.vue"
import DialogList from "./components/DialogList.vue"
import DialogContent from "./components/DialogContent.vue"
import UserMenu from "./components/UserMenu.vue"
import {ref} from "vue";
const dialogContentRef = ref()
const dialogListRef = ref()

const dialogRender = (id)=>{
  dialogContentRef.value.Render(id)
}
const setTitle = (title)=>{
  dialogListRef.value.SetTitle(title)
}
</script>
<template>
  <div class="app">
    <div class="slider">
      <DialogList ref="dialogListRef" @render="dialogRender" />
      <div class="user">
        <UserMenu/>
      </div>
    </div>
    <div class="main">
      <div class="top">
        <UserMenu/>
      </div>
      <DialogContent @title="setTitle" ref="dialogContentRef"/>
    </div>
    <Login/>
  </div>

</template>

<style  lang="scss">
.app{
  display: flex;
  .slider{
    width: 260px;
    height: 100vh;
    background: #202123;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 0.5em;
    color: #fff;
    .user{
      border-top:1px solid rgb(255 255 255 / 20%);
      padding-top: 10px;
    }
  }

  .main{
    flex: 1;
    background: #f7f7f8;
    height: 100vh;
    overflow-y: auto;
    .top{
      display: none;
    }
  }
  @media only screen and (max-width: 720px) {
    .slider {
      display: none;
    }
    .main .top{
      display: block;
      color: #fff;
      background: #202123;
    }
  }
}

</style>
