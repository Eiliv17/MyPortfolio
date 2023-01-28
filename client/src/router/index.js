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
      component: () => import('../views/HomeView.vue'),
      meta: {
        title: "Home | Alberto Fabro",
      }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue'),
      meta: {
        title: "About | Alberto Fabro",
      }
    },
    {
      path: '/projects',
      name: 'projects',
      component: () => import('../views/ProjectsView.vue'),
      meta: {
        title: "Projects | Alberto Fabro",
      }
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
      meta: {
        title: "Blog | Alberto Fabro",
      }
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
      component: () => import('../views/ContactView.vue'),
      meta: {
        title: "Contact | Alberto Fabro",
      }
    },
    {
        path: '/cookies',
        name: 'cookies',
        component: () => import('../views/CookiePolicyView.vue'),
        meta: {
            title: "Cookies | Alberto Fabro",
        }
    },
    { 
      path: '/:pathMatch(.*)*', 
      component: () => import('../views/Error404View.vue'),
      meta: {
        title: "Error 404 | Alberto Fabro",
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
    document.title = `${to.meta.title}`;
    next();
})

export default router
