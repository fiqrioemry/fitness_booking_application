import * as userService from "@/services/users";
import { useQuery } from "@tanstack/react-query";

export const useUsersQuery = (params) =>
  useQuery({
    queryKey: ["users", params],
    queryFn: () => userService.getAllUsers(params),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 5,
  });

export const useUserDetailQuery = (userId) =>
  useQuery({
    queryKey: ["user-detail", userId],
    queryFn: () => userService.getUserDetail(userId),
    enabled: !!userId,
  });

export const useUserStatsQuery = () =>
  useQuery({
    queryKey: ["users-stats"],
    queryFn: userService.getUserStats,
  });
