import { describe, it, expect, vi, beforeEach } from 'vitest'
import { useAuthStore } from '../../../src/stores/auth'
import { userJohn } from '../../fixtures/users'

vi.mock('../../../src/api/auth', () => ({
  authApi: {
    getMe: vi.fn(),
  },
}))

describe('Auth Store', () => {
  beforeEach(() => {
    localStorage.clear()
  })

  it('starts unauthenticated', () => {
    const store = useAuthStore()
    expect(store.isAuthenticated).toBe(false)
    expect(store.user).toBeNull()
  })

  it('sets token and marks as authenticated', () => {
    const store = useAuthStore()
    store.setToken('test-token')
    expect(store.isAuthenticated).toBe(true)
    expect(localStorage.getItem('token')).toBe('test-token')
  })

  it('logout clears state', () => {
    const store = useAuthStore()
    store.setToken('test-token')
    store.user = userJohn
    store.logout()
    expect(store.isAuthenticated).toBe(false)
    expect(store.user).toBeNull()
    expect(localStorage.getItem('token')).toBeNull()
  })

  it('fetchUser populates user from API', async () => {
    const { authApi } = await import('../../../src/api/auth')
    vi.mocked(authApi.getMe).mockResolvedValue(userJohn)

    const store = useAuthStore()
    store.setToken('test-token')
    await store.fetchUser()

    expect(store.user).toEqual(userJohn)
  })

  it('extractTokenFromHash reads URL hash', () => {
    const store = useAuthStore()
    const token = store.extractTokenFromHash('#token=jwt-value-here')
    expect(token).toBe('jwt-value-here')
  })
})
