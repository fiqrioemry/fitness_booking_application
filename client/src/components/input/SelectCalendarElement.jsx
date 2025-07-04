import { cn } from "@/lib/utils";
import { format } from "date-fns";
import { Button } from "@/components/ui/Button";
import { Calendar } from "@/components/ui/Calendar";
import { useEffect, useRef, useState } from "react";
import { useFormContext, Controller } from "react-hook-form";

export const SelectCalendarElement = ({
  name,
  label,
  rules = { required: true },
  mode = "future",
  ageLimit = 80,
}) => {
  const { control } = useFormContext();
  const [open, setOpen] = useState(false);
  const dropdownRef = useRef(null);

  const today = new Date();
  const minDate =
    mode === "past" ? new Date(today.getFullYear() - ageLimit, 0, 1) : today;
  const maxDate =
    mode === "past" ? today : new Date(today.getFullYear() + 5, 11, 31);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field: { value, onChange }, fieldState }) => {
        const selectedDate = value ? new Date(value) : undefined;

        return (
          <div className="space-y-1 relative" ref={dropdownRef}>
            {label && <label className="label">{label}</label>}

            <Button
              variant="outline"
              type="button"
              className={cn(
                "input text-left font-normal px-3 py-2 w-full",
                !value && "text-muted-foreground"
              )}
              onClick={() => setOpen(!open)}
            >
              {value ? format(new Date(value), "PPP") : "Pilih tanggal"}
            </Button>

            {open && (
              <div className="absolute z-50 mt-2 rounded-xl border bg-background shadow-lg p-2">
                <Calendar
                  mode="single"
                  selected={selectedDate}
                  onSelect={(date) => {
                    if (date) {
                      onChange(format(date, "yyyy-MM-dd"));
                      setOpen(false);
                    }
                  }}
                  initialFocus
                  fromDate={minDate}
                  toDate={maxDate}
                />
              </div>
            )}

            {fieldState.error && (
              <p className="error-message">{fieldState.error.message}</p>
            )}
          </div>
        );
      }}
    />
  );
};
