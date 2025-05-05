<script lang="ts" setup>
import { useMainEchoStore } from '/src/stores/MainEchoStore.ts'
import EchoSideBar from '/src/components/EchoSideBar.vue'
// define the store
const store = useMainEchoStore()
</script>

<template>
  <div>
    <EchoSideBar></EchoSideBar>
  </div>
</template>
