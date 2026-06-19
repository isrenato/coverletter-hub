<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useVacancyStore } from '../stores/vacancy'

const route = useRoute()
const router = useRouter()
const store = useVacancyStore()

onMounted(() => {
  store.fetchVacancy(route.params.id as string)
})

async function generateCoverLetter() {
  router.push({ name: 'cover-letters', query: { generate: route.params.id as string } })
}
</script>

<template>
  <div v-if="store.current">
    <div class="mb-6">
      <router-link to="/vacancies" class="text-blue-600 hover:underline text-sm">&larr; Back to Vacancies</router-link>
    </div>

    <div class="bg-white rounded-lg shadow p-6">
      <h1 class="text-2xl font-bold">{{ store.current.title }}</h1>
      <p class="text-lg text-gray-600">{{ store.current.company }}</p>
      <p class="text-sm text-gray-500 mb-4">{{ store.current.location }}</p>

      <div class="prose max-w-none mb-6">
        <p class="whitespace-pre-wrap">{{ store.current.description }}</p>
      </div>

      <button @click="generateCoverLetter" class="px-6 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 font-medium">
        Generate Cover Letter
      </button>
    </div>
  </div>
  <div v-else-if="store.loading" class="text-center py-8 text-gray-500">Loading...</div>
</template>
