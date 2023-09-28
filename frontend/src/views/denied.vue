<template>
  <div class="lock-flex">
    <svg-icon name="lock" style="width: 120px; height: 121px;"/>
    <div class="text-xl font-medium m-3" style="color: #20202099;">您暂无权限访问</div>
  </div>
</template>
<script setup lang="ts">

</script>
<style lang="scss" scoped>
.lock-flex {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  @apply mt-44;
}
</style>
