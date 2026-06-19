import type { CVProfile } from '../../src/types'

export const profileJohn: CVProfile = {
  id: 'c3d4e5f6-a7b8-9012-cdef-123456789012',
  user_id: 'a1b2c3d4-e5f6-7890-abcd-ef1234567890',
  full_name: 'John Doe',
  headline: 'Senior Software Engineer',
  summary: '10 years of experience building web applications.',
  experience: [{ title: 'Senior Engineer', company: 'TechCorp', start: '2020-01', end: 'present' }],
  education: [{ degree: 'BSc Computer Science', school: 'MIT', year: '2015' }],
  skills: ['Go', 'TypeScript', 'PostgreSQL', 'Docker'],
  languages: ['English', 'Dutch'],
  created_at: '2026-01-01T00:00:00Z',
  updated_at: '2026-01-01T00:00:00Z',
}
