<script>
import axios from 'axios'

export default {
  mounted() {
    axios.get('http://localhost:8080/gitrepos/fleet-default')
        .then((response) => {
          console.log(response.data)
          this.gitrepos = response.data.items
        })
  },
  data() {
    return {
      gitrepos: []
    };
  }
}

</script>

<template>
  <main>
    <h1>Git Repos</h1>
    <div v-if="!gitrepos.length">
      No GitRepos yet, please create one!
    </div>
    <div v-else>
      <table>
        <tr>
          <th>State</th>
          <th>Name</th>
          <th>Repo</th>
          <th>Clusters Ready</th>
          <th>Resources</th>
          <th>Age</th>
        </tr>
        <tr v-for="gitrepo in gitrepos" :key="gitrepos.name">
          <td>{{ gitrepo.state }}</td>
          <td>{{ gitrepo.name }}</td>
          <td>{{ gitrepo.repoName }}</td>
          <td>{{ gitrepo.clustersReady }}</td>
          <td>{{ gitrepo.resources }}</td>
          <td>{{ gitrepo.age }}</td>
        </tr>
      </table>
    </div>
  </main>
</template>

<style>
</style>