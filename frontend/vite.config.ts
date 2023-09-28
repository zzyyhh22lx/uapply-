import { defineConfig } from 'vite';
import path from 'path';
import { fileURLToPath } from "url";
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons';
import vueJsx from "@vitejs/plugin-vue-jsx";
import vue from '@vitejs/plugin-vue';

export default defineConfig({
    plugins:[
      vue(),
      vueJsx(),
      createSvgIconsPlugin({
        // 指定需要缓存的图标文件夹
        iconDirs: [path.resolve(process.cwd(), 'src/assets/img/icons')],
        // 指定symbolId格式
        symbolId: 'icon-[dir]-[name]',
      }),
    ],
    base: "./",
    resolve: {
        alias: {
          "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
})
