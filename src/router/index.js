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
            path: '/bookdance',
            name: 'bookdance',
            component: () => import('../views/BookDance.vue')

        },
        {
            path: '/bookpilates',
            name: 'bookpilates',
            component: () => import('../views/BookPilates.vue')

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
          {
            path: '/createClass',
            name: 'createClass',
            component: () => import('../views/CreateClass.vue')
          },
          {
            path: '/classSearch',
            name: 'classSearch',
            component: () => import('../views/ClassSearch.vue')
          },
          {
            path: '/login',
            name: 'login',
            component: () => import('../views/Login.vue')
          },
          {
            path: '/sign-up',
            name: 'sign-up',
            component: () => import('../views/SignUp.vue')
          }
    ]
})




export default router;