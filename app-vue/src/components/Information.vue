<template>
  <div class="information">
    <h1>Information about authenticated user</h1>
  </div>
  <p>UserInfo:</p>
  <div class="table-container" v-if="isAuthenticated">
    <table class="data-table">
      <thead>
      <tr>
        <th>Access Token</th>
        <th>Refresh Token</th>
        <th>User Info</th>
      </tr>
      </thead>
      <tbody>
      <tr>
        <td class="tooltip" @click="copyAccessToken">
          {{ truncatedAccessToken }}
          <span class="tooltip-text">Click to copy</span>
        </td>
        <td class="tooltip" @click="copyRefreshToken">
          {{ truncatedRefreshToken }}
          <span class="tooltip-text">Click to copy</span>
        </td>
        <td class="tooltip" @click="copyUserInfo">
          {{ truncatedUserInfo }}
          <span class="tooltip-text">Click to copy</span>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
  <div v-else>
    <p>You should login to retrieve user attributes.</p>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useAuthStore } from '@/store/AuthenticationStore'

const MAX_TABLE_SYMBOLS = 20

export default defineComponent({
  name: 'InformationView',
  setup() {
    const authenticationStore = useAuthStore()

    const truncatedAccessToken = computed(() => {
      const token = authenticationStore.accessToken
      return token ? token.slice(0, MAX_TABLE_SYMBOLS) : 'N/A'
    })

    const truncatedRefreshToken = computed(() => {
      const token = authenticationStore.refreshToken
      return token ? token.slice(0, MAX_TABLE_SYMBOLS) : 'N/A'
    })

    const truncatedUserInfo = computed(() => {
      const userInfo = authenticationStore.userInfo
      return userInfo ? JSON.stringify(userInfo).slice(0, MAX_TABLE_SYMBOLS) : 'N/A'
    })

    const isAuthenticated = computed(() => {
      return authenticationStore.accessToken != null && authenticationStore.accessToken !== "";
    })

    const copyAccessToken = async () => {
      if (authenticationStore.accessToken) {
        await navigator.clipboard.writeText(authenticationStore.accessToken)
      }
    }

    const copyRefreshToken = async () => {
      if (authenticationStore.refreshToken) {
        await navigator.clipboard.writeText(authenticationStore.refreshToken)
      }
    }

    const copyUserInfo = async () => {
      if (authenticationStore.userInfo) {
        await navigator.clipboard.writeText(JSON.stringify(authenticationStore.userInfo))
      }
    }

    return {
      truncatedAccessToken,
      truncatedRefreshToken,
      truncatedUserInfo,
      isAuthenticated,
      copyAccessToken,
      copyRefreshToken,
      copyUserInfo
    }
  }
})
</script>

<style scoped>
.information {
  padding: 20px;
}

.table-container {
  display: flex;
  justify-content: center;
  width: 100%;
  overflow-x: auto;
}

.data-table {
  width: auto;
  max-width: 100%;
  border-collapse: collapse;
  margin: 1rem 0;
}

.data-table th,
.data-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.data-table th {
  background-color: #f4f4f4;
  font-weight: bold;
}

.data-table td:nth-child(even) {
  background-color: #f9f9f9;
}

.data-table td:hover {
  background-color: #f1f1f1;
}

.tooltip {
  position: relative;
}

.tooltip .tooltip-text {
  visibility: hidden;
  width: 120px;
  background-color: #555;
  color: #fff;
  text-align: center;
  border-radius: 6px;
  padding: 5px;
  position: absolute;
  z-index: 1;
  bottom: 125%;
  left: 50%;
  margin-left: -60px;
  opacity: 0;
  transition: opacity 0.3s;
}

.tooltip .tooltip-text::after {
  content: '';
  position: absolute;
  top: 100%;
  left: 50%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: #555 transparent transparent transparent;
}

.tooltip:hover .tooltip-text {
  visibility: visible;
  opacity: 1;
}
</style>
