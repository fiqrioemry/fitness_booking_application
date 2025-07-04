import { toast } from "sonner";
import * as subcategoryService from "@/services/subcategories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useSubcategoriesQuery = () =>
  useQuery({
    queryKey: ["subcategories"],
    queryFn: subcategoryService.getAllSubcategories,
    keepPreviousData: true,
  });

export const useSubcategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["subcategory", id],
    queryFn: () => subcategoryService.getSubcategoryById(id),
    enabled: !!id,
  });

export const useSubcategoriesByCategoryQuery = (categoryId) =>
  useQuery({
    queryKey: ["subcategories", "category", categoryId],
    queryFn: () => subcategoryService.getSubcategoriesByCategory(categoryId),
    enabled: !!categoryId,
  });

export const useSubcategoryMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["subcategories"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: subcategoryService.createSubcategory,
      ...mutationOpts("Subcategory created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) =>
        subcategoryService.updateSubcategory(id, data),
      ...mutationOpts("Subcategory updated successfully"),
    }),

    deleteOptions: useMutation({
      mutationFn: subcategoryService.deleteSubcategory,
      ...mutationOpts("Subcategory deleted successfully"),
    }),
  };
};
