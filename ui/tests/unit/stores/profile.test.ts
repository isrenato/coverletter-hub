import { describe, it, expect, vi, beforeEach } from 'vitest'
import { useProfileStore } from '../../../src/stores/profile'
import { profileJohn } from '../../fixtures/profiles'

vi.mock('../../../src/api/profile', () => ({
  profileApi: {
    get: vi.fn(),
    update: vi.fn(),
    upload: vi.fn(),
  },
}))

describe('Profile Store', () => {
  it('starts with null profile', () => {
    const store = useProfileStore()
    expect(store.profile).toBeNull()
  })

  it('fetchProfile loads from API', async () => {
    const { profileApi } = await import('../../../src/api/profile')
    vi.mocked(profileApi.get).mockResolvedValue(profileJohn)

    const store = useProfileStore()
    await store.fetchProfile()

    expect(store.profile).toEqual(profileJohn)
  })

  it('updateProfile sends to API and updates local state', async () => {
    const { profileApi } = await import('../../../src/api/profile')
    const updated = { ...profileJohn, headline: 'Staff Engineer' }
    vi.mocked(profileApi.update).mockResolvedValue(updated)

    const store = useProfileStore()
    store.profile = profileJohn
    await store.updateProfile(updated)

    expect(store.profile?.headline).toBe('Staff Engineer')
  })
})
