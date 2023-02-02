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

  </main>
</template>

<style>

</style>