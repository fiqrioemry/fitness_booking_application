import { toast } from "sonner";
import * as category from "@/services/categories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCategoriesQuery = () =>
  useQuery({
    queryKey: ["categories"],
    queryFn: category.getAllCategories,
    keepPreviousData: true,
  });

export const useCategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["category", id],
    queryFn: () => category.getCategoryById(id),
    enabled: !!id,
  });

export const useCategoryMutation = () => {
  const qc = useQueryClient();

  const mutationOptions = (message) => ({
    onSuccess: (res) => {
      toast.success(res?.message || message);
      qc.invalidateQueries({ queryKey: ["categories"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: category.createCategory,
      ...mutationOptions("Category created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => category.updateCategory(id, data),
      ...mutationOptions("Category updated successfully"),
    }),

    deleteOptions: useMutation({
      mutationFn: category.deleteCategory,
      ...mutationOptions("Category deleted successfully"),
    }),
  };
};
