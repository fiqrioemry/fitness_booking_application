import { Skeleton } from "@/components/ui/Skeleton";
import { Loader } from "lucide-react";

export const SectionSkeleton = () => {
  return (
    <div className="section space-y-6 h-full flex items-center justify-center">
      <div className="flex flex-col items-center space-y-4">
        <Loader size={48} className="animate-spin text-primary" />
        <p className="text-gray-600 text-sm tracking-wide">
          Loading, please wait ...
        </p>
      </div>
    </div>
  );
};
