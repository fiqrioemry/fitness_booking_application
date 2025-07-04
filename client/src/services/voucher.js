import { authInstance } from ".";

// GET /api/vouchers
export const getAllVouchers = async () => {
  const res = await authInstance.get("/vouchers");
  return res.data;
};

// POST /api/vouchers
export const createVoucher = async (data) => {
  const res = await authInstance.post("/vouchers", data);
  return res.data;
};

// PoST /api/vouchers/apply
export const applyVoucher = async ({ userId, code, total }) => {
  const res = await authInstance.post("/vouchers/apply", {
    userId,
    code,
    total,
  });
  return res.data;
};

//  PUT /api/vouchers/:id
export const updateVoucher = async ({ id, data }) => {
  const res = await authInstance.put(`/vouchers/${id}`, data);
  return res.data;
};

//  DELETE /api/vouchers/:id
export const deleteVoucher = async (id) => {
  const res = await authInstance.delete(`/vouchers/${id}`);
  return res.data;
};
