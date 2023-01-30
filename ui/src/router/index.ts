import { createRouter, createWebHistory } from 'vue-router'
import GitRepos from '../views/GitRepos.vue'
import Bundles from '../views/Bundles.vue'
import Bundle from '../views/Bundle.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'gitrepos',
      component: GitRepos
    },
    {
      path: '/bundles',
      name: 'bundles',
      component: Bundles
    },
    {
      path: '/bundle/:name',
      name: 'bundle',
      component: Bundle,
      props: true
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
