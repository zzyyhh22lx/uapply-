// 配置文件：在导入任何以.vue结尾的文件时，将其视为Vue组件
declare module '*.vue' {
    import { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}
