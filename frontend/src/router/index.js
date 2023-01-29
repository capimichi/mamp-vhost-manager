import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/vhosts/single',
      name: 'vhost',
      component: () => import('../views/VhostSingleView.vue'),
    }
  ]
})

export default router
