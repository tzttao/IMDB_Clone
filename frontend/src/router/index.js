import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '/UserLogIn',
      component: () => import('../views/user/UserLogIn.vue')
    },
    {
      path: '/UserSignup',
      component: () => import('../views/user/UserSignUp.vue')
    },
    {
      path: '/signup',
      component: () => import('../components/UserContainer.vue')
    },
    {
      path: '/hello',
      component: () => import('../views/HelloWorld.vue')
    },
  ]
});

// 导航守卫
// 使用 router.beforeEach 注册一个全局前置守卫，判断用户是否登陆
router.beforeEach(async (to, from, next) => {
  if (to.path === '/UserLogIn') {
    next();
  } else {
    let token = localStorage.getItem('token');
    if (token === null || token === '') {
      next('/UserLogIn');
    } else {
      next();
    }
  }
});
export default router;
