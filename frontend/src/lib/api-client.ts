import { Exam, ExamSession, CreateSessionRequest, SubmitAnswerRequest, SubmitExamResponse, Answer } from '@/types/exam';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(
    endpoint: string,
    options?: RequestInit
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    
    const response = await fetch(url, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new Error(error.message || 'API request failed');
    }

    return response.json();
  }

  // Exam APIs
  async getExams(): Promise<Exam[]> {
    return this.request<Exam[]>('/api/exams');
  }

  async getExam(id: number): Promise<Exam> {
    return this.request<Exam>(`/api/exams/${id}`);
  }

  // Session APIs
  async createSession(data: CreateSessionRequest): Promise<ExamSession> {
    return this.request<ExamSession>('/api/exam-sessions', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async getSession(id: number): Promise<ExamSession> {
    return this.request<ExamSession>(`/api/exam-sessions/${id}`);
  }

  async submitAnswer(sessionId: number, data: SubmitAnswerRequest): Promise<Answer> {
    return this.request<Answer>(`/api/exam-sessions/${sessionId}/answers`, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async submitExam(sessionId: number): Promise<SubmitExamResponse> {
    return this.request<SubmitExamResponse>(`/api/exam-sessions/${sessionId}/submit`, {
      method: 'POST',
    });
  }

  async getResults(sessionId: number): Promise<SubmitExamResponse> {
    return this.request<SubmitExamResponse>(`/api/exam-sessions/${sessionId}/results`);
  }

  // Health check
  async healthCheck(): Promise<{ status: string; message: string }> {
    return this.request('/health');
  }
}

export const apiClient = new ApiClient(API_URL);
