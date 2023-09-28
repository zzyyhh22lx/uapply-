<template>
    <div class="login-warpper">
        登录注册
    </div>
</template>
<script setup>
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
</style>