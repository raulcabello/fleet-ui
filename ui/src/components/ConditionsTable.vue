<script setup>
import {gitRepoStore} from "@/store/gitRepo";
import {computed} from "vue";

const store = gitRepoStore()
const conditions = computed(() => {
  return store.getConditions;
});

</script>

<template>
  <table class="table table-striped table-hover">
    <thead>
    <tr>
      <th scope="col">Condition</th>
      <th scope="col">Status</th>
      <th scope="col">Updated</th>
      <th scope="col">Message</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="condition in conditions" :key="condition.type">
      <td>{{ condition.type }}</td>
      <td>{{ condition.status }}</td>
      <td>{{ condition.lastUpdateTime }}</td>
      <td>{{ condition.message }}</td>
    </tr>
    </tbody>
  </table>
</template>

<style scoped>
</style>
