<template>
  <div class="divider-title">
    <slot></slot>
  </div>
</template>
<script setup>

</script>
<style scoped lang="scss">
  .divider-title{
    width: 100%;
    height: 50px;
    display: flex;
    align-items: center;
    color: var(--van-primary-color);
    font-size: var(--van-notice-bar-font-size);
    padding: 0 22px;
  }
</style>
