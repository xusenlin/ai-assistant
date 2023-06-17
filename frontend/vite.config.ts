import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'




// https://vitejs.dev/config/
export default defineConfig({
  server: {
    base: "/",
    port: 3003,
    host: '0.0.0.0',
    open: true,
  },
  resolve: {
    alias: {
      '@': '/src/'
    },
  },

  plugins: [vue()],
})
