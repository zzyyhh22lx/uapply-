// tsx写法 无权限页面
import './index.scss';
import { defineComponent } from "vue";
export default defineComponent({
  name: "denied",
  setup(props, ctx) {
    return () => 
    <div class="lock-flex">
        <svg-icon name="lock" style="width: 120px; height: 121px;"/>
        <div class="text-xl font-medium m-3" style="color: #20202099;">您暂无权限访问</div>
    </div>;
  },
});