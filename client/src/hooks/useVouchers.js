import { toast } from "sonner";
import * as voucherService from "@/services/voucher";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useVouchersQuery = () =>
  useQuery({
    queryKey: ["vouchers"],
    queryFn: voucherService.getAllVouchers,
    staleTime: 1000 * 60 * 10,
  });

export const useVoucherMutation = () => {
  const qc = useQueryClient();

  const mutationOpts = (msg) => ({
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      qc.invalidateQueries({ queryKey: ["vouchers"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createVoucher: useMutation({
      mutationFn: voucherService.createVoucher,
      ...mutationOpts("Voucher created successfully"),
    }),

    updateVoucher: useMutation({
      mutationFn: ({ id, data }) => voucherService.updateVoucher({ id, data }),
      ...mutationOpts("Voucher updated successfully"),
    }),

    deleteVoucher: useMutation({
      mutationFn: voucherService.deleteVoucher,
      ...mutationOpts("Voucher deleted successfully"),
    }),

    applyVoucher: useMutation({
      mutationFn: voucherService.applyVoucher,
      ...mutationOpts("Voucher applied successfully"),
    }),
  };
};
