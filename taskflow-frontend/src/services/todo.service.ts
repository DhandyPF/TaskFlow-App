import { api } from "@/lib/api";
import { CreateTodoPayload, Todo, UpdateTodoPayload } from "@/types/todo";

export const todoService = {
  list: (params?: { is_completed?: boolean; sort_by?: string }) => {
    const query = new URLSearchParams();
    if (params?.is_completed !== undefined) {
      query.set("is_completed", String(params.is_completed));
    }
    if (params?.sort_by) {
      query.set("sort_by", params.sort_by);
    }
    const qs = query.toString();
    return api.get<Todo[]>(`/api/v1/todos${qs ? `?${qs}` : ""}`);
  },

  create: (payload: CreateTodoPayload) => api.post<Todo>("/api/v1/todos", payload),

  update: (id: string, payload: UpdateTodoPayload) =>
    api.put<Todo>(`/api/v1/todos/${id}`, payload),

  remove: (id: string) => api.delete<never>(`/api/v1/todos/${id}`),
};
