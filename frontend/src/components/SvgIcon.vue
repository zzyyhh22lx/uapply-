<template>
    <svg
      xmlns="http://www.w3.org/2000/svg"
      :width="width"
      :height="height"
      :aria-labelledby="name"
      class="udc-svg-icon aw-inline-block aw-align-middle aw--mt-0.5"
      aria-hidden="true"
      viewBox="0 0 16 16"
      role="presentation"
    >
      <title v-if="title">{{ title }}</title>
      <use v-if="!$slots.default" :href="symbolId" :fill="color" />
      <slot v-else />
    </svg>
</template>
  
<script lang="ts">
import { defineComponent, computed } from 'vue';
  
export default defineComponent({
    name: 'SvgIcon',
    props: {
      prefix: {
        type: String,
        default: 'icon',
      },
      name: {
        type: String,
        required: true,
      },
      title: {
        type: String,
        default: '',
      },
      color: {
        type: String,
        default: 'currentColor',
      },
      width: {
        type: String,
        default: '20px',
      },
      height: {
        type: String,
        default: '20px',
      },
    },
    setup(props) {
      const symbolId = computed(() => `#${props.prefix}-${props.name}`);
      return { symbolId };
    },
});
</script>
  