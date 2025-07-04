import { Controller, useFormContext } from "react-hook-form";

const InputTextElement = ({
  name,
  label,
  maxLength,
  type = "text",
  placeholder = "",
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
      render={({ field, fieldState }) => {
        const handleKeyDown = (e) => {
          if (isNumeric) {
            const allowedKeys = [
              "Backspace",
              "Tab",
              "Delete",
              "ArrowLeft",
              "ArrowRight",
            ];
            if (allowedKeys.includes(e.key)) return;
            if (!/^[0-9]$/.test(e.key)) {
              e.preventDefault();
            }
          }
        };

        return (
          <div className="space-y-1">
            {label && (
              <label
                htmlFor={name}
                className="block text-sm font-medium text-foreground"
              >
                {label}
              </label>
            )}

            <input
              id={name}
              type={type}
              {...field}
              onKeyDown={handleKeyDown}
              placeholder={placeholder}
              disabled={disabled}
              maxLength={maxLength}
              inputMode={isNumeric ? "numeric" : undefined}
              className={`input ${fieldState.error ? "input-error" : ""}`}
            />

            {fieldState.error && (
              <p className="error-message">{fieldState.error.message}</p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputTextElement };
