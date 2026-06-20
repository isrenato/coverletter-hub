# CoverLetter Hub

A multi-user SPA for generating tailored cover letters from LinkedIn vacancies using AI.

## Tech Stack

- **Backend:** Go 1.25, chi router, pgx/v5, PostgreSQL 16
- **Frontend:** Vue 3, TypeScript, Vite, Pinia, Tailwind CSS
- **AI:** Claude API (Anthropic) for CV parsing and cover letter generation
- **Auth:** LinkedIn OAuth → JWT
- **Task Runner:** [Task](https://taskfile.dev/) (go-task)

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- [Go](https://go.dev/dl/) 1.25+
- [Node.js](https://nodejs.org/) 20+
- [Task](https://taskfile.dev/installation/) (`brew install go-task`)

## Getting Started

### 1. Clone & configure environment

```bash
git clone git@github.com:isrenato/coverletter-hub.git
cd coverletter-hub
cp .env.example .env
```

Edit `.env` and set your credentials:

```env
POSTGRES_USER=coverletter
POSTGRES_PASSWORD=coverletter_secret
POSTGRES_DB=coverletter_hub
POSTGRES_PORT=5432
API_PORT=8080
UI_PORT=3000

# Optional — needed for full functionality
JWT_SECRET=your-jwt-secret
CLAUDE_API_KEY=sk-ant-xxxxx
LINKEDIN_CLIENT_ID=your-client-id
LINKEDIN_CLIENT_SECRET=your-client-secret
```

### 2. Install dependencies

```bash
cd api && go mod download && cd ..
cd ui && npm install && cd ..
```

### 3. Start development

```bash
task dev
```

This starts PostgreSQL (Docker), the Go API on `:8080`, and the Vue dev server on `:3000`.

## Task Commands

All commands are defined in `Taskfile.yml` and run via [Task](https://taskfile.dev/).

| Command | Description |
|---------|-------------|
| `task dev` | Start DB + API + UI for development |
| `task test:backend` | Run Go backend tests (requires Docker) |
| `task test:frontend` | Run Vue frontend tests |
| `task test:all` | Run all tests |
| `task build:api` | Build Go API binary to `api/bin/server` |
| `task build:ui` | Build Vue SPA to `ui/dist/` |
| `task docker:build` | Build Docker images for all services |
| `task docker:up` | Start all services in Docker |
| `task docker:down` | Stop all Docker services |

## Docker

### Build & run all services

```bash
task docker:build
task docker:up
```

This starts three containers:

| Container | Service | Port |
|-----------|---------|------|
| `coverletter-db` | PostgreSQL 16 | 5432 |
| `coverletter-api` | Go API | 8080 |
| `coverletter-ui` | Vue SPA (nginx) | 3000 |

### Stop services

```bash
task docker:down
```

### View logs

```bash
docker logs coverletter-api
docker logs coverletter-ui
docker logs coverletter-db
```

## Project Structure

```
coverletter-hub/
├── api/                          # Go backend
│   ├── cmd/server/main.go        # Entry point
│   ├── internal/
│   │   ├── auth/                 # LinkedIn OAuth, JWT, middleware
│   │   ├── config/               # Environment config
│   │   ├── coverletter/          # Cover letter generation & CRUD
│   │   ├── database/             # DB connection, migrations
│   │   ├── llm/                  # Claude API client
│   │   ├── model/                # Domain models
│   │   ├── parser/               # CV file parsing via Claude
│   │   ├── profile/              # CV profile management
│   │   ├── user/                 # User repository
│   │   └── vacancy/              # Vacancy CRUD
│   └── testutil/                 # Test helpers & fixtures
├── ui/                           # Vue 3 frontend
│   ├── src/
│   │   ├── api/                  # API client modules
│   │   ├── components/           # Reusable Vue components
│   │   ├── router/               # Vue Router with auth guards
│   │   ├── stores/               # Pinia stores
│   │   ├── types/                # TypeScript interfaces
│   │   └── views/                # Page components
│   └── tests/                    # Vitest unit tests & fixtures
├── docker-compose.yml            # Docker services
├── Taskfile.yml                  # Task runner commands
└── .env.example                  # Environment template
```

## API Endpoints

All endpoints under `/api/v1`. Auth endpoints are public; all others require JWT via `Authorization: Bearer <token>`.

### Authentication
| Method | Path | Description |
|--------|------|-------------|
| GET | `/auth/linkedin` | Initiate LinkedIn OAuth redirect |
| GET | `/auth/linkedin/callback` | Handle OAuth callback, issue JWT |
| GET | `/auth/me` | Get current user info |

### CV Profile
| Method | Path | Description |
|--------|------|-------------|
| GET | `/profile` | Get current user's CV profile |
| PUT | `/profile` | Update CV profile fields |
| POST | `/profile/upload` | Upload CV file (PDF/DOCX), parse via Claude |

### Vacancies
| Method | Path | Description |
|--------|------|-------------|
| GET | `/vacancies` | List vacancies (paginated) |
| POST | `/vacancies` | Create vacancy (manual entry) |
| GET | `/vacancies/:id` | Get vacancy details |
| PUT | `/vacancies/:id` | Update vacancy |
| DELETE | `/vacancies/:id` | Delete vacancy |

### Cover Letters
| Method | Path | Description |
|--------|------|-------------|
| POST | `/vacancies/:id/cover-letter` | Generate cover letter |
| GET | `/cover-letters` | List cover letters (filterable by status) |
| GET | `/cover-letters/:id` | Get cover letter details |
| PUT | `/cover-letters/:id` | Update edited text |
| PATCH | `/cover-letters/:id/status` | Approve or reject |

## Testing

```bash
# Run all tests
task test:all

# Backend only (requires Docker for testcontainers)
task test:backend

# Frontend only
task test:frontend
```

## License

Private — all rights reserved.
