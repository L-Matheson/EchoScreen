<script lang="ts" setup>
import Button from 'primevue/button';
import Checkbox from 'primevue/checkbox';
import InputText from 'primevue/inputtext';
import 'primeicons/primeicons.css';
import { ref } from 'vue';
import { useMainEchoStore } from './stores/MainEchoStore';
import EchoSideBar from './components/EchoSideBar.vue';
import EchoHeader from './components/EchoHeader.vue';
import DropdownList from './components/DropdownList.vue';
import './AppStyle.css'; // Import the custom CSS file

// define the store
const store = useMainEchoStore();
const loggedIn = ref(false);
const email = ref('');
const password = ref('');


async function handleLogin() {


  fetch('http://localhost:3000')
        .then(response => response.text())
        .then(data => {
        	console.log(data);
        });

  loggedIn.value = true;


}


</script>

<template>
  <div class="p-0 m-0">
    <div v-if="loggedIn" class="content-wrapper">
      <EchoHeader></EchoHeader>
      <EchoSideBar></EchoSideBar>

      <div class="content">
        <router-view></router-view>
      </div>
    </div>
    <div v-else>
      <div class="login-container">
        <div class="login-box">
          <h1 class="login-title">Login</h1>
          <form class="login-form">
            <label for="email" class="login-label">Email</label>
            <InputText v-model="email" id="email" type="email" placeholder="Enter your email" class="login-input" />

            <label for="password" class="login-label">Password</label>
            <InputText  v-model="password" id="password" type="password" placeholder="Enter your password" class="login-input" />

            <div class="login-actions">
              <Checkbox id="remember-me" :binary="true" class="login-checkbox" />
              <label for="remember-me" class="login-remember-label">Remember me</label>
            </div>

            <Button label="Login" class="login-button" @click="handleLogin()" />
          </form>
        </div>
      </div>
    </div>
  </div>
</template>