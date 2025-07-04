import { toast } from "sonner";
import * as classService from "@/services/class";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useClassesQuery = (params) =>
  useQuery({
    queryKey: ["classes", params],
    queryFn: () => classService.getAllClasses(params),
    refetchOnMount: true,
  });

export const useClassDetailQuery = (id) =>
  useQuery({
    queryKey: ["class", id],
    queryFn: () => classService.getClassById(id),
    enabled: !!id,
  });

export const useClassMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (loadingMsg, successMsg) => ({
    onMutate: async () => {
      toast.loading(loadingMsg, { id: "class-action" });
    },
    onSuccess: (res) => {
      toast.success(res?.message || successMsg, {
        id: "class-action",
      });
      qc.invalidateQueries({ queryKey: ["classes"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong", {
        id: "class-action",
      });
    },
  });
  return {
    createClass: useMutation({
      mutationFn: classService.createClass,
      ...mutationOpts("Creating class...", "Class created successfully"),
    }),

    updateClass: useMutation({
      mutationFn: ({ id, data }) => classService.updateClass(id, data),
      ...mutationOpts("Updating class...", "Class updated successfully"),
    }),

    deleteClass: useMutation({
      mutationFn: classService.deleteClass,
      ...mutationOpts("Deleting class...", "Class deleted successfully"),
    }),

    uploadGallery: useMutation({
      mutationFn: ({ id, images }) =>
        classService.uploadClassGallery(id, images),
      ...mutationOpts("Uploading gallery...", "Gallery uploaded successfully"),
    }),
  };
};
