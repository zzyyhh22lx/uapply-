<template>
    <div class="aside__warpper">
        <div class="aside__search">
            <el-input v-model="input2" class="w-50 m-2" placeholder="搜索岗位" :prefix-icon="Search"
                style="margin: 0; padding: 0;" />
            <button class="aside__more">查看</button>
        </div>
        <div class="aside__categories">
            <!-- 侧边栏分类 岗位类别 -->
            <div class="aside__categories--position">
                <div class="item__nav">
                    <div class="item__nav--catogrory">岗位类别</div>
                    <button @click="ShowPosition = !ShowPosition">
                        <el-icon>
                            <ArrowLeft />
                        </el-icon>
                    </button>
                </div>
                <div v-show="ShowPosition" class="item__content">
                    <el-checkbox lable="前端开发">前端开发</el-checkbox>
                    <el-checkbox lable="安卓开发">安卓开发</el-checkbox>
                    <el-checkbox lable="小程序开发">小程序开发</el-checkbox>
                </div>
            </div>
            <!-- 侧边栏分类 岗位类别 -->

            <!-- 侧边栏分类 社团地点 -->
            <div class="aside__categories--location">
                <div class="item__nav">
                    <div class="item__nav--catogrory">社团地点</div>
                    <button @click="ShowLocation = !ShowLocation">
                        <el-icon>
                            <ArrowLeft />
                        </el-icon>
                    </button>
                </div>
                <div v-show="ShowLocation" class="item__content">
                    <el-checkbox lable="四创">四创</el-checkbox>
                    <el-checkbox lable="五教">五教</el-checkbox>
                    <el-checkbox lable="观音山">观音山</el-checkbox>
                </div>
            </div>
            <!-- 侧边栏分类 社团地点 -->
        </div>
    </div>
</template>

<script setup lang="js">
import { ref } from 'vue'
import { Calendar, Search } from '@element-plus/icons-vue'
var ShowPosition = ref(false);
var ShowLocation = ref(false);
</script>

<style scoped lang="scss">
.aside__warpper {
    display: flex;
    flex-direction: column;
    background-color: #FFFF;
    border-radius: 10px;
    // border: solid 5px #ccc;
    width: 16vw;

    >* {
        padding: 20px 5.5% 0px;
    }

    .aside__search {
        display: flex;
        align-items: center;
        justify-content: space-between;

        .aside__more {
            color: #0052D9;
            white-space: nowrap;
            padding-left: 5%;
            font-size: 14px;
            cursor: pointer;
        }
    }

    .aside__categories {
        padding-left: 10%;

        >* {
            margin-bottom: 20px;

            .item__nav {
                display: flex;
                flex-direction: row;
                justify-content: space-between;

                .item__nav--catogrory {
                    font-size: 18px;
                }
            }

            .item__content {
                padding-left: 1%;
                display: flex;
                flex-direction: column;
                justify-content: space-between;
            }
        }
    }
}
</style>