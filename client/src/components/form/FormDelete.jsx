import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "@/components/ui/Dialog";
import { Trash2 } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

export const FormDelete = ({
  title,
  onDelete,
  description,
  buttonElement = (
    <Button variant="destructive" size="icon" type="button">
      <Trash2 className="w-4 h-4" />
    </Button>
  ),
  loading = false,
}) => {
  return (
    <Dialog>
      <DialogTrigger asChild>{buttonElement}</DialogTrigger>

      <DialogContent className="sm:max-w-md rounded-xl bg-background border border-border p-6 space-y-6">
        {loading ? (
          <SubmitLoading text="Deleting..." />
        ) : (
          <>
            <div className="text-center space-y-2">
              <DialogTitle className="text-2xl font-bold text-foreground">
                {title}
              </DialogTitle>
              <DialogDescription className="text-sm text-muted-foreground">
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
                  onClick={onDelete}
                >
                  Delete
                </Button>
              </DialogClose>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
