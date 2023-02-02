<script setup>
import {ref, onMounted, computed} from 'vue';
import axios from 'axios'
import moment from "moment";

const gitRepos = ref([])
const isAnyGitRepoSelected = computed(() => gitRepos.value.filter(r => r.checked).length)

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
    <div v-if="!gitRepos.length" class="fleet-empty mt-5">
      <h3>Welcome to Fleet Continuous Delivery</h3>
      <p>GitOps at scale. <a href="https://fleet.rancher.io/" target="_blank">Learn More</a></p>
      <p>You don't have any Git Repositories in your Workspaces</p>
      <router-link class="btn btn-outline-success" to="/creategitrepo">Add Repository</router-link>

    </div>
    <div v-else>
      <h4>Git Repos</h4>
      <hr class="mt-1 mb-1"/>
      <table class="table table-striped table-hover">
        <thead>
          <tr>
            <th scope="col"><input class="form-check-input" type="checkbox" value="" @change="selectAll($event)"></th>
            <th scope="col">State</th>
            <th scope="col">Name</th>
            <th scope="col">Repo</th>
            <th scope="col">Clusters Ready</th>
            <th scope="col">Resources</th>
            <th scope="col">Age</th>
          </tr>
        </thead>
        <tbody>
          <template v-for="gitRepo in gitRepos" :key="gitRepos.name">
            <tr>
              <th><input class="form-check-input" type="checkbox" value="" v-model="gitRepo.checked"></th>
              <td>{{ gitRepo.state }}</td>
              <td> <router-link :to="{ name: 'gitrepo', params: {name: gitRepo.name } }">{{ gitRepo.name }}</router-link></td>
              <td>{{ gitRepo.repoName }}</td>
              <td>{{ gitRepo.clustersReady }}</td>
              <td>{{ gitRepo.resourcesDesiredReady }} / {{ gitRepo.resourcesReady }}</td>
              <td>{{ moment(gitRepo.age, "YYYY-MM-DD hh:mm:ss").fromNow(true)}}</td>
            </tr>
          </template>
        </tbody>
      </table>
      <div class="mt-4">
        <a v-if="isAnyGitRepoSelected" class="btn btn-outline-secondary btn-sm me-3" @click="deleteGitRepos">Delete</a>
        <router-link class="btn btn-outline-success btn-sm" to="/creategitrepo">Add Repository</router-link>
      </div>
    </div>
    </main>
</template>

<style>
.fleet-empty{
  text-align: center;
}
</style>