import Vue from 'vue'
import Router from 'vue-router'
import HomePage from '../views/user/HomePage.vue'
import UserContainer from '../components/UserContainer'
Vue.use(Router)

const router = new Router({
  routes: [
    {
      path: '*',
      redirect: '/userContainer/HomePage'
    },
    {
      path: '/UserLogIn',//登录
      component: () => import('../views/user/UserLogIn.vue')
    },
    {
      path: '/UserSignup',//注册
      component: () => import('../views/user/UserSignUp.vue')
    },
    {
      path: '/userContainer',
      name: 'UserContainer',
      component: UserContainer,
      children: [
        {
          path: 'HomePage',//主页
          component: HomePage
        },
        {
          path: 'SearchResult',//搜索电影结果
          component: () => import('../views/movie/SearchResult.vue')
        },
        {
          path: 'Userprofile',//用户个人信息
          component: () => import('../views/user/Userprofile.vue')
        },
        {
          path: 'Movies',
          component: () => import('../views/movie/Movies.vue')
        },
        {
          path: 'TopMovies',
          component: () => import('../views/movie/TopMovies.vue')
        },
        {
          path: 'WishMovies',
          component: () => import('../views/user/WishMovies.vue')
        },
        {
          path: 'Actors',
          component: () => import('../views/movie/Actors.vue')
        },
        {
          path: 'Comments',
          component: () => import('../views/movie/Comments.vue')
        },
        {
          path: 'Rating',
          component: () => import('../views/movie/Rating.vue')
        },
      ]
    }
  ]
});

const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}
export default router;
