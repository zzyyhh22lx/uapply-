import App from "./App.vue";
import router from "./router";
import SvgIcon from '@/components/SvgIcon.vue';
import { createApp } from "vue";

import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'
import 'virtual:svg-icons-register';

const app = createApp(App);
app.use(router).component('svg-icon', SvgIcon).mount("#app");
