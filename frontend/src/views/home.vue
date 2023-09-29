<template>
    <div class="apply-home">
        <uapply-header ref="header"/>
        <div class="userhome" ref="userhome">
            <router-view></router-view>
        </div>
        <el-backtop />
        <div class="app-cpright-bar"><div class="app-cpright">© 2023.Powered by IST & 青协</div></div>
    </div>
</template>
<script setup lang="ts">
import UapplyHeader from './components/header/index.vue';
import { ElBacktop } from 'element-plus';
import { getCurrentInstance, ref, onMounted, onBeforeUnmount } from 'vue';
const vm = getCurrentInstance();
const userhome = ref();
vm?.proxy?.$event.on('clickCheckBtn', (isShow: boolean) => {
    isShow ? userhome.value.classList.add('opacity-0') : userhome.value.classList.remove('opacity-0');
});

const header = ref();
let lastScrollTop = 0;
const handleScroll = () => {
    const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    if (scrollTop > lastScrollTop) {
        // 向下滚动
        header.value?.setHeight(true);
    } else {
        // 向上滚动
        header.value?.setHeight(false);
    }
    lastScrollTop = scrollTop;
}
onMounted(() => {
    window.addEventListener("scroll", handleScroll);
})
onBeforeUnmount(() => {
    window.removeEventListener("scroll", handleScroll);
})
</script>
<style scoped lang="scss">
.apply-home {
    @apply bg-[#f5f5f5];
    padding-bottom: 70px;
}
.userhome { 
    @apply mt-16;
    transition: all .5s;
}
.app-cpright-bar {
    position: relative;
    width: 480px;
    height: 60px;
    bottom: 0;
    left: 50%;
    margin-left: -240px;
    line-height: 60px;
    font-size: 12px;
    color: rgba(23, 26, 29, 0.6);
    text-align: left;
    user-select:none;
    transform: scale(0.8);
    text-align: center;
}
</style>