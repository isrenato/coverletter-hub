import { describe, it, expect, vi } from 'vitest'
import { useVacancyStore } from '../../../src/stores/vacancy'
import { vacancyBackend, vacancyFrontend } from '../../fixtures/vacancies'

vi.mock('../../../src/api/vacancy', () => ({
  vacancyApi: {
    list: vi.fn(),
    get: vi.fn(),
    create: vi.fn(),
    update: vi.fn(),
    remove: vi.fn(),
  },
}))

describe('Vacancy Store', () => {
  it('starts with empty list', () => {
    const store = useVacancyStore()
    expect(store.vacancies).toEqual([])
    expect(store.total).toBe(0)
  })

  it('fetchVacancies loads from API', async () => {
    const { vacancyApi } = await import('../../../src/api/vacancy')
    vi.mocked(vacancyApi.list).mockResolvedValue({
      items: [vacancyBackend, vacancyFrontend],
      total: 2,
    })

    const store = useVacancyStore()
    await store.fetchVacancies()

    expect(store.vacancies).toHaveLength(2)
    expect(store.total).toBe(2)
  })

  it('createVacancy adds to list', async () => {
    const { vacancyApi } = await import('../../../src/api/vacancy')
    vi.mocked(vacancyApi.create).mockResolvedValue(vacancyBackend)

    const store = useVacancyStore()
    await store.createVacancy({
      title: 'Backend Engineer',
      company: 'StartupCo',
      description: 'Looking for BE dev',
      location: 'Amsterdam',
    })

    expect(vacancyApi.create).toHaveBeenCalled()
  })
})
