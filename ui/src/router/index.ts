import { createRouter, createWebHistory } from 'vue-router'
import GitRepos from '../views/GitRepos.vue'
import Bundles from '../views/Bundles.vue'
import Bundle from '../views/Bundle.vue'
import CreateGitRepo from "@/views/CreateGitRepo.vue";
import GitRepo from "@/views/GitRepo.vue";
import Clusters from "@/views/Clusters.vue";
import ClustersGroup from "@/views/ClustersGroup.vue";

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
      path: '/creategitrepo',
      name: 'creategitrepo',
      component: CreateGitRepo
    },
    {
      path: '/gitrepo/:name',
      name: 'gitrepo',
      component: GitRepo,
      props: true
    },
    {
      path: '/clusters',
      name: 'cluster',
      component: Clusters
    },
    {
      path: '/clustergroups',
      name: 'clustergroups',
      component: ClustersGroup
    }
  ]
})

export default router
