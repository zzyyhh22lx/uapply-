<template>
    <div class="login-page">
        <button @click="toggleState">切换{{ state === LOGINSTATE ? '注册' : '登录' }}</button>
        <form v-if="state === LOGINSTATE">
            账号 <input v-model="form.username">
            密码 <input v-model="form.password">
        </form>
        <form v-else>
            手机号 <input v-model="form.phoneNumber">
            账号 <input v-model="form.username">
            密码 <input v-model="form.password">
        </form>
        <button @click="loginIn">{{ state === LOGINSTATE ? '登录' : '注册' }}</button>
    </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { store } from '@/store/index';
/**表单类型*/
type FORM_TYPE = {
    phoneNumber: String,
    username: String,
    password: String,
}
const router = useRouter();

/**登录态*/
const LOGINSTATE = false;
/**注册态*/
const REGISTERSTATE = true;

/** 状态 --> 用于判断登录注册*/
const state = ref(LOGINSTATE);
/**表单数据*/
const form = ref<FORM_TYPE>({
    phoneNumber: '',
    username: '',
    password: '',
});

/***切换状态 */
const toggleState = () => {
    state.value = !state.value;
}
/**登录*/
const loginIn = () => {
    if(state.value === LOGINSTATE) {
        console.log('登录态');
    } else {
        console.log('注册态');
    }
    console.log(form.value);
    // 将登录信息存于vuex全局状态管理中
    store.commit('setUsername', form.value.username);
    // 路由跳转
    router.push('/home');
}

</script> 

<style scoped lang="scss">

</style>