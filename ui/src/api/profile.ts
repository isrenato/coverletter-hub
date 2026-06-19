import { apiClient } from './client'
import type { CVProfile } from '../types'

export const profileApi = {
  get: async (): Promise<CVProfile> => {
    const { data } = await apiClient.get<CVProfile>('/profile')
    return data
  },

  update: async (profile: Partial<CVProfile>): Promise<CVProfile> => {
    const { data } = await apiClient.put<CVProfile>('/profile', profile)
    return data
  },

  upload: async (file: File): Promise<CVProfile> => {
    const formData = new FormData()
    formData.append('cv', file)
    const { data } = await apiClient.post<CVProfile>('/profile/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    return data
  },
}
