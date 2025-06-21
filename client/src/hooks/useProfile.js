import { toast } from "sonner";
import * as profileService from "@/services/profile";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useProfileQuery = () =>
  useQuery({
    queryKey: ["user", "profile"],
    queryFn: profileService.getProfile,
    staleTime: 0,
  });

export const useProfileMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res.message || msg);
      qc.invalidateQueries({ queryKey: ["user", "profile"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    updateProfile: useMutation({
      mutationFn: profileService.updateProfile,
      ...mutationOpts("Profile updated successfully"),
    }),

    updateAvatar: useMutation({
      mutationFn: profileService.updateAvatar,
      ...mutationOpts("Avatar updated successfully"),
    }),
  };
};
