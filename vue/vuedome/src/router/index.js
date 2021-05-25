import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('@/components/HelloWorld')
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/view/register/Register')
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/view/login/Login')
    }
  ]
})
