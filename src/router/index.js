import { createRouter, createWebHistory } from 'vue-router'


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/HomePage.vue')
        },
        {
            path: '/queue',
            name: 'queue',
            component: () => import('../views/QueuePage.vue')
        },
        {
            path: '/joinqueue',
            name: 'joinqueue',
            component: () => import('../views/JoinQueue.vue')
        }
    ]
})

export default router
