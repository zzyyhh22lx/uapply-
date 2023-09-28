import App from "./App.vue";
import router from "./router";
import SvgIcon from '@/components/SvgIcon.vue';
import { createApp } from "vue";
import { getAccount } from '@/libs/request';
import type { LOGIN_RES } from '@/libs/request';
import { store } from "@/store/index";

import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'
import 'virtual:svg-icons-register';

const app = createApp(App);

/**
 * 根据是否返回account判断是否已登录
 */
getAccount().then((res) => {
    if(typeof res !== 'boolean' && (res as LOGIN_RES).data?.account) {
        store.commit('setUsername', (res as LOGIN_RES).data.account);
    }
    app.use(router).use(store).component('svg-icon', SvgIcon);
    router.isReady().then(() => app.mount("#app"));
})
