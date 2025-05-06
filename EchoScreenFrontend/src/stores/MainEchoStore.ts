import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useMainEchoStore = defineStore('mainEchoStore', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }
function getUser(){
  return {username: username, role: role}
}
  const username = ref('Test User')
  const role = ref('Admin')

  return { count, doubleCount, increment, username, role, getUser }
})
