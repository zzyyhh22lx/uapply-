<template>
    <div class="content__warpper">
        <div class="content__total">总共2个岗位</div>
        <div class="content__checked">
            <div class="content__checked--items">
                <span class="lable">
                    <span>前端开发</span>
                    <el-icon class="pointer">
                        <CloseBold />
                    </el-icon>
                </span>
                <span class="lable">
                    <span>后端开发</span>
                    <el-icon class="pointer">
                        <CloseBold />
                    </el-icon>
                </span>
            </div>
            <button class="content__checked--delete">清除全部</button>
        </div>

        <div class="content__content">
            <!-- 岗位介绍小豆腐块 -->
            <div class="offer__warpper">
                <div class="offer__title">前端开发方向</div>
                <div class="offer__lable">
                    <span>技术类</span>
                    <span>|</span>
                    <span>社团招新</span>
                </div>
                <div class="offer__location">桂电四创4102</div>
            </div>
            <!-- 岗位介绍小豆腐块 -->

            <!-- 岗位介绍小豆腐块 -->
            <div class="offer__warpper">
                <div class="offer__title">后端开发方向</div>
                <div class="offer__lable">
                    <span>技术类</span>
                    <span>|</span>
                    <span>社团招新</span>
                </div>
                <div class="offer__location">桂电四创4102</div>
            </div>
            <!-- 岗位介绍小豆腐块 -->
        </div>

        <el-pagination layout="prev, pager, next" :total="1000" />
    </div>
</template>

<script setup lang="js">

</script>

<style scoped lang="scss">
.pointer {
    cursor: pointer;
}

.content__warpper {
    display: flex;
    flex-direction: column;
    width: 56.5vw;
    background-color: #e8e8e8;
    padding: 0 5%;
    border-radius: 15px;

    >* {
        margin: 10px 0;
    }

    .content__total {
        margin-top: 22px;
        font-size: 16px;
    }

    .content__checked {
        display: flex;
        flex-direction: row;
        white-space: nowrap;
        justify-content: space-between;

        .content__checked--items {
            display: flex;
            flex-direction: row;
            align-items: center;
            white-space: nowrap;

            >* {
                padding: 0 6%
            }

            .lable {
                // padding: 10px 20px;
                display: flex;
                align-items: center;
                justify-content: space-between;
                background-color: #ffffff;
                margin-right: 10px;
                padding: 10px 8px 10px 15px;
                border-radius: 15px;

                span {
                    font-size: 16px;
                    margin-right: 5px;
                }

                el-icon {
                    cursor: pointer;
                }
            }
        }

        .content__checked--delete {
            padding-right: 3%;
            color: #0052D9;
        }
    }

    .content__content {
        >* {
            margin-bottom: 30px;
        }

        .offer__warpper {
            padding: 3px 4%;
            background-color: #ffffff;
            box-shadow: 0px 4px 60px 0px rgba(0, 0, 0, 0.04);
            border-radius: 15px;

            >* {
                padding: 8px 0;
            }

            .offer__title {
                padding-top: 20px;
                font-size: 24px;
            }

            .offer__lable {
                >* {
                    margin-right: 1%;
                    font-size: 18px;
                }
            }

            .offer__location{
                font-size: 14px;
            }
        }
    }

    .el-pagination {
        flex-direction: row-reverse;
    }
}
</style>