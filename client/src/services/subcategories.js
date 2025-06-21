import { publicInstance, authInstance } from ".";

// GET /api/subcategories
export const getAllSubcategories = async () => {
  const res = await publicInstance.get("/subcategories");
  return res.data;
};

// GET /api/subcategories/:id
export const getSubcategoryById = async (id) => {
  const res = await publicInstance.get(`/subcategories/${id}`);
  return res.data;
};

// GET /api/subcategories/category/:categoryId
export const getSubcategoriesByCategory = async (categoryId) => {
  const res = await publicInstance.get(`/subcategories/category/${categoryId}`);
  return res.data;
};

// POST /api/subcategories (Admin Only)
export const createSubcategory = async (data) => {
  const res = await authInstance.post("/subcategories", data);
  return res.data;
};

// PUT /api/subcategories/:id (Admin Only)
export const updateSubcategory = async (id, data) => {
  const res = await authInstance.put(`/subcategories/${id}`, data);
  return res.data;
};

// DELETE /api/subcategories/:id (Admin Only)
export const deleteSubcategory = async (id) => {
  const res = await authInstance.delete(`/subcategories/${id}`);
  return res.data;
};
