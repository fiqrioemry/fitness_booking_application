const dayOptions = [
  { label: "Sun", value: 0 },
  { label: "Mon", value: 1 },
  { label: "Tue", value: 2 },
  { label: "Wed", value: 3 },
  { label: "Thu", value: 4 },
  { label: "Fri", value: 5 },
  { label: "Sat", value: 6 },
];
import { FormLabel } from "@/components/ui/Form";
import { useFormContext, Controller } from "react-hook-form";

export const DaySelectorElement = ({ name, label }) => {
  const { control } = useFormContext();

  return (
    <div className="space-y-1">
      {label && <FormLabel>{label}</FormLabel>}

      <Controller
        control={control}
        name={name}
        render={({ field, fieldState }) => (
          <>
            <div className="flex flex-wrap gap-2">
              {dayOptions.map((day) => (
                <label
                  key={day.value}
                  className={`px-3 py-1 border rounded-full cursor-pointer text-sm ${
                    field.value?.includes(day.value)
                      ? "bg-green-500 text-white"
                      : "bg-gray-100 text-gray-700"
                  }`}
                >
                  <input
                    type="checkbox"
                    value={day.value}
                    checked={field.value?.includes(day.value)}
                    onChange={(e) => {
                      const checked = e.target.checked;
                      if (checked) {
                        field.onChange([...field.value, day.value]);
                      } else {
                        field.onChange(
                          field.value.filter((v) => v !== day.value)
                        );
                      }
                    }}
                    className="hidden"
                  />
                  {day.label}
                </label>
              ))}
            </div>
            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message}
              </p>
            )}
          </>
        )}
      />
    </div>
  );
};
