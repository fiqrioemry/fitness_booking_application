import { publicInstance, authInstance } from ".";

// GET /api/instructors
export const getAllInstructors = async () => {
  const res = await publicInstance.get("/instructors");
  return res.data;
};

// GET /api/instructors/:id
export const getInstructorById = async (id) => {
  const res = await publicInstance.get(`/instructors/${id}`);
  return res.data;
};

// POST /api/instructors (Admin Only)
export const createInstructor = async (data) => {
  const res = await authInstance.post("/instructors", data);
  return res.data;
};

// PUT /api/instructors/:id (Admin Only)
export const updateInstructor = async (id, data) => {
  const res = await authInstance.put(`/instructors/${id}`, data);
  return res.data;
};

// DELETE /api/instructors/:id (Admin Only)
export const deleteInstructor = async (id) => {
  const res = await authInstance.delete(`/instructors/${id}`);
  return res.data;
};
