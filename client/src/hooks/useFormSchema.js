import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

export function useFormSchema({ state, schema, action, shouldReset = false }) {
  const methods = useForm({
    resolver: zodResolver(schema),
    defaultValues: state,
    mode: "onChange",
  });

  const { reset } = methods;

  const handleSubmit = methods.handleSubmit(async (data) => {
    await action(data);
    if (shouldReset) reset();
  });

  return { methods, handleSubmit };
}
