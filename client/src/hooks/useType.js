import { toast } from "sonner";
import * as typeService from "@/services/type";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useTypesQuery = () =>
  useQuery({
    queryKey: ["types"],
    queryFn: typeService.getAllTypes,
    keepPreviousData: true,
  });

export const useTypeDetailQuery = (id) =>
  useQuery({
    queryKey: ["type", id],
    queryFn: () => typeService.getTypeById(id),
    enabled: !!id,
  });

export const useTypeMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["types"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: typeService.createType,
      ...mutationOpts("Type created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => typeService.updateType(id, data),
      ...mutationOpts("Type updated successfully"),
    }),

    deleteOptions: useMutation({
      mutationFn: typeService.deleteType,
      ...mutationOpts("Type deleted successfully"),
    }),
  };
};
