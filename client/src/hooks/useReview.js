import { toast } from "sonner";
import * as reviewService from "@/services/review";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useClassReviewsQuery = (classId) =>
  useQuery({
    queryKey: ["reviews", classId],
    queryFn: () => reviewService.getReviewsByClass(classId),
    enabled: !!classId,
  });

export const useCreateReviewMutation = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: reviewService.createReview,
    onSuccess: (res, variables) => {
      toast.success(res.message || "Review Submitted Successfully");
      qc.invalidateQueries({ queryKey: ["booking", variables.id] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
