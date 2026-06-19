import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { CoverLetter, GenerateResult } from '../types'
import { coverLetterApi } from '../api/coverletter'

export const useCoverLetterStore = defineStore('coverletter', () => {
  const coverLetters = ref<CoverLetter[]>([])
  const current = ref<CoverLetter | null>(null)
  const total = ref(0)
  const loading = ref(false)
  const error = ref<string | null>(null)
  const warning = ref<string | null>(null)

  async function fetchList(params?: { limit?: number; offset?: number; status?: string }) {
    loading.value = true
    error.value = null
    try {
      const result = await coverLetterApi.list(params)
      coverLetters.value = result.items || []
      total.value = result.total
    } catch {
      error.value = 'Failed to load cover letters'
    } finally {
      loading.value = false
    }
  }

  async function fetchOne(id: string) {
    loading.value = true
    try {
      current.value = await coverLetterApi.get(id)
    } catch {
      error.value = 'Failed to load cover letter'
    } finally {
      loading.value = false
    }
  }

  async function generate(vacancyId: string) {
    loading.value = true
    error.value = null
    warning.value = null
    try {
      const result: GenerateResult = await coverLetterApi.generate(vacancyId)
      current.value = result
      if (result.has_warning) {
        warning.value = result.warning || 'Recent cover letter exists for this vacancy'
      }
      return result
    } catch {
      error.value = 'Failed to generate cover letter'
    } finally {
      loading.value = false
    }
  }

  async function updateText(id: string, text: string) {
    try {
      current.value = await coverLetterApi.updateText(id, text)
    } catch {
      error.value = 'Failed to save changes'
    }
  }

  async function updateStatus(id: string, status: string) {
    try {
      current.value = await coverLetterApi.updateStatus(id, status)
    } catch {
      error.value = 'Failed to update status'
    }
  }

  return { coverLetters, current, total, loading, error, warning, fetchList, fetchOne, generate, updateText, updateStatus }
})
