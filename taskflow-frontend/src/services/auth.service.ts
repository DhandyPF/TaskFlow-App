import { api } from "@/lib/api";
import { LoginPayload, RegisterPayload, User } from "@/types/auth";

export const authService = {
  register: (payload: RegisterPayload) =>
    api.post<User>("/api/v1/auth/register", payload),

  login: async (payload: LoginPayload) => {
    const res = await api.post<never>("/api/v1/auth/login", payload);
    if (res.token) {
      localStorage.setItem("taskflow_token", res.token);
    }
    return res;
  },

  logout: () => {
    localStorage.removeItem("taskflow_token");
  },

  isAuthenticated: () => {
    if (typeof window === "undefined") return false;
    return Boolean(localStorage.getItem("taskflow_token"));
  },
};
