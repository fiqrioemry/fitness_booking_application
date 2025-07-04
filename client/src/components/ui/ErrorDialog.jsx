import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogFooter,
  DialogContent,
  DialogDescription,
} from "@/components/ui/Dialog";

import { useState, useEffect } from "react";
import { AlertTriangle } from "lucide-react";
import { Button } from "@/components/ui/Button";

const ErrorDialog = ({ open = true, onRetry }) => {
  const [visible, setVisible] = useState(open);

  useEffect(() => {
    if (!visible && onRetry) {
      setTimeout(() => {
        onRetry();
      }, 500);
    }
  }, [visible, onRetry]);

  return (
    <div className="min-h-[65vh] flex items-center justify-center bg-background px-4">
      <Dialog open={visible} onOpenChange={setVisible}>
        <DialogContent className="max-w-md sm:rounded-2xl sm:p-6 shadow-lg">
          <DialogHeader className="flex flex-col items-center text-center space-y-2">
            <div className="flex items-center justify-center w-16 h-16 rounded-full bg-red-100 text-red-600">
              <AlertTriangle className="w-8 h-8" />
            </div>
            <DialogTitle className="text-xl font-semibold text-red-600">
              Failed to Load Data
            </DialogTitle>
            <DialogDescription className="text-muted-foreground text-sm max-w-xs">
              An error occurred while fetching the data. Please check your
              internet connection and click the button below to reload the page.
            </DialogDescription>
          </DialogHeader>

          <DialogFooter className="mt-4 w-full">
            <Button className="w-full" onClick={() => setVisible(false)}>
              Try Again
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export { ErrorDialog };
