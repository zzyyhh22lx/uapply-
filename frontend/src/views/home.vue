<template>
    <div class="apply-home">
        <uapply-header />
        <router-view></router-view>
        <el-backtop />
    </div>
</template>
<script setup lan="ts">
import UapplyHeader from './components/header/header';
import { ElBacktop } from 'element-plus';
</script>
<style scoped lang="scss">
.apply-home {
    @apply bg-[#f5f5f5];
}
</style>