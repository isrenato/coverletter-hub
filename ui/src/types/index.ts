export interface User {
  id: string
  linkedin_id: string
  email: string
  name: string
  created_at: string
  updated_at: string
}

export interface CVProfile {
  id: string
  user_id: string
  full_name: string
  headline: string
  summary: string
  experience: Experience[]
  education: Education[]
  skills: string[]
  languages: string[]
  created_at: string
  updated_at: string
}

export interface Experience {
  title: string
  company: string
  start: string
  end: string
}

export interface Education {
  degree: string
  school: string
  year: string
}

export interface Vacancy {
  id: string
  user_id: string
  title: string
  company: string
  description: string
  location: string
  linkedin_url?: string
  source: string
  created_at: string
}

export interface CoverLetter {
  id: string
  vacancy_id: string
  cv_profile_id: string
  generated_text: string
  edited_text: string
  status: 'draft' | 'approved' | 'rejected'
  generated_at: string
  approved_at?: string
  created_at: string
  updated_at: string
}

export interface GenerateResult extends CoverLetter {
  has_warning?: boolean
  warning?: string
}

export interface PaginatedResponse<T> {
  items: T[]
  total: number
}
