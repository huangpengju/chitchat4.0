# chitchat-hot-list

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Customize configuration （自定义配置）

Vite配置参考  
官网（英文）：https://vitejs.dev/config/  
官网（中文）：https://cn.vitejs.dev/config/  
博文：https://blog.csdn.net/to_the_Future/article/details/131147098

## Project Setup （项目设置）

```sh
npm install
```

### Compile and Hot-Reload for Development （编译和热重新加载以进行开发）

```sh
npm run dev
```

### Compile and Minify for Production （编译和缩小生产）

```sh
npm run build
```

### axios （vue3 + axios 与 go 交互）
官网：https://www.axios-http.cn/docs/intro  
官网：http://www.axios-js.com/zh-cn/docs/ 
博文：
```
npm install axios
```

### vue-router (Vue.js 的官方路由) 
官网：https://router.vuejs.org/zh/introduction.html  
参考博文：https://zhuanlan.zhihu.com/p/379580221
```
npm install vue-router@4    
```

### tailwindcss
官网（英文）：https://tailwindcss.com/  
官网（中文）：https://www.tailwindcss.cn/
参考博文1：https://blog.csdn.net/to_the_Future/article/details/131093102
参考博文2：https://blog.csdn.net/agonie201218/article/details/125762819
```
// 安装Tailwind CSS和它的相关依赖
npm install -D tailwindcss postcss autoprefixer  // 方式一
npm install -D tailwindcss@latest postcss@latest autoprefixer@latest    // 方式二

// 创建Tailwind CSS的配置文件
npx tailwindcss init    // 方式一
npx tailwindcss init -p    // 方式二
```

### Headless UI 与 Tailwind CSS完美集成。
官网：https://headlessui.com/
官网：https://github.com/tailwindlabs/headlessui
```
npm install @headlessui/vue@latest
```

### 使用第三方图标 heroicons 使用方法：
官网：https://heroicons.com/
参考博文：https://blog.csdn.net/yongyafang123/article/details/125978200
```
npm install @heroicons/vue
```