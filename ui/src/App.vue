<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Menu from "@/components/Menu.vue";
</script>

<template>
  <Menu />
  <div class="content-wrapper">
    <RouterView />
  </div>
</template>

<style scoped>

@media (min-width: 1024px) {
  .content-wrapper{
    width: 1024px;
    margin: 1.5rem auto;
  }
}
</style>
