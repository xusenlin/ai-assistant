<script setup>
import Login from "./components/Login.vue"
import DialogContent from "./components/DialogContent.vue"
import UserMenu from "./components/UserMenu.vue"
</script>
<template>
  <el-container class="app">
    <el-header>
      <div class="container head">
        <div style="display:flex;flex: 1">
          <div class="title">
            Ai 助手
          </div>
        </div>
        <UserMenu/>
      </div>
    </el-header>
    <el-main>
      <div class="container">
        <DialogContent/>
      </div>
    </el-main>
    <Login/>
  </el-container>
</template>

<style  lang="scss">

.container {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}
.el-tabs__nav-wrap::after{
  height: 0;
  background: none;
}

.el-tabs__item{
  height: 46px;
}
.el-main{
  padding: 0;
}
.el-header{
  padding: 0;
}
.head {
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: solid 1px #e4e7ed;
  color: #606266;
}
.menu{
  margin-left: 30px;flex: 1;
  display: flex;
  align-items: flex-end;
}

.footer {
  width: 100%;
  text-align: center;
  color: #aa0000;
  font-size: 12px
}

</style>
