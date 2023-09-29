<template>
    <div class="uapply-box">
        <cv-minimap></cv-minimap>
        <div class="resume_wrapper">
            <div class="detail_content">
                <div class="info_item">
                    <div class="title bold">具体信息</div>
                    <div class="list">
                        <div class="item item_half">
                            <div class="label">姓名</div>
                            <div v-if="!status" class="text">林浩扬</div>
                            <el-input placeholder="Please input" v-else/>
                        </div>
                        <div class="item item_half">
                            <div class="label">性别</div>
                            <div class="text">男</div>
                        </div>
                        <div class="item item_half">
                            <div class="label">年龄</div>
                            <div v-if="!status"  class="text">19</div>
                            <el-input placeholder="Please input" v-else/>
                        </div>
                        <div class="item item_half">
                            <div class="label">手机号</div>
                            <div v-if="!status" class="text">+86 17266677788</div>
                            <el-input placeholder="Please input" v-else/>
                        </div>
                        <div class="item item_half">
                            <div class="label">邮箱</div>
                            <div v-if="!status" class="text">2744732798@qq.com</div>
                            <el-input placeholder="Please input" v-else/>
                        </div>
                        
                        <div class="item item_half">
                            <div class="label">感兴趣的方向</div> 
                            <div class="text">后端 or 前端 儿</div>
                        </div>
                        <div class="item item_half">
                            <div class="label">投递岗位</div>
                            <div class="text">软件开发-前端开发方向</div>
                        </div>
                        <div class="item item_half">
                            <div class="label">感兴趣的部门</div>
                            <div class="text">IST软件部</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="fixed_box isFixed myfooter" ref="myfooter" id="footer">
                <div class="center_box">
                    <template v-if="!status">
                        <div class="resume_btn primary mr-12" @click="edit">编辑</div>
                        <div class="resume_btn disabled">保存</div>
                    </template>
                    <template v-else>
                        <div class="resume_btn primary mr-12" @click="cancel">取消</div>
                        <div class="resume_btn" @click="save">保存</div>
                    </template>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
import CvMinimap from './CvMinimap.vue';
import { onMounted, onBeforeUnmount, ref } from 'vue';
import { ElInput } from 'element-plus';

/** 编辑态 */
const status = ref(false);
const edit = () => {
    status.value = true;
}

const cancel = () => {
    status.value = false;
}
const save = () => {
    
}

const myfooter = ref();
const handleScroll = () => {
    const footer = myfooter.value;
    const windowHeight = window.innerHeight;
    const scrollHeight = document.documentElement.scrollHeight;
    const scrollTop =
        window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;
    if (scrollHeight - (scrollTop + windowHeight) < 80) {
        // 距离底部小于 100px
        footer?.classList.add('absolute__myfooter');
        footer?.classList.remove('fixed__myfooter');
    } else {
        footer?.classList.add('fixed__myfooter');
        footer?.classList.remove('absolute__myfooter');
    }
}
onMounted(() => {
    handleScroll();
    window.addEventListener("scroll", handleScroll);
})
onBeforeUnmount(() => {
    window.removeEventListener('scroll', handleScroll);
})
</script>
<style lang="scss" scoped>
@import './usercard.scss';
.uapply-box {
    @apply m-4 px-12;
    position: relative;
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

.resume_btn:hover {
    opacity: 0.8;
}
.disabled {
    color: #999999!important;
    border: 1px solid #cccccc!important;
    cursor:  not-allowed!important;
}
.disabled:hover {
    opacity: 1!important;
}
.myfooter {
    @apply px-16;
}
.fixed__myfooter {
    position: fixed !important;
    bottom: 0!important;
    padding-left: 64px;
    padding-right: 64px;
}
.absolute__myfooter {
    position: absolute!important;
    bottom: -130px!important;
    padding-left: 3rem;
    padding-right: 3rem;
}
@media (max-width: 750px) {
    .uapply-box {
        padding: 0;
    }
    .myfooter {
        @apply px-4;
    }
    .fixed__myfooter {
        @apply px-4;
    }
    .absolute__myfooter {
        @apply px-0;
    }
}
</style>