<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useVacancyStore } from '../stores/vacancy'
import VacancyCard from '../components/vacancy/VacancyCard.vue'
import VacancyForm from '../components/vacancy/VacancyForm.vue'

const store = useVacancyStore()
const showForm = ref(false)

onMounted(() => store.fetchVacancies())

async function handleCreate(data: any) {
  await store.createVacancy(data)
  showForm.value = false
}
</script>

<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Vacancies</h1>
      <button @click="showForm = !showForm" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
        {{ showForm ? 'Cancel' : 'Add Vacancy' }}
      </button>
    </div>

    <VacancyForm v-if="showForm" @submit="handleCreate" class="mb-6" />

    <div v-if="store.loading" class="text-center py-8 text-gray-500">Loading...</div>

    <div v-else class="space-y-4">
      <VacancyCard
        v-for="v in store.vacancies"
        :key="v.id"
        :vacancy="v"
        @delete="store.deleteVacancy($event)"
      />
      <p v-if="store.vacancies.length === 0" class="text-center text-gray-500 py-8">No vacancies yet</p>
    </div>
  </div>
</template>
