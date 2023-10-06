<template>
    <div class="title_box">
        <div class="title">
            <span>{{ props.title }}</span>
        </div>
        <div class="id_text">（ID：{{ getLocalStorage(ID_NAME) }}）</div>
    </div>
</template>
<script setup lang="ts">
import { getLocalStorage } from '@/libs/utils';
import { ID_NAME } from '@/libs/constant';
const props = defineProps({
    title: {
        type: String,
        required: true,
    }
})
</script>
<style lang="scss" scoped>
.title_box {
    display: flex;
    align-items: flex-end;
}
.title_box .title {
    position: relative;
    height: 46px;
    line-height: 40px;
    font-family: "TencentSansW7";
    font-size: 32px;
    font-weight: normal;
}

.title_box .title:before {
    display: block;
    position: absolute;
    left: 0;
    bottom: 0;
    z-index: 0;
    content: "";
    width: 128px;
    height: 12px;
    background: url(/src/assets/img/b.png) no-repeat center;
    background-size: cover;
    user-select: none;
}
.title_box .id_text {
    margin-left: 8px;
    color: rgba(0, 0, 0, 0.4);
    font-size: 14px;
    font-weight: 400;
    line-height: 22px;
}
</style>