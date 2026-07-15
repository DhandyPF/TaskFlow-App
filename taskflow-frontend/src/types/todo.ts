export type Priority = "LOW" | "MEDIUM" | "HIGH";

export interface Todo {
  id: string;
  title: string;
  description: string;
  is_completed: boolean;
  priority: Priority;
  due_date: string | null;
  created_at: string;
  updated_at: string;
}

export interface CreateTodoPayload {
  title: string;
  description: string;
  priority: Priority;
  due_date: string | null;
}

export interface UpdateTodoPayload {
  title?: string;
  description?: string;
  is_completed?: boolean;
  priority?: Priority;
  due_date?: string | null;
}
