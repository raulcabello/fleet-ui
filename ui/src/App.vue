<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Menu from "@/components/Menu.vue";
import {configStore} from "@/store/config";

const config = configStore()
</script>

<template>
  <Menu />
  <div class="content-wrapper">
    <RouterView />
    <div v-if="config.loading" class="overlay-container">
        <div class="spinner-border spinner" role="status">
          <span class="visually-hidden">Loading...</span>
        </div>
    </div>
  </div>
</template>

<style scoped>

@media (min-width: 1024px) {
  .content-wrapper{
    width: 1024px;
    margin: 1.5rem auto;
  }
  .overlay-container{
    position:absolute;
    top:0;
    background: white;
    width: 100%;
    height: 100%;
  }
  .spinner{
    top: 80px;
    left: 496px;
  }
}
</style>
