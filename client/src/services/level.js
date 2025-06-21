import { publicInstance, authInstance } from ".";

// GET /api/levels
export const getAllLevels = async () => {
  const res = await publicInstance.get("/levels");
  return res.data;
};

// GET /api/levels/:id
export const getLevelById = async (id) => {
  const res = await publicInstance.get(`/levels/${id}`);
  return res.data;
};

// POST /api/levels (Admin Only)
export const createLevel = async (data) => {
  const res = await authInstance.post("/levels", data);
  return res.data;
};

// PUT /api/levels/:id (Admin Only)
export const updateLevel = async (id, data) => {
  const res = await authInstance.put(`/levels/${id}`, data);
  return res.data;
};

// DELETE /api/levels/:id (Admin Only)
export const deleteLevel = async (id) => {
  const res = await authInstance.delete(`/levels/${id}`);
  return res.data;
};
