import { Skeleton } from "@/components/ui/Skeleton";

export const ClassesSkeleton = () => {
  return (
    <section className="section py-24 text-foreground">
      {/* Heading */}
      <div className="bg-primary text-primary-foreground rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <Skeleton className="h-8 w-60 mx-auto" />
        <Skeleton className="h-4 w-80 mx-auto" />
      </div>

      {/* Filter Bar */}
      <div className="sticky top-4 z-10 bg-card text-foreground border border-border shadow-sm rounded-xl p-4 mb-8">
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
          {Array.from({ length: 5 }).map((_, i) => (
            <Skeleton key={i} className="h-10 w-full rounded-md" />
          ))}
        </div>
      </div>

      {/* Grid Class Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        {Array.from({ length: 8 }).map((_, i) => (
          <div
            key={i}
            className="border rounded-xl overflow-hidden shadow-sm flex flex-col h-full"
          >
            {/* Image */}
            <Skeleton className="h-48 w-full" />

            {/* Title & Description */}
            <div className="p-4 space-y-2">
              <Skeleton className="h-5 w-3/4" />
              <Skeleton className="h-4 w-full" />
            </div>

            {/* Footer */}
            <div className="px-4 pb-4 mt-auto">
              <Skeleton className="h-3 w-24" />
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};
