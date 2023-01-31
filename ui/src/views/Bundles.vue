<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios'

const bundles = ref([])

onMounted(() => {
  axios.get('http://localhost:8080/bundles/fleet-default')
      .then((response) => {
        console.log(response.data)
        bundles.value = response.data.items
      })
})
</script>

<template>
  <main>
    <h1>Bundles</h1>
    <div v-if="!bundles.length">
      No Bundles!
    </div>
    <div v-else>
      <table>
        <tr>
          <th>State</th>
          <th>Name</th>
          <th>Deployments</th>
          <th>Age</th>
        </tr>
        <tr v-for="bundle in bundles" :key="bundle.name">
          <td>{{ bundle.state }}</td>
          <td> <router-link :to="{ name: 'bundle', params: {name: bundle.name } }">{{ bundle.name }}</router-link></td>
          <td>{{ bundle.deployments }}</td>
          <td>{{ bundle.age }}</td>
        </tr>
      </table>
    </div>
  </main>
</template>

<style>

</style>