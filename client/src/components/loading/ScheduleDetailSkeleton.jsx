import { Skeleton } from "@/components/ui/Skeleton";

export const ScheduleDetailSkeleton = () => {
  return (
    <section className="section py-24 text-foreground space-y-8">
      {/* Title & Date */}
      <div className="space-y-4 text-center">
        <Skeleton className="h-6 w-48 mx-auto" />
        <Skeleton className="h-4 w-40 mx-auto" />
      </div>

      {/* Content */}
      <div className="border rounded-xl p-6 flex flex-col md:flex-row gap-6 items-start">
        {/* Image */}
        <Skeleton className="w-full md:w-64 aspect-square rounded-lg" />

        {/* Detail */}
        <div className="flex-1 space-y-4 w-full">
          <div className="flex justify-between items-center">
            <Skeleton className="h-4 w-24" />
            <Skeleton className="h-4 w-20" />
          </div>
          <div className="flex justify-between items-center">
            <Skeleton className="h-4 w-24" />
            <Skeleton className="h-4 w-20" />
          </div>
          <div className="flex justify-between items-center">
            <Skeleton className="h-4 w-24" />
            <Skeleton className="h-4 w-28" />
          </div>

          {/* Instructor */}
          <div className="flex items-center gap-4 pt-4 border-t mt-4">
            <Skeleton className="w-10 h-10 rounded-full" />
            <div className="space-y-2">
              <Skeleton className="h-4 w-32" />
            </div>
          </div>
        </div>
      </div>

      {/* Button */}
      <div className="text-center">
        <Skeleton className="h-10 w-40 mx-auto rounded-md" />
      </div>
    </section>
  );
};
