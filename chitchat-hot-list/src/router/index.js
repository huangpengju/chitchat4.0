import {createRouter,createWebHashHistory} from 'vue-router'

// 定义一些路由
// 每个路由都需要映射到一个组件。
const routes = [
    {
        path: '/',  // 当前路由规则匹配的 hash 地址（需要和 router-link标签里的地址一样）
        name: 'Home', // 给路由规则起一个别名，命名路由
        component: () => import("views/Home.vue"),   // 当前路由规则对应要展示的组件
        redirect: '/index',
        children: [
          {
            path: '/index',
            name: 'Index',
            component: () => import("views/home/Main.vue")
          }
        ]
    },
    // {
    //   path:'/login',
    //   name:'Login',
    //   component: () => import("views/auth/Login.vue")
    // }
]

const router =  createRouter({
    history:createWebHashHistory(import.meta.env.VITE_BASE),
    routes
})

export default router