# Deployment Guide

## Prerequisites

### Development Environment
- Docker & Docker Compose
- Go 1.21+ (for local development)
- Node.js 20+ (for local development)
- PostgreSQL client (optional, for database access)

### Production Deployment
- Docker Hub account
- Server with Docker installed
- Domain name (optional)

## Environment Variables

### Required GitHub Secrets for CI/CD
Set these in your GitHub repository settings:

```bash
DOCKER_USERNAME=your-dockerhub-username
DOCKER_PASSWORD=your-dockerhub-password-or-token
```

### Application Environment Variables

Create `.env` file in the root directory:

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=examuser
DB_PASSWORD=exampass
DB_NAME=examdb

# MinIO/S3
MINIO_ENDPOINT=localhost:9000
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin
MINIO_USE_SSL=false
MINIO_BUCKET=exam-images

# Backend
BACKEND_PORT=8080
GIN_MODE=debug

# Frontend
FRONTEND_PORT=3000
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Local Development

### 1. Clone Repository
```bash
git clone https://github.com/phuonglv-empire/thidanhgianangluc.git
cd thidanhgianangluc
```

### 2. Start Services with Docker Compose
```bash
# Copy environment file
cp .env.example .env

# Start all services
make docker-up

# Or manually:
docker-compose up -d
```

This will start:
- PostgreSQL (port 5432)
- MinIO (port 9000, console 9001)
- Backend API (port 8080)
- Frontend (port 3000)

### 3. Seed Database (Optional)
```bash
make seed

# Or manually:
cd backend && go run cmd/seeder/main.go
```

### 4. Access Applications
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- API Health: http://localhost:8080/health
- MinIO Console: http://localhost:9001 (minioadmin/minioadmin)

### 5. View Logs
```bash
make docker-logs

# Or for specific service:
docker-compose logs -f backend
docker-compose logs -f frontend
```

## Local Development (Without Docker)

### Backend
```bash
cd backend

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Run database (requires Docker)
docker-compose up -d postgres minio

# Run backend
make run-backend
# Or: go run cmd/api/main.go
```

### Frontend
```bash
cd frontend

# Install dependencies
npm install

# Copy environment file
cp .env.example .env

# Run development server
make run-frontend
# Or: npm run dev
```

## Production Deployment

### Method 1: Using Docker Compose on Server

1. SSH into your server
2. Install Docker and Docker Compose
3. Clone the repository
4. Update environment variables in `.env`
5. Run services:

```bash
docker-compose up -d
```

### Method 2: Using Pre-built Docker Images

1. Pull images from Docker Hub:
```bash
docker pull your-username/exam-backend:latest
docker pull your-username/exam-frontend:latest
```

2. Create `docker-compose.prod.yml`:
```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: always

  minio:
    image: minio/minio:latest
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: ${MINIO_ROOT_USER}
      MINIO_ROOT_PASSWORD: ${MINIO_ROOT_PASSWORD}
    volumes:
      - minio_data:/data
    restart: always

  backend:
    image: your-username/exam-backend:latest
    environment:
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: ${DB_USER}
      DB_PASSWORD: ${DB_PASSWORD}
      DB_NAME: ${DB_NAME}
      MINIO_ENDPOINT: minio:9000
      MINIO_ACCESS_KEY: ${MINIO_ROOT_USER}
      MINIO_SECRET_KEY: ${MINIO_ROOT_PASSWORD}
      SERVER_PORT: 8080
    ports:
      - "8080:8080"
    depends_on:
      - postgres
      - minio
    restart: always

  frontend:
    image: your-username/exam-frontend:latest
    environment:
      NEXT_PUBLIC_API_URL: ${NEXT_PUBLIC_API_URL}
    ports:
      - "3000:3000"
    depends_on:
      - backend
    restart: always

volumes:
  postgres_data:
  minio_data:
```

3. Start services:
```bash
docker-compose -f docker-compose.prod.yml up -d
```

### Method 3: Manual Deployment

#### Backend
```bash
# Build
cd backend
CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Copy to server and run
./api
```

#### Frontend
```bash
# Build
cd frontend
npm run build

# Copy .next folder and run
npm start
```

## CI/CD Pipeline

The GitHub Actions workflow automatically:
1. Builds Docker images on push to `main` or `develop`
2. Runs tests
3. Pushes images to Docker Hub with tags:
   - `latest` (for main branch)
   - Branch name
   - Commit SHA

### Triggering a Deployment

```bash
# Push to main branch
git checkout main
git merge develop
git push origin main

# Images will be automatically built and pushed to Docker Hub
```

## Health Checks

### Backend Health
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

### Database Connection
```bash
docker exec -it exam-postgres psql -U examuser -d examdb -c "SELECT 1"
```

### MinIO Health
```bash
curl http://localhost:9000/minio/health/live
```

## Backup and Restore

### Database Backup
```bash
docker exec exam-postgres pg_dump -U examuser examdb > backup.sql
```

### Database Restore
```bash
cat backup.sql | docker exec -i exam-postgres psql -U examuser examdb
```

### MinIO Backup
```bash
docker run --rm \
  --volumes-from exam-minio \
  -v $(pwd):/backup \
  alpine tar czf /backup/minio-backup.tar.gz /data
```

## Monitoring

### View Application Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend
docker-compose logs -f postgres
```

### Database Queries
```bash
# Connect to database
docker exec -it exam-postgres psql -U examuser examdb

# List tables
\dt

# Count exams
SELECT COUNT(*) FROM exams;

# Count sessions
SELECT COUNT(*) FROM exam_sessions;
```

## Troubleshooting

### Backend won't start
1. Check database connection:
   ```bash
   docker-compose logs postgres
   ```
2. Verify environment variables in `.env`
3. Check backend logs:
   ```bash
   docker-compose logs backend
   ```

### Frontend won't start
1. Check if backend is running
2. Verify `NEXT_PUBLIC_API_URL` in `.env`
3. Check frontend logs:
   ```bash
   docker-compose logs frontend
   ```

### Database connection issues
1. Ensure PostgreSQL is running:
   ```bash
   docker-compose ps postgres
   ```
2. Test connection:
   ```bash
   docker exec exam-postgres pg_isready
   ```

### Port conflicts
If ports are already in use, update in `.env`:
```bash
DB_PORT=5433
BACKEND_PORT=8081
FRONTEND_PORT=3001
MINIO_PORT=9001
```

## Security Considerations

### Production Checklist
- [ ] Change default database credentials
- [ ] Change MinIO root credentials
- [ ] Use strong JWT secret (for future auth)
- [ ] Enable HTTPS/TLS
- [ ] Set up firewall rules
- [ ] Regular database backups
- [ ] Keep Docker images updated
- [ ] Use environment-specific configs
- [ ] Enable rate limiting
- [ ] Add authentication/authorization

## Scaling

### Horizontal Scaling
- Use load balancer (nginx, HAProxy)
- Run multiple backend instances
- Use managed PostgreSQL (AWS RDS, Google Cloud SQL)
- Use S3 instead of MinIO for production

### Vertical Scaling
- Increase Docker container resources
- Optimize database queries
- Add database indexes
- Enable caching (Redis)

## Support

For issues:
1. Check logs: `docker-compose logs -f`
2. Review environment variables
3. Consult README.md and TODO.md
4. Create an issue on GitHub
