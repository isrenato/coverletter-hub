import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Vacancy } from '../types'
import { vacancyApi } from '../api/vacancy'

export const useVacancyStore = defineStore('vacancy', () => {
  const vacancies = ref<Vacancy[]>([])
  const current = ref<Vacancy | null>(null)
  const total = ref(0)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchVacancies(params?: { limit?: number; offset?: number }) {
    loading.value = true
    error.value = null
    try {
      const result = await vacancyApi.list(params)
      vacancies.value = result.items || []
      total.value = result.total
    } catch {
      error.value = 'Failed to load vacancies'
    } finally {
      loading.value = false
    }
  }

  async function fetchVacancy(id: string) {
    loading.value = true
    try {
      current.value = await vacancyApi.get(id)
    } catch {
      error.value = 'Failed to load vacancy'
    } finally {
      loading.value = false
    }
  }

  async function createVacancy(data: Partial<Vacancy>) {
    loading.value = true
    try {
      const created = await vacancyApi.create(data)
      vacancies.value.unshift(created)
      total.value++
      return created
    } catch {
      error.value = 'Failed to create vacancy'
    } finally {
      loading.value = false
    }
  }

  async function deleteVacancy(id: string) {
    try {
      await vacancyApi.remove(id)
      vacancies.value = vacancies.value.filter(v => v.id !== id)
      total.value--
    } catch {
      error.value = 'Failed to delete vacancy'
    }
  }

  return { vacancies, current, total, loading, error, fetchVacancies, fetchVacancy, createVacancy, deleteVacancy }
})
