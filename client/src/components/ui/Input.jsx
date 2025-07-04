import * as React from "react";

const Input = ({
  label,
  name,
  error,
  register,
  placeholder,
  className = "",
  type = "text",
  ...rest
}) => {
  return (
    <div className="space-y-1 w-full">
      {label && (
        <label htmlFor={name} className="label">
          {label}
        </label>
      )}
      <input
        id={name}
        name={name}
        type={type}
        placeholder={placeholder}
        {...(register ? register(name) : {})}
        {...rest}
        className={`input ${error ? "input-error" : ""} ${className}`}
      />
      {error && <p className="error-message">{error.message}</p>}
    </div>
  );
};

export { Input };
