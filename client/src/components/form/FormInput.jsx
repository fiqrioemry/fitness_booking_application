import { SubmitButton } from "./SubmitButton";
import { FormProvider } from "react-hook-form";
import { useFormSchema } from "@/hooks/useFormSchema";

export const FormInput = ({
  action,
  state,
  schema,
  text,
  className,
  isLoading,
  children,
  shouldReset,
}) => {
  const { methods, handleSubmit } = useFormSchema({
    state,
    schema,
    action,
    shouldReset,
  });

  return (
    <FormProvider {...methods}>
      <form onSubmit={handleSubmit} className="grid-cols-2 space-y-2 space-y-4">
        {children}
        <SubmitButton
          text={text}
          className={className}
          isLoading={isLoading}
          disabled={!methods.formState.isValid}
        />
      </form>
    </FormProvider>
  );
};
