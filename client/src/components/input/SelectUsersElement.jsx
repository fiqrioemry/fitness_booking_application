import { useUsersQuery } from "@/hooks/useUsers";
import { useEffect, useRef, useState } from "react";
import { Controller, useFormContext } from "react-hook-form";

export const SelectUsersElement = ({
  name,
  label = "Select User",
  rules = { required: true },
  placeholder = "Search user...",
  role = "customer",
  disabled = false,
}) => {
  const { control } = useFormContext();
  const [query, setQuery] = useState("");
  const [open, setOpen] = useState(false);
  const inputRef = useRef(null);

  const { data, isLoading } = useUsersQuery({
    q: query,
    role,
    limit: 20,
    page: 1,
  });

  const users = data?.data || [];

  // Auto-close dropdown if click outside
  useEffect(() => {
    const handleClickOutside = (e) => {
      if (
        inputRef.current &&
        !inputRef.current.parentElement?.contains(e.target)
      ) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const selectedUser = users.find((u) => u.id === field.value);
        const selectedName = selectedUser?.fullname || "";

        return (
          <div className="space-y-1 relative">
            {label && (
              <label htmlFor={name} className="label">
                {label}
              </label>
            )}

            <input
              id={name}
              ref={inputRef}
              type="text"
              className={`input ${
                disabled ? "bg-red-500 text-muted-foreground" : ""
              }`}
              placeholder={placeholder}
              value={open ? query : selectedName}
              disabled={disabled}
              onFocus={() => {}}
              onChange={(e) => {
                setQuery(e.target.value);
                setOpen(true);
              }}
            />

            {open && users.length > 0 && (
              <ul className="absolute z-10 w-full bg-background border border-gray-300 mt-1 rounded-md max-h-60 overflow-y-auto shadow-md">
                {users.map((user) => (
                  <li
                    key={user.id}
                    onClick={() => {
                      field.onChange(user.id);
                      setQuery(user.fullname);
                      setOpen(false);
                    }}
                    className="px-4 py-2 cursor-pointer hover:bg-muted"
                  >
                    {user.fullname} ({user.email})
                  </li>
                ))}
              </ul>
            )}

            {open && !isLoading && users.length === 0 && (
              <div className="absolute z-10 w-full bg-background border mt-1 p-2 text-sm text-muted-foreground rounded-md shadow">
                No users found.
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
