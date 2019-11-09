import Vue from 'vue'
import VueRouter from 'vue-router'
import Voting from '../views/Voting.vue'
import Welcome from '../views/Welcome.vue'
import AdminHome from '../views/AdminHome.vue'
import store from '@/store'

Vue.use(VueRouter)

const routes = [
  {
    path: '/voting',
    name: 'voting',
    component: Voting,
    beforeEnter: (to: any, from: any, next: any) => {
      if (store.state.authenticated) {
        next()
      } else {
        next("/")
      }
    },
  },
  {
    path: '/',
    name: 'welcome',
    component: Welcome,
    beforeEnter: (to: any, from: any, next: any) => {
      if (!store.state.authenticated) {
        next()
      } else {
        next("/voting")
      }
    },
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  },
  {
    path: '/admin',
    name: 'admin',
    component: AdminHome,
    beforeEnter: (to: any, from: any, next: any) => {
      if (store.state.authenticated && store.state.role === "admin") {
        next()
      } else {
        next("/")
      }
    },
    meta: {
      requiresAuth: true,
      is_admin: true
    }
  }
]
const router = new VueRouter({
  routes
})
export default router
