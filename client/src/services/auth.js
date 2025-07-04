import { publicInstance, authInstance } from ".";

// POST /api/auth/send-otp
export const sendOTP = async (data) => {
  const res = await publicInstance.post("/auth/send-otp", data);
  return res.data;
};

// POST /api/auth/verify-otp
export const verifyOTP = async (data) => {
  const res = await publicInstance.post("/auth/verify-otp", data);
  return res.data;
};

// POST /api/auth/register
export const register = async (data) => {
  const res = await publicInstance.post("/auth/register", data);
  return res.data;
};

// POST /api/auth/login
export const login = async (data) => {
  const res = await publicInstance.post("/auth/login", data);
  return res.data;
};

// POST /api/auth/logout
export const logout = async () => {
  const res = await authInstance.post("/auth/logout");
  return res.data;
};

// POST /api/auth/refresh-token
export const refreshToken = async () => {
  const res = await authInstance.post("/auth/refresh-token");
  return res.data;
};

// GET /api/auth/me
export const getMe = async () => {
  const res = await authInstance.get("/auth/me");
  return res.data;
};
