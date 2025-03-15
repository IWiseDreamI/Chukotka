import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '@/views/user/HomeView.vue'
import TermView from '@/views/user/TermView.vue'
import AboutView from '@/views/user/AboutView.vue'
import VillageView from '@/views/user/VillageView.vue'
import VillagesView from '@/views/user/VillagesView.vue'
import GuidanceView from '@/views/user/GuidanceView.vue'
import DistrictView from '@/views/user/DistrictView.vue'
import DistrictsView from '@/views/user/DistrictsView.vue'
import MaterialView from "@/views/user/MaterialView.vue"
import MaterialsView from "@/views/user/MaterialsView.vue"


import AuthView from '@/views/admin/AuthView.vue'
import AdminView from '@/views/admin/AdminView.vue'
import EntityView from '@/views/admin/EntitiesView.vue'
import AddEntityView from '@/views/admin/AddEntityView.vue'
import EditEntityView from '@/views/admin/EditEntityView.vue'
import { api } from './entities/api/axios'

const routes = [
  { name: 'main', path: '/', component: HomeView },
  { name: 'auth', path: '/auth', component: AuthView },
  { name: 'about', path: '/about', component: AboutView },

  { name: 'term', path: '/guidance/:id', component: TermView },
  { name: 'guidance', path: '/guidance', component: GuidanceView },

  { name: 'materials', path: '/materials', component: MaterialsView },
  { name: 'material', path: '/materials/:id', component: MaterialView },
  
  { name: 'districts', path: '/districts', component: DistrictsView },
  { name: 'district', path: '/districts/:id', component: DistrictView },

  { name: 'villages', path: '/villages', component: VillagesView },
  { name: 'village', path: '/villages/:id', component: VillageView },

  // Admin routes
  {
    path: '/admin', component: AdminView,
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
  {
    path: '/admin/:entity', component: EntityView,
    meta: { admin: true }
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach(async (to) => {
  if (to.meta.admin) {
    const token = localStorage.getItem('token');
    if (token) {
      try {
        const response = await api.get('/is-admin', {
          headers: { Authorization: token }
        });
        if (response.data.valid) {
          return true;
        } else {
          return { name: 'auth' };
        }
      } catch (error) {
        return { name: 'auth' };
      }
    } else {
      return { name: 'auth' };
    }
  } else {
    return true;
  }
});

export default router;
