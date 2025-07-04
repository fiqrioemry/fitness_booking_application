import {
  Command,
  CommandGroup,
  CommandItem,
  CommandList,
} from "@/components/ui/Command";
import { X } from "lucide-react";
import { useState } from "react";
import { Input } from "@/components/ui/Input";
import { Badge } from "@/components/ui/Badge";
import { Skeleton } from "@/components/ui/Skeleton";
import { useSelectOptions } from "@/hooks/useSelectOptions";
import { Controller, useFormContext } from "react-hook-form";

const MultiSelectElement = ({ name, label, data = "class", rules = {} }) => {
  const { control } = useFormContext();
  const [query, setQuery] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const { data: results, isLoading } = useSelectOptions(data);

  if (isLoading) return <Skeleton className="h-10 w-full rounded-md" />;

  const optionsRaw = Array.isArray(results) ? results : [];
  const options = optionsRaw.map((item) => ({
    label: item.title,
    value: item.id,
  }));

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const selectedValues = field.value || [];

        const filteredOptions = options.filter(
          (opt) =>
            typeof opt.label === "string" &&
            opt.label.toLowerCase().includes(query.toLowerCase()) &&
            !selectedValues.includes(opt.value)
        );

        const removeTag = (val) => {
          field.onChange(selectedValues.filter((v) => v !== val));
        };

        const addTag = (val) => {
          field.onChange([...selectedValues, val]);
          setQuery("");
        };

        return (
          <div className="space-y-2">
            {label && (
              <label
                htmlFor={name}
                className="block text-sm font-medium text-foreground"
              >
                {label}
              </label>
            )}

            <div className="flex flex-wrap gap-2 min-h-12 border border-input rounded-md p-2 bg-background">
              {selectedValues.length === 0 && (
                <span className="text-sm text-muted-foreground">
                  No selection
                </span>
              )}
              {selectedValues.map((val) => {
                const matched = options.find((opt) => opt.value === val);
                return (
                  <Badge
                    key={val}
                    variant="secondary"
                    className="flex items-center gap-1 pr-1"
                  >
                    {matched?.label || val}
                    <X
                      className="w-3 h-3 cursor-pointer ml-1"
                      onClick={() => removeTag(val)}
                    />
                  </Badge>
                );
              })}
            </div>

            <div className="relative">
              <Command className="border border-input  rounded-md w-full shadow-sm">
                <Input
                  placeholder="Search options..."
                  value={query}
                  onChange={(e) => {
                    setQuery(e.target.value);
                    setIsOpen(true);
                  }}
                  onFocus={() => setIsOpen(true)}
                  onBlur={() => setTimeout(() => setIsOpen(false), 150)}
                  className="border-none h-12 focus:ring-0 px-3 py-2"
                />
                {isOpen && (
                  <CommandList className="absolute z-10 w-full bg-popover border border-border rounded-md shadow-lg max-h-40 overflow-auto">
                    <CommandGroup>
                      {filteredOptions.map((opt) => (
                        <CommandItem
                          key={opt.value}
                          value={opt.value}
                          onSelect={() => addTag(opt.value)}
                        >
                          {opt.label}
                        </CommandItem>
                      ))}
                      {filteredOptions.length === 0 && (
                        <div className="text-muted-foreground text-sm px-3 py-2">
                          No options found
                        </div>
                      )}
                    </CommandGroup>
                  </CommandList>
                )}
              </Command>
            </div>

            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { MultiSelectElement };
