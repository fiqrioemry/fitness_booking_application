import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableHeader,
} from "@/components/ui/Table";
import { useEffect } from "react";
import { CirclePlus } from "lucide-react";
import { formatRupiah } from "@/lib/utils";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { Loading } from "@/components/ui/Loading";
import { useVouchersQuery } from "@/hooks/useVouchers";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Card, CardContent } from "@/components/ui/Card";
import { SectionTitle } from "@/components/header/SectionTitle";
import { VoucherUpdate } from "@/components/admin/vouchers/VoucherUpdate";
import { VoucherDelete } from "@/components/admin/vouchers/VoucherDelete";
import { truncateText } from "@/lib/utils";

const VouchersList = () => {
  const {
    data: vouchers = [],
    isLoading,
    isError,
    refetch,
  } = useVouchersQuery();
  const navigate = useNavigate();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section">
      <SectionTitle
        title="Vouchers Management"
        description="View, add, and manage training packages available for purchase by
          users."
      />

      <div className="flex justify-end mt-4">
        <Button size="nav" onClick={() => navigate("/admin/vouchers/add")}>
          <CirclePlus className="w-4 h-4 mr-2" />
          New Voucher
        </Button>
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          <div className="hidden md:block max-w-8xl w-full">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Code</TableHead>
                  <TableHead>Description</TableHead>
                  <TableHead>Type</TableHead>
                  <TableHead>Discount</TableHead>
                  <TableHead>Max Discount</TableHead>
                  <TableHead>Quota</TableHead>
                  <TableHead>Expired At</TableHead>
                  <TableHead>Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {vouchers.map((voucher) => (
                  <TableRow key={voucher.id}>
                    <TableCell>{voucher.code}</TableCell>
                    <TableCell>
                      <p className="text-sm text-muted-foreground">
                        {truncateText(voucher.description, 20)}
                      </p>
                    </TableCell>
                    <TableCell>{voucher.discountType}</TableCell>
                    <TableCell>
                      {voucher.discountType === "percentage"
                        ? `${voucher.discount}%`
                        : `${formatRupiah(voucher.discount)}`}
                    </TableCell>
                    <TableCell>
                      {voucher.maxDiscount
                        ? `${formatRupiah(voucher.maxDiscount)}`
                        : "-"}
                    </TableCell>
                    <TableCell>{voucher.quota}</TableCell>
                    <TableCell>{voucher.expiredAt}</TableCell>
                    <TableCell className="space-x-2">
                      <VoucherDelete voucher={voucher} />
                      <VoucherUpdate voucher={voucher} />
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          {/* Mobile View */}
          <div className="md:hidden w-full space-y-4 p-4">
            {vouchers.map((voucher) => (
              <div
                key={voucher.id}
                className="border border-border rounded-lg p-4 shadow-sm bg-background"
              >
                <div className="space-y-1 mb-3">
                  <h3 className="text-base font-semibold text-foreground">
                    {voucher.code}
                  </h3>
                  <p className="text-sm text-muted-foreground">
                    {truncateText(voucher.description, 10)}
                  </p>
                </div>

                <div className="text-sm text-muted-foreground space-y-1 mb-3">
                  <p>
                    <span className="font-medium text-foreground">Type:</span>{" "}
                    {voucher.discountType}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Discount:
                    </span>{" "}
                    {voucher.discountType === "percentage"
                      ? `${voucher.discount}%`
                      : `Rp ${voucher.discount.toLocaleString("id-ID")}`}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Max Discount:
                    </span>{" "}
                    {voucher.maxDiscount
                      ? `Rp ${voucher.maxDiscount.toLocaleString("id-ID")}`
                      : "-"}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">Quota:</span>{" "}
                    {voucher.quota}
                  </p>
                  <p>
                    <span className="font-medium text-foreground">
                      Expired:
                    </span>{" "}
                    {voucher.expiredAt}
                  </p>
                </div>

                <div className="flex justify-end gap-2">
                  <VoucherDelete voucher={voucher} />
                  <VoucherUpdate voucher={voucher} />
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </section>
  );
};

export default VouchersList;
