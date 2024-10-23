<template>
  <nav class="nav-bar-font">
    <router-link to="/" class="nav-link">Home</router-link>
    <router-link to="/users" class="nav-link">Users</router-link>
    <router-link to="/information" class="nav-link">Information</router-link>
    <a v-if="authenticationStore.userInfo" class="nav-link" @click="handleLogout()">Logout</a>
    <a v-else class="nav-link" @click="handleLogin()">Login</a>
  </nav>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useAuthStore, logout, login } from '@/store/AuthenticationStore'

export default defineComponent({
  name: 'NavBar',
  methods: { logout, login },
  setup() {
    const authenticationStore = useAuthStore()

    const handleLogin = async () => {
      try {
        await login().then(it => {
          authenticationStore.loadFromStorage()
        })
      } catch (error) {
        console.error("Login method failed: ", error)
      }
    }

    const handleLogout = async () => {
      await logout().then(it => {
        authenticationStore.clearAuthentication()
      })
    }

    return {
      authenticationStore,
      handleLogin,
      handleLogout
    }
  }
})
</script>

<style scoped>
nav {
  padding: 2rem;
  background-color: #2c3e50;
  display: flex;
  justify-content: center;
  gap: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
  font-weight: 700;
  font-size: 1.5rem;
  font-family: 'Arial', sans-serif;
  cursor: default;
  transition: color 0.3s ease;
}

.router-link-exact-active {
  color: #ffffff; /* Active link color */
  border-bottom: 2px solid #ffffff; /* Add a bottom border for active link */
}
</style>
