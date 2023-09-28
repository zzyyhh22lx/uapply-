<template>
    <div class="login-page">
        <button @click="toggleState">切换{{ state === LOGINSTATE ? '注册' : '登录' }}</button>
        <form v-if="state === LOGINSTATE">
            账号 <input v-model="form.account">
            密码 <input v-model="form.password">
        </form>
        <form v-else>
            手机号 <input v-model="form.phoneNumber">
            账号 <input v-model="form.account">
            密码 <input v-model="form.password">
        </form>
        <button @click="loginIn">{{ state === LOGINSTATE ? '登录' : '注册' }}</button>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { store } from '@/store/index';
import { Login, Register } from '@/libs/request';
import { setLocalStorage } from '@/libs/utils';
import { TOKEN_NAME } from '@/libs/constant';
import { LOGINSTATE, REGISTERSTATE } from './constant';
import type FORM_TYPE from './constant';

const router = useRouter();

/** 状态 --> 用于判断登录注册*/
const state = ref(LOGINSTATE);
/**表单数据*/
const form = ref<FORM_TYPE>({
    phoneNumber: '',
    account: '',
    password: '',
});

/** 将数据置空 */
const initForm = () => {
    const data = form.value;
    Object.keys(data).forEach(element => {
        data[element] = '';
    });
}
/***切换状态 */
const toggleState = () => {
    state.value = !state.value;
    initForm();
}
/**登录*/
const loginIn = async () => {
    let res;
    if(state.value === LOGINSTATE) {
        console.log('登录态');
        res = await Login(form.value);
    } else {
        console.log('注册态');
        res = await Register(form.value);
    }
    // 将登录信息存于vuex全局状态管理中
    store.commit('setUsername', res.data.account);
    setLocalStorage(TOKEN_NAME, res.data.token);
    // 路由跳转
    router.push('/home');
}

</script> 

<style scoped lang="scss">

</style>