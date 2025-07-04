import qs from "qs";
import { authInstance } from ".";

//   GET /api/admin/users
export const getAllUsers = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/admin/users?${queryString}`);
  return res.data;
};

//  GET /api/admin/users/:id
export const getUserDetail = async (id) => {
  const res = await authInstance.get(`/admin/users/${id}`);
  return res.data;
};

//  GET /api/admin/users/stats
export const getUserStats = async () => {
  const res = await authInstance.get("/admin/users/stats");
  return res.data;
};
