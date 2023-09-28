<template>
    <div v-html="transformStr('<div>测试</div>')" class="myButton"></div>
</template>
<script setup lan="ts">
import { transformStr } from '@/libs/utils.ts';

</script>
<style lang="scss" scoped>
.myButton {
    /**
        这里的@apply相当于add class样式
        p-4 -> padding: 16px
    */
    @apply p-4 bg-[pink];
}
</style>
