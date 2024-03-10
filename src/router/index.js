import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ClassList from '../views/ClassList.vue'
import UserProfile from '../views/UserProfile.vue'
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
    path: '/profile',
    name: 'profile',
    component: UserProfile
  },
  {
    path: '/queue',
    name: 'queue',
    component: UserQueue
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
