import { apiClient } from './client'
import type { CoverLetter, GenerateResult, PaginatedResponse } from '../types'

export const coverLetterApi = {
  generate: async (vacancyId: string): Promise<GenerateResult> => {
    const { data } = await apiClient.post<GenerateResult>(`/vacancies/${vacancyId}/cover-letter`)
    return data
  },

  list: async (params?: { limit?: number; offset?: number; status?: string }): Promise<PaginatedResponse<CoverLetter>> => {
    const { data } = await apiClient.get<PaginatedResponse<CoverLetter>>('/cover-letters', { params })
    return data
  },

  get: async (id: string): Promise<CoverLetter> => {
    const { data } = await apiClient.get<CoverLetter>(`/cover-letters/${id}`)
    return data
  },

  updateText: async (id: string, editedText: string): Promise<CoverLetter> => {
    const { data } = await apiClient.put<CoverLetter>(`/cover-letters/${id}`, { edited_text: editedText })
    return data
  },

  updateStatus: async (id: string, status: string): Promise<CoverLetter> => {
    const { data } = await apiClient.patch<CoverLetter>(`/cover-letters/${id}/status`, { status })
    return data
  },
}
