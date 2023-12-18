import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../pages/Home.vue')
  },
  {
    path: '/stats',
    name: 'Stats',
    component: () => import('../pages/Stats.vue')
  },
]
const router = createRouter({
  history: createWebHistory(),
  routes
})
export default router