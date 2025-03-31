<!-- src/App.vue -->
<template>
  <div id="app">
    <NavBar />
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import NavBar from '@/components/NavBar.vue';
import { useAuthStore } from '@/store/AuthenticationStore'

export default defineComponent({
  name: 'App',
  components: {
    NavBar,
  },
  setup() {
    const authenticationStore = useAuthStore()

    onMounted(() => {
      authenticationStore.loadFromStorage()
    })
  }
});
</script>

<style>
/* Remove default margin and padding */
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin: 0; /* Ensure no margin */
  padding: 0; /* Ensure no padding */
}

body {
  margin: 0;
}
</style>
