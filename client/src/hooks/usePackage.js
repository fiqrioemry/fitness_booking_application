import { toast } from "sonner";
import * as packageService from "@/services/package";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const usePackagesQuery = (params) =>
  useQuery({
    queryKey: ["packages", params],
    queryFn: () => packageService.getAllPackages(params),
    refetchOnMount: true,
    staleTime: 1000 * 60 * 15,
  });

export const usePackageDetailQuery = (id) =>
  useQuery({
    queryKey: ["package", id],
    queryFn: () => packageService.getPackageById(id),
    enabled: !!id,
  });

export const usePackageMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["packages"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createPackage: useMutation({
      mutationFn: packageService.createPackage,
      ...mutationOpts("Package created successfully"),
    }),

    updatePackage: useMutation({
      mutationFn: ({ id, data }) => packageService.updatePackage(id, data),
      ...mutationOpts("Package updated successfully"),
    }),

    deletePackage: useMutation({
      mutationFn: packageService.deletePackage,
      ...mutationOpts("Package deleted successfully"),
    }),
  };
};
