import { Switch } from "@/components/ui/Switch";
import { Controller, useFormContext } from "react-hook-form";

const SwitchElement = ({
  name,
  label,
  disabled = false,
  defaultChecked = false,
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      render={({ field }) => (
        <div className="flex items-center gap-3">
          <Switch
            disabled={disabled}
            checked={field.value ?? defaultChecked}
            onCheckedChange={(checked) => field.onChange(checked)}
          />
          {label && (
            <label className="text-sm font-medium cursor-pointer select-none">
              {label}
            </label>
          )}
        </div>
      )}
    />
  );
};

export { SwitchElement };
