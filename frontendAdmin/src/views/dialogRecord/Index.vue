<template>
  <div class="box">
    <ActionBar @reset="resetParams" @refresh="refreshTable">
      <template #left>
      </template>
      <template #right>
        <el-input v-model="searchParams.keyword" placeholder="请输入关键词"  clearable />
      </template>
    </ActionBar>
    <el-table :data="tableData" style="width: 100%" max-height="calc(100vh - 267px)">
      <el-table-column prop="ID" width="80" label="ID"/>
      <el-table-column prop="Username" label="用户"/>
      <el-table-column prop="Question" label="问题" show-overflow-tooltip/>
      <el-table-column prop="Response" label="回答" show-overflow-tooltip/>
      <el-table-column prop="UserIP" label="IP"/>
      <el-table-column prop="CreatedAt" label="CreatedAt"/>
    </el-table>
    <Pagination ref="paginationRef" :params="searchParams" :reqFunc="dialogList" @pageData="setTableData" />
  </div>
</template>

<script setup lang="ts">
import { dialogList } from "@/api/dialog.ts"
import usePagination from "@/compositionApi/pagination.ts"



const {searchParams, tableData, paginationRef, setTableData, refreshTable,resetParams} =  usePagination()



</script>

<style lang="scss" scoped>
.box {
  width: 100%;
  height: 100%;
  padding: 10px 0;
}
</style>
