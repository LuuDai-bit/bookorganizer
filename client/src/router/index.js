import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/LoginView.vue'
import Signup from '../views/SignupView.vue'
import PageNotFound from '../views/PageNotFoundView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/:pathMatch(.*)',
      name: 'pageNotFound',
      component: PageNotFound
    }
  ]
})

const authRouteNames = ['login', 'signup']

const isAuthRouteNames = (routeName) => {
  return authRouteNames.includes(routeName)
}

router.beforeEach((to, from, next) => {
  let token = localStorage.getItem('token')
  if(!token && !isAuthRouteNames(to.name)) {
    next({ name: 'login' })
  } else {
    if(isAuthRouteNames(to.name) && token) {
      next({ name: 'home' })
    } else {
      next()
    }
  }
})

export default router
