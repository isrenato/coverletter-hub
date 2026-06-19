<script setup lang="ts">
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

function logout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <nav class="bg-white shadow">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16">
        <div class="flex space-x-8">
          <router-link to="/dashboard" class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900">
            Dashboard
          </router-link>
          <router-link to="/profile" class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-500 hover:text-gray-900">
            Profile
          </router-link>
          <router-link to="/vacancies" class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-500 hover:text-gray-900">
            Vacancies
          </router-link>
          <router-link to="/cover-letters" class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-500 hover:text-gray-900">
            Cover Letters
          </router-link>
        </div>
        <div class="flex items-center space-x-4">
          <span class="text-sm text-gray-600">{{ auth.user?.name }}</span>
          <button @click="logout" class="text-sm text-red-600 hover:text-red-800">Logout</button>
        </div>
      </div>
    </div>
  </nav>
</template>
