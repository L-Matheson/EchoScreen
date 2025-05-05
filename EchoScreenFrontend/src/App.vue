<script lang="ts" setup>
import { useMainEchoStore } from '/src/stores/MainEchoStore.ts'
import EchoSideBar from '/src/components/EchoSideBar.vue'
// define the store
const store = useMainEchoStore()
</script>

<template>
  <div class="p-0 m-0" style="margin: 0px; padding: 0px;">
    <EchoSideBar></EchoSideBar>
  </div>
</template>

<style>


</style>
