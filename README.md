# CoverLetter Hub

A multi-user SPA for generating tailored cover letters from LinkedIn vacancies using AI.

## Tech Stack

- **Backend:** Go 1.22+, chi, pgx/v5, PostgreSQL 16
- **Frontend:** Vue 3, TypeScript, Pinia, Tailwind CSS
- **AI:** Claude API (Anthropic) for CV parsing and cover letter generation
- **Auth:** LinkedIn OAuth → JWT

## Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.22+
- Node 20+

### Development

1. Copy environment files:
   ```bash
   cp .env.example .env
   cp api/.env.example api/.env
   cp ui/.env.example ui/.env
   ```

2. Start PostgreSQL:
   ```bash
   docker compose up -d db
   ```

3. Run API:
   ```bash
   cd api && go run ./cmd/server/
   ```

4. Run UI:
   ```bash
   cd ui && npm install && npm run dev
   ```

### Production

```bash
docker compose up -d
```

## Project Structure

```
coverletter-hub/
├── api/          # Go backend
├── ui/           # Vue 3 frontend
├── docs/         # Design specs & plans (not committed)
└── docker-compose.yml
```

## API Endpoints

All endpoints under `/api/v1`. Auth endpoints are public; all others require JWT.

| Method | Path | Description |
|--------|------|-------------|
| GET | /auth/linkedin | Initiate LinkedIn OAuth |
| GET | /auth/linkedin/callback | OAuth callback |
| GET | /auth/me | Current user info |
| GET/PUT | /profile | CV profile |
| POST | /profile/upload | Upload CV (PDF/DOCX) |
| GET/POST | /vacancies | List/create vacancies |
| GET/PUT/DELETE | /vacancies/:id | Vacancy CRUD |
| POST | /vacancies/:id/cover-letter | Generate cover letter |
| GET | /cover-letters | List cover letters |
| GET/PUT | /cover-letters/:id | Get/edit cover letter |
| PATCH | /cover-letters/:id/status | Approve/reject |

## Testing

```bash
# Backend (requires Docker for testcontainers)
cd api && go test ./... -v

# Frontend
cd ui && npm test
```
