import axios from 'axios'
import { useAuthStore } from '@/store/AuthenticationStore'

const client = axios.create({
  baseURL: 'http://localhost:9902',
  headers: {
    'Content-Type': 'application/json'
  }
});

client.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore()
    const token = authStore.accessToken;

    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    console.log("axios.config", config.headers)

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
);

client.interceptors.response.use(
  (response) => {
    return response
  },
  async (error) => {
    const originalRequest = error.config
    const authStore = useAuthStore()
    let refreshToken = authStore.refreshToken
    if (refreshToken == null || refreshToken !== "") {
      return
    }

    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;

      try {
        await  authStore.makeRefresh()
        const newToken = authStore.accessToken;

        originalRequest.headers.Authorization = `Bearer ${newToken}`
        return client(originalRequest)
      } catch (refreshErr) {
        await authStore.clearAuthentication();
        return Promise.reject(refreshErr)
      }
    }
  }
)

export default client
