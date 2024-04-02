import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: () => import('../views/HomePage.vue'),
            meta: { showNav: true }
        },
        {
            path: '/joinqueue',
            name: 'joinqueue',
            component: () => import('../views/JoinQueue.vue'),
            meta: { showNav: true }
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
            },
            meta: { showNav: true }

        },
        {   
            path: '/book',
            name: 'book',
            component: () => import('../views/ClassList.vue'),
            meta: { showNav: true }
        },
        {
            path: '/bookyoga',
            name: 'bookyoga',
            component: () => import('../views/BookYoga.vue'),
            meta: { showNav: true }

        },
        
        {
            path: '/bookdance',
            name: 'bookdance',
            component: () => import('../views/BookDance.vue'),
            meta: { showNav: true }

        },
        {
            path: '/bookpilates',
            name: 'bookpilates',
            component: () => import('../views/BookPilates.vue'),
            meta: { showNav: true }

        },
        {
          path: '/profile',
          name: 'profile',
          component: () => import('../views/ProfileDefaultPage.vue'),
          meta: { showNav: true }
        },
        {
            path: '/profileEditInfo',
            name: 'profileEditInfo',
            component: () => import('../views/ProfileEditInfoPage.vue'),
            meta: { showNav: true }
        },
        {
          path: '/profileBookings',
          name: 'profileBookings',
          component: () => import('../views/ProfileBookingsPage.vue'),
          meta: { showNav: true }
        },
        {
            path: '/profileRefer',
            name: 'profileRefer',
            component: () => import('../views/ProfileReferPage.vue'),
            meta: { showNav: true }
        },
        {
            path: '/profileLogout',
            name: 'profileLogout',
            component: () => import('../views/ProfileLogoutPage.vue'),
            meta: { showNav: true }
        },
        {
          path: '/createClass',
          name: 'createClass',
          component: () => import('../views/CreateClass.vue'),
          meta: { showNav: true }
        },
        {
          path: '/classSearch',
          name: 'classSearch',
          component: () => import('../views/ClassSearch.vue'),
          meta: { showNav: true }
        },
        {
          path: '/login',
          name: 'Login',
          component: () => import('../views/Login.vue'),
          meta: {showNav: false}
        },
        {
          path: '/sign-up',
          name: 'SignUp',
          component: () => import('../views/SignUp.vue'),
          meta: { showNav: false }
      },
      {
        path: "/gymManagement",
        name: "GymManagement",
        component: () => import("../views/GymManagement.vue"),
        meta: {showNav: true}
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

// Global beforeEach guard
router.beforeEach((to, from, next) => {
  // List of public pages that don't require auth
  const publicPages = ['/login', '/sign-up'];
  // Assume 'isLoggedIn' state is stored in localStorage
  const isLoggedIn = JSON.parse(localStorage.getItem('loggedIn') || 'false');

  if (!publicPages.includes(to.path) && !isLoggedIn) {
    // Redirect to the login page if not logged in and trying to access a restricted page
    return next('/login');
  }
  // For all other cases, just continue as normal
  next();
});


export default router;