<script lang="ts" setup>
import '/src/components/componentStyling/DropDownList.css'
import { useMainEchoStore } from '/src/stores/MainEchoStore.ts'
import Button from 'primevue/button'

const mainStore = useMainEchoStore()

defineProps({
  buttons: Array<{name: string, icon: string}>
})
</script>

<template>
  <div class="dropdown-background">
    
    
      <div v-for="button in buttons" :key="button" >

       <button style="cursor: pointer;"  class="button">
        <i :class="button.icon" ></i>

        {{ button.name }} 
    </button> 
      </div>
  
  </div>
</template>
