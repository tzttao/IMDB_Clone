import Vue from 'vue'
import Router from 'vue-router'
import Login from "../components/Login";
import SignUp from "../components/SignUp";
import LogIn from "../components/Login";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/signUp',
      name: 'signUp',
      component: SignUp
    },
  ]
})
