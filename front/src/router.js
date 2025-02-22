import { createWebHistory, createRouter } from 'vue-router'

import HomeView from './HomeView.vue'
import NewView from './NewView.vue'
import ViewView from './ViewView.vue'
import EditView from './EditView.vue'

const routes = [
  { path: '/', component: HomeView, name: 'home' },
  { path: '/new', component: NewView, name: 'new' },
  { path: '/:slug([a-z]*)', component: ViewView, name: 'view' },
  { path: '/:slug([a-z]*)/edit', component: EditView, name: 'edit' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router