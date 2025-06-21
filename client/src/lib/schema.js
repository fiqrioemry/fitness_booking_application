import { z } from "zod";

const imageItemSchema = z
  .any()
  .optional()
  .refine((file) => !file || file instanceof File || typeof file === "string", {
    message: "Input must be a file or a valid URL",
  })
  .refine(
    (file) =>
      !file || typeof file === "string" || file.type?.startsWith("image/"),
    { message: "File must be an image" }
  )
  .refine(
    (file) => !file || typeof file === "string" || file.size <= 2 * 1024 * 1024,
    { message: "Image size must be <= 2MB" }
  );

const requiredNumberField = (label, { min, max } = {}) =>
  z.preprocess(
    (val) => (val === "" || val === null ? undefined : val),
    z
      .number({
        required_error: `${label} is required`,
        invalid_type_error: `${label} must be a number`,
      })
      .refine(
        (val) =>
          (min === undefined || val >= min) &&
          (max === undefined || val <= max),
        {
          message:
            min !== undefined && max !== undefined
              ? `${label} must be between ${min} and ${max}`
              : min !== undefined
              ? `${label} must be at least ${min}`
              : max !== undefined
              ? `${label} must be at most ${max}`
              : `${label} must be valid`,
        }
      )
  );

export const sendOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
});

export const checkoutSchema = z.object({
  verificationCode: z.string().min(1, "verification code is required"),
});

export const verifyOTPSchema = z.object({
  email: z.string().email("Invalid email address"),
  otp: z.string().min(6, "OTP code must be at least 6 characters"),
});

export const registerSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  fullname: z.string().min(1, "Full name is required"),
});

export const loginSchema = z.object({
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  rememberMe: z.boolean().optional(),
});

// profiles
export const profileSchema = z.object({
  fullname: z.string().min(6, "Fullname must be at least 6 characters"),
  birthday: z.string().optional(),
  gender: z.string().optional(),
  phone: z.string().optional(),
  bio: z.string().optional(),
});

export const avatarSchema = z.object({
  avatar: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
});

export const packageSchema = z.object({
  name: z.string().min(6, "Package name must be at least 6 characters"),
  description: z.string().min(20, "Description must be at least 20 characters"),
  price: z.number().positive("Price must be greater than 0"),
  credit: z.number().positive("Credit must be greater than 0"),
  discount: z
    .number()
    .min(0, "Discount cannot be negative")
    .max(100, "Discount cannot exceed 100")
    .optional(),
  classIds: z
    .array(z.string().uuid("Invalid class ID"))
    .min(1, "At least one class must be selected"),
  expired: z.number().min(1, "Expiry duration is required"),
  additional: z.array(z.string()).optional(),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
  isActive: z.boolean().optional(),
});

export const optionSchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
});

export const subcategorySchema = z.object({
  name: z.string().min(2, "Name must be at least 2 characters"),
  categoryId: z.string().min(1, "Category is required"),
});

export const locationSchema = z.object({
  name: z.string().min(2, "Location name is required"),
  address: z.string().min(1, "Address is required"),
  geoLocation: z.string().min(1, "Geolocation is required"),
});

export const instructorSchema = z.object({
  userId: z.string().min(1, "User is required"),
  experience: z.number().min(0, "experience is required"),
  specialties: z.string().min(1, "User is required"),
  certifications: z.string().min(1, "User is required"),
});

export const midtransNotificationSchema = z.object({
  transaction_status: z.string(),
  order_id: z.string(),
  payment_type: z.string(),
  fraud_status: z.string(),
});

export const createPaymentSchema = z.object({
  packageId: z.string().min(1, "Package is required"),
});

export const markAttendanceSchema = z.object({
  bookingId: z.string().min(1, "Booking ID is required"),
  status: z.enum(["attended", "absent", "cancelled"]),
});

export const bookingSchema = z.object({
  classScheduleId: z.string().min(1, "Class Schedule is required"),
});

export const updateRecuringScheduleSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  dayOfWeeks: z.array(z.number().int().min(0).max(6), {
    required_error: "Days must be an array of valid weekdays (0-6)",
  }),
  startHour: z.number().min(0).max(23),
  startMinute: z.number().max(59),
  capacity: z.number().positive(),
  color: z.string().optional(),
  capacity: requiredNumberField("Capacity", { min: 1 }),
});

export const classSchema = z.object({
  title: z.string().min(6, "Title must be at least 6 characters"),
  duration: requiredNumberField("Duration", { min: 15, max: 180 }),
  description: z.string().min(1, "Description is required"),
  additional: z.array(z.string()).optional(),
  typeId: z.string().min(1, "Type is required"),
  levelId: z.string().min(1, "Level is required"),
  locationId: z.string().min(1, "Location is required"),
  categoryId: z.string().min(1, "Category is required"),
  subcategoryId: z.string().min(1, "Subcategory is required"),
  image: imageItemSchema.refine((val) => !!val, {
    message: "Image is required",
  }),
  images: z.array(imageItemSchema).optional(),
  isActive: z.boolean().optional(),
});

