import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/Dialog";
import { formatRupiah } from "@/lib/utils";
import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { Skeleton } from "@/components/ui/Skeleton";
import { useParams, useNavigate } from "react-router-dom";
import { useMyPaymentDetailQuery } from "@/hooks/usePayment";

export const MyTransactionDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { data, isLoading } = useMyPaymentDetailQuery(id);

  return (
    <Dialog open={true} onOpenChange={() => navigate(-1)}>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>Transaction Details</DialogTitle>
          <DialogDescription>
            Invoice and payment details for the selected transaction.
          </DialogDescription>
        </DialogHeader>

        {isLoading ? (
          <div className="space-y-3">
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
            <Skeleton className="w-full h-6" />
          </div>
        ) : (
          <div className="space-y-5 text-sm">
            <div>
              <p className="text-muted-foreground text-xs mb-1">
                Invoice Number
              </p>
              <h3 className="font-medium text-base">{data.invoiceNumber}</h3>
            </div>

            <div className="grid grid-cols-2 gap-2">
              <p>
                <span className="text-muted-foreground">User:</span>{" "}
                <span className="font-medium">{data.fullname}</span>
              </p>
              <p>
                <span className="text-muted-foreground">Email:</span>{" "}
                <span>{data.email}</span>
              </p>
              <p>
                <span className="text-muted-foreground">Package:</span>{" "}
                <span>{data.packageName}</span>
              </p>
              <p>
                <span className="text-muted-foreground">Payment Method:</span>{" "}
                <span>{data.paymentMethod}</span>
              </p>
              <p>
                <span className="text-muted-foreground">Status:</span>{" "}
                <Badge
                  variant={
                    data.status === "success"
                      ? "default"
                      : data.status === "failed"
                      ? "destructive"
                      : "secondary"
                  }
                >
                  {data.status}
                </Badge>
              </p>
              <p>
                <span className="text-muted-foreground">Paid At:</span>{" "}
                <span>{data.paidAt || "-"}</span>
              </p>
            </div>

            <hr className="my-2 border-border" />

            <div className="grid gap-2 text-right">
              <div className="flex justify-between">
                <span className="text-muted-foreground">Base Price</span>
                <span>{formatRupiah(data.basePrice)}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-muted-foreground">Tax</span>
                <span>{formatRupiah(data.tax)}</span>
              </div>
              {data.voucherDiscount > 0 && (
                <div className="flex justify-between">
                  <span className="text-muted-foreground">
                    Discount {data.voucherCode && `(${data.voucherCode})`}
                  </span>
                  <span>- {formatRupiah(data.voucherDiscount)}</span>
                </div>
              )}
              <hr className="border-muted my-1" />
              <div className="flex justify-between font-bold text-lg">
                <span>Total</span>
                <span>{formatRupiah(data.total)}</span>
              </div>
            </div>
          </div>
        )}

        <DialogClose asChild>
          <Button variant="outline" className="w-full mt-4">
            Close
          </Button>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
};
