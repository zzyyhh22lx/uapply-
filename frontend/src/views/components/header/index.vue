<template>
    <nav id="top" ref="header">
      <div class="logo">
        <img :src="logoSrc" :title="logoTitle" />
      </div>
      <div class="nav-text">
        <input type="checkbox" id="check" />
        <label for="check" class="checkbtn" id="uapply-checkbtn" ref="uapplyCheckbtn">
          <i class="fa fa-bars"></i>
        </label>
        <ul class="nav-titles">
          <li
            v-for="(obj, index) in navTitles"
            :key="index"
            :class="['nav-title', { active: activeIndex === index, disabled: obj.disabled }]"
            @click="!obj.disabled && setActive(index, obj.name)"
            :title="obj.disabled ? '该功能尚未开放, 敬请期待' : ''"
          >
            {{ obj.label }}
          </li>
        </ul>
      </div>
    </nav>
    <!-- absolute撑开 -->
    <div style="height: 70px; width: 100%"></div>
</template>

<script lang="ts">
import { getCurrentInstance, defineComponent, ref, onMounted, onBeforeUnmount } from "vue";
import { useRouter, useRoute } from "vue-router";
import { store } from "@/store/index";


export default defineComponent({
  name: "ApplyHeader",
  setup() {
    const router = useRouter();
    const route = useRoute();
    const vm = getCurrentInstance();

    const navTitles = [
      {
        name: "recommand",
        label: "为你推荐",
        index: 0,
        disabled: true,
      },
      {
        name: "uapply",
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
      },
    ];
    const activeIndex = ref(1);
    const uapplyCheckBtn = ref<HTMLElement | null>(null);
    const header = ref<HTMLElement | null>(null);
    const setActive = (index: number, name: string) => {
      activeIndex.value = index;
      document.getElementById('uapply-checkbtn')?.click();
      router.push({
        name,
      });
    };
    const goClickCheckBtn = () => {
      const isShow = !store.state.isShow;
      store.commit("setShowType", isShow);
      vm?.proxy?.$event.emit("clickCheckBtn", isShow);
    };
    const handleClick = () => {
      const currentWidth = window.innerWidth;
      if (currentWidth < 750) {
        goClickCheckBtn();
      }
    };
    const setHeight = (isShow: boolean) => {
        isShow ? (header.value as HTMLElement).classList.add('opacity-0') : (header.value as HTMLElement).classList.remove('opacity-0');
    }

    onMounted(() => {
      activeIndex.value = navTitles.find((item, index) => item.name === route.name)?.index as number;
      uapplyCheckBtn.value?.addEventListener("click", handleClick);
    });
    onBeforeUnmount(() => {
      uapplyCheckBtn.value?.removeEventListener("click", handleClick);
    });

    const logoSrc = "/src/assets/img/icons/person.svg";
    const logoTitle = `您好、尊贵的用户@${store.state.account}`;

    return {
      logoSrc,
      logoTitle,
      navTitles,
      activeIndex,
      goClickCheckBtn,
      setActive,
      uapplyCheckBtn,
      header,
      setHeight,
    };
  },
});
</script>

<style lang="scss" scoped>
@import "./index.scss";
</style>