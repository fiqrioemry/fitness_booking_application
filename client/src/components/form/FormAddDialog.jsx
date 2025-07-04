import {
  Dialog,
  DialogTitle,
  DialogTrigger,
  DialogContent,
  DialogDescription,
} from "@/components/ui/Dialog";
import { PlusCircle } from "lucide-react";
import { Button } from "@/components/ui/Button";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { ScrollArea } from "@/components/ui/ScrollArea";
import { SubmitLoading } from "@/components/ui/SubmitLoading";
import { SubmitButton } from "@/components/form/SubmitButton";
import { useState, useCallback, useMemo, useEffect } from "react";

export function FormAddDialog({
  title,
  state,
  schema,
  action,
  children,
  loading = false,
  shouldReset = true,
  buttonElement = (
    <Button type="button">
      <PlusCircle className="w-4 h-4 mr-2" />
      <span>Add new</span>
    </Button>
  ),
}) {
  const [isOpen, setIsOpen] = useState(false);
  const [showConfirmation, setShowConfirmation] = useState(false);

  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
    mode: "onChange",
  });

  const { formState, reset, handleSubmit } = methods;

  const isFormDirty = useMemo(() => formState.isDirty, [formState.isDirty]);

  const resetAndCloseDialog = useCallback(() => {
    reset();
    setIsOpen(false);
  }, [reset]);

  const handleCancel = useCallback(() => {
    if (isFormDirty) setShowConfirmation(true);
    else resetAndCloseDialog();
  }, [isFormDirty, resetAndCloseDialog]);

  const handleConfirmation = useCallback(
    (confirmed) => {
      if (confirmed) resetAndCloseDialog();
      setShowConfirmation(false);
    },
    [resetAndCloseDialog]
  );

  useEffect(() => {
    if (state) reset(state);
  }, [state, reset]);

  const handleSave = useCallback(
    async (data) => {
      await action(data);
      if (formState.isValid && shouldReset) reset();
      setIsOpen(false);
    },
    [action, formState.isValid, reset, shouldReset]
  );

  return (
    <>
      {/* Main Dialog */}
      <Dialog
        open={isOpen}
        onOpenChange={(open) => (!open ? handleCancel() : setIsOpen(open))}
      >
        <DialogTrigger asChild>{buttonElement}</DialogTrigger>

        <DialogContent className="sm:max-w-lg overflow-hidden rounded-xl p-0 bg-card border border-border">
          {loading ? (
            <SubmitLoading />
          ) : (
            <FormProvider {...methods}>
              <form
                onSubmit={handleSubmit(handleSave)}
                className="flex flex-col max-h-[70vh]"
              >
                {/* Header */}
                <div className="border-b border-border p-4 ">
                  <DialogTitle className="text-lg font-semibold text-center text-foreground">
                    {title}
                  </DialogTitle>
                  <DialogDescription className="text-sm text-muted-foreground text-center">
                    Submit button will activate when all required fields are
                    filled.
                  </DialogDescription>
                </div>

                {/* Scrollable Form Content */}
                <ScrollArea className="flex-1">
                  <div className="space-y-4 p-4">{children}</div>
                </ScrollArea>

                {/* Footer */}
                <div className="border-t border-border px-6 py-4 flex justify-end ">
                  <SubmitButton
                    text="Save Changes"
                    isLoading={loading}
                    disabled={!formState.isValid || !formState.isDirty}
                  />
                </div>
              </form>
            </FormProvider>
          )}
        </DialogContent>
      </Dialog>

      {/* Confirmation Dialog */}
      <Dialog open={showConfirmation} onOpenChange={setShowConfirmation}>
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
              variant="secondary"
              className="w-32"
              onClick={() => handleConfirmation(false)}
            >
              Cancel
            </Button>
            <Button
              variant="destructive"
              className="w-32"
              onClick={() => handleConfirmation(true)}
            >
              Discard
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </>
  );
}
