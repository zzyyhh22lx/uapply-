<template>
    <div class="apply-home">
        <router-view></router-view>
        <apply-admin />
        <el-button>按钮2</el-button>
    </div>
</template>
<script setup lan="ts">
import ApplyAdmin from './components/admin/admin.vue';
import { ElButton } from 'element-plus';
</script>
<style scoped lang="scss">
</style>