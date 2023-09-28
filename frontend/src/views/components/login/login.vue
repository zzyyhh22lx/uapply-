<template>
    <div class="app-page-bannerview page-wrapper">
        <div class="banner-area" style="background-image: url(/src/assets/img/bg.png);">
            <div class="banner-link" data-spm-anchor-id="0.0.0.i1.756374356BbmAJ"></div>
        </div>
        <div class="app-area">
            <div class="app-box">
                <div class="module-comm-page-title">{{ statement }}账号</div>
                <template v-if="state === LOGINSTATE">
                    <div class="module-register-input-box base-comp-mobile-box base-comp-input-box">
                        <div class="base-comp-mobile-box-input myinput">
                            <input
                                autocorrect="off"
                                autocapitalize="off"
                                autocomplete="off"
                                maxlength="24"
                                placeholder="请输入账号"
                                v-model="form.account"
                            />
                        </div>
                    </div>
                    <div class="module-register-input-box base-comp-mobile-box base-comp-input-box mybox">
                        <div class="base-comp-mobile-box-input myinput">
                            <input
                                autocorrect="off"
                                autocapitalize="off"
                                autocomplete="off"
                                maxlength="24"
                                placeholder="请输入密码"
                                v-model="form.password"
                            />
                        </div>
                    </div>
                </template>
                <div class="module-register-input-box base-comp-mobile-box base-comp-input-box" v-else>
                    <div class="base-comp-mobile-box-select">
                        <div class="base-comp-mobile-box-select-value">
                            <span class="dingtalk-login-iconfont dingtalk-login-icon-turn-down">+86</span>
                        </div>
                        <select>
                            <option v-for="code in countryCodes" :key="code" :value="'+' + code">
                                +{{ code }}
                            </option>
                        </select>
                        <div class="base-comp-mobile-box-input">
                            <input
                                type="tel"
                                autocorrect="off"
                                autocapitalize="off"
                                autocomplete="off"
                                maxlength="12"
                                placeholder="请输入手机号"
                                v-model="form.phoneNumber"
                                data-spm-anchor-id="0.0.0.i3.75637435da9ArK"
                            />
                        </div>
                    </div>
                </div>
               
                <div class="base-comp-button base-comp-button-type-primary base-comp-button-disabled" @click="loginIn">确认</div>
                <div class="module-register-op-line">
                    <div class="module-register-op-item"><span>{{ statement }}即代表我已阅读并同意协议</span></div>
                </div>
                <div class="module-footer">
                    <div class="module-footer-item" data-spm-anchor-id="0.0.0.i1.75637435da9ArK" @click="toggleState">{{ statement }}账号</div>
                </div>
            </div>
            <div class="app-cpright-bar"><div class="app-cpright">© 2023.Powered by IST & 青协</div></div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { store } from '@/store/index';
import { Login, Register } from '@/libs/request';
import { setLocalStorage } from '@/libs/utils';
import { TOKEN_NAME } from '@/libs/constant';
import { LOGINSTATE, countryCodes } from './constant';
import type { FORM_TYPE } from './constant';

const router = useRouter();

/** 状态 --> 用于判断登录注册*/
const state = ref(LOGINSTATE);
const statement = computed(()=>{
  return state.value === LOGINSTATE ? '登录' : '注册'
});
/**表单数据*/
const form = ref<FORM_TYPE>({
    phoneNumber: '',
    account: '',
    password: '',
});

/** 将数据置空 */
const initForm = () => {
  const data = form.value;
  for (const key in data) {
    data[key as keyof FORM_TYPE] = "";
  }
};

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
        res = await Login({ ...form.value });
    } else {
        console.log('注册态');
        res = await Register({ ...form.value });
    }
    // 将登录信息存于vuex全局状态管理中
    store.commit('setUsername', res.data.account);
    setLocalStorage(TOKEN_NAME, res.data.token);
    // 路由跳转
    router.push('/home');
}

</script> 

<style scoped lang="scss">
@import './index.scss';

</style>