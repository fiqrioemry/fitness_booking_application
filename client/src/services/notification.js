import { authInstance } from ".";

export const getNotificationSettings = async () => {
  const res = await authInstance.get("/notifications/settings");
  return res.data;
};

export const updateNotificationSetting = async (payload) => {
  const res = await authInstance.put("/notifications/settings", payload);
  return res.data;
};

export const markAllNotificationsAsRead = async () => {
  const res = await authInstance.patch(`/notifications/read`);
  return res.data;
};

export const getAllBrowserNotifications = async () => {
  const res = await authInstance.get("/notifications");
  return res.data;
};

export const sendPromoNotification = async (payload) => {
  const res = await authInstance.post("/notifications/broadcast", payload);
  return res.data;
};
