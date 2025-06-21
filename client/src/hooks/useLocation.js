import { toast } from "sonner";
import * as locationService from "@/services/location";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useLocationsQuery = () =>
  useQuery({
    queryKey: ["locations"],
    queryFn: locationService.getAllLocations,
    keepPreviousData: true,
  });

export const useLocationDetailQuery = (id) =>
  useQuery({
    queryKey: ["location", id],
    queryFn: () => locationService.getLocationById(id),
    enabled: !!id,
  });

export const useLocationMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["locations"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: locationService.createLocation,
      ...mutationOpts("Location created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => locationService.updateLocation(id, data),
      ...mutationOpts("Location updated successfully"),
    }),

    deleteOptions: useMutation({
      mutationFn: locationService.deleteLocation,
      ...mutationOpts("Location deleted successfully"),
    }),
  };
};
