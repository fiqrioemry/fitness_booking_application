import { Badge } from "@/components/ui/Badge";
import { Button } from "@/components/ui/Button";
import { formatDate, formatHour } from "@/lib/utils";
import { Card, CardContent } from "@/components/ui/Card";
import { useLocation, useNavigate } from "react-router-dom";
import { CalendarIcon, ClockIcon, MapPinIcon } from "lucide-react";
import { ClassAccessInfo } from "./ClassAccessInfo";

export const UpcomingClassSchedules = ({ schedule }) => {
  const location = useLocation();
  const navigate = useNavigate();

  const handleSeeDetail = (id) => {
    navigate(`/instructor/schedules/${id}/attendance`, {
      state: { backgroundLocation: location },
    });
  };

  const handleStartClass = (id) => {
    navigate(`/instructor/schedules/${id}/open`, {
      state: { backgroundLocation: location },
    });
  };

  return (
    <Card>
      <div className="grid grid-cols-1 sm:grid-cols-[220px_1fr]">
        <img
          src={schedule.classImage}
          alt={schedule.className}
          className="w-full h-48 sm:h-full object-cover rounded-t-sm sm:rounded-l-sm sm:rounded-t-none"
        />

        <CardContent className="p-5 flex flex-col justify-between gap-5">
          <div className="w-full">
            <div className="flex justify-between items-start">
              <h3 className="text-xl font-semibold">{schedule.className}</h3>
              <Badge variant={schedule.isOpen ? "success" : "secondary"}>
                {schedule.isOpen ? "Started" : "Not Started"}
              </Badge>
            </div>

            <div className="mt-2 space-y-2 text-sm text-muted-foreground">
              <div className="flex items-center gap-2">
                <CalendarIcon className="w-4 h-4" />
                {formatDate(schedule.date)}
              </div>
              <div className="flex items-center gap-2">
                <ClockIcon className="w-4 h-4" />
                {formatHour(schedule.startHour, schedule.startMinute)} –{" "}
                {formatHour(
                  schedule.startHour,
                  schedule.startMinute + schedule.duration
                )}
              </div>
              <div className="flex items-center gap-2">
                <MapPinIcon className="w-4 h-4" />
                {schedule.location}
              </div>
              <div className="flex items-center gap-2">
                <span className="text-sm font-medium">Participants:</span>
                <Badge variant="outline">
                  {schedule.bookedCount} / {schedule.capacity}
                </Badge>
              </div>
            </div>
          </div>

          <div className="flex gap-3 justify-end flex-wrap">
            <Button
              variant="outline"
              onClick={() => handleSeeDetail(schedule.id)}
            >
              View Participants
            </Button>

            {!schedule.isOpen && (
              <Button onClick={() => handleStartClass(schedule.id)}>
                Start Class
              </Button>
            )}
            <ClassAccessInfo schedule={schedule} />
          </div>
        </CardContent>
      </div>
    </Card>
  );
};
