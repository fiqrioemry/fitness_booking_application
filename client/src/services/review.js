import { publicInstance, authInstance } from ".";

// POST /api/reviews
export const createReview = async ({ id, data }) => {
  const res = await authInstance.post(`/reviews/${id}`, data);
  return res.data;
};

// GET /api/reviews/:classId
export const getReviewsByClass = async (classId) => {
  const res = await publicInstance.get(`/reviews/${classId}`);
  return res.data;
};
