<template>
    <div class="resume_wrapper" style="min-height: 218px">
        <div data-v-1f011456="" class="top">
            <div data-v-1f011456="" class="title_box">
                <div data-v-1f011456="" class="title">
                    <span data-v-1f011456="">我的简历</span>
                </div>
                <div data-v-1f011456="" class="id_text">（ID：2643428）</div> <!---->
            </div>
            <div data-v-1f011456="" class="tip_line">
                <div data-v-1f011456="" class="tip_text">最近更新时间：2023-03-29</div>
                <div data-v-1f011456="" class="progress_box">
                    <div data-v-1f011456="" class="el-tooltip tip_icon_question" aria-describedby="el-tooltip-678" tabindex="0"></div>
                    <span data-v-1f011456="" class="text">简历完善程度：</span>
                    <div data-v-1f011456="" role="progressbar" aria-valuenow="62" aria-valuemin="0" aria-valuemax="100" class="el-progress el-progress--line">
                        <div class="el-progress-bar">
                            <div class="el-progress-bar__outer" style="height: 6px; background-color: rgb(235, 238, 245);">
                                <div class="el-progress-bar__inner" style="width: 62%;"><!----></div>
                            </div>
                        </div>
                        <div class="el-progress__text" style="font-size: 14.4px; color: rgb(96, 98, 102);">62%</div>
                    </div>
                </div>
            </div>
        </div>
        <div data-v-1f011456="" class="info_card"  id="cv_minimap">
            <div data-v-1f011456="" class="img default">
                <div data-v-1f011456="" class="default-photo"></div>
            </div>
            <div data-v-1f011456="" class="info_card_content">
                <div data-v-1f011456="" class="top"><div data-v-1f011456="" class="name">
                    <span data-v-1f011456=""></span>林浩扬
                </div>
                <label data-v-1f011456="" class="el-checkbox is-checked">
                        <span class="el-checkbox__label">欢迎加入IST & 青协</span>
                </label>
            </div>
            <div data-v-1f011456="" class="infoList">
                <div data-v-1f011456="" class="item">
                    <div data-v-1f011456="" class="top">
                        <div data-v-1f011456="" class="label">性别</div>
                        <div data-v-1f011456="" class="text">男</div>
                    </div>
                    <div data-v-1f011456="" class="bottom">
                        <div data-v-1f011456="" class="label">手机号</div>
                        <div data-v-1f011456="" class="text">+86 17266677788</div>
                    </div></div>
                    <div data-v-1f011456="" class="item">
                        <div data-v-1f011456="" class="top">
                            <div data-v-1f011456="" class="label">年龄</div>
                            <div data-v-1f011456="" class="text">19</div></div>
                            <div data-v-1f011456="" class="bottom">
                                <div data-v-1f011456="" class="label">邮箱</div>
                                <div data-v-1f011456="" class="text">2744732798@qq.com</div>
                            </div></div>
                            <div data-v-1f011456="" class="item">
                                <div data-v-1f011456="" class="top">
                                    <div data-v-1f011456="" class="label">专业</div>
                                    <div data-v-1f011456="" class="text">软件工程</div>
                                </div>
                                <div data-v-1f011456="" class="bottom">
                                    <div data-v-1f011456="" class="label">QQ</div>
                                    <div data-v-1f011456="" class="text">2744737298</div>
                                </div>
                            </div>
                            <div data-v-1f011456="" class="item">
                            <div data-v-1f011456="" class="top">
                                <div data-v-1f011456="" class="label">感兴趣的方向</div>
                                <div data-v-1f011456="" class="text">后端 or 前端 儿</div>
                            </div>
                            <div data-v-1f011456="" class="bottom">
                                <div data-v-1f011456="" class="label">微信</div>
                                <div data-v-1f011456="" class="text">_lhYxz_</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    </div> 
</template>
<script lang="ts" setup>
import { onMounted, onBeforeUnmount, getCurrentInstance } from 'vue';
import { store } from '@/store/index';
const vm = getCurrentInstance();

/**
 * 缩放
 */
function updateZoom() {
  const targetWidth = 1400;
  const currentWidth = window.innerWidth;
  let zoom = Math.ceil((currentWidth / targetWidth) * 10) / 10;
  zoom = zoom >= 1 ? 1 : zoom;
  const wrapper = document.querySelector("#cv_minimap");
  if (wrapper) {
    (wrapper as HTMLElement).style.zoom = zoom;
  }
  // 显示
  if(currentWidth > 750) {
    vm?.proxy?.$event.emit('clickCheckBtn', false);
  } else {
    vm?.proxy?.$event.emit('clickCheckBtn', store.state.isShow);
  }
}

// 初始化时调用 updateZoom 函数
onMounted(() => {
    updateZoom();
    // 监听窗口大小变化并更新缩放比例
    window.addEventListener("resize", updateZoom);
})
// 卸载
onBeforeUnmount(() => {
    window.removeEventListener("resize", updateZoom);
})
</script>
<style lang="scss" scoped>
@import './index.scss';
</style>