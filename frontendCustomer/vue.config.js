const { defineConfig } = require('@vue/cli-service')

const autoprefixer = require('autoprefixer');
const pxtoviewport = require('postcss-px-to-viewport-8-plugin');
const pagesConfig = require("./page.config.js");

module.exports = defineConfig({
  publicPath:"./",
  devServer:{
    allowedHosts: "all",
  },
  transpileDependencies: true,
  pages: pagesConfig,
  css: {
    loaderOptions: {
      postcss: {
        postcssOptions: {plugins:[
            autoprefixer(),
            pxtoviewport({
              viewportWidth: 375
            })
          ]}
      }
    }
  },
  // chainWebpack: config => {
  //   config.optimization.splitChunks({
  //     cacheGroups: {
  //       vue: {
  //         name: 'vue',
  //         test: /[\\/]node_modules[\\/](vue|vue-router|vuex)[\\/]/,
  //         chunks: 'all',
  //         priority: 20
  //       },
  //       vendors: {
  //         name: 'vendors',
  //         test: /[\\/]node_modules[\\/]/,
  //         chunks: 'all',
  //         priority: 10
  //       },
  //       commons: {
  //         name: 'commons',
  //         test: /[\\/]src[\\/]/,
  //         chunks: 'all',
  //         minChunks: 2,
  //         priority: 5,
  //         reuseExistingChunk: true
  //       }
  //     }
  //   })
  // }
})
