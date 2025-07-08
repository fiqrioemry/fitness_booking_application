import qs from "qs";
import { authInstance, publicInstance } from ".";

// GET /api/schedules
export const getAllClassSchedules = async () => {
  const res = await publicInstance.get("/schedules");
  return res.data;
};

// GET /api/schedules/status
export const getAllClassSchedulesWithStatus = async () => {
  const res = await authInstance.get("/schedules/status");
  return res.data;
};

// GET /api/schedules/:id
export const getClassScheduleDetail = async (id) => {
  const res = await authInstance.get(`/schedules/${id}`);
  return res.data;
};

export const createClassSchedule = async (data) => {
  const url = data.isRecurring
    ? "/admin/schedules/recurring"
    : "/admin/schedules";
  const res = await authInstance.post(url, data);
  return res.data;
};

// PUT /api/schedules/:id
export const updateClassSchedule = async (id, data) => {
  const res = await authInstance.put(`/admin/schedules/${id}`, data);
  return res.data;
};

// DELETE /api/schedules/:id
export const deleteClassSchedule = async (id) => {
  const res = await authInstance.delete(`/admin/schedules/${id}`);
  return res.data;
};

// POST /api/schedule-templates
export const createTemplate = async (data) => {
  const res = await authInstance.post("/admin/schedule-templates", data);
  return res.data;
};

// GET /api/schedule-templates
export const getAllRecuringSchedule = async () => {
  const res = await authInstance.get(`/admin/schedule-templates`);
  return res.data;
};

// PUT /api/schedule-templates/:id
export const updateScheduleTemplate = async (id, data) => {
  const res = await authInstance.put(`/admin/schedule-templates/${id}`, data);
  return res.data;
};

// DELETE /api/schedule-templates/:id
export const deleteScheduleTemplate = async (id) => {
  const res = await authInstance.delete(`/admin/schedule-templates/${id}`);
  return res.data;
};

// PUT /api/schedule-templates/:id/run
export const runScheduleTemplate = async (id) => {
  const res = await authInstance.post(`/admin/schedule-templates/${id}/run`);
  return res.data;
};
// PUT /api/schedule-templates/:id/stop
export const stopScheduleTemplate = async (id) => {
  const res = await authInstance.post(`/admin/schedule-templates/${id}/stop`);
  return res.data;
};

// GET /api/instructor/schedules
export const getInstructorSchedules = async (params) => {
  const queryString = qs.stringify(params, { skipNulls: true });
  const res = await authInstance.get(`/instructor/schedules?${queryString}`);
  return res.data;
};

// PATCH /api/instructor/schedules/:id/open
export const openClassSchedule = async ({ id, data }) => {
  const res = await authInstance.patch(
    `/instructor/schedules/${id}/open`,
    data
  );
  return res.data;
};

// GET /api/instructor/schedules/:id/attendance
export const getClassAttendances = async (id) => {
  const res = await authInstance.get(`/instructor/schedules/${id}/attendance`);
  return res.data;
};
