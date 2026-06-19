<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useVacancyStore } from '../stores/vacancy'
import { useCoverLetterStore } from '../stores/coverletter'
import StatusBadge from '../components/coverletter/StatusBadge.vue'

const auth = useAuthStore()
const vacancyStore = useVacancyStore()
const clStore = useCoverLetterStore()

const stats = computed(() => ({
  totalVacancies: vacancyStore.total,
  totalCoverLetters: clStore.total,
  drafts: clStore.coverLetters.filter(cl => cl.status === 'draft').length,
  approved: clStore.coverLetters.filter(cl => cl.status === 'approved').length,
}))

onMounted(() => {
  vacancyStore.fetchVacancies({ limit: 5 })
  clStore.fetchList({ limit: 5 })
})
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">Welcome, {{ auth.user?.name }}</h1>

    <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
      <div class="bg-white rounded-lg shadow p-6 text-center">
        <p class="text-3xl font-bold text-blue-600">{{ stats.totalVacancies }}</p>
        <p class="text-sm text-gray-500">Vacancies</p>
      </div>
      <div class="bg-white rounded-lg shadow p-6 text-center">
        <p class="text-3xl font-bold text-purple-600">{{ stats.totalCoverLetters }}</p>
        <p class="text-sm text-gray-500">Cover Letters</p>
      </div>
      <div class="bg-white rounded-lg shadow p-6 text-center">
        <p class="text-3xl font-bold text-yellow-600">{{ stats.drafts }}</p>
        <p class="text-sm text-gray-500">Drafts</p>
      </div>
      <div class="bg-white rounded-lg shadow p-6 text-center">
        <p class="text-3xl font-bold text-green-600">{{ stats.approved }}</p>
        <p class="text-sm text-gray-500">Approved</p>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
      <div>
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-lg font-semibold">Recent Vacancies</h2>
          <router-link to="/vacancies" class="text-blue-600 text-sm hover:underline">View all</router-link>
        </div>
        <div class="space-y-3">
          <router-link
            v-for="v in vacancyStore.vacancies"
            :key="v.id"
            :to="`/vacancies/${v.id}`"
            class="block bg-white rounded shadow p-4 hover:shadow-md transition"
          >
            <p class="font-medium">{{ v.title }}</p>
            <p class="text-sm text-gray-500">{{ v.company }} &middot; {{ v.location }}</p>
          </router-link>
          <p v-if="vacancyStore.vacancies.length === 0" class="text-gray-500 text-sm">No vacancies yet</p>
        </div>
      </div>

      <div>
        <div class="flex justify-between items-center mb-4">
          <h2 class="text-lg font-semibold">Recent Cover Letters</h2>
          <router-link to="/cover-letters" class="text-blue-600 text-sm hover:underline">View all</router-link>
        </div>
        <div class="space-y-3">
          <router-link
            v-for="cl in clStore.coverLetters"
            :key="cl.id"
            :to="`/cover-letters/${cl.id}`"
            class="block bg-white rounded shadow p-4 hover:shadow-md transition"
          >
            <div class="flex justify-between">
              <p class="text-sm text-gray-700 line-clamp-1">{{ cl.generated_text.substring(0, 80) }}...</p>
              <StatusBadge :status="cl.status" />
            </div>
          </router-link>
          <p v-if="clStore.coverLetters.length === 0" class="text-gray-500 text-sm">No cover letters yet</p>
        </div>
      </div>
    </div>
  </div>
</template>
