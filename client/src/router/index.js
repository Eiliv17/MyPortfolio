import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  scrollBehavior(to, from, savedPosition) {
    return { top: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/projects',
      name: 'projects',
      component: () => import('../views/ProjectsView.vue'),
    },
    {
      path: '/projects/:id',
      name: 'projects-article',
      component: () => import('../views/ProjectsArticleView.vue'),
      props: true
    },
    {
      path: '/blog',
      name: 'blog',
      component: () => import('../views/BlogView.vue'),
    },
    {
      path: '/blog/:id',
      name: 'blog-article',
      component: () => import('../views/BlogArticleView.vue'),
      props: true
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('../views/ContactView.vue')
    },
    {
        path: '/cookies',
        name: 'cookies',
        component: () => import('../views/CookiePolicyView.vue')
    },
    { 
      path: '/:pathMatch(.*)*', 
      component: () => import('../views/Error404View.vue')
    }
  ]
})

export default router
