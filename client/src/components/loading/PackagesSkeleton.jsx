import { Skeleton } from "@/components/ui/Skeleton";

export const PackagesSkeleton = () => {
  return (
    <section className="section py-24 text-foreground">
      {/* Header */}
      <div className="bg-primary text-primary-foreground rounded-xl shadow px-6 py-10 text-center space-y-2 mb-10">
        <Skeleton className="h-8 w-52 mx-auto" />
        <Skeleton className="h-4 w-80 mx-auto" />
      </div>

      {/* Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        {Array.from({ length: 6 }).map((_, i) => (
          <div
            key={i}
            className="border rounded-xl overflow-hidden shadow-sm flex flex-col"
          >
            <Skeleton className="h-44 w-full" />

            <div className="p-5 flex flex-col gap-3">
              <Skeleton className="h-5 w-32 mx-auto" />
              <Skeleton className="h-4 w-48 mx-auto" />

              <div className="text-sm text-muted-foreground space-y-1">
                <Skeleton className="h-4 w-32" />
                <Skeleton className="h-6 w-40" />
              </div>

              <div className="space-y-2 mt-2">
                <Skeleton className="h-4 w-full" />
                <Skeleton className="h-4 w-5/6" />
                <Skeleton className="h-4 w-2/3" />
              </div>
            </div>

            <div className="p-5 mt-auto">
              <Skeleton className="h-10 w-full rounded-md" />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};
