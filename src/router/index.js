// import Vue from 'vue'
// import VueRouter from 'vue-router'

// import { createRouter, createWebHistory } from 'vue-router'

// import App from './App.vue'
// import Home from '../views/home.vue'
// import Queue from '../views/queue.vue'

// Vue.use(VueRouter)

// const routes = [
//     { path: '/', component: Home },
//     { path: '/queue', component: Queue }

// ]

// const router = new VueRouter({
//     history: createWebHistory(import.meta.env.BASE_URL),
//     routes
// })

// new Vue({
//   router,
//   render: h => h(App)
// }).$mount('#app')


import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/home.vue')
        },
        {
            path: '/queue',
            name: 'queue',
            component: () => import('../views/queue.vue')
        }
    ]
})

export default router