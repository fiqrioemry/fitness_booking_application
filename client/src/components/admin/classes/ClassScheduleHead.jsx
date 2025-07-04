import { Button } from "@/components/ui/Button";
import { ArrowLeft, ArrowRight } from "lucide-react";
import { format, startOfWeek, endOfWeek } from "date-fns";

const ClassScheduleHead = ({ onNavigate, date }) => {
  const end = endOfWeek(date, { weekStartsOn: 6 });
  const start = startOfWeek(date, { weekStartsOn: 6 });

  return (
    <div className="flex items-center justify-between p-2">
      <div className="flex gap-2">
        <Button
          onClick={() => onNavigate("TODAY")}
          className="text-sm font-medium"
        >
          Today
        </Button>
        <Button
          onClick={() => onNavigate("PREV")}
          className="text-sm font-medium"
        >
          <ArrowLeft />
        </Button>
        <Button
          onClick={() => onNavigate("NEXT")}
          className="text-sm font-medium"
        >
          <ArrowRight />
        </Button>
      </div>
      <h2 className="text-center font-semibold text-lg">
        Event Schedule: {format(start, "eeee, d MMMM")} -{" "}
        {format(end, "eeee, d MMMM yyyy")}
      </h2>
      <div className="w-14" />
    </div>
  );
};

export default ClassScheduleHead;
