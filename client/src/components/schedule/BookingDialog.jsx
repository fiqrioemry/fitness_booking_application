import {
  DialogTitle,
  DialogHeader,
  DialogContent,
} from "@/components/ui/Dialog";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useCreateBookingMutation } from "@/hooks/useBooking";
import { useUserClassPackagesQuery } from "@/hooks/useUserPackage";

export const BookingDialog = ({ schedule, openDialog, setOpenDialog }) => {
  const navigate = useNavigate();

  const [selectedPackageId, setSelectedPackageId] = useState(null);

  const bookingMutation = useCreateBookingMutation();

  const { data: userPackages = [], isLoading } = useUserClassPackagesQuery(
    schedule?.classId,
    {
      enabled: !!schedule?.classId && openDialog,
    }
  );

  const handleBooking = () => {
    if (!selectedPackageId || !schedule?.id) return;
    bookingMutation.mutate(
      {
        scheduleId: schedule.id,
        packageId: selectedPackageId,
      },
      {
        onSuccess: () => {
          setOpenDialog(false);
          navigate("/profile/bookings");
        },
      }
    );
  };

  return (
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Choose a Package</DialogTitle>
      </DialogHeader>

      {isLoading ? (
        <p>Loading packages...</p>
      ) : userPackages && userPackages.length > 0 ? (
        <div className="space-y-4">
          {userPackages.map((p) => (
            <div
              key={p.id}
              onClick={() => setSelectedPackageId(p.packageId)}
              className={`border p-3 rounded cursor-pointer transition ${
                selectedPackageId === p.packageId
                  ? "border-primary bg-muted"
                  : ""
              }`}
            >
              <div className="font-semibold">{p.packageName}</div>
              <div className="text-sm text-muted-foreground">
                Credit: {p.remainingCredit} • Exp: {formatDateTime(p.expiredAt)}
              </div>
            </div>
          ))}
          <Button
            className="mt-4 w-full"
            onClick={handleBooking}
            disabled={!selectedPackageId}
          >
            Confirm Booking
          </Button>
        </div>
      ) : (
        <div className="space-y-4">
          <p className="text-muted-foreground text-sm">
            You don’t have a package for this class. Please buy one:
          </p>
          {schedule.packages?.map((p) => (
            <div
              key={p.id}
              onClick={() => navigate(`/packages/${p.id}`)}
              className="border p-3 rounded cursor-pointer hover:bg-muted transition flex gap-4"
            >
              <img
                src={p.image}
                alt={p.name}
                className="w-20 h-20 object-cover rounded"
              />
              <div>
                <div className="font-semibold">{p.name}</div>
                <div className="text-sm text-muted-foreground">
                  {formatRupiah(p.price)}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </DialogContent>
  );
};
