// import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'

import router from './router'   // 导入路由
import './style/index.css'      // 导入 Tailwind Css 样式文件
import 'element-plus/dist/index.css'

createApp(App).use(router).mount('#app')
