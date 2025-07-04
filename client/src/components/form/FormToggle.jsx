import {
  Dialog,
  DialogTitle,
  DialogClose,
  DialogContent,
  DialogTrigger,
  DialogDescription,
} from "@/components/ui/Dialog";
import { Play, StopCircle } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

export const FormToggle = ({
  title,
  description,
  onToggle,
  type = "start",
  loading = false,
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>
        {type === "start" ? (
          <Button className="w-full" type="button">
            <Play className="w-4 h-4" />
            <span className="capitalize">start</span>
          </Button>
        ) : (
          <Button variant="destructive" className="w-full" type="button">
            <StopCircle className="w-4 h-4" />
            <span className="capitalize">Stop</span>
          </Button>
        )}
      </DialogTrigger>

      <DialogContent className="sm:max-w-md rounded-xl p-6 space-y-6">
        {loading ? (
          <SubmitLoading text="Processing..." />
        ) : (
          <>
            <div className="text-center space-y-2">
              <DialogTitle className="text-2xl font-bold text-foreground">
                {title}
              </DialogTitle>
              <DialogDescription className="text-muted-foreground">
                {description}
              </DialogDescription>
            </div>

            <div className="flex justify-center gap-4 pt-4">
              <DialogClose asChild>
                <Button variant="secondary" className="w-32">
                  Cancel
                </Button>
              </DialogClose>

              <DialogClose asChild>
                <Button
                  variant="destructive"
                  className="w-32"
                  onClick={onToggle}
                >
                  Yes, Execute
                </Button>
              </DialogClose>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
