import { Skeleton } from "@/components/ui/Skeleton";

export const DashboardSkeleton = () => {
  return (
    <div className="p-6 space-y-6">
      {/* Header */}
      <div>
        <Skeleton className="h-6 w-48 mb-4" />
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <Skeleton key={i} className="h-24 w-full rounded-lg" />
          ))}
        </div>
      </div>

      {/* Revenue Chart */}
      <div>
        <Skeleton className="h-6 w-48 mb-4" />
        <div className="bg-background rounded-xl shadow p-6 space-y-4">
          <div className="flex justify-between items-center">
            <Skeleton className="h-6 w-24" />
            <Skeleton className="h-10 w-[120px] rounded-md" />
          </div>
          <Skeleton className="h-56 w-full rounded-md" />
        </div>
      </div>

      {/* Recent Transactions */}
      <div>
        <Skeleton className="h-6 w-48 mb-4" />
        <div className="border rounded-lg shadow-sm overflow-hidden">
          <Skeleton className="h-12 w-full" />
          <div className="space-y-2 p-4">
            {Array.from({ length: 2 }).map((_, i) => (
              <Skeleton key={i} className="h-16 w-full rounded-md" />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};
