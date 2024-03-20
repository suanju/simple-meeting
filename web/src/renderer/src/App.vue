<script setup lang="ts">

const cross = () => {
  window.electron.ipcRenderer.send("close")
}
const minus = () => {
  window.electron.ipcRenderer.send("minimize")
}
</script>

<template>
  <div>
    <!-- 顶部菜单 -->
    <div class="h-8  w-full flex flex-row-reverse" style="-webkit-app-region: drag;">
      <div class="i-system-uicons-cross h-6 w-6 bg-gray-500 m-2" style="-webkit-app-region: no-drag;" @click="cross">
      </div>
      <div class="i-system-uicons-minus h-6 w-6 bg-gray-500 m-2" style="-webkit-app-region: no-drag;" @click="minus">
      </div>
    </div>
    <div>
      <router-view></router-view>
    </div>
  </div>
</template>