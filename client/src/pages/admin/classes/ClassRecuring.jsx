import { useEffect } from "react";
import { Badge } from "@/components/ui/Badge";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/Button";
import { Loading } from "@/components/ui/Loading";
import { CalendarClock, PlusCircle } from "lucide-react";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SectionTitle } from "@/components/header/SectionTitle";
import { useRecurringTemplatesQuery } from "@/hooks/useSchedules";
import { RunTemplate } from "@/components/admin/classes/RunTemplate";
import { StopTemplate } from "@/components/admin/classes/StopTemplate";
import { format, addMonths, parseISO, isValid, isBefore } from "date-fns";
import { DeleteTemplate } from "@/components/admin/classes/DeleteTemplate";
import { UpdateTemplate } from "@/components/admin/classes/UpdateTemplate";

const weekdays = [
  "Sunday",
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
];

const ClassRecuring = () => {
  const { data, isLoading, isError, refetch } = useRecurringTemplatesQuery();
  const navigate = useNavigate();

  useEffect(() => {
    refetch();
  }, []);

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const templates = data || [];

  return (
    <section className="section max-w-7xl mx-auto px-4 py-10 text-foreground space-y-6">
      <SectionTitle
        title="Recurring Schedule Templates"
        description="Manage, activate, and generate recurring class schedules."
      />

      {templates.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-16 text-center">
          <img
            src="/no-bookings.webp"
            alt="No templates"
            className="w-72 h-72 mb-6 object-contain"
          />
          <h3 className="text-lg font-semibold text-foreground">
            No templates yet
          </h3>
          <p className="text-muted-foreground mb-6">
            You haven’t created any recurring schedule template. Let’s start
            now!
          </p>
          <Button onClick={() => navigate("/admin/classes/schedules")}>
            <PlusCircle className="w-4 h-4 mr-2" />
            Add Template
          </Button>
        </div>
      ) : (
        <div className="flex flex-col gap-4">
          {templates.map((t) => {
            const createdAt = isValid(parseISO(t.createdAt))
              ? parseISO(t.createdAt)
              : new Date();
            const endDate = isValid(parseISO(t.endDate))
              ? parseISO(t.endDate)
              : null;
            const nextGen = addMonths(createdAt, 1);

            const showNextGen = endDate && isBefore(nextGen, endDate);
            const endDateStr = endDate ? format(endDate, "yyyy-MM-dd") : "-";

            return (
              <div
                key={t.id}
                className="border rounded-xl p-5 bg-background shadow-sm hover:shadow transition flex flex-col md:flex-row md:items-center md:justify-between gap-4"
              >
                <div className="space-y-1 md:max-w-2xl">
                  <div className="flex items-center gap-2">
                    <CalendarClock className="w-4 h-4 text-primary" />
                    <h3 className="font-semibold text-base truncate">
                      {t.className}
                    </h3>
                    <Badge variant={t.isActive ? "success" : "outline"}>
                      {t.isActive ? "Active" : "Inactive"}
                    </Badge>
                  </div>
                  <ul className="text-sm text-muted-foreground space-y-0.5 mt-1">
                    <li>
                      <strong>Instructor:</strong> {t.instructorName}
                    </li>
                    <li>
                      <strong>Capacity:</strong> {t.capacity}
                    </li>
                    <li>
                      <strong>Time:</strong> {t.startHour}:
                      {t.startMinute.toString().padStart(2, "0")}
                    </li>
                    <li>
                      <strong>Days:</strong>{" "}
                      {t.dayOfWeeks.map((d) => weekdays[d]).join(", ")}
                    </li>
                    <li>
                      <strong>End Date:</strong> {endDateStr}
                    </li>
                    {showNextGen && (
                      <li>
                        <strong>Next Generation:</strong>{" "}
                        {format(nextGen, "yyyy-MM-dd")}
                      </li>
                    )}
                  </ul>
                </div>

                <div className="flex flex-wrap gap-2 md:justify-end">
                  {t.isActive ? (
                    <StopTemplate template={t} />
                  ) : (
                    <>
                      <RunTemplate template={t} />
                      <UpdateTemplate template={t} />
                      <DeleteTemplate template={t} />
                    </>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      )}
    </section>
  );
};

export default ClassRecuring;
