import { toast } from "sonner";
import * as booking from "@/services/booking";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCreateBookingMutation = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: booking.createBooking,
    onSuccess: (res) => {
      toast.success(res.message || "Booking created successfully");
      qc.invalidateQueries({ queryKey: ["bookings"] });
      qc.invalidateQueries({ queryKey: ["schedules", "with-status"] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to create booking");
    },
  });
};

export const useMyBookingsQuery = (params) =>
  useQuery({
    queryKey: ["bookings", params],
    queryFn: () => booking.getMyBookings(params),
    refetchOnMount: true,
    staleTime: 1000 * 60 * 60,
  });

export const useBookingDetailQuery = (id) =>
  useQuery({
    queryKey: ["booking", id],
    queryFn: () => booking.getBookingDetail(id),
    enabled: !!id,
  });

export const useCheckinBookingMutation = () => {
  const qc = useQueryClient();

  return useMutation({
    mutationFn: booking.checkinBooking,
    onSuccess: (_, v) => {
      toast.success("Check-in successful");
      qc.invalidateQueries({ queryKey: ["booking", v.id] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to check-in");
    },
  });
};

export const useCheckoutBookingMutation = () => {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: booking.checkoutBooking,
    onSuccess: (_, v) => {
      toast.success("Check-out successful");
      qc.invalidateQueries({ queryKey: ["booking", v.id] });
    },
    onError: (error) => {
      toast.error(error?.response?.data?.message || "Failed to check-out");
    },
  });
};
