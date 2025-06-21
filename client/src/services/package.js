import qs from "qs";
import { buildFormData } from "@/lib/utils";
import { publicInstance, authInstance } from ".";

export const getAllPackages = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await publicInstance.get(`/packages?${queryString}`);
  return res.data;
};

export const getPackageById = async (id) => {
  const res = await publicInstance.get(`/packages/${id}`);
  return res.data;
};

export const createPackage = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/packages", formData);
  return res.data;
};

export const updatePackage = async (id, data) => {
  const formData = buildFormData(data);
  const res = await authInstance.put(`/packages/${id}`, formData);
  return res.data;
};

export const deletePackage = async (id) => {
  const res = await authInstance.delete(`/packages/${id}`);
  return res.data;
};
