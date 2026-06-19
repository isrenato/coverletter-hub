import { apiClient } from './client'
import type { User } from '../types'

export const authApi = {
  getMe: async (): Promise<User> => {
    const { data } = await apiClient.get<User>('/auth/me')
    return data
  },

  getLinkedInAuthURL: (): string => {
    const baseURL = import.meta.env.VITE_API_BASE_URL || ''
    return `${baseURL}/api/v1/auth/linkedin`
  },
}
