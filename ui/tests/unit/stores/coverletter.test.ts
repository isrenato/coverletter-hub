import { describe, it, expect, vi } from 'vitest'
import { useCoverLetterStore } from '../../../src/stores/coverletter'
import { coverLetterDraft, coverLetterApproved } from '../../fixtures/coverletters'

vi.mock('../../../src/api/coverletter', () => ({
  coverLetterApi: {
    generate: vi.fn(),
    list: vi.fn(),
    get: vi.fn(),
    updateText: vi.fn(),
    updateStatus: vi.fn(),
  },
}))

describe('CoverLetter Store', () => {
  it('starts with empty state', () => {
    const store = useCoverLetterStore()
    expect(store.coverLetters).toEqual([])
    expect(store.current).toBeNull()
  })

  it('fetchList loads from API', async () => {
    const { coverLetterApi } = await import('../../../src/api/coverletter')
    vi.mocked(coverLetterApi.list).mockResolvedValue({
      items: [coverLetterDraft, coverLetterApproved],
      total: 2,
    })

    const store = useCoverLetterStore()
    await store.fetchList()

    expect(store.coverLetters).toHaveLength(2)
  })

  it('generate creates draft and sets current', async () => {
    const { coverLetterApi } = await import('../../../src/api/coverletter')
    const result = { ...coverLetterDraft, has_warning: false }
    vi.mocked(coverLetterApi.generate).mockResolvedValue(result)

    const store = useCoverLetterStore()
    await store.generate('vacancy-id')

    expect(store.current).toBeTruthy()
    expect(store.current?.status).toBe('draft')
  })

  it('updateStatus changes status', async () => {
    const { coverLetterApi } = await import('../../../src/api/coverletter')
    const approved = { ...coverLetterDraft, status: 'approved' as const }
    vi.mocked(coverLetterApi.updateStatus).mockResolvedValue(approved)

    const store = useCoverLetterStore()
    store.current = coverLetterDraft
    await store.updateStatus(coverLetterDraft.id, 'approved')

    expect(store.current?.status).toBe('approved')
  })
})
