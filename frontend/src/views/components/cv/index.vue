<template>
  <div class="uapply-box">
    <mini-cv
        :form="miniform"
    />
    <div class="resume_wrapper">
      <div class="detail_content">
        <div class="info_item">
          <div class="title bold">具体信息</div>
          <div class="list">
            <template v-for="field in fields" :key="index">
              <div class="item item_half">
                <div class="label">{{ field.label }}</div>
                <div v-if="!status" class="text">{{ formatFieldValue(field) }}</div>
                <component
                  v-else
                  :is="field.component"
                  v-model="form[field.model]"
                  :placeholder="field.placeholder"
                  :options="field.options"
                />
              </div>
            </template>
          </div>
        </div>
      </div>
      <div ref="myfooter" id="footer" :class="['fixed_box', footerClass]">
        <div class="center_box">
          <template v-if="!status">
            <div class="resume_btn primary mr-12" @click="edit">编辑</div>
            <div class="resume_btn disabled">保存</div>
          </template>
          <template v-else>
            <div class="resume_btn primary mr-12" @click="cancel">取消</div>
            <div class="resume_btn" @click="save">保存</div>
          </template>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import MiniCv from "./minicv.vue";
import { onMounted, onBeforeUnmount, ref, computed } from "vue";
import { getUserCV, commitUserCV } from "@/libs/request";
import { ElInput } from "element-plus";
import { options } from "./constant";
import type { CV_TYPE } from "@/libs/request";
import MySelect from './select/MySelect.vue';
import type { OPTIONS_TYPE } from './constant';
import { ElMessage } from 'element-plus';

const form = ref<CV_TYPE>({
  user_id: undefined,
  name: "",
  age: undefined,
  sex: undefined,
  major: undefined,
  interest: undefined,
  phone: undefined,
  email: undefined,
  init: undefined,
  qq: undefined,
  wc: undefined,
});
const miniform = ref<CV_TYPE>(form.value);

type Field = {
  label: string;
  model: keyof CV_TYPE;
  component: any;
  placeholder: string;
  options?: OPTIONS_TYPE[];
};
const fields: Field[] = [
  { label: "姓名", model: "name", component: ElInput, placeholder: "Please input" },
  { label: "性别", model: "sex", component: MySelect, placeholder: "Select", options: options },
  { label: "年龄", model: "age", component: ElInput, placeholder: "Please input" },
  { label: "手机号", model: "phone", component: ElInput, placeholder: "Please input" },
  { label: "邮箱", model: "email", component: ElInput, placeholder: "Please input" },
  { label: "感兴趣的方向", model: "interest", component: MySelect, placeholder: "Select", options: options },
  { label: "专业", model: "major", component: MySelect, placeholder: "Select", options: options },
  { label: "微信", model: "wc", component: ElInput, placeholder: "Select" },
  { label: "QQ", model: "qq", component: ElInput, placeholder: "Select" },
];

const status = ref(false);
const edit = () => {
  status.value = true;
};
const cancel = () => {
  status.value = false;
};
const save = async () => {
  // 保存逻辑
  console.log(form.value);
  try {
    await commitUserCV(form.value);
    ElMessage.success('上传成功');
  } catch(e) {
    console.log(e);
  }
};

const formatFieldValue = (field: Field) => {
  if(!form.value[field.model]) {
    return '未填';
  }
  if (field.model === "sex") {
    return form.value[field.model] === 1 ? "男" : "女";
  }
  return form.value[field.model];
};

const myfooter = ref<HTMLElement | null>(null);
const footerClass = computed(() => {
  return myfooter.value?.classList.contains("fixed__myfooter") ? "isFixed" : "myfooter";
});

const handleScroll = () => {
  const footer = myfooter.value;
  const windowHeight = window.innerHeight;
  const scrollHeight = document.documentElement.scrollHeight;
  const scrollTop =
    window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;
  if (scrollHeight - (scrollTop + windowHeight) < 80) {
    // 距离底部小于 100px
    footer?.classList.add("absolute__myfooter");
    footer?.classList.remove("fixed__myfooter");
  } else {
    footer?.classList.add("fixed__myfooter");
        footer?.classList.remove('absolute__myfooter');
    }
}
onMounted(async () => {
    try {
      const data = await getUserCV();
      form.value = data.data;
      miniform.value = form.value;
    } catch(e) {
      console.log(e);
    }
    handleScroll();
    window.addEventListener("scroll", handleScroll);
})
onBeforeUnmount(() => {
    window.removeEventListener('scroll', handleScroll);
})
</script>
<style lang="scss" scoped>
@import './cv.scss';
.myfooter {
    @apply px-16;
}
.fixed__myfooter {
    position: fixed !important;
    bottom: 0!important;
    padding-left: 64px;
    padding-right: 64px;
}
.absolute__myfooter {
    position: absolute!important;
    bottom: -130px!important;
    padding-left: 3rem;
    padding-right: 3rem;
}
@media (max-width: 750px) {
    .uapply-box {
        padding: 0;
    }
    .myfooter {
        @apply px-4;
    }
    .fixed__myfooter {
        @apply px-4;
    }
    .absolute__myfooter {
        @apply px-0;
    }
}

@media (min-width: 600px) {
    ::v-deep(.el-input) {
        width: 200px;
    }
}
@media (min-width: 1000px) {
    ::v-deep(.el-input) {
        width: 300px;
    }
}
</style>