<script setup>
//TODO view not implemented yet!
import { ref, onMounted } from 'vue';
import axios from 'axios'

const props = defineProps(['name'])
const bundle = ref([])

onMounted(() => {
  axios.get('http://localhost:8080/bundles/fleet-default/'+props.name)
      .then((response) => {
        bundle.value = response.data
      })
})
</script>

<template>
  <main>
    <div v-if="!bundle">
      No Bundle yet, please create one!
    </div>
    <div v-else>
    <h1>Bundle: {{bundle.name}}</h1>
      <table>
        <tr>
          <th>State</th>
          <th>API Version</th>
          <th>Kind</th>
          <th>Name</th>
          <th>Namespace</th>
        </tr>
        <tr v-for="resource in bundle.resources" :key="bundle.name">
          <td>{{ resource.state }}</td>
          <td>{{ resource.apiVersion }}</td>
          <td>{{ resource.kind }}</td>
          <td>{{ resource.name }}</td>
          <td>{{ resource.namespace }}</td>
        </tr>
      </table>
    </div>
  </main>
</template>

<style>

</style>