import { Controller, useFormContext } from "react-hook-form";

const SelectElement = ({
  name,
  label,
  placeholder = "Select an option",
  options = [],
  disabled = false,
  rules = { required: true },
  isNumeric = false,
}) => {
  const { control } = useFormContext();

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
            value={field.value ?? ""}
            onChange={(e) => {
              const value = e.target.value;
              field.onChange(isNumeric ? Number(value) : value);
            }}
            disabled={disabled}
            className={`input bg-background text-foreground border border-input focus:ring focus:ring-ring disabled:bg-muted disabled:text-muted-foreground ${
              fieldState.error ? "input-error" : ""
            }`}
          >
            <option value="">{placeholder}</option>
            {options.map((option) => (
              <option
                key={option.id || option.value || option}
                value={option.id || option.value || option}
              >
                {option.name || option.label || option}
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

export { SelectElement };
