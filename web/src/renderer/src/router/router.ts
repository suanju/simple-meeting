import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '@render/views/home.vue'
import Room from '@render/views/room.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/room/:room_id',
    name: 'Room',
    component: Room
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


export default router
