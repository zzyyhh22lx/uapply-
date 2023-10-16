<template>
    <el-select>
        <el-option 
            v-for="item in props.options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        />
    </el-select>
</template>
<script setup lang="ts">
import { ElSelect, ElOption } from "element-plus";
import { PropType } from 'vue';
import type { OPTIONS_TYPE } from '../constant';

const props = defineProps({
    options: {
        type: Object as PropType<OPTIONS_TYPE[]>,
        required: true,
    }
})
</script>