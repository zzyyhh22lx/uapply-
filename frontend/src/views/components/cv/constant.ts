import { ref } from 'vue';
import type { CV_TYPE } from "@/libs/request";

export const options = [
    {
      value: 'Option1',
      label: 'Option1',
    },
    {
      value: 'Option2',
      label: 'Option2',
      disabled: true,
    },
]

export type OPTIONS_TYPE = {
  value: number | string,
  label: string,
  disabled?: boolean,
};

export const form = ref<CV_TYPE>({
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