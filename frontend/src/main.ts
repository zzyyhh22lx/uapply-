import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'
import 'virtual:svg-icons-register';
import SvgIcon from '@/components/SvgIcon.vue';

const app = createApp(App);
app.use(router).component('svg-icon', SvgIcon).mount("#app");
