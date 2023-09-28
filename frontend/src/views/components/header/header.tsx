import { defineComponent, ref } from "vue";
import { useRouter } from 'vue-router';
import { store } from '@/store/index';

import "./index.scss";

export default defineComponent({
  name: "ApplyHeader",
  setup() {
    const router = useRouter();
    const navTitles = [{
      name: 'uapply',
      label: "招新报名"
    }, {
      name: "view-cv",
      label: "我的简历"
    }, 
    {
      name: "view-pc",
      label: "流程查看",
    }];
    const activeIndex = ref(0);
    const checkbtn = ref();
    const setActive = (index: number, name: string) => {
      activeIndex.value = index;
      // 触发点击label事件
      document.getElementById('uapply-checkbtn')?.click();
      router.push({
        name
      })
    };

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
                  class={`nav-title ${activeIndex.value === index ? "active" : ""}`}
                  onClick={() => setActive(index, obj.name)}
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