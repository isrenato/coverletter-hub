<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useCoverLetterStore } from '../stores/coverletter'
import CoverLetterCard from '../components/coverletter/CoverLetterCard.vue'

const store = useCoverLetterStore()
const route = useRoute()

onMounted(async () => {
  if (route.query.generate) {
    await store.generate(route.query.generate as string)
  }
  await store.fetchList()
})
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">Cover Letters</h1>

    <div v-if="store.warning" class="mb-4 p-4 bg-yellow-50 border border-yellow-200 rounded text-yellow-800">
      {{ store.warning }}
    </div>

    <div v-if="store.loading" class="text-center py-8 text-gray-500">Loading...</div>

    <div v-else class="space-y-4">
      <CoverLetterCard v-for="cl in store.coverLetters" :key="cl.id" :cover-letter="cl" />
      <p v-if="store.coverLetters.length === 0" class="text-center text-gray-500 py-8">No cover letters yet</p>
    </div>
  </div>
</template>
