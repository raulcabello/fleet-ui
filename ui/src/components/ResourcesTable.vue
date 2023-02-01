<script setup>
import {gitRepoStore} from "@/store/store";
import {computed} from "vue";

const store = gitRepoStore()
const resources = computed(() => {
  return store.getResources;
});

</script>

<template>
  <table class="table table-striped table-hover">
    <thead>
    <tr>
      <th scope="col">State</th>
      <th scope="col">Cluster</th>
      <th scope="col">API Version</th>
      <th scope="col">Kind</th>
      <th scope="col">Name</th>
      <th scope="col">Namespace</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="resource in resources" :key="resource.name">
      <td>{{ resource.state }}</td>
      <td>{{ resource.cluster }}</td>
      <td>{{ resource.apiVersion }}</td>
      <td>{{ resource.kind }}</td>
      <td>{{ resource.name }}</td>
      <td>{{ resource.namespace }}</td>
    </tr>
    </tbody>
  </table>
</template>

<style scoped>
</style>
