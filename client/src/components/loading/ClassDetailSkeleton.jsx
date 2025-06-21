import { Skeleton } from "@/components/ui/Skeleton";

export const ClassDetailSkeleton = () => {
  return (
    <section className="section py-24 text-foreground space-y-8">
      {/* Back Button */}
      <div>
        <Skeleton className="h-4 w-32" />
      </div>

      {/* Header */}
      <div className="flex flex-col lg:flex-row gap-8 items-start">
        {/* Image */}
        <Skeleton className="w-full lg:w-1/2 h-64 rounded-2xl" />

        {/* Info */}
        <div className="w-full space-y-4">
          <Skeleton className="h-8 w-64" />
          <Skeleton className="h-4 w-full" />
          <Skeleton className="h-4 w-5/6" />
          <div className="flex flex-wrap gap-2">
            {Array.from({ length: 4 }).map((_, i) => (
              <Skeleton key={i} className="h-6 w-20 rounded-full" />
            ))}
          </div>
          <div className="space-y-2">
            <Skeleton className="h-4 w-1/2" />
            <Skeleton className="h-4 w-1/3" />
            <Skeleton className="h-4 w-2/3" />
            <Skeleton className="h-4 w-1/2" />
          </div>
        </div>
      </div>

      {/* Gallery */}
      <div className="space-y-4">
        <Skeleton className="h-6 w-32" />
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <Skeleton key={i} className="w-full h-60 rounded-xl" />
          ))}
        </div>
      </div>

      {/* Reviews */}
      <div className="space-y-4">
        <Skeleton className="h-6 w-24" />
        {Array.from({ length: 2 }).map((_, i) => (
          <div key={i} className="border rounded-xl p-4 space-y-2">
            <div className="flex items-center gap-4">
              <Skeleton className="h-4 w-24" />
              <div className="flex gap-1">
                {Array.from({ length: 5 }).map((_, j) => (
                  <Skeleton key={j} className="h-4 w-4 rounded-full" />
                ))}
              </div>
            </div>
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-3 w-32" />
          </div>
        ))}
      </div>
    </section>
  );
};
