import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ClassList from '../views/ClassList.vue'
import UserQueue from '../views/UserQueue.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/classList',
    name: 'classList',
    component: ClassList
  },
  {
    path: '/queue',
    name: 'queue',
    component: UserQueue
  },
  {
    path: '/profile',
    name: 'profile',
    component: () => import('../views/ProfileDefaultPage.vue')
  },
  {
      path: '/profileEditInfo',
      name: 'profileEditInfo',
      component: () => import('../views/ProfileEditInfoPage.vue')
  },
  {
    path: '/profileBookings',
    name: 'profileBookings',
    component: () => import('../views/ProfileBookingsPage.vue')
  },
  {
      path: '/profileRefer',
      name: 'profileRefer',
      component: () => import('../views/ProfileReferPage.vue')
  },
  {
      path: '/profileLogout',
      name: 'profileLogout',
      component: () => import('../views/ProfileLogoutPage.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
