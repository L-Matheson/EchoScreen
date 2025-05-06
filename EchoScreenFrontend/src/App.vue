<script lang="ts" setup>
import { useMainEchoStore } from './stores/MainEchoStore'
import EchoSideBar from './components/EchoSideBar.vue'
import EchoHeader from './components/EchoHeader.vue'
// define the store
const store = useMainEchoStore()
</script>

<template>
  <div class="p-0 m-0">
    
    <div class="content-wrapper">
    <EchoHeader></EchoHeader>
    <EchoSideBar></EchoSideBar>
<div class="content">
  content
</div>

    </div>
  </div>
</template>

<style>
.content{
  border-radius: 25px 0px 0px 0px; 
  padding: 20px; 
background-color: #f1f5f9;
min-height: 100vh - 70px;
width: 100%;
display: flex;
box-shadow: rgba(0, 0, 0, 0.06) 0px 2px 4px 0px inset;
}
.content-wrapper{
margin-top: 70px;
margin-left: 70px;
min-height: 100vh;
display: flex;

}
</style>