export const uploadGallerySchema = z.object({
  images: z.array(imageItemSchema).min(1, "Image is required"),
});

export const scheduleSchema = z
  .object({
    isRecurring: z.boolean().optional().default(false),
    classId: z.string().min(1, "Class is required"),
    instructorId: z.string().min(1, "Instructor is required"),
    capacity: requiredNumberField("Capacity", { min: 1 }),
    color: z.string().optional(),
    date: z
      .string()
      .optional()
      .refine((val) => !val || !isNaN(Date.parse(val)), {
        message: "Date must be a valid date",
      }),

    startHour: z.number().min(0).max(23),
    startMinute: z.number().max(59),
    dayOfWeeks: z.array(z.number().int().min(0).max(6)).optional().default([]),
    endDate: z
      .string()
      .optional()
      .refine((val) => !val || !isNaN(Date.parse(val)), {
        message: "End Date must be a valid date",
      }),
  })
  .superRefine((data, ctx) => {
    if (data.isRecurring) {
      if (!data.dayOfWeeks || data.dayOfWeeks.length === 0) {
        ctx.addIssue({
          path: ["dayOfWeeks"],
          code: "custom",
          message: "At least one day must be selected for recurring schedule",
        });
      }

      if (!data.endDate) {
        ctx.addIssue({
          path: ["endDate"],
          code: "custom",
          message: "End date is required for recurring schedule",
        });
      }
    } else {
      if (!data.date) {
        ctx.addIssue({
          path: ["date"],
          code: "custom",
          message: "Date is required for one-time schedule",
        });
      }
    }
  });

export const updateScheduleSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  color: z.string().optional(),
  date: z
    .string()
    .optional()
    .refine((val) => !val || !isNaN(Date.parse(val)), {
      message: "Date must be a valid date",
    }),
  startHour: z.number().min(0).max(23),
  startMinute: z.number().max(59),
  capacity: requiredNumberField("Capacity", { min: 1 }),
});

export const updateTemplateSchema = z.object({
  classId: z.string().min(1, "Class is required"),
  instructorId: z.string().min(1, "Instructor is required"),
  capacity: requiredNumberField("Capacity", { min: 1 }),
  startHour: z.number().min(0).max(23),
  startMinute: z.number().max(59),
  dayOfWeeks: z.array(z.number().int().min(0).max(6)).optional().default([]),
  endDate: z
    .string()
    .optional()
    .refine((val) => !val || !isNaN(Date.parse(val)), {
      message: "End Date must be a valid date",
    }),
});

export const createReviewSchema = z.object({
  rating: z
    .number({
      required_error: "Rating is required",
      invalid_type_error: "Rating must be a number",
    })
    .min(1, "Minimum rating is 1")
    .max(5, "Maximum rating is 5"),
  comment: z.string().min(10, "Comment must be at least 10 characters"),
});

export const createVoucherSchema = z
  .object({
    code: z.string().min(1, "Code is required"),
    description: z.string().min(1, "Description is required"),
    discountType: z.enum(["fixed", "percentage"], {
      required_error: "Discount type is required",
      invalid_type_error: "Please select a valid discount type",
    }),
    isReusable: z.boolean().optional(),
    discount: z
      .number({ invalid_type_error: "Discount must be a number" })
      .gt(0, "Discount must be greater than 0"),
    maxDiscount: z
      .number({ invalid_type_error: "Max discount must be a number" })
      .optional()
      .nullable(),
    quota: z
      .number({ invalid_type_error: "Quota must be a number" })
      .gt(0, "Quota must be greater than 0"),
    expiredAt: z.string().refine((val) => !isNaN(Date.parse(val)), {
      message: "Expired date must be valid (YYYY-MM-DD)",
    }),
  })
  .superRefine((val, ctx) => {
    if (val.discountType === "percentage" && val.discount > 100) {
      ctx.addIssue({
        code: "custom",
        path: ["discount"],
        message: "Percentage discount cannot exceed 100",
      });
    }
  });

export const notificationSchema = z.object({
  title: z.string().min(3, "Title is required"),
  message: z
    .string()
    .min(5, "Message is required")
    .max(200, "Maximum 200 characters allowed"),
  typeCode: z.enum(["system_message", "class_reminder", "promo_offer"]),
});

export const openClassSchema = z.object({
  verificationCode: z
    .string()
    .min(6, "Verification code min. 6 char")
    .max(6, "Verification code cannot exceed 6 char"),
  zoomLink: z.string().url("Invalid Zoom URL").optional(),
});
