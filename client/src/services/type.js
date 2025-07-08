import { publicInstance, authInstance } from ".";

// GET /api/types
export const getAllTypes = async () => {
  const res = await publicInstance.get("/types");
  return res.data;
};

// GET /api/types/:id
export const getTypeById = async (id) => {
  const res = await publicInstance.get(`/types/${id}`);
  return res.data;
};

// POST /api/admin/types (Admin Only)
export const createType = async (data) => {
  const res = await authInstance.post("/admin/types", data);
  return res.data;
};

// PUT /api/admin/types/:id (Admin Only)
export const updateType = async (id, data) => {
  const res = await authInstance.put(`/admin/types/${id}`, data);
  return res.data;
};

// DELETE /api/admin/types/:id (Admin Only)
export const deleteType = async (id) => {
  const res = await authInstance.delete(`/admin/types/${id}`);
  return res.data;
};
