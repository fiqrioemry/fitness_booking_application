import qs from "qs";
import { authInstance } from ".";

// GET /api/payments?q=&page=&limit=&status=&sort=
export const getAllUserPayments = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/payments?${queryString}`);
  return res.data;
};

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

// GET /api/payments/:id
export const getPaymentDetail = async (id) => {
  const res = await authInstance.get(`/payments/${id}`);
  return res.data;
};

// GET /api/payments/me/:id
export const getMyPaymentDetail = async (id) => {
  const res = await authInstance.get(`/payments/me/${id}`);
  return res.data;
};
