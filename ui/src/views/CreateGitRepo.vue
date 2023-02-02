<script setup>
import {reactive, ref} from "vue";
import axios from "axios";
import router from "@/router";
import {configStore} from "@/store/config";

const allTargets = [
  {"clusterSelector" : {
      "matchExpressions":[
        {
          "key": "provider.cattle.io",
          "operator": "NotIn",
          "values": ["harvester"]
        }
      ]
    }
  }
]

const gitRepo = reactive({"value": {
    name: "",
    repoURL: "",
    gitSecretSelected: "none",
    gitSecretName: "",
    helmSecretSelected: "none",
    helmSecretName: "",
    tls: "valid",
    paths: []
  }
})

const config = configStore()

function addPath() {
  gitRepo.value.paths.push({ value: '' })
}

function removePath(index) {
  gitRepo.value.paths.splice(index, 1)
}

function createGitRepo() {
  axios.post("http://"+config.url+"/gitrepo", {
    name: gitRepo.value.name,
    repoUrl: gitRepo.value.repoURL,
    paths: gitRepo.value.paths.map(path => path.value),
    targets: allTargets
  }).then(function (response) {
    router.push("/gitrepo/"+gitRepo.value.name)
  })
  .catch(function (error) {
    alert("err "+error);
  });
}

function close() {
  router.push("/")
}
</script >

<!-- TODO: refactor in smaller components -->
<template>
  <main>
    <h4>Create GitRepo</h4>
    <hr class="mt-1 mb-1"/>
    <form>
      <div class="row mt-2">
        <div class="col-4">
          <div class="form-group ">
            <label for="name">Name</label>
            <input type="text" class="form-control form-control-sm" id="name" v-model="gitRepo.value.name">
          </div>
        </div>
        <div class="col">
          <div class="form-group">
            <label for="description">Description</label>
            <input type="text" class="form-control form-control-sm" id="description">
          </div>
        </div>
      </div>
      <div class="row mt-2">
        <div class="col-5">
          <div class="form-group">
            <label for="repoUrl">Repository URL</label>
            <input type="text" class="form-control form-control-sm" id="repoUrl" v-model="gitRepo.value.repoURL">
          </div>
        </div>
        <div class="col">
          <div class="row">
            <div class="col-3">
              <div class="form-group">
                <label for="watch">Watch</label>
                <select id="watch" class="form-select form-select-sm" aria-label="Default select example">
                  <option selected value="branch">Branch</option>
                  <option value="revision">Revision</option>
                </select>
              </div>
            </div>
            <div class="col">
              <div class="form-group">
                <label for="branch">Branch name</label>
                <input type="text" class="form-control form-control-sm" id="branch">
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="row mt-2" v-if="false">
        <div class="col-4">
          <div class="form-group">
            <label for="gitSecret">Git Authentication</label>
            <select v-model="gitRepo.value.gitSecretSelected" id="gitSecret" class="form-select form-select-sm">
              <option selected value="none">None</option>
              <option value="secret">Use secret</option>
            </select>
          </div>
        </div>
        <div class="col">
          <div v-if="gitRepo.value.gitSecretSelected !== 'none'" class="form-group">
            <label for="gitSecretName">Secret name</label>
            <input type="text" class="form-control form-control-sm" id="form-group" v-model="gitRepo.value.gitSecretName">
          </div>
        </div>
      </div>
      <!-- TODO -->
      <div class="row mt-2"  v-if="false">
        <div class="col-4">
          <div class="form-group">
            <label for="helmSecret">Helm Authentication</label>
            <select v-model="gitRepo.value.helmSecretSelected" id="helmSecret" class="form-select form-select-sm">
              <option selected value="none">None</option>
              <option value="secret">Use secret</option>
            </select>
          </div>
        </div>
        <div class="col">
          <div v-if="gitRepo.value.helmSecretSelected !== 'none'" class="form-group">
            <label for="gitSecretName">Secret name</label>
            <input type="text" class="form-control form-control-sm" id="helmSecretSelect" v-model="gitRepo.value.helmSecretName">
          </div>
        </div>
      </div>
      <div class="row mt-2"  v-if="false">
        <div class="col-4">
          <div class="form-group">
            <label for="helmSecret">TLS Certificate Verification</label>
            <select v-model="gitRepo.value.tls" id="tls" class="form-select form-select-sm">
              <option selected value="valid">Require a valid certificate</option>
              <option value="certificates">Specify certificates to be accepted</option>
              <option value="none">Accept any certificate (insecure)</option>
            </select>
          </div>
        </div>
        <div class="col">
          <div v-if="gitRepo.value.tls === 'certificates'" class="form-group">
            <label for="tlsCert">Certificates</label>
            <input type="text" class="form-control form-control-sm" id="tlsCert">
          </div>
        </div>
      </div>
      <div class="row mt-4">
      <h5>Paths</h5>
        <div class="form-group" v-if="!gitRepo.value.paths.length">
          The root of the repo is used by default. To use one or more different directories, add them here.
        </div>
        <div class="form-group" v-else v-for="(path, index) in gitRepo.value.paths">
          <div class="row">
            <div class="col-10">
              <div class="form-group">
                <label for="path">Path</label>
                <input v-model="path.value" type="text" class="form-control form-control-sm" id="path">
              </div>
            </div>
            <div class="col align-self-end" >
              <a @click="removePath(index)" class="btn btn-outline-secondary btn-sm" >Remove</a>
            </div>
          </div>
        </div>
      </div>
      <a @click="addPath" class="mt-2 btn btn-outline-secondary btn-sm">Add Path</a>
      <hr class="mt-3 mb-3"/>

      <div class="row mt-2">
        <h5>Targets</h5>
        <div class="row mt-1">
          <div class="col-4">
            <div class="form-group">
              <label for="target">Deploy to:</label>
              <select id="target" class="form-select form-select-sm">
                <option value="none">No Clusters</option>
                <option selected value="all">All Clusters in the Workspace</option>
              </select>
            </div>
          </div>
        </div>
      </div>
      <div class="row mt-2">
        <div class="col">
          <div class="form-group">
            <label for="sa">Service Account Name</label>
            <input type="text" class="form-control form-control-sm" id="sa">
          </div>
        </div>
        <div class="col">
          <div class="form-group">
            <label for="targetNs">Target Namespace</label>
            <input type="text" class="form-control form-control-sm" id="targetNs">
          </div>
        </div>
      </div>
      <a type="submit" @click="close" class="mt-4 mb-2 me-3 btn btn-outline-secondary btn-sm">Cancel</a>
      <a type="submit" @click="createGitRepo" class="mt-4 mb-2 btn btn-outline-success btn-sm">Create</a>
    </form>
  </main>
</template>

<style>

</style>