import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import './assets/apply.scss';
import "./assets/tailwind.css";
import 'element-plus/dist/index.css'

const app = createApp(App);
app.use(router).mount("#app");
