import { useEffect, useState } from "react";
import { Controller, useFormContext } from "react-hook-form";

export const InputDateElement = ({
  name,
  label,
  rules = { required: true },
  mode = "past",
  ageLimit = 80,
}) => {
  const { control, setValue } = useFormContext();
  const [day, setDay] = useState("");
  const [month, setMonth] = useState("");
  const [year, setYear] = useState("");
  const [hasInteracted, setHasInteracted] = useState(false);

  const today = new Date();
  const startDay = today.getDate();
  const startMonth = today.getMonth() + 1;
  const startYearVal = today.getFullYear();

  const endYear = mode === "future" ? startYearVal + 5 : startYearVal;
  const startYear = mode === "future" ? startYearVal : startYearVal - ageLimit;

  const getDaysInMonth = (month, year) => {
    if (!month || !year) return 31;
    return new Date(year, month, 0).getDate();
  };

  useEffect(() => {
    if (!day && !month && !year) {
      const initialValue = control._formValues?.[name];
      if (initialValue) {
        const [y, m, d] = initialValue.split("-").map(Number);
        setYear(y);
        setMonth(m);
        setDay(d);
      } else {
        setDay(startDay);
        setMonth(startMonth);
        setYear(startYearVal);
      }
    }
  }, [
    day,
    month,
    year,
    control._formValues,
    name,
    startDay,
    startMonth,
    startYearVal,
  ]);

  useEffect(() => {
    const isValid =
      Number(day) >= 1 &&
      Number(month) >= 1 &&
      Number(month) <= 12 &&
      Number(year) > 1900;

    if (isValid) {
      const formatted = `${year}-${String(month).padStart(2, "0")}-${String(
        day
      ).padStart(2, "0")}`;
      setValue(name, formatted, {
        shouldValidate: true,
        shouldDirty: true,
      });
    } else {
      setValue(name, "", {
        shouldValidate: true,
        shouldDirty: true,
      });
    }
  }, [day, month, year, name, setValue]);

  const daysInMonth = getDaysInMonth(month, year);
  const isSameMonthAndYear =
    Number(month) === startMonth && Number(year) === startYearVal;

  const minValidDay = mode === "future" && isSameMonthAndYear ? startDay : 1;
  const maxValidDay =
    mode === "past" && isSameMonthAndYear && hasInteracted
      ? startDay
      : daysInMonth;

  const days = Array.from(
    { length: maxValidDay - minValidDay + 1 },
    (_, i) => i + minValidDay
  );

  const months = [
    "Januari",
    "Februari",
    "Maret",
    "April",
    "Mei",
    "Juni",
    "Juli",
    "Agustus",
    "September",
    "Oktober",
    "November",
    "Desember",
  ];

  const years = Array.from(
    { length: endYear - startYear + 1 },
    (_, i) => startYear + i
  );

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ fieldState }) => (
        <div className="space-y-1">
          {label && (
            <label htmlFor={name} className="label">
              {label}
            </label>
          )}
          <div className="flex gap-2">
            {/* Day */}
            <select
              value={day}
              onChange={(e) => setDay(Number(e.target.value))}
              className={`input ${fieldState.error ? "input-error" : ""}`}
            >
              <option value="">Tanggal</option>
              {days.map((d) => (
                <option key={d} value={d}>
                  {d}
                </option>
              ))}
            </select>

            {/* Month */}
            <select
              value={month}
              onChange={(e) => {
                setMonth(Number(e.target.value));
                setHasInteracted(true);
              }}
              className={`input ${fieldState.error ? "input-error" : ""}`}
            >
              <option value="">Bulan</option>
              {months.map((m, i) => {
                const monthVal = i + 1;
                const isDisabledFuture =
                  mode === "future" &&
                  Number(year) === startYearVal &&
                  monthVal < startMonth;
                const isDisabledPast =
                  mode === "past" &&
                  Number(year) === startYearVal &&
                  monthVal > startMonth;
                return (
                  <option
                    key={m}
                    value={monthVal}
                    disabled={isDisabledFuture || isDisabledPast}
                  >
                    {m}
                  </option>
                );
              })}
            </select>

            {/* Year */}
            <select
              value={year}
              onChange={(e) => {
                setYear(Number(e.target.value));
                setHasInteracted(true);
              }}
              className={`input ${fieldState.error ? "input-error" : ""}`}
            >
              <option value="">Tahun</option>
              {years.map((y) => (
                <option key={y} value={y}>
                  {y}
                </option>
              ))}
            </select>
          </div>

          {fieldState.error && (
            <p className="error-message">{fieldState.error.message}</p>
          )}
        </div>
      )}
    />
  );
};
