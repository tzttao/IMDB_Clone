// import Vue from 'vue'
// import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
//
// Vue.use(Router)
//
// // export default new Router({
// //   routes: [
// //
// //   ]
// // })
//
//
//
//
// const routes = [
//   {
//     path: '/',
//     name: 'HelloWorld',
//     component: HelloWorld
//   },
//   // {
//   //   path: '/',
//   //   name: 'signup',
//   //   component: () => import('../components/Signup.vue')
//   // },
//   {
//     path: '/login',
//     name: 'login',
//     component: () => import('../components/Login.vue')
//   },
//   {
//     path: '/forgot-password',
//     name: 'forgot-password',
//     component: () => import('../components/ForgotPassword.vue')
//   }
// ]
//
// const router = new VueRouter({
//   mode: 'history',
//   base: process.env.BASE_URL,
//   routes
// })
//
// export default router
import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  //   {
  //   path: '/',
  //   name: 'HelloWorld',
  //     component: () => import('../components/HelloWorld.vue')
  // },
  {
    path: '/',
    name: 'signup',
    component: () => import('../components/Signup.vue')
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../components/Login.vue')
  },
  {
    path: '/forgot-password',
    name: 'forgot-password',
    component: () => import('../components/ForgotPassword.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
