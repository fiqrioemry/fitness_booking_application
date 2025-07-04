import { Controller, useFormContext } from "react-hook-form";

const InputTextareaElement = ({
  name,
  label,
  rows = 4,
  maxLength,
  placeholder = "",
  disabled = false,
  rules = { required: true },
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
          <textarea
            id={name}
            {...field}
            placeholder={placeholder}
            disabled={disabled}
            rows={rows}
            maxLength={maxLength}
            className={`input resize-none ${
              fieldState.error ? "input-error" : ""
            }`}
          />
          {fieldState.error && (
            <p className="error-message">{fieldState.error.message}</p>
          )}
        </div>
      )}
    />
  );
};

export { InputTextareaElement };
