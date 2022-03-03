// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
//引入axios  by zhengkai.blog.csdn.net
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css'
import '@/assets/css/main.css'
import qs from 'qs'
Vue.prototype.$axios = axios
axios.defaults.baseURL = '/api'  //自动附加在所有axios请求前面，则可以省略/api，直接写'/xxxx/xxx'。否则需要设置'/api/xxxx/xxx'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
