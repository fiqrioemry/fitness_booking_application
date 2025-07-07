import { useSearchParams } from "react-router-dom";
import { Skeleton } from "@/components/ui/Skeleton";
import { useSelectOptions } from "@/hooks/useSelectOptions";

export const SearchFilterSelection = ({ paramKey, label, data }) => {
  const { data: options, isLoading, isError } = useSelectOptions(data);
  const [searchParams, setSearchParams] = useSearchParams();
  const selected = searchParams.get(paramKey) || "";

  const handleChange = (e) => {
    const value = e.target.value;
    if (value) {
      searchParams.set(paramKey, value);
    } else {
      searchParams.delete(paramKey);
    }
    setSearchParams(searchParams);
  };

  if (isLoading) return <Skeleton className="h-9 rounded-md" />;
  if (isError)
    return (
      <div className="text-sm text-muted-foreground">Error loading options</div>
    );

  return (
    <div className="space-y-1">
      {label && (
        <label htmlFor={paramKey} className="label">
          {label}
        </label>
      )}
      <select
        id={paramKey}
        value={selected}
        onChange={handleChange}
        className="input bg-background text-foreground border border-input focus:ring focus:ring-ring disabled:bg-muted disabled:text-muted-foreground"
      >
        <option value="">All</option>
        {options && Array.isArray(options) && options.length > 0 ? (
          options.map((opt, idx) => (
            <option key={opt?.id || `option-${idx}`} value={opt?.id || ""}>
              {opt?.name || "Unknown"}
            </option>
          ))
        ) : (
          <option disabled>No options available</option>
        )}
      </select>
    </div>
  );
};
