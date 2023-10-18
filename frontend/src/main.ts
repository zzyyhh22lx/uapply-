import App from "./App.vue";
import router from "./router";
import SvgIcon from '@/components/SvgIcon.vue';
import { createApp } from "vue";
import { getAccount } from '@/libs/request';
import type { LOGIN_RES } from '@/libs/request';
import { store } from "@/store/index";
import { EventEmitter } from "@/libs/eventbus";
import ElementPlus from "element-plus"

import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'
import 'virtual:svg-icons-register';
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

const app = createApp(App);

app.use(ElementPlus)

//使用element-plus icon库
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// 事件总线
const eventBus = new EventEmitter();
app.config.globalProperties.$event = eventBus;




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





