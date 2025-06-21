import { useState } from "react";
import { X } from "lucide-react";
import { Badge } from "@/components/ui/Badge";
import { Controller, useFormContext } from "react-hook-form";

const InputTagsElement = ({
  name,
  label,
  placeholder = "Enter a tag and press Enter",
  disabled = false,
  rules = { required: true },
}) => {
  const { control } = useFormContext();
  const [inputValue, setInputValue] = useState("");

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const addTag = () => {
          const value = inputValue.trim();
          if (value && !(field.value || []).includes(value)) {
            field.onChange([...(field.value || []), value]);
          }
          setInputValue("");
        };

        const removeTag = (tag) => {
          const updated = (field.value || []).filter((t) => t !== tag);
          field.onChange(updated);
        };

        const handleKeyDown = (e) => {
          if (e.key === "Enter") {
            e.preventDefault();
            addTag();
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

            <div className="flex flex-wrap gap-2 p-2 border border-border bg-background rounded-md transition focus-within:ring-1 focus-within:ring-primary">
              {(field.value || []).map((val) => (
                <Badge
                  key={val}
                  variant="secondary"
                  className="flex items-center gap-1 pr-1"
                >
                  {val}
                  <X
                    className="w-3 h-3 cursor-pointer ml-1"
                    onClick={() => removeTag(val)}
                  />
                </Badge>
              ))}

              <input
                id={name}
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                onKeyDown={handleKeyDown}
                placeholder={placeholder}
                disabled={disabled}
                className="flex-1 resize-none h-10 bg-transparent text-sm outline-none placeholder:text-muted-foreground"
              />
            </div>

            {fieldState.error && (
              <p className="text-destructive text-xs mt-1">
                {fieldState.error.message}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { InputTagsElement };
