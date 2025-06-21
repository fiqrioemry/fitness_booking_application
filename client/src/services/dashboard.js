import { authInstance } from ".";

// GET /admin/dashboard/summary
export const getDashboardSummary = async () => {
  const res = await authInstance.get("/admin/dashboard/summary");
  return res.data;
};

// GET /admin/dashboard/revenue?range=daily|monthly|yearly
export const getRevenueStats = async (range = "daily") => {
  const res = await authInstance.get("/admin/dashboard/revenue", {
    params: { range },
  });
  return res.data;
};
