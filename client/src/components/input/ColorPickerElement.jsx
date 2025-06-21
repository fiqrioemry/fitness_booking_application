import { useFormContext, Controller } from "react-hook-form";

const COLOR_OPTIONS = [
  "#f87171", // red
  "#fb923c", // orange
  "#facc15", // yellow
  "#4ade80", // green
  "#60a5fa", // blue
  "#a78bfa", // purple
  "#f472b6", // pink
  "#94a3b8", // gray
];

export const ColorPickerElement = ({ name = "colorCode", label = "Warna" }) => {
  const { control } = useFormContext();

  return (
    <Controller
      name={name}
      control={control}
      defaultValue={COLOR_OPTIONS[0]}
      render={({ field }) => (
        <div className="space-y-1">
          {label && (
            <label className="block text-sm font-medium text-gray-700">
              {label}
            </label>
          )}
          <div className="flex flex-wrap gap-3 mt-1">
            {COLOR_OPTIONS.map((color) => (
              <label
                key={color}
                className="cursor-pointer w-8 h-8 rounded-full border-2"
                style={{
                  backgroundColor: color,
                  borderColor: field.value === color ? "#000" : "transparent",
                }}
              >
                <input
                  type="radio"
                  name={name}
                  value={color}
                  checked={field.value === color}
                  onChange={() => field.onChange(color)}
                  className="sr-only"
                />
              </label>
            ))}
          </div>
        </div>
      )}
    />
  );
};
