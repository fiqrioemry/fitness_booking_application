import { authInstance } from ".";
import { buildFormData } from "@/lib/utils";

// GET /api/user/profile
export const getProfile = async () => {
  const res = await authInstance.get("/users/me");
  return res.data;
};

// PUT /api/user/profile
export const updateProfile = async (data) => {
  const res = await authInstance.put("/users/me", data);
  return res.data;
};

// PUT /api/user/profile/avatar
export const updateAvatar = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put("/users/me/avatar", formData);
  return res.data;
};
