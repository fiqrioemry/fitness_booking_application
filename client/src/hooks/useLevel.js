import { toast } from "sonner";
import * as levelService from "@/services/level";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useLevelsQuery = () =>
  useQuery({
    queryKey: ["levels"],
    queryFn: levelService.getAllLevels,
    keepPreviousData: true,
  });

export const useLevelDetailQuery = (id) =>
  useQuery({
    queryKey: ["level", id],
    queryFn: () => levelService.getLevelById(id),
    enabled: !!id,
  });

export const useLevelMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["levels"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: levelService.createLevel,
      ...mutationOpts("Level created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => levelService.updateLevel(id, data),
      ...mutationOpts("Level updated successfully"),
    }),

    deleteOptions: useMutation({
      mutationFn: levelService.deleteLevel,
      ...mutationOpts("Level deleted successfully"),
    }),
  };
};
