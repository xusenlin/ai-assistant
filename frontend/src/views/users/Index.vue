<template>
  <div class="box">
    <ActionBar @reset="resetParams" @refresh="refreshTable">
      <template #left>
<!--        <el-button type="primary" @click="addSensitiveWords()">添加</el-button>-->
      </template>
      <template #right>
<!--        <el-input v-model="searchParams.id" placeholder="请输入ID"  clearable />-->
      </template>
    </ActionBar>
    <el-table :data="tableData" style="width: 100%" max-height="calc(100vh - 267px)">
      <el-table-column prop="ID" label="ID"/>
      <el-table-column prop="Username" label="Username"/>
      <el-table-column prop="IsAdmin" label="是否是管理员"/>
      <el-table-column prop="TokenConsumed" label="共消耗Token数"/>
      <el-table-column prop="RemainingDialogueCount" label="剩余对话次数"/>
      <el-table-column prop="CreatedAt" label="CreatedAt"/>
      <el-table-column fixed="right"  label="操作" width="460">
        <template #default="scope">

          <el-button v-if="scope.row.Username!=='Admin'" :type="scope.row.Status==0?'success':'warning'" @click.prevent="setStatus(scope.row)">
            {{ scope.row.Status ==0 ? "启用":"禁用"}}
          </el-button>

          <el-button type="primary" @click.prevent="setRemainingDialogueCount(scope.row.ID)">
            设置对话次数
          </el-button>
          <el-button v-if="scope.row.Username!=='Admin'" type="danger" @click.prevent="deleteRow(scope.row)">
            删除
          </el-button>
          <el-button  type="warning" @click.prevent="setPassword(scope.row.ID)">
            重置密码
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination ref="paginationRef" :params="searchParams" :reqFunc="userList" @pageData="setTableData" />
  </div>
</template>

<script setup lang="ts">
import { userList } from "@/api/user.ts"
import usePagination from "@/compositionApi/pagination.ts"
import useExtraAction from "./extraAction"


const {searchParams, tableData, paginationRef, setTableData, refreshTable,resetParams} =  usePagination()
const { deleteRow,setPassword,setStatus,setRemainingDialogueCount } = useExtraAction(refreshTable)




</script>

<style lang="scss" scoped>
.box {
  width: 100%;
  height: 100%;
  padding: 10px 0;
}
</style>
