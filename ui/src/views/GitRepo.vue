<script setup>
import {onMounted, computed} from 'vue';
import BundleTable from "@/components/BundleTable.vue";
import {gitRepoStore} from "@/store/gitRepo";
import ResourcesTable from "@/components/ResourcesTable.vue";
import ConditionsTable from "@/components/ConditionsTable.vue";

const props = defineProps(['name'])
const store = gitRepoStore()

const displayResourcesReady = computed(() => {
  if (store.gitRepo.resourceCount === undefined) {
    return ""
  } else {
    return store.gitRepo.resourceCount.ready+"/"+store.gitRepo.resourceCount.desiredReady
  }
});

onMounted(() => {
  store.fetchGitRepo(props.name)
})
</script>

<template>
  <main>
    <div >
      <h4>Git Repo: {{store.gitRepo.name}}</h4>
      <div class="mt-4">
        <nav>
          <div class="nav nav-tabs" id="nav-tab" role="tablist">
            <button class="nav-link active" id="nav-home-tab" data-bs-toggle="tab" data-bs-target="#nav-home" type="button" role="tab" aria-controls="nav-home" aria-selected="true">Bundles {{store.gitRepo.displayBundlesReady}}</button>
            <button class="nav-link" id="nav-profile-tab" data-bs-toggle="tab" data-bs-target="#nav-profile" type="button" role="tab" aria-controls="nav-profile" aria-selected="false">Resources {{displayResourcesReady}}</button>
            <button class="nav-link" id="nav-contact-tab" data-bs-toggle="tab" data-bs-target="#nav-contact" type="button" role="tab" aria-controls="nav-contact" aria-selected="false">Conditions</button>
          </div>
        </nav>
        <div class="tab-content" id="nav-tabContent">
          <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-home-tab">
            <BundleTable/>
          </div>
          <div class="tab-pane fade" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">
            <ResourcesTable/>
          </div>
          <div class="tab-pane fade" id="nav-contact" role="tabpanel" aria-labelledby="nav-contact-tab">
            <ConditionsTable/>
          </div>
        </div>
      </div>
    </div>
    </main>
</template>

<style>
.fleet-empty{
  text-align: center;
}
</style>