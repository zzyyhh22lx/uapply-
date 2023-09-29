<template>
    <div class="resume_wrapper" id="cv_minimap" style="min-height: 218px">
    <div data-v-1f011456="" class="info_card">
        <div data-v-1f011456="" class="img default">
            <div data-v-1f011456="" class="default-photo"></div>
        </div>
        <div data-v-1f011456="" class="info_card_content">
            <div data-v-1f011456="" class="top"><div data-v-1f011456="" class="name">
                <span data-v-1f011456=""></span>林浩扬
            </div>
            <label data-v-1f011456="" class="el-checkbox is-checked">
                    <span class="el-checkbox__label">会当凌绝顶，一览众山小</span>
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
                            <div data-v-1f011456="" class="label">个人证件</div>
                            <div data-v-1f011456="" class="text">中国-居民身份证 - 4452242xxxxxxxxx</div></div>
                            <div data-v-1f011456="" class="bottom">
                                <div data-v-1f011456="" class="label">邮箱</div>
                                <div data-v-1f011456="" class="text">2744732798@qq.com</div>
                            </div></div>
                            <div data-v-1f011456="" class="item">
                                <div data-v-1f011456="" class="top">
                                    <div data-v-1f011456="" class="label">紧急联系人</div>
                                    <div data-v-1f011456="" class="text">-</div>
                                </div>
                                <div data-v-1f011456="" class="bottom">
                                    <div data-v-1f011456="" class="label">QQ</div>
                                    <div data-v-1f011456="" class="text">2744737298</div>
                                </div>
                            </div>
                            <div data-v-1f011456="" class="item">
                            <div data-v-1f011456="" class="top">
                                <div data-v-1f011456="" class="label">紧急联系人电话</div>
                                <div data-v-1f011456="" class="text">-</div>
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
import { onMounted } from 'vue';

function updateZoom() {
  const targetWidth = 1400;
  const currentWidth = window.innerWidth;
  const zoom = Math.ceil((currentWidth / targetWidth) * 10) / 10;
  const wrapper = document.querySelector("#cv_minimap");
  if (wrapper) {
    wrapper.style.zoom = zoom;
  }
}

// 初始化时调用 updateZoom 函数
onMounted(() => {
    updateZoom();
    // 监听窗口大小变化并更新缩放比例
    window.addEventListener("resize", updateZoom);
})

</script>
<style lang="scss" scoped>
.resume_wrapper {
    padding: 0 50px;
}
@media (max-width: 750px) {
    .resume_wrapper {
        display: none;
    }
}
.resume_wrapper .info_card {
    display: flex;
    width: 100%;
    height: 240px;
    margin-top: 38px;
    border-radius: 24px;
    overflow: hidden;
}
.resume_wrapper .info_card .default {
    background: url(/src/assets/img/transparent.png) no-repeat center;
    background-size: cover;
}
.resume_wrapper .info_card .img {
    width: 240px;
    height: 240px;
}
.resume_wrapper .info_card .default .default-photo {
    width: 130px;
    height: 154px;
    background: url(/src/assets/img/icons/user.svg) no-repeat center;
    margin: 40.5px 0 0 39px;
}
html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, dl, dt, dd, ol, ul, li, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video {
    margin: 0;
    padding: 0;
    border: 0;
    font-size: 100%;
    font: inherit;
    vertical-align: baseline;
}
.resume_wrapper .info_card .info_card_content {
    flex: 1;
}

.resume_wrapper .info_card .info_card_content > .top {
    display: flex;
    align-items: center;
    position: relative;
    height: 74px;
    background: #ffffff;
}
@media screen and (max-width: 1060px){
.resume_wrapper .info_card .info_card_content > .top:before {
    width: 20px;
}}
.resume_wrapper .info_card .info_card_content > .top:before {
    display: block;
    content: "";
    position: absolute;
    top: 0;
    left: -14px;
    width: 40px;
    height: 100%;
    background: #ffffff;
    transform: skewX(-6deg);
}
@media screen and (max-width: 1060px){
.resume_wrapper .info_card .info_card_content > .top .name {
    padding-left: 20px;
}}
.resume_wrapper .info_card .info_card_content > .top .name {
    box-sizing: border-box;
    padding-left: 70px;
    font-size: 26px;
    color: #0052D9;
}
.resume_wrapper .info_card .info_card_content > .top .el-checkbox {
    margin-left: auto;
    margin-right: 40px;
    color: rgba(0, 0, 0, 0.8);
    font-weight: 400;
}
.el-checkbox:last-of-type {
    margin-right: 0;
}
.el-checkbox {
    color: #606266;
    font-weight: 500;
    font-size: 14px;
    cursor: pointer;
    user-select: none;
    margin-right: 30px;
}
.el-checkbox, .el-checkbox__input {
    display: inline-block;
    position: relative;
    white-space: nowrap;
}
.el-button, .el-checkbox {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
}
.el-checkbox__input {
    cursor: pointer;
    outline: 0;
    line-height: 1;
    vertical-align: middle;
}
.el-checkbox, .el-checkbox__input {
    display: inline-block;
    position: relative;
    white-space: nowrap;
}
.resume_wrapper .info_card .info_card_content > .top .el-checkbox .el-checkbox__input.is-checked + .el-checkbox__label {
    color: rgba(0, 0, 0, 0.8);
}
.el-checkbox__input.is-checked+.el-checkbox__label {
    color: #409EFF;
}
.el-checkbox__label {
    display: inline-block;
    padding-left: 10px;
    line-height: 19px;
    font-size: 14px;
}
@media screen and (max-width: 1060px){
.resume_wrapper .info_card .info_card_content .infoList {
    padding: 24px 20px;
}}
.resume_wrapper .info_card .info_card_content .infoList {
    position: relative;
    display: flex;
    justify-content: space-between;
    height: 166px;
    box-sizing: border-box;
    padding: 24px 70px;
    background: #0052D9;
    line-height: 24px;
}
@media screen and (max-width: 1060px){
.resume_wrapper .info_card .info_card_content .infoList::before {
    width: 36px;
}}
.resume_wrapper .info_card .info_card_content .infoList:before {
    display: block;
    content: "";
    position: absolute;
    top: 0;
    left: -26px;
    width: 40px;
    height: 100%;
    background: #0052D9;
    transform: skewX(-6deg);
}
.resume_wrapper .info_card .info_card_content .infoList .label {
    color: rgba(255, 255, 255, 0.6);
    font-size: 14px;
    font-weight: 400;
    text-align: left;
    line-height: 22px;
    white-space: nowrap;
}
.resume_wrapper .info_card .info_card_content .infoList .bottom {
    margin-top: 22px;
}
.resume_wrapper .info_card .info_card_content .infoList .label {
    color: rgba(255, 255, 255, 0.6);
    font-size: 14px;
    font-weight: 400;
    text-align: left;
    line-height: 22px;
    white-space: nowrap;
}

.resume_wrapper .info_card .info_card_content .infoList .text {
    margin-top: 4px;
    white-space: nowrap;
    color: #ffffff;
    font-size: 16px;
    font-weight: 400;
    font-family: "PingFang SC";
    text-align: left;
    line-height: 24px;
}
</style>