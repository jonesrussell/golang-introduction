import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/tutorial/:id',
      name: 'tutorial',
      component: () => import('../views/TutorialView.vue'),
      props: true,
    },
    {
      path: '/tutorial/:id/section/:sectionIndex',
      name: 'tutorial-section',
      component: () => import('../views/TutorialView.vue'),
      props: true,
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue'),
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/',
    },
  ],
})

export default router
