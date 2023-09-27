// 配置文件：在导入任何以.vue结尾的文件时，将其视为Vue组件
declare module "*.vue" {
  import { defineComponent } from "vue";
  const component: ReturnType<typeof defineComponent>;
  export default component;
}