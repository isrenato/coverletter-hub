<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth'

const auth = useAuthStore()
const router = useRouter()

onMounted(async () => {
  const token = auth.extractTokenFromHash(window.location.hash)
  if (token) {
    auth.setToken(token)
    window.location.hash = ''
  }
  if (auth.isAuthenticated) {
    await auth.fetchUser()
  }
})
</script>

<template>
  <router-view />
</template>
