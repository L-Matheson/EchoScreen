<script lang="ts" setup>
import '/src/components/componentStyling/EchoSideBar.css'
import { useMainEchoStore } from '/src/stores/MainEchoStore.ts'
import 'primeicons/primeicons.css'
import Drawer from 'primevue/drawer'
import Button from 'primevue/button'
import InputGroup from 'primevue/inputgroup'
import InputText from 'primevue/inputtext'
import Tooltip from 'primevue/tooltip'

const mainStore = useMainEchoStore()

function getTooltip(name: string) {
  return {
    value: name,
    pt: {
      arrow: {
        style: {
          borderBottomColor: 'var(--p-primary-color)',
        },
      },
      text: '!bg-primary !text-primary-contrast !font-medium',
    },
  }
}

</script>

<template>
  <div class="side-bar-background">
    <div class="side-bar-buttons">
      <div>
        <Button
          rounded
          class="p-button-text"
          v-tooltip.right="getTooltip('Dashboard')"
          style="height: 50px; width: 50px"
        >
          <i class="pi pi-home" style="font-size: 24px"></i>
        </Button>
      </div>
      <div>
        <Button rounded class="p-button-text" v-tooltip.right="getTooltip('Create Client View')" style="height: 50px; width: 50px">
          <i class="pi pi-pen-to-square" style="font-size: 24px"></i>
        </Button>
      </div>
      <div>
        <Button rounded   v-tooltip.right="getTooltip('View Message Groups')" class="p-button-text" style="height: 50px; width: 50px">
          <i class="pi pi-globe" style="font-size: 24px"></i>
        </Button>
      </div>
      <div>
        <Button rounded    v-tooltip.right="getTooltip('Calendar')" class="p-button-text"  style="height: 50px; width: 50px">
          <i class="pi pi-calendar" style="font-size: 24px"></i>
        </Button>
      </div>
      <div>
        <Button rounded   v-tooltip.right="getTooltip('Current Client View')" class="p-button-text" style="height: 50px; width: 50px">
          <i class="pi pi-camera" style="font-size: 24px"></i>
        </Button>
      </div>
    </div>
  </div>
</template>
