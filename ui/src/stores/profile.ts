import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { CVProfile } from '../types'
import { profileApi } from '../api/profile'

export const useProfileStore = defineStore('profile', () => {
  const profile = ref<CVProfile | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchProfile() {
    loading.value = true
    error.value = null
    try {
      profile.value = await profileApi.get()
    } catch (e: any) {
      if (e.response?.status !== 404) {
        error.value = 'Failed to load profile'
      }
    } finally {
      loading.value = false
    }
  }

  async function updateProfile(data: Partial<CVProfile>) {
    loading.value = true
    error.value = null
    try {
      profile.value = await profileApi.update(data)
    } catch {
      error.value = 'Failed to update profile'
    } finally {
      loading.value = false
    }
  }

  async function uploadCV(file: File) {
    loading.value = true
    error.value = null
    try {
      profile.value = await profileApi.upload(file)
    } catch {
      error.value = 'Failed to upload CV'
    } finally {
      loading.value = false
    }
  }

  return { profile, loading, error, fetchProfile, updateProfile, uploadCV }
})
