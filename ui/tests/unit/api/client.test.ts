import { describe, it, expect, vi, beforeEach } from 'vitest'
import axios from 'axios'

vi.mock('axios', () => {
  const instance = {
    interceptors: {
      request: { use: vi.fn() },
      response: { use: vi.fn() },
    },
    defaults: { headers: { common: {} } },
    get: vi.fn(),
    post: vi.fn(),
    put: vi.fn(),
    patch: vi.fn(),
    delete: vi.fn(),
  }
  return {
    default: { create: vi.fn(() => instance) },
  }
})

describe('API Client', () => {
  beforeEach(() => {
    vi.resetModules()
    if (typeof localStorage !== 'undefined') {
      localStorage.clear()
    }
  })

  it('creates axios instance with base URL', async () => {
    const { apiClient } = await import('../../../src/api/client')
    expect(apiClient).toBeDefined()
    expect(axios.create).toHaveBeenCalledWith(
      expect.objectContaining({ baseURL: expect.any(String) })
    )
  })

  it('registers request interceptor for JWT', async () => {
    const { apiClient } = await import('../../../src/api/client')
    expect(apiClient.interceptors.request.use).toHaveBeenCalled()
  })

  it('registers response interceptor for 401 handling', async () => {
    const { apiClient } = await import('../../../src/api/client')
    expect(apiClient.interceptors.response.use).toHaveBeenCalled()
  })
})
