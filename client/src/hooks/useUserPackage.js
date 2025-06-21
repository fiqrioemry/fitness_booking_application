import { useQuery } from "@tanstack/react-query";
import * as userPackages from "@/services/userpackages";

export const useUserPackagesQuery = (params) =>
  useQuery({
    queryKey: ["user-packages", params],
    queryFn: () => userPackages.getUserPackages(params),
    staleTime: 0,
  });

export const useUserClassPackagesQuery = (id) =>
  useQuery({
    queryKey: ["user-packages", id],
    queryFn: () => userPackages.getUserPackagesByClassID(id),
    enabled: !!id,
  });
