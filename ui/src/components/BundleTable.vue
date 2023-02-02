<script setup>
import {gitRepoStore} from "@/store/store";
import {computed} from "vue";
import moment from "moment";

const store = gitRepoStore()
const bundles = computed(() => {
  return store.getBundles;
});

</script>

<template>
  <table class="table table-striped table-hover">
    <thead>
    <tr>
      <th scope="col">State</th>
      <th scope="col">Name</th>
      <th scope="col">Deployments</th>
      <th scope="col">Date</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="bundle in bundles" :key="bundle.name">
      <td>{{ bundle.state }}</td>
      <td>{{ bundle.name }}</td>
      <td>{{ bundle.deployments }}</td>
      <td>{{ moment(bundle.date, "YYYY-MM-DD hh:mm:ss").fromNow(true)}}</td>
    </tr>
    </tbody>
  </table>
</template>

<style scoped>
</style>
