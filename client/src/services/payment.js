import qs from "qs";
import { authInstance } from ".";

// GET /api/payments/me?q=&page=&limit=&status=&sort=
export const getMyPayments = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/payments/me?${queryString}`);
  return res.data;
};

// POST /api/payments
export const createPayment = async (data) => {
  const res = await authInstance.post("/payments", data);
  return res.data;
};

// GET /api/admin/payments/:id
export const getPaymentDetail = async (id) => {
  const res = await authInstance.get(`/admin/payments/${id}`);
  console.log("Payment Detail Response:", res.data);
  return res.data;
};

// GET /api/admin/payments/me/:id
export const getMyPaymentDetail = async (id) => {
  const res = await authInstance.get(`/payments/me/${id}`);
  return res.data;
};

// GET /api/admin/payments?q=&page=&limit=&status=&sort=
export const getAllUserPayments = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/admin/payments?${queryString}`);
  return res.data;
};
