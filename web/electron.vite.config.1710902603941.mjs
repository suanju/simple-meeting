// electron.vite.config.ts
import { resolve } from "path";
import UnoCSS from "unocss/vite";
import { defineConfig, externalizeDepsPlugin } from "electron-vite";
import vue from "@vitejs/plugin-vue";
var alias = {
  "@render": resolve("src/renderer/src"),
  "@main": resolve("src/main"),
  "@res": resolve("resources")
};
var electron_vite_config_default = defineConfig({
  main: {
    plugins: [externalizeDepsPlugin()]
  },
  preload: {
    plugins: [externalizeDepsPlugin()]
  },
  renderer: {
    resolve: {
      alias
    },
    plugins: [vue(), UnoCSS()]
  }
});
export {
  electron_vite_config_default as default
};
