import { toast } from "sonner";
import { ArrowLeft } from "lucide-react";
import { useEffect, useState } from "react";
import { Badge } from "@/components/ui/Badge";
import { Input } from "@/components/ui/Input";
import { Button } from "@/components/ui/Button";
import { useAuthStore } from "@/store/useAuthStore";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { useParams, useNavigate } from "react-router-dom";
import { usePackageDetailQuery } from "@/hooks/usePackage";
import { useStripePayment } from "@/hooks/useStripePayment";
import { PackageDetailSkeleton } from "@/components/loading/PackageDetailSkeleton";

const PackageDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const [voucherCode, setVoucherCode] = useState("");
  const [voucherInfo, setVoucherInfo] = useState(null);

  const { applyVoucher } = useVoucherMutation();
  const { data: pkg, isLoading, isError } = usePackageDetailQuery(id);
  const { handleBuyNow, isPending } = useStripePayment(id, voucherCode);

  useEffect(() => {
    if (!isLoading && (isError || !pkg?.id)) {
      navigate("/not-found", { replace: true });
    }
  }, [isLoading, isError, pkg, navigate]);

  if (isLoading || !pkg?.id) return <PackageDetailSkeleton />;

  const discountedPrice =
    pkg.Discount > 0 ? pkg.price * (1 - pkg.Discount / 100) : pkg.price;
  const voucherDiscount = voucherInfo?.discountValue || 0;
  const priceAfterVoucher = discountedPrice - voucherDiscount;
  const tax = priceAfterVoucher * 0.1;
  const finalTotal = priceAfterVoucher + tax;

  const handleApplyVoucher = () => {
    if (!voucherCode.trim()) return;
    applyVoucher.mutate(
      { userId: user?.id, code: voucherCode, total: priceAfterVoucher },
      {
        onSuccess: (res) => {
          setVoucherInfo(res);
          toast.success(`Voucher "${res.code}" applied`);
        },
        onError: () => {
          setVoucherInfo(null);
        },
      }
    );
  };

  return (
    <>
      <section className="section py-24 text-foreground">
        <div className="mb-6">
          <button
            onClick={() => navigate(-1)}
            className="flex items-center text-sm text-muted-foreground hover:text-primary transition"
          >
            <ArrowLeft className="w-4 h-4 mr-1" />
            Back to Packages
          </button>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-10 items-start">
          {/* LEFT */}
          <div className="md:col-span-2 space-y-5">
            <img
              src={pkg.image}
              alt={pkg.name}
              className="rounded-xl w-full h-[400px] object-cover border"
            />

            <h2 className="text-3xl font-bold">{pkg.name}</h2>
            <p className="text-muted-foreground text-sm">{pkg.description}</p>

            <div className="mt-2 flex items-center gap-2 flex-wrap">
              <Badge variant="outline">
                {pkg.credit} Credits • Rp {pkg.price.toLocaleString("id-ID")}
              </Badge>
              {pkg.Discount > 0 && (
                <Badge className="bg-green-600 text-white">
                  {pkg.Discount}% Discount
                </Badge>
              )}
            </div>

            <ul className="list-disc text-muted-foreground pl-5">
              {pkg.additional?.map((item, idx) => (
                <li key={idx}>{item}</li>
              ))}
            </ul>
          </div>

          {/* RIGHT */}
          <div className="bg-card border shadow-md rounded-2xl p-5 sticky top-24 space-y-4">
            <h3 className="text-xl font-semibold mb-1">Checkout</h3>

            <div className="text-sm flex justify-between text-muted-foreground">
              <span>Base Price</span>
              <span>Rp {pkg.price.toLocaleString("id-ID")}</span>
            </div>

            {pkg.Discount > 0 && (
              <div className="text-sm text-red-600 flex justify-between">
                <span>Discount</span>
                <span>-{pkg.Discount}%</span>
              </div>
            )}

            <div className="text-sm flex justify-between text-muted-foreground">
              <span>Tax (10%)</span>
              <span>Rp {tax.toLocaleString("id-ID")}</span>
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Promo Code</label>
              <div className="flex items-center gap-2">
                <Input
                  value={voucherCode}
                  onChange={(e) => setVoucherCode(e.target.value)}
                  placeholder="Enter code"
                  className="flex-1"
                />
                <Button
                  size="sm"
                  onClick={handleApplyVoucher}
                  disabled={applyVoucher.isPending}
                >
                  Apply
                </Button>
              </div>
              {voucherInfo && (
                <p className="text-xs text-green-600">
                  Applied: {voucherInfo.code} (Save Rp{" "}
                  {voucherDiscount.toLocaleString("id-ID")})
                </p>
              )}
            </div>

            <hr />

            <div className="text-base font-semibold flex justify-between">
              <span>Total</span>
              <span className="text-primary">
                Rp {finalTotal.toLocaleString("id-ID")}
              </span>
            </div>

            <Button
              size="lg"
              disabled={isPending}
              className="w-full mt-2"
              onClick={handleBuyNow}
            >
              {isPending ? "Processing..." : "Buy Now"}
            </Button>
          </div>
        </div>
      </section>
    </>
  );
};

export default PackageDetail;
