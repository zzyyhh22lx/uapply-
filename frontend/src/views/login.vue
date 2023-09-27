<template>
    <div>
        用户名：<el-input v-model="form.username"/>
    </div>
    <div>
        密码：<el-input  v-model="form.password"/>
    </div>
    <el-button @click="loginIn">登录</el-button>
</template>
<script lang="ts" setup>
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