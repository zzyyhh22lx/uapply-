<template>
    <div class="login-warpper">
        <div class="login-box">
            <div class="login-selector">
                <el-button class="no-border" id="status-login">登录</el-button>
                <el-button class="no-border" id="status-regist">注册</el-button>
            </div>
            <div class="login-name">
                <p class="margin">用户名：</p>
                <el-input v-model="form.username" class="margin inline"/>
            </div>
            <div class="login-psw">
                <p class="margin">密码：</p>
                <el-input  class="margin" v-model="form.password" type="password" show-password/>
            </div>
            <a href="" >忘记密码？</a>
            <br>
            <div class="login-submit__warpper">
                <el-button @click="loginIn" class="login-submit margin">登录</el-button>
            </div>
        </div>
    </div>
</template>
<script setup>

import { ElInput, ElButton } from 'element-plus';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { store } from '@/store/index';

const router = useRouter();
/**
 * 表单数据
 */
const form = ref({
    username: '',
    password: '',
});

/**
 * 登录
 */
const loginIn = () => {
    store.commit('setUsername', form.value.username);
    console.log(form.value, '登录成功');
    router.push('/home');
}

</script> 

<style scoped>
.login-warpper{
    display: flex;    
    justify-content: center;
    align-items: center;
}

.login-box{
    margin-top: 25vh;
    padding: 30px;
    width: 600px; 
    border: 1px solid #ccc;
    border-radius: 5px;
}


.login-selector{
    display: flex;
    justify-content: center;
}

.login-submit__warpper{
    display: flex;
    justify-content: center;
}

.login-submit{
    width: 30%;
}

#status-login{
    border-right: solid 1px #ccc;
}
.no-border{
    border: 0px;
    margin: 0px;
}

.margin{
    margin-top: 10px;
}

</style>