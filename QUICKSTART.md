# Quick Start Guide

Get the Online Exam System running in 5 minutes!

## Prerequisites

- Docker Desktop installed
- Git installed
- 8GB RAM available

## Step-by-Step Setup

### 1. Clone the Repository
```bash
git clone https://github.com/phuonglv-empire/thidanhgianangluc.git
cd thidanhgianangluc
```

### 2. Configure Environment
```bash
cp .env.example .env
```

### 3. Start All Services
```bash
docker-compose up -d
```

Wait 30-60 seconds for all services to start.

### 4. Seed the Database (Optional)
```bash
docker-compose exec backend go run cmd/seeder/main.go
```

### 5. Access the Applications

| Service | URL | Credentials |
|---------|-----|-------------|
| Frontend | http://localhost:3000 | - |
| Backend API | http://localhost:8080 | - |
| API Health | http://localhost:8080/health | - |
| MinIO Console | http://localhost:9001 | minioadmin / minioadmin |

## Test the API

### Health Check
```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "ok",
  "message": "Exam API is running"
}
```

### Get All Exams
```bash
curl http://localhost:8080/api/exams
```

### Get Exam Details
```bash
curl http://localhost:8080/api/exams/1
```

### Start Exam Session
```bash
curl -X POST http://localhost:8080/api/exam-sessions \
  -H "Content-Type: application/json" \
  -d '{
    "exam_id": 1,
    "user_id": "test-user-001"
  }'
```

### Submit Answer
```bash
curl -X POST http://localhost:8080/api/exam-sessions/1/answers \
  -H "Content-Type: application/json" \
  -d '{
    "question_id": 1,
    "user_answer": "f'\''(x) = 2x + 2"
  }'
```

### Submit Exam
```bash
curl -X POST http://localhost:8080/api/exam-sessions/1/submit
```

### Get Results
```bash
curl http://localhost:8080/api/exam-sessions/1/results
```

## Common Commands

### View Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
```

### Stop Services
```bash
docker-compose down
```

### Restart Services
```bash
docker-compose restart
```

### Rebuild After Code Changes
```bash
docker-compose up -d --build
```

### Access Database
```bash
docker-compose exec postgres psql -U examuser examdb
```

### Clean Up Everything
```bash
docker-compose down -v  # Removes volumes too
make clean
```

## Development Workflow

### Backend Development
```bash
cd backend

# Install dependencies
go mod download

# Run locally (requires Docker for database)
docker-compose up -d postgres minio
go run cmd/api/main.go

# Build
make build-backend
```

### Frontend Development
```bash
cd frontend

# Install dependencies
npm install

# Run dev server
npm run dev

# Build
npm run build
```

## Troubleshooting

### Port Already in Use
Update ports in `.env`:
```bash
BACKEND_PORT=8081
FRONTEND_PORT=3001
DB_PORT=5433
```

### Backend Won't Start
Check logs:
```bash
docker-compose logs backend
```

Verify database is running:
```bash
docker-compose ps postgres
```

### Frontend Won't Connect to Backend
1. Check backend is running: `curl http://localhost:8080/health`
2. Verify `NEXT_PUBLIC_API_URL` in `frontend/.env`
3. Check CORS settings in backend

### Database Issues
Reset database:
```bash
docker-compose down -v
docker-compose up -d postgres
make seed
```

## Next Steps

1. **Explore the API**: Use the endpoints above to test functionality
2. **Read Documentation**: 
   - `README.md` - Overview and architecture
   - `TODO.md` - Development roadmap
   - `DEPLOYMENT.md` - Deployment guide
   - `PROJECT_SUMMARY.md` - What's been built
3. **Start Developing**:
   - Frontend: Implement UI components in `frontend/src/components/`
   - Backend: Add features in respective layers
4. **Run Tests**: `make test` (when tests are added)

## Getting Help

- Check logs: `docker-compose logs -f [service]`
- Review documentation files
- Check GitHub issues
- Verify environment variables

## Quick Reference

### Useful URLs
```
Frontend:      http://localhost:3000
Backend API:   http://localhost:8080
API Docs:      http://localhost:8080/api
Health Check:  http://localhost:8080/health
MinIO Console: http://localhost:9001
```

### Default Credentials
```
Database:
  User: examuser
  Pass: exampass
  DB:   examdb

MinIO:
  User: minioadmin
  Pass: minioadmin
```

### Key Directories
```
backend/cmd/api/           - Main application
backend/internal/domain/   - Business entities
backend/internal/usecase/  - Business logic
frontend/src/app/          - Pages
frontend/src/components/   - React components
frontend/src/lib/          - Utilities
```

---

**You're all set!** 🎉

The exam system is now running. Start exploring the API or begin frontend development.
