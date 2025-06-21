import qs from "qs";
import { authInstance } from ".";

// GET /api/user-packages
export const getUserPackages = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/user-packages?${queryString}`);
  return res.data;
};

// GET /api/user-packages/class
export const getUserPackagesByClassID = async (id) => {
  const res = await authInstance.get(`/user-packages/class/${id}`);
  return res.data;
};
