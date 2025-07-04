import { toast } from "sonner";
import * as paymentService from "@/services/payment";
import { useMutation, useQuery } from "@tanstack/react-query";

export const useAllPaymentsQuery = (params) =>
  useQuery({
    queryKey: ["payments", params],
    queryFn: () => paymentService.getAllUserPayments(params),
    refetchOnMount: true,
  });

export const useMyPaymentsQuery = (params) =>
  useQuery({
    queryKey: ["my-payments", params],
    queryFn: () => paymentService.getMyPayments(params),
    refetchOnMount: true,
  });

export const useCreatePaymentMutation = () =>
  useMutation({
    mutationFn: paymentService.createPayment,
    onSuccess: (res) => {
      toast.success(res?.message || "Payment created successfully");
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to create payment");
    },
  });

export const usePaymentDetailQuery = (id) =>
  useQuery({
    queryKey: ["payment-detail", id],
    queryFn: () => paymentService.getPaymentDetail(id),
    enabled: !!id,
  });

export const useMyPaymentDetailQuery = (id) =>
  useQuery({
    queryKey: ["my-payment-detail", id],
    queryFn: () => paymentService.getMyPaymentDetail(id),
    enabled: !!id,
  });
