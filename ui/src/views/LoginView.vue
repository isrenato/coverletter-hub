<script setup lang="ts">
import { authApi } from '../api/auth'

function loginWithLinkedIn() {
  window.location.href = authApi.getLinkedInAuthURL()
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50">
    <div class="max-w-md w-full space-y-8 p-8 bg-white rounded-lg shadow">
      <div class="text-center">
        <h1 class="text-3xl font-bold text-gray-900">CoverLetter Hub</h1>
        <p class="mt-2 text-gray-600">Generate tailored cover letters with AI</p>
      </div>
      <button
        @click="loginWithLinkedIn"
        class="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 font-medium"
      >
        Sign in with LinkedIn
      </button>
    </div>
  </div>
</template>
