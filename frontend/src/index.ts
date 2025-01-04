import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import DistrictView from '@/views/DistrictView.vue'
import DistrictsView from '@/views/DistrictsView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/about', component: AboutView },
  { path: '/districts', component: DistrictsView },
  { path: '/districts/:name', component: DistrictView },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router