import './assets/main.css'
import 'virtual:uno.css'


import { createApp } from 'vue'
import App from './App.vue'
import router from '@render/router/router'

const app = createApp(App)
app.use(router)
app.mount('#app')
