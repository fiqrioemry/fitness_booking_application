import { Loader2 } from "lucide-react";
import { Button } from "@/components/ui/Button";

export const SubmitButton = ({
  text = "Submit",
  className = "",
  disabled = false,
  isLoading = false,
  ...props
}) => {
  return (
    <Button
      type="submit"
      className="w-full"
      disabled={disabled || isLoading}
      {...props}
    >
      {isLoading ? <Loader2 className="animate-spin h-4 w-4 mr-2" /> : text}
    </Button>
  );
};
