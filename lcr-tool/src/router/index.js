import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Members from '@/components/Members'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/members',
      name: 'Members',
      component: Members
    }
  ]
})
