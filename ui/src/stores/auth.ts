import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { User } from '../types'
import { authApi } from '../api/auth'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(null)

  const isAuthenticated = computed(() => !!token.value)

  function setToken(newToken: string) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
  }

  async function fetchUser() {
    if (!token.value) return
    user.value = await authApi.getMe()
  }

  function extractTokenFromHash(hash: string): string | null {
    const match = hash.match(/token=([^&]+)/)
    return match ? match[1] : null
  }

  return { token, user, isAuthenticated, setToken, logout, fetchUser, extractTokenFromHash }
})
