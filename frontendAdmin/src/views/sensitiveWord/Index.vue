<template>
  <div class="box">
    <ActionBar @reset="resetParams" @refresh="refreshTable">
      <template #left>
        <el-button type="primary" @click="addSensitiveWords()">添加</el-button>
        <el-button type="success" @click="clickMigrate()">敏感词迁移</el-button>
      </template>
      <template #right>
<!--        <el-input v-model="searchParams.id" placeholder="请输入ID"  clearable />-->
      </template>
    </ActionBar>
    <el-table :data="tableData" style="width: 100%" max-height="calc(100vh - 267px)">
      <el-table-column prop="ID" label="ID"/>
      <el-table-column prop="Name" label="Name"/>
      <el-table-column prop="CreatedAt" label="CreatedAt"/>
      <el-table-column fixed="right"  label="操作">
        <template #default="scope">
          <el-button type="danger" @click.prevent="deleteRow(scope.row)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <Pagination ref="paginationRef" :params="searchParams" :reqFunc="sensitiveWordsList" @pageData="setTableData" />
  </div>
</template>

<script setup lang="ts">
import { sensitiveWordsList } from "@/api/sensitiveWords.ts"
import usePagination from "@/compositionApi/pagination.ts"
import useExtraAction from "./extraAction"


const {searchParams, tableData, paginationRef, setTableData, refreshTable,resetParams} =  usePagination()
const { deleteRow,clickMigrate,addSensitiveWords } = useExtraAction(refreshTable)




</script>

<style lang="scss" scoped>
.box {
  width: 100%;
  height: 100%;
  padding: 10px 0;
}
</style>
