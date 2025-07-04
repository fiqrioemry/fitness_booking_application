import { publicInstance, authInstance } from ".";

// GET /api/locations
export const getAllLocations = async () => {
  const res = await publicInstance.get("/locations");
  return res.data;
};

// GET /api/locations/:id
export const getLocationById = async (id) => {
  const res = await publicInstance.get(`/locations/${id}`);
  return res.data;
};

// POST /api/locations
export const createLocation = async (data) => {
  const res = await authInstance.post("/locations", data);
  return res.data;
};

// PUT /api/locations/:id
export const updateLocation = async (id, data) => {
  const res = await authInstance.put(`/locations/${id}`, data);
  return res.data;
};

// DEL /api/locations/:id
export const deleteLocation = async (id) => {
  const res = await authInstance.delete(`/locations/${id}`);
  return res.data;
};
