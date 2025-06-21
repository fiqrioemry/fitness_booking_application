import { Controller, useFormContext } from "react-hook-form";

const InputNumberElement = ({
  name,
  label,
  placeholder = "",
  disabled = false,
  rules = { required: true },
  max,
  min,
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleKeyDown = (e) => {
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
        };

        const handleChange = (e) => {
          const value = e.target.value;
          if (/^\d{0,15}$/.test(value)) {
            field.onChange(value === "" ? null : Number(value));
          }
        };

        return (
          <div className="space-y-1">
            {label && (
              <label htmlFor={name} className="label">
                {label}
              </label>
            )}
            <input
              id={name}
              type="text"
              inputMode="numeric"
              placeholder={placeholder}
              disabled={disabled}
              min={min}
              max={max}
              onKeyDown={handleKeyDown}
              onChange={handleChange}
              value={field.value ?? ""}
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

export { InputNumberElement };
