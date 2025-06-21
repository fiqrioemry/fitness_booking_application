import { Loader } from "lucide-react";

const Loading = () => {
  return (
    <div className="h-screen w-full flex flex-col items-center justify-center bg-background animate-fadeIn">
      <div className="flex flex-col items-center space-y-4">
        <Loader size={48} className="animate-spin text-primary" />
        <p className="text-gray-600 text-sm tracking-wide">
          Loading, please wait ...
        </p>
      </div>
    </div>
  );
};

export { Loading };
