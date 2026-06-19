import type { CoverLetter } from '../../src/types'

export const coverLetterDraft: CoverLetter = {
  id: 'f6a7b8c9-d0e1-2345-fabc-456789012345',
  vacancy_id: 'd4e5f6a7-b8c9-0123-defa-234567890123',
  cv_profile_id: 'c3d4e5f6-a7b8-9012-cdef-123456789012',
  generated_text: 'Dear Hiring Manager,\n\nI am writing to express my interest...',
  edited_text: '',
  status: 'draft',
  generated_at: '2026-06-01T12:00:00Z',
  created_at: '2026-06-01T12:00:00Z',
  updated_at: '2026-06-01T12:00:00Z',
}

export const coverLetterApproved: CoverLetter = {
  id: 'a7b8c9d0-e1f2-3456-abcd-567890123456',
  vacancy_id: 'e5f6a7b8-c9d0-1234-efab-345678901234',
  cv_profile_id: 'c3d4e5f6-a7b8-9012-cdef-123456789012',
  generated_text: 'Dear Hiring Manager,\n\nI am excited to apply...',
  edited_text: 'Dear Hiring Manager,\n\nI am thrilled to apply...',
  status: 'approved',
  generated_at: '2026-01-15T09:00:00Z',
  approved_at: '2026-01-15T10:00:00Z',
  created_at: '2026-01-15T09:00:00Z',
  updated_at: '2026-01-15T10:00:00Z',
}
