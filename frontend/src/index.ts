import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '@/views/user/HomeView.vue'
import TermView from '@/views/user/TermView.vue'
import AboutView from '@/views/user/AboutView.vue'
import GuidanceView from '@/views/user/GuidanceView.vue'
import DistrictView from '@/views/user/DistrictView.vue'
import DistrictsView from '@/views/user/DistrictsView.vue'

import AuthView from '@/views/admin/AuthView.vue'
import AdminView from '@/views/admin/AdminView.vue'
import EntityView from '@/views/admin/EntitiesView.vue'
import AddEntityView from '@/views/admin/AddEntityView.vue'
import EditEntityView from '@/views/admin/EditEntityView.vue'


const routes = [
  { name: 'main', path: '/', component: HomeView },
  { name: 'auth', path: '/auth', component: AuthView },
  { name: 'about', path: '/about', component: AboutView },
  { name: 'guidance', path: '/guidance', component: GuidanceView},
  { name: 'term', path: '/guidance/:id', component: TermView},
  { name: 'districts', path: '/districts', component: DistrictsView },
  { name: 'district', path: '/districts/:id', component: DistrictView },
  {
    path: '/admin', component: AdminView,
    meta: { admin: true }
  },
  {
    path: '/admin/:entity', component: EntityView,
    meta: { admin: true }
  },
  {
    path: '/admin/:entity/add', component: AddEntityView,
    meta: { admin: true }
  },
  {
    path: '/admin/:entity/edit/:id', component: EditEntityView,
    meta: { admin: true }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to) => {
  if(to.meta.admin) {
    if (localStorage.getItem('token')) {
      return true
    } else {
      return {
        name: 'auth'
      }
    }
  }
  else {
    return true
  }
})

export default router