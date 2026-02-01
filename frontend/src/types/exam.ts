// Domain types matching backend models

export type QuestionType = 'multiple_choice' | 'essay' | 'true_false';

export interface Exam {
  id: number;
  title: string;
  description: string;
  duration: number; // minutes
  content_html: string;
  is_active: boolean;
  questions?: Question[];
  created_at: string;
  updated_at: string;
}

export interface Question {
  id: number;
  exam_id: number;
  content_html: string;
  type: QuestionType;
  options: any; // JSONB - flexible structure
  answer: any; // JSONB - correct answer(s)
  explanation: string;
  points: number;
  order: number;
  created_at: string;
  updated_at: string;
}

export type SessionStatus = 'in_progress' | 'completed' | 'abandoned';

export interface ExamSession {
  id: number;
  user_id: string;
  exam_id: number;
  exam?: Exam;
  start_time: string;
  end_time?: string;
  status: SessionStatus;
  score?: number;
  total_points: number;
  answers?: Answer[];
  created_at: string;
  updated_at: string;
}

export interface Answer {
  id: number;
  session_id: number;
  question_id: number;
  question?: Question;
  user_answer: any; // JSONB - user's answer
  is_correct?: boolean;
  points?: number;
  created_at: string;
  updated_at: string;
}

// Request/Response types
export interface CreateSessionRequest {
  exam_id: number;
  user_id: string;
}

export interface SubmitAnswerRequest {
  question_id: number;
  user_answer: any;
}

export interface SubmitExamResponse {
  session: ExamSession;
  score: number;
  total_points: number;
  correct_answers: number;
  total_questions: number;
}
