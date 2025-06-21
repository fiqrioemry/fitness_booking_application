// src/lib/utils.js
import {
  format,
  parse,
  getDay,
  formatDuration,
  intervalToDuration,
} from "date-fns";
import { clsx } from "clsx";
import id from "date-fns/locale/id";
import { twMerge } from "tailwind-merge";
import { dateFnsLocalizer } from "react-big-calendar";

export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

export const formatRupiah = (number) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(number);
};

export const formatDateTime = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleString("id-ID", {
    year: "numeric",
    month: "short",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
  });
};

export const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString("id-ID", {
    year: "numeric",
    month: "short",
    day: "2-digit",
  });
};

export function formatTime(hour, minute) {
  const h = String(hour).padStart(2, "0");
  const m = String(minute).padStart(2, "0");
  return `${h}:${m}`;
}
export const formatHour = (hour, minute) => {
  const date = new Date();
  date.setHours(hour);
  date.setMinutes(minute);
  return date.toLocaleString("id-ID", {
    hour: "2-digit",
    minute: "2-digit",
  });
};

export const localizer = dateFnsLocalizer({
  format,
  parse,
  startOfWeek: () => new Date(),
  getDay,
  locales: { id },
});

export const getTimeLeft = (startTime) => {
  const seconds = (startTime - new Date()) / 1000;
  if (seconds > 0) {
    const duration = intervalToDuration({ start: 0, end: seconds * 1000 });
    return formatDuration(duration, {
      format: ["days", "hours", "minutes", "seconds"],
    });
  }
  return "Ongoing or passed";
};

export const isAttendanceWindow = (startTime) => {
  const now = new Date();
  const startWindow = new Date(startTime.getTime() - 15 * 60000);
  const endWindow = new Date(startTime.getTime() + 30 * 60000);
  return now >= startWindow && now <= endWindow;
};

export const buildDateTime = (dateStr, hour, minute) => {
  if (!dateStr || hour === undefined || minute === undefined) return null;

  const [year, month, day] = dateStr.split("-").map(Number);
  return new Date(year, month - 1, day, hour, minute, 0);
};

export const truncateText = (text, maxLength) => {
  return text.length > maxLength ? `${text.slice(0, maxLength)}...` : text;
};

export const buildFormData = (data) => {
  const formData = new FormData();

  Object.entries(data).forEach(([key, value]) => {
    if (value === undefined || value === null) return;

    if (Array.isArray(value)) {
      if (value.length > 0 && value[0] instanceof File) {
        value.forEach((file) => {
          formData.append(key, file);
        });
      } else {
        value.forEach((item) => {
          formData.append(`${key}`, item);
        });
      }
    } else if (value instanceof File) {
      formData.append(key, value);
    } else {
      formData.append(key, value);
    }
  });

  return formData;
};
