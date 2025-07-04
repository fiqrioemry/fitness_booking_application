import {
  Carousel,
  CarouselContent,
  CarouselItem,
} from "@/components/ui/Carousel";
import { format, addDays } from "date-fns";
import { Skeleton } from "@/components/ui/Skeleton";

export const SchedulesSkeleton = () => {
  const today = new Date();
  const dateRange = Array.from({ length: 14 }, (_, i) => addDays(today, i));

  return (
    <section className="section py-24 text-foreground">
      {/* Mobile: Carousel */}
      <div className="block md:hidden mb-6">
        <Carousel opts={{ align: "start" }} className="w-full">
          <CarouselContent>
            {dateRange.map((date) => (
              <CarouselItem
                key={format(date, "yyyy-MM-dd")}
                className="basis-auto px-2"
              >
                <Skeleton className="w-14 h-16 rounded-lg" />
              </CarouselItem>
            ))}
          </CarouselContent>
        </Carousel>
      </div>

      {/* Desktop: Static Horizontal Layout */}
      <div className="hidden md:flex items-center justify-center gap-4 mb-6">
        {dateRange.map((date) => (
          <Skeleton
            key={format(date, "yyyy-MM-dd")}
            className="w-14 h-16 rounded-lg"
          />
        ))}
      </div>

      {/* Summary */}
      <div className="text-sm text-muted-foreground mb-6 text-center">
        <Skeleton className="h-4 w-64 mx-auto" />
      </div>

      {/* Schedule List Skeleton */}
      <div className="space-y-4">
        {Array.from({ length: 3 }).map((_, i) => (
          <div
            key={i}
            className="bg-card border border-border rounded-xl px-6 py-8 shadow-sm flex flex-col md:flex-row md:items-center md:justify-between"
          >
            {/* Left */}
            <div className="flex flex-col space-y-2">
              <Skeleton className="h-4 w-40" />
              <Skeleton className="h-6 w-48" />
              <Skeleton className="h-4 w-32" />
            </div>

            {/* Right */}
            <div className="mt-4 md:mt-0 flex flex-col items-end gap-2">
              <Skeleton className="h-4 w-28" />
              <Skeleton className="h-9 w-28 rounded-md" />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};
