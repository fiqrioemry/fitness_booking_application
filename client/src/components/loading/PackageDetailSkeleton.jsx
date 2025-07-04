import { Skeleton } from "@/components/ui/Skeleton";

export const PackageDetailSkeleton = () => {
  return (
    <section className="section py-24 text-foreground">
      {/* Back Button */}
      <div className="mb-6">
        <Skeleton className="h-4 w-32" />
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-10 items-start">
        {/* LEFT: Image + Info */}
        <div className="md:col-span-2 space-y-5">
          <Skeleton className="w-full h-[400px] rounded-xl" />

          <Skeleton className="h-8 w-64" />
          <Skeleton className="h-4 w-80" />

          <div className="flex gap-2 flex-wrap mt-2">
            <Skeleton className="h-6 w-40 rounded-full" />
            <Skeleton className="h-6 w-24 rounded-full" />
          </div>

          <div className="space-y-2 pl-5">
            <Skeleton className="h-4 w-60" />
            <Skeleton className="h-4 w-52" />
            <Skeleton className="h-4 w-48" />
          </div>
        </div>

        {/* RIGHT: Checkout */}
        <div className="bg-card border shadow-md rounded-2xl p-5 space-y-4 sticky top-24">
          <Skeleton className="h-6 w-40" />

          <div className="space-y-2">
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-5/6" />
          </div>

          <div className="space-y-2">
            <Skeleton className="h-4 w-1/2" />
            <Skeleton className="h-10 w-full rounded-md" />
          </div>

          <Skeleton className="h-px w-full bg-border" />

          <div className="flex justify-between items-center">
            <Skeleton className="h-5 w-20" />
            <Skeleton className="h-5 w-32" />
          </div>

          <Skeleton className="h-10 w-full mt-2 rounded-md" />
        </div>
      </div>
    </section>
  );
};
