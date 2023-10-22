import App from "./App.vue";
import router from "./router";
import SvgIcon from '@/components/SvgIcon.vue';
import TopTitle from '@/components/TopTitle.vue';
import { createApp } from "vue";
import { getAccount } from '@/libs/request';
import type { LOGIN_RES } from '@/libs/request';
import { store } from "@/store/index";
import { EventEmitter } from "@/libs/eventbus";

import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'
import 'virtual:svg-icons-register';

const app = createApp(App);
// 事件总线
const eventBus = new EventEmitter();
app.config.globalProperties.$event = eventBus;

const InHome = () => {
    app.use(router).use(store).component('svg-icon', SvgIcon).component('top-title', TopTitle);
    router.isReady().then(() => app.mount("#app"));
}

/**
 * 根据是否返回account判断是否已登录
 */
(async () => {
   try {
        const res = await getAccount();
        if(typeof res !== 'boolean' && (res as LOGIN_RES).data?.account) {
            store.commit('setUsername', (res as LOGIN_RES).data.account);
        }
        InHome();
   } catch (e) {
        console.log(e);
        InHome();
   }
})()
