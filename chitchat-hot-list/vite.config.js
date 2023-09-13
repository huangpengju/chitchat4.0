import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/ （官方文档）
export default ({mode})=>{
  // command： "serve" || "build" (开发环境值为 serve，生产环境值为 build)
  // console.log("command=",command)  
  // console.log("mode=",mode)
  // process.cwd：返回node进程的工作目录 Function: wrappedCwd
  // console.log("process.cwd=",process.cwd)
  // 根据当前工作目录中的 `mode` 加载 .env 文件
  // 设置第三个参数为 '' 来加载所有环境变量，而不管是否有 `VITE_` 前缀。
  const env = loadEnv(mode, process.cwd(), '')

  return defineConfig ({
    plugins: [
    // 配置需要使用的插件列表
      vue(),
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],
    resolve: {
        // 别名设置
      alias: {
        '@':resolve(__dirname,'src'),
        'views':resolve(__dirname,'src/views')
        // '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },

    // 开发服务器配置 server（本地运行配置，以及反向代理配置）
    server:{
    	// host指定服务器应该监听哪个 IP 地址。 如果将此设置为 0.0.0.0 或者 true 将监听所有地址，包括局域网和公网地址。
      host:env.VITE_HOST,
      // 开发服务器端口配置。（注意：如果配置的端口已经被使用，Vite 会自动尝试下一个可用的端口，要看项目运行时最终生成的端口地址。）
      port:env.VITE_PORT,
      // open 项目运行完毕是否自动在默认浏览器打开
      open:true,
      // 是否使用 https ，需要证书
      https:false,
      // proxy 代理配置
      // proxy:{
        // '/api':{
          // target:env.VITE_SERVER, //后端服务地址
          // changeOrigin:true,  // 是否跨域
          // ws:true, // 支持 websocket
        // }
      // }
    }
  });
}
