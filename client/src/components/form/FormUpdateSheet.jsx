import {
  Sheet,
  SheetHeader,
  SheetTitle,
  SheetContent,
  SheetTrigger,
  SheetDescription,
} from "@/components/ui/Sheet";
import { Pencil } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { ScrollArea } from "@/components/ui/ScrollArea";
import { SubmitLoading } from "@/components/ui/SubmitLoading";
import { SubmitButton } from "@/components/form/SubmitButton";
import { useState, useCallback, useEffect, useMemo } from "react";
import { Dialog, DialogTitle, DialogContent } from "@/components/ui/Dialog";

export function FormUpdateSheet({
  title,
  state,
  schema,
  action,
  children,
  loading = false,
  shouldReset = true,
  buttonElement = (
    <Button variant="edit" size="icon" type="button">
      <Pencil className="w-4 h-4" />
    </Button>
  ),
}) {
  const [open, setOpen] = useState(false);
  const [showConfirm, setShowConfirm] = useState(false);

  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
    mode: "onChange",
  });

  const { formState, reset, handleSubmit } = methods;
  const isDirty = useMemo(() => formState.isDirty, [formState.isDirty]);

  const resetAndClose = useCallback(() => {
    reset();
    setOpen(false);
  }, [reset]);

  const handleCancel = useCallback(() => {
    if (isDirty) setShowConfirm(true);
    else resetAndClose();
  }, [isDirty, resetAndClose]);

  const handleConfirmClose = (confirmed) => {
    if (confirmed) resetAndClose();
    setShowConfirm(false);
  };

  useEffect(() => {
    if (state) reset(state);
  }, [state, reset]);

  const handleSave = useCallback(
    async (data) => {
      await action({ id: state.id, data });
      if (formState.isValid && shouldReset) reset();
      setOpen(false);
    },
    [action, formState.isValid, reset, shouldReset, state.id]
  );

  return (
    <>
      <Sheet
        open={open}
        onOpenChange={(val) => (!val ? handleCancel() : setOpen(val))}
      >
        <SheetTrigger asChild>{buttonElement}</SheetTrigger>

        <SheetContent className="w-full sm:max-w-lg flex flex-col p-0">
          {loading ? (
            <SubmitLoading />
          ) : (
            <FormProvider {...methods}>
              <form
                onSubmit={handleSubmit(handleSave)}
                className="flex flex-col h-full"
              >
                {/* Header */}
                <div className="border-b p-4">
                  <SheetHeader>
                    <SheetTitle className="text-lg font-semibold">
                      {title}
                    </SheetTitle>
                    <SheetDescription className="text-sm text-muted-foreground">
                      Submit button will activate when you make changes.
                    </SheetDescription>
                  </SheetHeader>
                </div>

                {/* Content */}
                <ScrollArea className="flex-1 py-4">
                  <div className="space-y-4 px-4">{children}</div>
                </ScrollArea>

                {/* Footer */}
                <div className="border-t p-4 flex justify-end gap-2">
                  <SubmitButton
                    text="Save Changes"
                    isLoading={loading}
                    disabled={!formState.isValid || !formState.isDirty}
                  />
                </div>
              </form>
            </FormProvider>
          )}
        </SheetContent>
      </Sheet>

      {/* Optional Confirmation Sheet */}
      {showConfirm && (
        <Dialog open={showConfirm} onOpenChange={setShowConfirm}>
          <DialogContent className="sm:max-w-md p-6 rounded-xl space-y-6 bg-card border border-border">
            <div className="text-center">
              <DialogTitle className="text-xl font-semibold text-foreground">
                Unsaved Changes
              </DialogTitle>
              <p className="mt-2 text-sm text-muted-foreground">
                You have made changes. Are you sure you want to discard them?
              </p>
            </div>

            <div className="flex justify-center gap-4">
              <Button
                variant="outline"
                className="w-32"
                onClick={() => handleConfirmClose(false)}
              >
                Keep Editing
              </Button>
              <Button
                variant="destructive"
                className="w-32"
                onClick={() => handleConfirmClose(true)}
              >
                Discard
              </Button>
            </div>
          </DialogContent>
        </Dialog>
      )}
    </>
  );
}
