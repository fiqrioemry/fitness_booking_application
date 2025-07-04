import { Skeleton } from "@/components/ui/Skeleton";

export const ClassAttendanceSkeleton = () => {
  return (
    <div className="mt-6 space-y-4">
      {Array.from({ length: 4 }).map((_, i) => (
        <div
          key={i}
          className="flex items-center gap-4 border border-muted rounded-xl px-4 py-3"
        >
          <Skeleton className="w-10 h-10 rounded-full" />
          <div className="flex-1 space-y-1">
            <Skeleton className="w-32 h-3" />
            <Skeleton className="w-44 h-2" />
          </div>
          <Skeleton className="w-14 h-4 rounded-md" />
        </div>
      ))}
    </div>
  );
};
