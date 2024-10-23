import { defineStore } from 'pinia'
import { SecureStorage } from '@/store/SecureStore'
import { UserManager, WebStorageStateStore } from 'oidc-client-ts'
import axios from 'axios'

// CONSTANTS BLOCK
const ls = new SecureStorage()
const USER_INFO = 'userInfo'
const ACCESS_TOKEN = 'accessToken'
const REFRESH_TOKEN = 'refreshToken'
// END CONSTANTS BLOCK

// OIDC-CLIENT CONFIGURATION
const config = {
  authority: 'http://localhost:8080/realms/master',
  client_id: 'default-client',
  redirect_uri: 'http://localhost:9093/callback',
  response_type: 'code',
  scope: 'openid',
  client_secret: 'Qo5FVV8XEwqyon2MLXdaTjmoIjgguxmZ',
  post_logout_redirect_uri: 'http://localhost:9093/',
  userStore: new WebStorageStateStore({ store: ls }),
  automaticSilentRenew: false
}

const userManager = new UserManager(config)
export const login = () => {
  return userManager.signinRedirect()
}
export const logout = () => {
  let promise = userManager.signoutRedirect()
  ls.clear()
  return promise
}
export const handleCallback = () => {
  return userManager.signinRedirectCallback()
}
// END OIDC-CLIENT CONFIGURATION

// DEFINE STORE
interface Authentication {
  accessToken: string | null;
  refreshToken: string | null;
  userInfo: object | null;
}

export const useAuthStore = defineStore('auth', {
  state: (): Authentication => ({
    accessToken: null,
    refreshToken: null,
    userInfo: null
  }),

  actions: {
    setUserInfo(userInfo: object) {
      this.userInfo = userInfo
      ls.setItem(USER_INFO, JSON.stringify(userInfo))
    },

    setAccessToken(accessToken: string) {
      this.accessToken = accessToken
      ls.setItem(ACCESS_TOKEN, accessToken)
    },

    setRefreshToken(refreshToken: string | null) {
      this.refreshToken = refreshToken
      ls.setItem(REFRESH_TOKEN, refreshToken ? refreshToken : '')
    },

    clearAuthentication() {
      this.accessToken = null
      this.refreshToken = null
      this.userInfo = null
      ls.clear()
    },

    async makeRefresh() {
      try {
        const response = await axios.post(
          'http://localhost:8080/realms/master/protocol/openid-connect/token',
          {
            refreshToken: this.refreshToken
          })

        this.setAccessToken(response.data.accessToken)
        this.setRefreshToken(response.data.refreshToken)
      } catch (error) {
        throw new Error('Failed to refresh token')
      }
    },

    loadFromStorage() {
      const userInfo = ls.getItem(USER_INFO)
      this.userInfo = userInfo ? JSON.parse(userInfo) : null
      this.accessToken = ls.getItem(ACCESS_TOKEN)
      this.refreshToken = ls.getItem(REFRESH_TOKEN)
    }
  }
})
// END DEFINE STORE
