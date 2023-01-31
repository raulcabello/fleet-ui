<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="menu-wrapper">
    <div class="fleet">
      <img height="30" src="./assets/logo.svg" alt="Logo"> <span class="title">Fleet</span>
    </div>
    <nav>
      <RouterLink class="menu-item" active-class="active" to="/">Git Repos</RouterLink>
      <RouterLink class="menu-item" active-class="active" to="/bundles">Bundles</RouterLink>
      <RouterLink class="menu-item" active-class="active" to="/about">Clusters</RouterLink>
      <RouterLink class="menu-item" active-class="active" to="/about">Cluster Groups</RouterLink>
    </nav>
  </div>
  <div class="content-wrapper">
  <RouterView />
  </div>
</template>

<style scoped>

@media (min-width: 1024px) {
  .menu-wrapper {
    text-align: left;
    font-size: 1rem;
    padding: 1rem 1rem;
    width: 180px;
    height: 100%;
    border-right: green;
    border-right-style: solid;
    border-right-width: 1px;
  }
  .menu-wrapper nav {
    width: 100%;
    margin-top: 1rem;
  }
  .menu-wrapper a {
    display: block;
    width: 100%;
    margin: 0.5rem 0;
  }
  .menu-item {
    color: black;
  }
  .active {
    color: green;
    text-decoration: underline;
  }
  .content-wrapper{
    width: 100%;
    margin: 2rem;
  }
  .title{
    height: 30px;
    line-height: 30px;
    margin-left: 0.5rem;
    font-weight: bold;
  }
  .fleet{
    display: flex;
  }
}
</style>
