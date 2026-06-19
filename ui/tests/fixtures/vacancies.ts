import type { Vacancy } from '../../src/types'

export const vacancyBackend: Vacancy = {
  id: 'd4e5f6a7-b8c9-0123-defa-234567890123',
  user_id: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
  title: 'Backend Engineer',
  company: 'StartupCo',
  description: 'We are looking for a backend engineer proficient in Go and PostgreSQL.',
  location: 'Amsterdam, NL',
  linkedin_url: 'https://linkedin.com/jobs/12345',
  source: 'manual',
  created_at: '2026-06-01T00:00:00Z',
}

export const vacancyFrontend: Vacancy = {
  id: 'e5f6a7b8-c9d0-1234-efab-345678901234',
  user_id: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
  title: 'Frontend Developer',
  company: 'DesignLab',
  description: 'React/Vue developer needed for our design tools platform.',
  location: 'Remote',
  source: 'manual',
  created_at: '2026-06-10T00:00:00Z',
}
