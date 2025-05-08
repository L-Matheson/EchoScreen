<script lang="ts" setup>
import { useMainEchoStore } from '../stores/MainEchoStore'
import './styling/Dashboard.css' 


const store = useMainEchoStore()
</script>

<template>
  <div class="p-0 m-0">
    <div class="top-cards">
      <div class="single-card">
        <i class="pi pi-chart-bar card-icon"></i>
        <h2 class="card-title">Active Screens</h2>
        <p class="card-content">0</p>
      </div>
      <div class="single-card">
        <i class="pi pi-chart-line card-icon"></i>
        <h2 class="card-title">Active Groups</h2>
        <p class="card-content">0</p>
      </div>
      <div class="single-card">
        <i class="pi pi-chart-pie card-icon"></i>
        <h2 class="card-title">Card 3</h2>
        <p class="card-content">Content for Card 3</p>
      </div>
      <div class="single-card">
        <i class="pi pi-cog card-icon"></i>
        <h2 class="card-title">Card 4</h2>
        <p class="card-content">Content for Card 4</p>
      </div>
    </div>
  </div>
</template>

<style></style>
