<template>
  <div class="sample-list">
    <!--    <Search-->
    <!--        v-model="keyword"-->
    <!--        shape="round"-->
    <!--        placeholder="请输入搜索关键词"-->
    <!--    />-->
    <div class="e" v-if="reportList.length===0">
      暂无查询数据
    </div>
    <div class="list" v-else>
      <CellGroup style="margin-bottom: 10px" inset v-for="(r,i) in reportList" :key="i">
        <div class="van-cell" style="flex-direction: column">
          <div>
            <Icon name="newspaper-o"/>
            {{ (r.reportInfo && r.reportInfo.fileName) ? r.reportInfo.fileName : "暂无报告" }}
          </div>
          <div style="display: flex;justify-content: space-between;margin-top: 10px">
            <div style="color: #999">
              {{ r.sampleId }}
              <!--              <span style="margin-right: 10px">{{r.userTel }}</span>-->
            </div>
            <div style="display:flex;">
              <div class="btn" @click="bindSample(r.sampleId)">解绑</div>
              <div class="btn" v-if="r.reportInfo" @click="downloadPDF(r.reportInfo.reportUrl)">下载</div>
              <div class="btn" v-if="r.reportInfo" @click="viewPDF(r.reportInfo.reportUrl)">查看</div>
              <!--              <Button round type="success" size="small" @click="reportView(r.reportUrl)">预览</Button>-->
            </div>
          </div>
        </div>
      </CellGroup>
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue"
import {CellGroup, Icon, showToast, showConfirmDialog} from "vant"
import {getSamplesByUser, DownloadReport, bindOrUnBindSampleUser, getPdfViewUrl} from "@/api/inex.js"

const reportList = ref([
  // {fileName:"asdsadad.pdf",userName:"userName",userTel:"userTel",reportUrl:"reportUrl"},
  // {fileName:"asdsadad.pdf",userName:"userName",userTel:"userTel",reportUrl:"reportUrl"}
])

const getSampleList = () => {
  getSamplesByUser({}).then(r => {
    reportList.value = r || [];
  }).catch(() => {
  })
}
getSampleList()
const downloadPDF = (fileName) => {
  window.open(DownloadReport(fileName))
}
const viewPDF = async (fileName) => {
  const x = await getPdfViewUrl({fileName})
  window.location.href = x
  console.log(x)
  // window.open(x)
}
const bindSample = async (id) => {
  showConfirmDialog({
    title: '提示',
    message:
        '确认解绑此样本吗？',
  }).then(() => {
    return bindOrUnBindSampleUser({bind: false, sampleId: id})
  }).then(() => {
    showToast("解绑成功")
    getSampleList()
  }).catch(() => {
    // on cancel
  });

}
</script>

<style scoped lang="scss">
.sample-list {
  display: flex;
  width: 100%;
  min-height: calc(100vh - 188px);
  flex-direction: column;

  .e {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .list {
    width: 100%;
    flex: 1;
    margin: 0px 0 20px 0;
    overflow-y: auto;

    .btn {
      margin-left: 20px;
      color: var(--van-primary-color);
    }
  }
}
</style>
