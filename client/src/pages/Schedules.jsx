import {
  Carousel,
  CarouselItem,
  CarouselContent,
} from "@/components/ui/Carousel";
import {
  useSchedulesQuery,
  useSchedulesWithStatusQuery,
} from "@/hooks/useSchedules";
import { Link } from "react-router-dom";
import { scheduleTitle } from "@/lib/constant";
import { Button } from "@/components/ui/Button";
import { useState, useMemo, useRef } from "react";
import { useAuthStore } from "@/store/useAuthStore";
import { format, isSameDay, addDays } from "date-fns";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useDocumentTitle } from "@/hooks/useDocumentTitle";
import { SchedulesSkeleton } from "@/components/loading/ScheduleSkeleton";

const Schedules = () => {
  const { user } = useAuthStore();
  useDocumentTitle(scheduleTitle);
  const todayRef = useRef(new Date());
  const today = todayRef.current;
  const [selectedDate, setSelectedDate] = useState(today);

  const {
    data = [],
    isLoading,
    isError,
    refetch,
  } = user === null ? useSchedulesQuery() : useSchedulesWithStatusQuery();

  const dateRange = useMemo(() => {
    return Array.from({ length: 14 }, (_, i) => addDays(today, i));
  }, [today]);

  const filteredSchedules = useMemo(() => {
    return data
      .map((item) => {
        const start = new Date(item.date);
        start.setHours(item.startHour, item.startMinute, 0, 0);
        const end = new Date(start.getTime() + 60 * 60 * 1000);
        return { ...item, startTime: start, endTime: end };
      })
      .filter((item) => isSameDay(item.startTime, selectedDate));
  }, [data, selectedDate]);

  if (isLoading) return <SchedulesSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <section className="section py-24 text-foreground">
      {/* Mobile: Carousel */}
      <div className="block md:hidden">
        <Carousel opts={{ align: "start" }} className="w-full mb-6">
          <CarouselContent>
            {dateRange.map((date) => {
              const isSelected = isSameDay(date, selectedDate);
              return (
                <CarouselItem
                  key={format(date, "yyyy-MM-dd")}
                  className="basis-auto px-2"
                >
                  <button
                    onClick={() => setSelectedDate(date)}
                    aria-label={`Select ${format(date, "EEEE, dd MMMM")}`}
                    className={`w-14 h-16 flex flex-col items-center justify-center rounded-lg transition font-medium flex-shrink-0 ${
                      isSelected
                        ? "bg-primary text-primary-foreground"
                        : "bg-muted hover:bg-accent hover:text-background text-muted-foreground"
                    }`}
                  >
                    <span className="text-xs">{format(date, "EEE")}</span>
                    <span className="text-lg font-bold">
                      {format(date, "dd")}
                    </span>
                  </button>
                </CarouselItem>
              );
            })}
          </CarouselContent>
        </Carousel>
      </div>

      {/* Desktop: Static Horizontal Layout */}
      <div className="hidden md:flex items-center justify-center gap-4 mb-6">
        {dateRange.map((date) => {
          const isSelected = isSameDay(date, selectedDate);
          return (
            <button
              key={format(date, "yyyy-MM-dd")}
              onClick={() => setSelectedDate(date)}
              aria-label={`Select ${format(date, "EEEE, dd MMMM")}`}
              className={`w-14 h-16 flex flex-col items-center justify-center rounded-lg transition font-medium ${
                isSelected
                  ? "bg-primary text-primary-foreground"
                  : "bg-muted hover:bg-accent hover:text-background text-muted-foreground"
              }`}
            >
              <span className="text-xs">{format(date, "EEE")}</span>
              <span className="text-lg font-bold">{format(date, "dd")}</span>
            </button>
          );
        })}
      </div>

      {/* Summary */}
      <div className="text-sm text-muted-foreground mb-6 text-center">
        <strong className="text-foreground">
          {format(selectedDate, "EEEE, dd MMM")}
        </strong>{" "}
        • {filteredSchedules.length} classes
      </div>

      {/* Schedule List */}
      <div className="space-y-4">
        {filteredSchedules.length > 0 ? (
          filteredSchedules.map((s) => (
            <div
              key={s.id}
              className="bg-card border border-border rounded-xl px-6 py-8 shadow-sm flex flex-col md:flex-row md:items-center md:justify-between"
            >
              {/* Left */}
              <div className="flex flex-col space-y-2">
                <p className="text-sm font-semibold">
                  {format(s.startTime, "h:mm a")} •{" "}
                  {Math.round((s.endTime - s.startTime) / 60000)} mins
                </p>
                <p className="font-medium text-base">{s.className}</p>
                <p className="text-sm text-muted-foreground">
                  {s.instructorName}
                </p>
              </div>

              {/* Right */}
              <div className="mt-4 md:mt-0 md:text-right flex flex-col items-end gap-2">
                <p className="text-sm text-muted-foreground">
                  {s.bookedCount} / {s.capacity} Slot
                </p>
                {s.isBooked ? (
                  <Link to="/profile/bookings">
                    <Button
                      variant="outline"
                      className="text-green-700 border-green-500 bg-green-100 hover:bg-green-200"
                    >
                      Booked
                    </Button>
                  </Link>
                ) : (
                  <>
                    {Date.now() > s.startTime.getTime() ? (
                      <Button
                        disabled
                        variant="outline"
                        className="opacity-60 cursor-not-allowed"
                      >
                        Closed
                      </Button>
                    ) : (
                      <Link to={`/schedules/${s.id}`}>
                        <Button>Book Now</Button>
                      </Link>
                    )}
                  </>
                )}
              </div>
            </div>
          ))
        ) : (
          <p className="text-center text-muted-foreground text-lg mt-20">
            📭 No schedules available for this day.
          </p>
        )}
      </div>
    </section>
  );
};

export default Schedules;
