<script>
import axios from 'axios'

export default {
  mounted() {
    axios.get('http://localhost:8080/bundles/fleet-default')
        .then((response) => {
          console.log(response.data)
          this.bundles = response.data.items
        })
  },
  data() {
    return {
      bundles: []
    };
  }
}

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
          <td>{{ bundle.name }}</td>
          <td>{{ bundle.deployments }}</td>
          <td>{{ bundle.age }}</td>
        </tr>
      </table>
    </div>
  </main>
</template>

<style>

</style>