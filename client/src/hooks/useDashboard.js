import { useQuery } from "@tanstack/react-query";
import * as dashboardService from "@/services/dashboard";

export const useDashboardSummary = () =>
  useQuery({
    queryKey: ["dashboard", "summary"],
    queryFn: dashboardService.getDashboardSummary,
  });

export const useRevenueStats = (range = "daily") =>
  useQuery({
    queryKey: ["dashboard", "revenue", range],
    queryFn: () => dashboardService.getRevenueStats(range),
  });
