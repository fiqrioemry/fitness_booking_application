import { Star } from "lucide-react";
import { Controller, useFormContext } from "react-hook-form";

const InputRatingElement = ({
  name,
  label,
  rules = { required: true },
  maxRating = 5,
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
            <label className="block text-sm font-medium text-foreground">
              {label}
            </label>
          )}

          <div className="flex gap-1">
            {[...Array(maxRating)].map((_, i) => {
              const value = i + 1;
              return (
                <Star
                  key={value}
                  size={24}
                  className={`cursor-pointer transition ${
                    value <= field.value
                      ? "fill-yellow-400 stroke-yellow-400"
                      : "stroke-muted"
                  }`}
                  onClick={() => field.onChange(value)}
                />
              );
            })}
          </div>

          {fieldState.error && (
            <p className="error-message text-sm text-red-500">
              {fieldState.error.message}
            </p>
          )}
        </div>
      )}
    />
  );
};

export { InputRatingElement };
