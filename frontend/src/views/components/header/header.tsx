import { defineComponent, ref, onMounted } from "vue";
import { useRouter, useRoute } from 'vue-router';
import { store } from '@/store/index';

import "./index.scss";

export default defineComponent({
  name: "ApplyHeader",
  setup() {
    const router = useRouter();
    const route = useRoute();
    const navTitles = [
      {
        name: "recommand",
        label: "为你推荐",
        index: 0,
        disabled: true
      },
      {
        name: 'uapply',
        label: "招新报名",
        index: 1,
      },
      {
        name: "view-cv",
        label: "我的简历",
        index: 2,
      }, 
      {
        name: "view-pc",
        label: "流程查看",
        index: 3,
      }];
    const activeIndex = ref(1);
    const setActive = (index: number, name: string) => {
      activeIndex.value = index;
      // 触发点击label事件
      document.getElementById('uapply-checkbtn')?.click();
      router.push({
        name
      })
    };
    onMounted(() => {
      activeIndex.value = navTitles.find((item, index) => item.name === route.name)?.index as number;
    });
    return () => (
      <div>
        <nav id="top">
          <div class="logo">
            <img src="/src/assets/img/icons/person.svg" title={`您好、尊贵的用户@${store.state.account}`}/>
          </div>
          <div class="nav-text">
            <input type="checkbox" id="check" />
            <label for="check" class="checkbtn" id="uapply-checkbtn">
              <i class="fa fa-bars"></i>
            </label>
            <ul class="nav-titles">
              {navTitles.map((obj, index) => (
                <li
                  class={`nav-title ${activeIndex.value === index ? "active" : ""} ${
                    obj.disabled ? "disabled" : ""
                  }`}
                  onClick={() => !obj.disabled && setActive(index, obj.name)}
                  title={`${obj.disabled ? "该功能尚未开放, 敬请期待": ""}`}
                >
                  {obj.label}
                </li>
              ))}
            </ul>
          </div>
        </nav>
        {/* absolute撑开 */}
        <div style="height: 70px; width: 100%"></div>
      </div>
    );
  },
});