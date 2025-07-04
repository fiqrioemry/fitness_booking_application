import { Skeleton } from "@/components/ui/Skeleton";

export const AboutSkeleton = () => {
  return (
    <section className="section py-24 space-y-20 text-foreground">
      {/* Hero Section */}
      <div className="bg-primary text-primary-foreground rounded-xl shadow-md px-6 py-10 text-center space-y-2 mb-8">
        <Skeleton className="h-8 w-80 mx-auto" />
        <Skeleton className="h-4 w-[32rem] mx-auto" />
      </div>

      {/* Our Mission */}
      <div className="text-center space-y-4">
        <Skeleton className="h-6 w-40 mx-auto" />
        <Skeleton className="h-4 w-[28rem] mx-auto" />
      </div>

      {/* Why Choose Us */}
      <div className="space-y-6">
        <Skeleton className="h-6 w-60 mx-auto" />
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 text-center">
          {[...Array(3)].map((_, i) => (
            <div key={i} className="p-6 border rounded-xl space-y-4">
              <Skeleton className="w-8 h-8 mx-auto rounded-full" />
              <Skeleton className="h-4 w-32 mx-auto" />
              <Skeleton className="h-4 w-[80%] mx-auto" />
              <Skeleton className="h-4 w-[60%] mx-auto" />
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};
