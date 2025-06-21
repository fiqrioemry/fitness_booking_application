import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogTrigger,
  DialogContent,
  DialogDescription,
} from "@/components/ui/Dialog";
import { toast } from "sonner";
import { Copy } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { formatDate, formatHour } from "@/lib/utils";

export const ClassAccessInfo = ({ schedule }) => {
  const handleCopy = () => {
    navigator.clipboard.writeText(schedule.zoomLink);
    toast.success("Zoom link copied to clipboard");
  };

  return (
    <Dialog>
      <DialogTrigger>
        <Button variant="outline">Access Info</Button>
      </DialogTrigger>
      <DialogContent className="max-w-md">
        <DialogHeader>
          <DialogTitle>🔒 Access Information</DialogTitle>
          <DialogDescription>
            This is your access info to join the class
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-3">
          <img
            src={schedule.classImage}
            alt={schedule.className}
            className="w-full h-40 object-cover rounded-lg"
          />

          <div>
            <h3 className="text-center">{schedule.className}</h3>
            <p className="text-sm text-muted-foreground">
              👤 {schedule.instructorName}
            </p>
            <p className="text-sm text-muted-foreground">
              🗓️ {formatDate(schedule.date)} · 🕒{" "}
              {formatHour(schedule.startHour, schedule.startMinute)} WIB · ⏱️{" "}
              {schedule.duration} mins
            </p>
          </div>

          <div className="bg-muted rounded-md p-4">
            <p className="text-sm font-medium">🔑 Verification Code:</p>
            <p className="text-xl font-bold tracking-widest text-primary">
              {schedule.verificationCode}
            </p>
          </div>

          {!schedule.zoomLink ? (
            <div className="bg-muted rounded-md p-4 space-y-2">
              <p className="text-sm font-medium">🔗 Zoom Link:</p>
              <div className="flex items-center gap-2">
                <p className="text-sm break-all flex-1">{schedule.zoomLink}</p>
                <Button size="icon" variant="ghost" onClick={handleCopy}>
                  <Copy className="w-4 h-4" />
                </Button>
              </div>
            </div>
          ) : (
            <div className="bg-yellow-100 text-yellow-800 text-sm p-3 rounded-md">
              🚫 This is an offline classes.
            </div>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
};
