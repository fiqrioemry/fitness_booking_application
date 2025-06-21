import { Loader2 } from "lucide-react";

const SubmitLoading = ({ text = "Saving...", className = "" }) => {
  return (
    <div
      className={`flex flex-col items-center justify-center py-10 ${className}`}
    >
      <Loader2 className="animate-spin w-8 h-8 text-primary mb-3" />
      <p className="text-gray-600 text-sm">{text}</p>
    </div>
  );
};

export { SubmitLoading };
