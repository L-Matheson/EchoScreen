<script lang="ts" setup>
import '/src/components/componentStyling/EchoHeader.css'
import { useMainEchoStore } from '/src/stores/MainEchoStore.ts'
import 'primeicons/primeicons.css'
import Drawer from 'primevue/drawer'
import Button from 'primevue/button'
import InputGroup from 'primevue/inputgroup'
import InputText from 'primevue/inputtext'
import InputGroupAddon from 'primevue/inputgroupaddon'
import { ref, onMounted } from 'vue'
import Avatar from 'primevue/avatar'
import DropdownList from '../components/DropdownList.vue'

const mainStore = useMainEchoStore()
const visible = ref(false)
const user = ref({ username: '', role: '' })
const isDropdownVisible = ref(false)

onMounted(() => {
  loadData()
})

async function loadData() {
  user.value = mainStore.getUser()
}

</script>

<template>
  <div class="header-background">
    <!-- replace this icon with the offical logo of echoscreen when obtained -->
    <i class="pi pi-desktop" style="font-size: 42px; padding-left: 14px; padding-top: 14px"></i>

    <div class="action-buttons">
      <Button
        @click="visible = true"
        style="max-height: 35px; max-width: 35px; min-height: 35px; min-width: 35px"
        rounded
      >
        <i class="pi pi-cog" style="font-size: 20px"></i
      ></Button>

      <InputGroup>
        <InputText placeholder="Search" />
        <InputGroupAddon>
          <Button icon="pi pi-search" severity="secondary" variant="text" />
        </InputGroupAddon>
      </InputGroup>

      <div
        class="user-profile-button"
        style="cursor: pointer"
        @focus="isDropdownVisible = true"
        @blur="isDropdownVisible = false"
        @click="isDropdownVisible = !isDropdownVisible"
      >
        <Avatar label="V" style="background-color: #ece9fc; color: #2a1261" shape="circle" />

        <div style="text-align: left">
          <div
            style="
              width: 80px;
              height: 17px;
              font-size: 14px;
              font-family: sans-serif;
              font-weight: 600;
            "
          >
            {{ user.username }}
          </div>
          <div style="width: 80px; height: 17px; font-size: 14px; font-family: sans-serif">
            {{ user.role }}
          </div>
          <div class="dropdown-menu">
            <DropdownList
              v-if="isDropdownVisible"
              :buttons="[
                { name: 'Test', icon: 'pi pi-user' },
                { name: 'Test2', icon: 'pi pi-cog' },
              ]"
              class="dropdown-menu"
            ></DropdownList>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="card flex justify-center">
    <Drawer v-model:visible="visible" style="background-color: white" position="right">
      <template #header>
        <div class="flex items-center gap-2">
          <span class="font-bold">Settings</span>
        </div>
      </template>
      <p>Add styling options here</p>
      <template #footer>
        <div class="drawer-footer">
          <Button label="Temp" icon="pi pi-user" class="flex-auto" outlined></Button>
          <Button
            label="Temp"
            icon="pi pi-sign-out"
            class="flex-auto"
            severity="danger"
            text
          ></Button>
        </div>
      </template>
    </Drawer>
  </div>
</template>
