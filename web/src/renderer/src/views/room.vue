<template>
    <div>
      123123
    </div>
</template>


<script setup lang="ts">
import { useRoute } from 'vue-router';
import io from 'socket.io-client';

const socket = io("http://localhost:3000",{ timeout: 60000 });
const route = useRoute()

socket.on("test",(msg: any) => {
  console.log(msg)
})

setInterval(() => {
  socket.emit("test","fadawd")
})
console.log(route.params.room_id)
</script>