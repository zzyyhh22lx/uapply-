import { defineComponent, ref } from "vue";
export default defineComponent({
  name: "ApplyHeader",
  setup(props, ctx) {
    const text = ref("我是header");
    return () => <div>{text.value}</div>;
  },
});