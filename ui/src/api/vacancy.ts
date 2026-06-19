import { apiClient } from './client'
import type { Vacancy, PaginatedResponse } from '../types'

export const vacancyApi = {
  list: async (params?: { limit?: number; offset?: number }): Promise<PaginatedResponse<Vacancy>> => {
    const { data } = await apiClient.get<PaginatedResponse<Vacancy>>('/vacancies', { params })
    return data
  },

  get: async (id: string): Promise<Vacancy> => {
    const { data } = await apiClient.get<Vacancy>(`/vacancies/${id}`)
    return data
  },

  create: async (vacancy: Partial<Vacancy>): Promise<Vacancy> => {
    const { data } = await apiClient.post<Vacancy>('/vacancies', vacancy)
    return data
  },

  update: async (id: string, vacancy: Partial<Vacancy>): Promise<Vacancy> => {
    const { data } = await apiClient.put<Vacancy>(`/vacancies/${id}`, vacancy)
    return data
  },

  remove: async (id: string): Promise<void> => {
    await apiClient.delete(`/vacancies/${id}`)
  },
}
