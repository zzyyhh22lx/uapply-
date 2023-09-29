<template>
    <div class="uapply-box">
        <cv-minimap></cv-minimap>
        <div class="resume_wrapper">
            <div class="detail_content">
                <div data-v-1f011456="" class="info_item">
                    <div data-v-1f011456="" class="title bold">具体信息</div>
                    <div data-v-1f011456="" class="list">
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">姓名</div>
                            <div data-v-1f011456="" class="text">林浩扬</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">性别</div>
                            <div data-v-1f011456="" class="text">男</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">年龄</div>
                            <div data-v-1f011456="" class="text">19</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">手机号</div>
                            <div data-v-1f011456="" class="text">+86 17266677788</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">邮箱</div>
                            <div data-v-1f011456="" class="text">2744732798@qq.com</div>
                        </div>
                        
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">感兴趣的方向</div> 
                            <div data-v-1f011456="" class="text">后端 or 前端 儿</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">投递岗位</div>
                            <div data-v-1f011456="" class="text">软件开发-前端开发方向</div>
                        </div>
                        <div data-v-1f011456="" class="item item_half">
                            <div data-v-1f011456="" class="label">感兴趣的部门</div>
                            <div data-v-1f011456="" class="text">IST软件部</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import CvMinimap from './CvMinimap.vue';
</script>
<style lang="scss" scoped>
@import './usercard.scss';
.uapply-box {
    @apply m-4 px-12;
    position: relative;
}
@media (max-width: 750px) {
    .uapply-box {
        padding: 0;
    }
}
.detail_content {
    position: relative;
    margin-top: 40px;
    margin-bottom: 40px;
    box-sizing: border-box;
    padding: 78px 64px 0 88px;
    border-radius: 24px;
    background: #ffffff;
}
</style>