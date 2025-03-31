<template>
  <h1>Home</h1>
  <div>
    <p>
      Simple application for make requests to resource-service (person-service).
    </p>
  </div>
</template>

<script lang="ts">
import { useAuthStore } from '@/store/AuthenticationStore'

export default {
  name: 'HomeView',
  setup() {
    const authenticationStore = useAuthStore()

    return {
      authenticationStore
    }
  }
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
