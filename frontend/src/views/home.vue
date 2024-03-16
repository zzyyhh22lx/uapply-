<template>
    <div class="apply-home">
        <uapply-header ref="header"/>
        <div class="userhome" ref="userhome">
            <router-view></router-view>
        </div>
        <div class="app-cpright-bar"><div class="app-cpright">© 2023.Powered by IST & 青协</div></div>
        <el-backtop :right="64" :bottom="100" />
        <div style="background-color: transparent; height: 40px;"></div>
    </div>
</template>
<script setup lang="ts">
import UapplyHeader from './components/header/index.vue';
import { ElBacktop } from 'element-plus';
import { getCurrentInstance, ref, onMounted, onBeforeUnmount } from 'vue';
import { store } from '@/store/index';

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

const resizeShow = () => {
  const currentWidth = window.innerWidth;
  // 显示
  if(currentWidth > 750) {
    vm?.proxy?.$event.emit('clickCheckBtn', false);
  } else {
    vm?.proxy?.$event.emit('clickCheckBtn', store.state.isShow);
  }
}

onMounted(() => {
    window.addEventListener("scroll", handleScroll);
    window.addEventListener("resize", resizeShow);
})
onBeforeUnmount(() => {
    window.removeEventListener("scroll", handleScroll);
    window.removeEventListener("resize", resizeShow);
})
</script>
<style scoped lang="scss">
.apply-home {
    @apply bg-[#f5f5f5];
    padding-bottom: 70px;
    width: 100%;
    overflow-x: hidden;
}
.userhome { 
    @apply mt-16;
    transition: all .5s;
}
@media (max-width: 750px) {
    .userhome {
        @apply mt-6;
    }
}
.app-cpright-bar {
    position: relative;
    width: 480px;
    height: 60px;
    bottom: -100px;
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