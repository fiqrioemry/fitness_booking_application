import {
  Clock,
  XCircle,
  CheckCircle,
  AlertTriangle,
  CalendarCheck,
} from "lucide-react";
import {
  Card,
  CardTitle,
  CardHeader,
  CardContent,
  CardDescription,
} from "@/components/ui/Card";
import { formatDate } from "@/lib/utils";
import { Badge } from "@/components/ui/Badge";
import { differenceInDays, isBefore } from "date-fns";

export const PackageCard = ({ pkgs }) => {
  const today = new Date();

  return (
    <div className="grid gap-6">
      {pkgs.map((item) => {
        const expiredDate = new Date(item.expiredAt);
        const daysLeft = differenceInDays(expiredDate, today);

        const isExpired = isBefore(expiredDate, today);
        const isCreditEmpty = item.remainingCredit === 0;
        const isActive = !isExpired && !isCreditEmpty;

        return (
          <Card
            key={item.id}
            className={`transition-shadow ${
              isActive ? "bg-card" : "bg-muted opacity-60"
            } py-4`}
          >
            <CardHeader className="items-start text-left p-5">
              <div className="flex justify-between w-full">
                <div>
                  <CardTitle className="text-base text-foreground">
                    {item.packageName}
                  </CardTitle>
                  <CardDescription>
                    Purchased on :{formatDate(item.purchasedAt)}
                  </CardDescription>
                </div>

                <div className="text-right space-y-3">
                  <Badge
                    variant={isActive ? "default" : "destructive"}
                    className="text-xs"
                  >
                    {item.remainingCredit} CREDIT
                  </Badge>

                  <div className="text-xs mt-1 flex items-center justify-end gap-1">
                    {isActive ? (
                      <>
                        <CheckCircle className="w-4 h-4 text-green-500" />
                        <span className="text-green-600">Active</span>
                      </>
                    ) : isExpired ? (
                      <>
                        <XCircle className="w-4 h-4 text-red-500" />
                        <span className="text-red-500">Expired</span>
                      </>
                    ) : (
                      <>
                        <AlertTriangle className="w-4 h-4 text-yellow-500" />
                        <span className="text-yellow-600">No Credit</span>
                      </>
                    )}
                  </div>
                </div>
              </div>
            </CardHeader>

            <CardContent className="pt-0 pb-5">
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-sm text-muted-foreground">
                <p className="flex items-center gap-2">
                  <CalendarCheck className="w-4 h-4" />
                  Expires on: {formatDate(expiredDate)}
                </p>
                <p className="flex items-center gap-2">
                  <Clock className="w-4 h-4" />
                  {isExpired ? (
                    <span>Expired {Math.abs(daysLeft)} day(s) ago</span>
                  ) : (
                    <span>Time left: {daysLeft} day(s)</span>
                  )}
                </p>
              </div>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
};
