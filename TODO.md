# TODO - Online Exam System (HSA, VACT, THPT)

## Project Overview
Building an online exam system focused on User Portal experience with Clean Architecture (Backend) and Atomic Design (Frontend).

**Tech Stack:**
- Backend: Golang (Gin/Fiber), GORM, PostgreSQL
- Frontend: Next.js 16+ (App Router), TailwindCSS, Zustand, shadcn UI
- Infrastructure: Docker, Docker Hub, GitHub Actions, S3/MinIO

---

## Phase 1: Infrastructure Setup
### 1.1 Project Structure
- [x] Create TODO.md
- [ ] Create backend directory structure (Clean Architecture)
- [ ] Create frontend directory structure (Atomic Design)
- [ ] Set up .gitignore files
- [ ] Create README documentation

### 1.2 Docker Configuration
- [ ] Create backend/Dockerfile
- [ ] Create frontend/Dockerfile
- [ ] Create docker-compose.yml (PostgreSQL, MinIO, Backend, Frontend)
- [ ] Create .env.example files

### 1.3 Development Environment
- [ ] Configure MinIO for local S3 storage
- [ ] Set up PostgreSQL with initial database
- [ ] Configure environment variables
- [ ] Create database initialization scripts

---

## Phase 2: Backend Core (Golang + Clean Architecture)
### 2.1 Project Foundation
- [ ] Initialize Go module
- [ ] Set up Clean Architecture layers:
  - [ ] Domain (Entities)
  - [ ] Repository (Data Access)
  - [ ] UseCase (Business Logic)
  - [ ] Delivery (HTTP Handlers)
- [ ] Configure GORM with PostgreSQL
- [ ] Set up routing (Gin or Fiber)

### 2.2 Database Schema & Models
- [ ] Create Exam entity (ID, Title, Description, Duration, Content HTML)
- [ ] Create Question entity (ID, Content HTML, Type, Options JSONB, Answer JSONB, ExamID)
- [ ] Create ExamSession entity (ID, UserID, ExamID, StartTime, EndTime, Status)
- [ ] Create Answer entity (ID, SessionID, QuestionID, UserAnswer JSONB, IsCorrect)
- [ ] Create migration files
- [ ] Set up database seeder for testing

### 2.3 Core APIs - Exam Management
- [ ] GET /api/exams - List available exams
- [ ] GET /api/exams/:id - Get exam details with questions
- [ ] POST /api/exam-sessions - Start an exam session
- [ ] GET /api/exam-sessions/:id - Get session details
- [ ] POST /api/exam-sessions/:id/answers - Submit answer for a question
- [ ] POST /api/exam-sessions/:id/submit - Submit entire exam
- [ ] GET /api/exam-sessions/:id/results - Get exam results

### 2.4 S3/MinIO Integration
- [ ] Configure MinIO client
- [ ] Create image upload utility
- [ ] Image URL generation for exam content

### 2.5 Middleware & Error Handling
- [ ] CORS middleware
- [ ] Request logging
- [ ] Error handling middleware
- [ ] Request validation

---

## Phase 3: Frontend Runner (Next.js + Atomic Design)
### 3.1 Project Foundation
- [ ] Initialize Next.js 16+ with App Router
- [ ] Configure TailwindCSS
- [ ] Set up shadcn UI components
- [ ] Configure Zustand for state management
- [ ] Set up TypeScript strict mode

### 3.2 Atomic Design Structure
- [ ] Create atoms (Button, Input, Card, Badge, etc.)
- [ ] Create molecules (QuestionCard, Timer, AnswerOption, etc.)
- [ ] Create organisms (ExamHeader, QuestionList, NavigationPanel, etc.)
- [ ] Create templates (ExamLayout, ResultLayout, etc.)
- [ ] Create pages (Exam List, Exam Runner, Results, etc.)

### 3.3 State Management (Zustand)
- [ ] Create exam store (current exam, questions)
- [ ] Create session store (session ID, answers, time remaining)
- [ ] Create UI store (current question index, navigation)

### 3.4 Core Features - User Flow
- [ ] Exam List Page (/exams)
  - [ ] Display available exams
  - [ ] Show exam details (title, duration, question count)
  - [ ] Start exam button
- [ ] Exam Runner Page (/exams/:id/take)
  - [ ] Question navigation panel
  - [ ] Question display with HTML content rendering
  - [ ] Answer selection/input
  - [ ] Timer countdown
  - [ ] Save answer functionality
  - [ ] Submit exam confirmation
- [ ] Exam Results Page (/exams/:sessionId/results)
  - [ ] Display score and statistics
  - [ ] Show correct/incorrect answers
  - [ ] Review all questions with explanations

### 3.5 API Integration
- [ ] Create API client with fetch/axios
- [ ] Implement exam fetching
- [ ] Implement session management
- [ ] Implement answer submission
- [ ] Handle loading and error states

### 3.6 UI/UX Enhancements
- [ ] Responsive design for mobile/tablet/desktop
- [ ] Loading states and skeletons
- [ ] Toast notifications
- [ ] Confirmation dialogs
- [ ] Keyboard shortcuts for navigation

---

## Phase 4: DevOps & CI/CD
### 4.1 GitHub Actions
- [ ] Create .github/workflows/deploy.yml
- [ ] Configure Docker Hub secrets (DOCKER_USERNAME, DOCKER_PASSWORD)
- [ ] Set up multi-stage Docker builds
- [ ] Configure build triggers (push to main)

### 4.2 Docker Hub Deployment
- [ ] Build backend image
- [ ] Build frontend image
- [ ] Push images with version tags
- [ ] Push images with latest tag
- [ ] Add image size optimization

### 4.3 Deployment Documentation
- [ ] Document environment variables
- [ ] Create deployment guide
- [ ] Add health check endpoints
- [ ] Create backup/restore procedures

---

## Phase 5: Testing & Quality Assurance
### 5.1 Backend Testing
- [ ] Unit tests for use cases
- [ ] Integration tests for repositories
- [ ] API endpoint tests
- [ ] Database migration tests

### 5.2 Frontend Testing
- [ ] Component unit tests
- [ ] Integration tests for pages
- [ ] E2E tests for exam flow

### 5.3 Performance & Security
- [ ] API rate limiting
- [ ] Input validation and sanitization
- [ ] SQL injection prevention (GORM parameterization)
- [ ] XSS prevention for HTML content
- [ ] CSRF protection

---

## Technical Requirements & Constraints
1. **Content Storage**: Exam content stored as HTML Text, images via S3 URL
2. **Focus**: 100% on User Flow, no Admin UI for now
3. **Flexibility**: Use JSONB for question configuration (multiple choice, essay, etc.)
4. **Clean Architecture**: Separate concerns in backend (Domain, Repository, UseCase, Delivery)
5. **Atomic Design**: Component hierarchy in frontend (Atoms → Molecules → Organisms → Templates → Pages)

---

## Current Status
**Last Updated**: 2026-02-01
**Phase**: Phase 1 - Infrastructure Setup
**Next Task**: Create project directory structure
