import { buildDateTime } from "@/lib/utils";
import { DeleteClassSchedule } from "./DeleteClassSchedule";
import { UpdateClassSchedule } from "./UpdateClassSchedule";
import { Dialog, DialogContent, DialogTitle } from "@/components/ui/Dialog";

export const ClassScheduleDetail = ({ open, onClose, event }) => {
  if (!event) return null;

  const { title, start, end, resource } = event;

  const classStart = buildDateTime(
    resource?.date,
    resource?.startHour,
    resource?.startMinute
  );

  const now = new Date();
  const tenMinutesBeforeClass = new Date(classStart.getTime() - 10 * 60000);
  const isEditAllowed = now < tenMinutesBeforeClass;

  const isPast = classStart < now;
  const hasBooking = resource?.bookedCount > 0;

  return (
    <Dialog open={open} onOpenChange={onClose}>
      <DialogContent>
        <DialogTitle>{title}</DialogTitle>
        <p>
          <strong>Date :</strong> {new Date(start).toLocaleString()} -{" "}
          {new Date(end).toLocaleTimeString()}
        </p>
        <p>
          <strong>Instructor :</strong> {resource?.instructorName}
        </p>
        <p>
          <strong>Capacity :</strong> {resource?.bookedCount || 0}/
          {resource?.capacity}
        </p>
        {isEditAllowed && (
          <div className="mt-4 flex gap-4 justify-end">
            {!hasBooking && (
              <DeleteClassSchedule onClose={onClose} schedule={resource} />
            )}
            <UpdateClassSchedule onClose={onClose} schedule={resource} />
          </div>
        )}
      </DialogContent>
    </Dialog>
  );
};
