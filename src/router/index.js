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

        },
        {
            path: '/book',
            name: 'book',
            component: () => import('../views/ClassList.vue')
        },
        {
            path: '/bookyoga',
            name: 'bookyoga',
            component: () => import('../views/BookYoga.vue')
        },
        {
          path: '/login',
          name: 'Login',
          component: () => import('../views/LoginPage.vue'),
        },
        {
          path: '/sign-up',
          name: 'SignUp',
          component: () => import('../views/SignUp.vue'),
        }
    ],
    scrollBehavior(to, from, savedPosition) {
      if (to.hash) {
        return {
          el: to.hash,
          behavior: 'smooth',
        };
      } else if (savedPosition) {
        return savedPosition;
      } else {
        return { top: 0 };
      }
    }
})

export default router
