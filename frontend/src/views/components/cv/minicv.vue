<template>
    <div class="resume_wrapper">
        <div class="top">
            <div class="title_box">
                <div class="title">
                    <span>我的简历</span>
                </div>
                <div class="id_text">（ID：{{ getLocalStorage(ID_NAME) }}）</div> <!---->
            </div>
            <div class="tip_line">
                <div class="tip_text">最近更新时间：2023-03-29</div>
                <div class="progress_box">
                    <div class="el-tooltip tip_icon_question" aria-describedby="el-tooltip-678" tabindex="0"></div>
                    <span class="text">简历完善程度：</span>
                    <div role="progressbar" aria-valuenow="62" aria-valuemin="0" aria-valuemax="100" class="el-progress el-progress--line">
                        <div class="el-progress-bar">
                            <div class="el-progress-bar__outer" style="height: 6px; background-color: rgb(235, 238, 245);">
                                <div class="el-progress-bar__inner" :style="{ width: `${proportion}%` }"><!----></div>
                            </div>
                        </div>
                        <div class="el-progress__text" style="font-size: 14.4px; color: rgb(96, 98, 102);">{{ proportion }}%</div>
                    </div>
                </div>
            </div>
        </div>
        <div class="info_card"  id="cv_minimap">
            <div class="img default">
                <div class="default-photo"></div>
            </div>
            <div class="info_card_content">
                <div class="top">
                    <div class="name">
                        <span></span>{{ form.name }}
                    </div>
                    <label class="el-checkbox is-checked">
                            <span class="el-checkbox__label">欢迎加入IST & 青协</span>
                    </label>
                </div>
                <div class="infoList">
                    <div class="item">
                        <div class="top">
                            <div class="label">性别</div>
                            <div class="text">{{ typeof form.sex === 'number' ? form.sex === 1 ? '男' : '女' : '未填' }}</div>
                        </div>
                        <div class="bottom">
                            <div class="label">手机号</div>
                            <div class="text">{{ form.phone || '未填' }}</div>
                        </div>
                    </div>
                    <div class="item">
                        <div class="top">
                            <div class="label">年龄</div>
                            <div class="text">{{ form.age || '未填' }}</div>
                        </div>
                        <div class="bottom">
                            <div class="label">邮箱</div>
                            <div class="text">{{ form.email || '未填' }}</div>
                        </div>
                    </div>
                        <div class="item">
                            <div class="top">
                                <div class="label">专业</div>
                                <div class="text">{{ form.major || '未填' }}</div>
                            </div>
                            <div class="bottom">
                                <div class="label">QQ</div>
                                <div class="text">{{ form.qq || '未填' }}</div>
                            </div>
                        </div>
                        <div class="item">
                            <div class="top">
                                <div class="label">感兴趣的方向</div>
                                <div class="text">{{ form.interest || '未填' }}</div>
                            </div>
                            <div class="bottom">
                                <div class="label">微信</div>
                                <div class="text">{{ form.wc || '未填' }}</div>
                            </div>
                        </div>
                </div>
            </div>
        </div>
    </div> 
</template>
<script lang="ts" setup>
import { onMounted, onBeforeUnmount, PropType, ref } from 'vue';
import type { CV_TYPE } from "@/libs/request";
import { getLocalStorage } from '@/libs/utils';
import { ID_NAME } from '@/libs/constant';
const props = defineProps({
    form: {
        type: Object as PropType<CV_TYPE>,
        required: true,
    }
})

const proportion = ref(0);
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
}

// 初始化时调用 updateZoom 函数
onMounted(() => {
    updateZoom();
    // 监听窗口大小变化并更新缩放比例
    window.addEventListener("resize", updateZoom);
    // 计算简历完善比例
    let i = 0, j = 0;
    Object.keys(props.form).forEach((item: string) => {
        const key = item as keyof CV_TYPE;
        i++;
        if (props.form[key]) j++;
    });
    proportion.value = Math.floor((j / i) * 100);
})
// 卸载
onBeforeUnmount(() => {
    window.removeEventListener("resize", updateZoom);
})
</script>
<style lang="scss" scoped>
@import './minicv.scss';
</style>