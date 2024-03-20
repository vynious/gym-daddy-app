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
            path: '/joinqueue',
            name: 'joinqueue',
            component: () => import('../views/JoinQueue.vue')
        },
        {
            path: '/queue',
            name: 'queue',
            component: () => import('../views/QueuePage.vue'),
            beforeEnter: (to, from, next) => {
                const hasJoinedQueue = false; // supposed to take boolean from backend
                if (hasJoinedQueue) {
                    next({name: 'joinqueue'});
                } else {
                    next();
                }
             }
        }
    ]
})

export default router
