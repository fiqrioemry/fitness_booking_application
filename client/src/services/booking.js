import qs from "qs";
import { authInstance } from ".";

// POST /api/bookings
export const createBooking = async (data) => {
  const res = await authInstance.post("/bookings", data);
  return res.data;
};

// GET /api/bookings
export const getMyBookings = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/bookings?${queryString}`);
  return res.data;
};

// GET /api/bookings/:id
export const getBookingDetail = async (id) => {
  const res = await authInstance.get(`/bookings/${id}`);
  return res.data;
};

// POST /api/bookings/:id/check-in
export const checkinBooking = async ({ id }) => {
  const res = await authInstance.post(`/bookings/${id}/check-in`);
  return res.data;
};

// POST /api/bookings/:id/check-out
export const checkoutBooking = async ({ id, data }) => {
  const res = await authInstance.post(`/bookings/${id}/check-out`, data);
  return res.data;
};
