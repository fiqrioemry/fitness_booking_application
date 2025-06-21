import { Skeleton } from "@/components/ui/Skeleton";

export const BookedDetailSkeleton = () => (
  <div className="space-y-4 mt-4">
    <Skeleton className="w-full h-48" />
    <Skeleton className="w-3/4 h-6" />
    <Skeleton className="w-1/2 h-5" />
    <Skeleton className="w-full h-5" />
    <Skeleton className="w-full h-5" />
    <Skeleton className="w-full h-5" />
    <div className="flex gap-2">
      <Skeleton className="h-10 w-1/2 rounded-md" />
      <Skeleton className="h-10 w-1/2 rounded-md" />
    </div>
  </div>
);
