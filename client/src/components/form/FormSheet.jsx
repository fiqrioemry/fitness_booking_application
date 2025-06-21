import {
  Sheet,
  SheetHeader,
  SheetTitle,
  SheetContent,
  SheetDescription,
} from "@/components/ui/Sheet";
import { useEffect, useCallback } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, FormProvider } from "react-hook-form";
import { ScrollArea } from "@/components/ui/ScrollArea";
import { SubmitButton } from "@/components/form/SubmitButton";
import { SubmitLoading } from "@/components/ui/SubmitLoading";

export function FormSheet({
  title,
  state,
  schema,
  action,
  children,
  open,
  setOpen,
  loading = false,
  shouldReset = true,
}) {
  const methods = useForm({
    defaultValues: state,
    resolver: zodResolver(schema),
    mode: "onChange",
  });

  const { formState, reset, handleSubmit } = methods;

  const resetAndCloseSheet = useCallback(() => {
    reset();
    setOpen(false);
  }, [reset, setOpen]);

  const handleCancel = useCallback(() => {
    resetAndCloseSheet();
  }, [resetAndCloseSheet]);

  useEffect(() => {
    if (state) reset(state);
  }, [state, reset]);

  const handleSave = useCallback(
    async (data) => {
      await action(data);
      if (formState.isValid && shouldReset) reset();
      setOpen(false);
    },
    [action, formState.isValid, reset, shouldReset, setOpen]
  );

  return (
    <Sheet
      open={open}
      onOpenChange={(val) => (!val ? handleCancel() : setOpen(val))}
    >
      <SheetContent className="w-full sm:max-w-lg flex flex-col p-0">
        {loading ? (
          <SubmitLoading />
        ) : (
          <FormProvider {...methods}>
            <form
              onSubmit={handleSubmit(handleSave)}
              className="flex flex-col h-full"
            >
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

              <ScrollArea className="flex-1 py-4">
                <div className="space-y-4 px-4">{children}</div>
              </ScrollArea>

              <div className="border-t p-4 gap-4 flex justify-end">
                <SubmitButton
                  text="Save Changes"
                  isLoading={loading}
                  disabled={!formState.isValid}
                />
              </div>
            </form>
          </FormProvider>
        )}
      </SheetContent>
    </Sheet>
  );
}
