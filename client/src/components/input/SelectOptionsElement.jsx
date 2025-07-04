import { Skeleton } from "@/components/ui/Skeleton";
import { useSelectOptions } from "@/hooks/useSelectOptions";
import { Controller, useFormContext } from "react-hook-form";

const SelectOptionsElement = ({
  name,
  label,
  data = "category",
  disabled = false,
  rules = { required: true },
  placeholder = "Select an option",
}) => {
  const { control } = useFormContext();
  const { data: options = [], isLoading } = useSelectOptions(data);

  if (isLoading) return <Skeleton className="h-9 rounded-md" />;

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label htmlFor={name} className="label">
              {label}
            </label>
          )}

          <select
            id={name}
            {...field}
            disabled={disabled}
            className={`input bg-background text-foreground border-input focus:ring focus:ring-ring disabled:bg-muted disabled:text-muted-foreground ${
              fieldState.error ? "input-error" : ""
            }`}
          >
            <option value="">{placeholder}</option>
            {options.map((option) => (
              <option
                key={option.id || option.value || option}
                value={option.id || option.value || option}
              >
                {option.name || option.title || option.fullname || option}
              </option>
            ))}
          </select>

          {fieldState.error && (
            <p className="error-message">{fieldState.error.message}</p>
          )}
        </div>
      )}
    />
  );
};

export { SelectOptionsElement };
