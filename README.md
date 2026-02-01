# Online Exam System (HSA, VACT, THPT)

A modern online exam platform for practicing HSA, VACT, and THPT exams with a focus on user experience.

## 🚀 Tech Stack

### Backend
- **Language**: Golang 1.21+
- **Framework**: Gin (HTTP router)
- **ORM**: GORM
- **Database**: PostgreSQL 15
- **Storage**: MinIO (S3-compatible)

### Frontend
- **Framework**: Next.js 15+ (App Router)
- **UI**: TailwindCSS + shadcn UI
- **State Management**: Zustand
- **Language**: TypeScript

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **CI/CD**: GitHub Actions
- **Registry**: Docker Hub

## 📁 Project Structure

```
├── backend/                    # Golang backend
│   ├── cmd/api/               # Application entry point
│   ├── internal/
│   │   ├── domain/            # Domain entities
│   │   ├── repository/        # Data access layer
│   │   ├── usecase/           # Business logic
│   │   └── delivery/http/     # HTTP handlers
│   ├── pkg/
│   │   ├── database/          # Database utilities
│   │   ├── logger/            # Logging utilities
│   │   └── s3/                # S3/MinIO client
│   └── scripts/               # Database scripts
├── frontend/                   # Next.js frontend
│   └── src/
│       ├── app/               # App Router pages
│       ├── components/        # React components (Atomic Design)
│       │   ├── atoms/         # Basic components
│       │   ├── molecules/     # Composite components
│       │   ├── organisms/     # Complex components
│       │   └── templates/     # Page templates
│       ├── lib/               # Utilities & API client
│       ├── stores/            # Zustand stores
│       └── types/             # TypeScript types
├── .github/workflows/         # CI/CD workflows
├── docker-compose.yml         # Local development setup
└── TODO.md                    # Detailed task tracking
```

## 🏗️ Architecture

### Backend - Clean Architecture
- **Domain Layer**: Business entities and rules
- **Repository Layer**: Database operations and data access
- **UseCase Layer**: Application business logic
- **Delivery Layer**: HTTP handlers and routing

### Frontend - Atomic Design
- **Atoms**: Basic UI elements (buttons, inputs)
- **Molecules**: Simple component groups
- **Organisms**: Complex UI sections
- **Templates**: Page layouts
- **Pages**: Actual routes and views

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.21+ (for local development)
- Node.js 20+ (for local development)

### Setup & Run with Docker

1. Clone the repository:
```bash
git clone https://github.com/phuonglv-empire/thidanhgianangluc.git
cd thidanhgianangluc
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Start all services:
```bash
docker-compose up -d
```

4. Access the application:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- MinIO Console: http://localhost:9001

### Local Development

#### Backend
```bash
cd backend
cp .env.example .env
go mod download
go run cmd/api/main.go
```

#### Frontend
```bash
cd frontend
cp .env.example .env
npm install
npm run dev
```

## 🔧 Configuration

### Environment Variables

See `.env.example` for all available configuration options.

Key variables:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` - Database connection
- `MINIO_ENDPOINT`, `MINIO_ACCESS_KEY`, `MINIO_SECRET_KEY` - S3 storage
- `NEXT_PUBLIC_API_URL` - Frontend API endpoint

## 📦 Docker Deployment

### Building Images
```bash
# Backend
docker build -t exam-backend ./backend

# Frontend
docker build -t exam-frontend ./frontend
```

### Push to Docker Hub
```bash
docker tag exam-backend your-username/exam-backend:latest
docker push your-username/exam-backend:latest

docker tag exam-frontend your-username/exam-frontend:latest
docker push your-username/exam-frontend:latest
```

## 🤖 CI/CD

GitHub Actions automatically:
- Builds Docker images on push to `main` or `develop`
- Runs tests
- Pushes images to Docker Hub with version tags

### Required GitHub Secrets
- `DOCKER_USERNAME`: Docker Hub username
- `DOCKER_PASSWORD`: Docker Hub password/token

## 📝 Database Schema

### Core Entities
- **Exams**: Exam metadata and content
- **Questions**: Individual questions with JSONB options
- **ExamSessions**: User exam attempts
- **Answers**: User responses to questions

## 🎯 Features

### Current Features
- ✅ Clean Architecture backend structure
- ✅ Docker containerization
- ✅ CI/CD pipeline
- ✅ Database migrations
- ✅ Basic API endpoints
- ✅ Next.js App Router setup
- ✅ TailwindCSS styling

### Planned Features (See TODO.md)
- 🔄 Exam listing API
- 🔄 Exam session management
- 🔄 Answer submission
- 🔄 Real-time timer
- 🔄 Results & statistics
- 🔄 Exam runner UI

## 📖 API Documentation

### Endpoints

#### Health Check
- `GET /health` - Service health status

#### Exams
- `GET /api/exams` - List available exams
- `GET /api/exams/:id` - Get exam details with questions

#### Exam Sessions
- `POST /api/exam-sessions` - Start a new exam session
- `GET /api/exam-sessions/:id` - Get session details
- `POST /api/exam-sessions/:id/answers` - Submit answer
- `POST /api/exam-sessions/:id/submit` - Submit entire exam
- `GET /api/exam-sessions/:id/results` - Get exam results

## 🧪 Testing

### Backend
```bash
cd backend
go test ./...
```

### Frontend
```bash
cd frontend
npm test
```

## 📄 License

This project is private and proprietary.

## 🤝 Contributing

For development workflow and task tracking, see [TODO.md](./TODO.md)

## 📞 Support

For issues and questions, please create an issue in the repository.