<script setup>
import {ref, onMounted} from 'vue';
import axios from 'axios'

const gitRepos = ref([])

function deleteGitRepos() {
  console.log(gitRepos.value.filter(r => r.checked).map(r => r.name))
  axios.delete('http://localhost:8080/gitrepos/', {
    data: gitRepos.value.filter(r => r.checked).map(r => r.name),
  }).then((response)=> {
    location.reload()
  })
}

function selectAll(e) {
  gitRepos.value.forEach(function (gitRepo) {
    gitRepo.checked = e.target.checked
  })
}

onMounted(() => {
  axios.get('http://localhost:8080/gitrepos/fleet-default')
      .then((response) => {
        gitRepos.value = response.data.items
      })
})
</script>

<template>
  <main>
    <h1>Git Repos</h1>

    <div v-if="!gitRepos.length">
      No GitRepos yet, please create one!
    </div>
    <div v-else>
      <a class="btn btn-secondary" @click="deleteGitRepos">Delete</a>
      <table>
        <tr>
          <th><input class="form-check-input" type="checkbox" value="" @change="selectAll($event)"></th>
          <th>State</th>
          <th>Name</th>
          <th>Repo</th>
          <th>Clusters Ready</th>
          <th>Resources</th>
          <th>Age</th>
        </tr>
        <tr v-for="gitRepo in gitRepos" :key="gitRepos.name">
          <th><input class="form-check-input" type="checkbox" value="" v-model="gitRepo.checked"></th>
          <td>{{ gitRepo.state }}</td>
          <td>{{ gitRepo.name }}</td>
          <td>{{ gitRepo.repoName }}</td>
          <td>{{ gitRepo.clustersReady }}</td>
          <td>{{ gitRepo.resources }}</td>
          <td>{{ gitRepo.age }}</td>
        </tr>
      </table>
    </div>
    <br/>
    <router-link to="/creategitrepo">Add Repository</router-link>
  </main>
</template>

<style>
</style>