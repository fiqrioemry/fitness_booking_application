// src/services/index.js
import axios from "axios";
import * as auth from "./auth";

export const publicInstance = axios.create({
  baseURL: import.meta.env.VITE_API_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

export const authInstance = axios.create({
  baseURL: import.meta.env.VITE_API_SERVICES,
  withCredentials: true,
  headers: {
    "X-API-Key": import.meta.env.VITE_API_KEY,
  },
});

authInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response?.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url.includes("/auth/refresh-token")
    ) {
      originalRequest._retry = true;
      try {
        await auth.refreshToken();
        return authInstance(originalRequest);
      } catch (refreshError) {
        useAuthStore.getState().clearUser();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);
