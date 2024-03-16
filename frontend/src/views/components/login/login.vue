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
                                class="bg-[#f2f2f6]"
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
                                class="bg-[#f2f2f6]"
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
                                class="bg-[#f2f2f6]"
                                type="tel"
                                autocorrect="off"
                                autocapitalize="off"
                                autocomplete="off"
                                autofocus
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
                    <div class="module-register-op-item"><span>{{ statement }}即代表我已同意协议</span></div>
                </div>
                <div class="module-footer">
                    <div class="module-footer-item" data-spm-anchor-id="0.0.0.i1.75637435da9ArK" @click="toggleState">{{ state === LOGINSTATE ? '注册' : '登录' }}账号</div>
                </div>
            </div>
            <div class="app-cpright-bar"><div class="app-cpright">© 2023.Powered by IST & 青协</div></div>
        </div>
    </div>
</template>
<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage } from 'element-plus';
import { store } from '@/store/index';
import { Login, Register, getAccount } from '@/libs/request';
import { setLocalStorage } from '@/libs/utils';
import { TOKEN_NAME, ID_NAME } from '@/libs/constant';
import { LOGINSTATE, countryCodes } from './constant';
import type { LOGIN_RES } from '@/libs/request';
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
    try {
        if(state.value === LOGINSTATE) {
            console.log('登录态');
            res = await Login({ ...form.value });
        } else {
            console.log('注册态');
            res = await Register({ ...form.value });
        }
        // 将登录信息存于vuex全局状态管理中
        store.commit('setUsername', res.data.account);
        setLocalStorage(TOKEN_NAME, res.data.token as string);
        res = await getAccount();
        setLocalStorage(ID_NAME, String((res as LOGIN_RES).data.id));
        ElMessage.success(`${statement.value}成功`);
        // 路由跳转
        router.push('/home');
    } catch(e: any) {
        ElMessage.error(`${statement.value}失败 -- ${e.msg ? e.msg : e}`);
    }
}

</script> 

<style scoped lang="scss">
@import './index.scss';
.base-comp-mobile-box-select-value {
    position: relative;
}
.base-comp-mobile-box-select-value::after {
  content: "";
  position: absolute;
  top: 50%;
  right: 5px; /* 根据需要调整右侧距离 */
  width: 1px;
  height: 16px;
  background-color: rgba(126, 134, 142, 0.16);
  transform: translateY(-50%);
}
</style>