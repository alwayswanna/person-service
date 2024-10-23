<template>
  <div>Loading...</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useAuthStore, handleCallback } from '@/store/AuthenticationStore';

export default defineComponent({
  name: 'CallbackView',
  setup() {
    const authStore = useAuthStore();

    handleCallback().then((user) => {
      if (user) {
        authStore.setAccessToken(user.access_token);
        authStore.setRefreshToken(user.refresh_token ? user.refresh_token : null)
        authStore.setUserInfo(user)
        window.location.href = '/';
      }
    });

    return {};
  },
});
</script>
