import { Skeleton } from "@/components/ui/Skeleton";

export const OurClassSkeleton = () => {
  return (
    <section className="py-20 px-4">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-7xl mx-auto">
        {[1, 2, 3].map((_, i) => (
          <div
            key={i}
            className="bg-card text-foreground border border-border rounded-xl shadow hover:shadow-xl transition"
          >
            <Skeleton className="w-full h-56 rounded-t-xl" />
            <div className="p-6 space-y-4">
              <Skeleton className="h-5 w-3/4" />
              <Skeleton className="h-4 w-2/3" />
              <div className="flex justify-between items-center">
                <Skeleton className="h-5 w-20" />
                <Skeleton className="h-8 w-24 rounded-md" />
              </div>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};
