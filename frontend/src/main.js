// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import VueResource from 'vue-resource'
import 'element-ui/lib/theme-chalk/index.css'
import axios from "axios"
import qs from 'qs'
import dayjs from 'dayjs'
import * as echarts from 'echarts'

Vue.prototype.echarts = echarts
Vue.prototype.$axios = axios
Vue.prototype.$qs = qs
Vue.prototype.$dayjs = dayjs;
//全局注册，使用方法为:this.qs
//axios.defaults.baseURL = '/api'
//允许跨域携带session信息了
axios.defaults.withCredentials = true;
//配置前后端数据交互的请求头：
//axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
axios.interceptors.request.use(function (request) {
  let token = localStorage.getItem('token')
  console.log("request token:" + token)
  if (token) {
    // 添加headers
    request.headers.token = `${token}`;
    request.headers['content-type'] = 'application/x-www-form-urlencoded;charset=UTF-8';
  }
  return request;
}, function (error) {
  return Promise.reject(error);
})

axios.interceptors.response.use(function (response) { // ①10010 token过期（30天） ②10011 token无效
  console.log("response:" + response.data)
  if (response.data.status === -1) {
    localStorage.clear()
    router.push({
      path: 'HomePage'
    })
    router.go(0)
  } else if (response.data.token) {
    localStorage.setItem('token', response.data.token)
  }
  return response
}, function (error) {
  return Promise.reject(error)
})
//Vue.prototype.HOST='/api'
Vue.use(ElementUI)
Vue.use(VueResource)

Vue.config.productionTip = false

//设置全局表单提交格式
//Vue.http.options.emulateJOSN=true;
/* eslint-disable no-new */
new Vue({
  el: '#app', router, components: {App}, template: '<App/>'
})
