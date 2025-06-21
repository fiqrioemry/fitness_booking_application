export const registerState = {
  otp: "",
  email: "",
  password: "",
  fullname: "",
};

export const sendOTPState = {
  email: "",
};

export const verifyOTPState = {
  otp: "",
  email: "",
};
export const getLoginState = (rememberMe = false) => ({
  email: "",
  password: "",
  rememberMe,
});

export const profileState = {
  bio: "",
  phone: "",
  gender: "",
  birthday: "",
  fullname: "",
};

export const checkoutState = {
  verificationCode: "",
};

export const classState = {
  title: "",
  duration: 0,
  additional: [],
  typeId: "",
  levelId: "",
  locationId: "",
  categoryId: "",
  isActive: true,
  description: "",
  subcategoryId: "",
  image: undefined,
  images: undefined,
};

export const updateClassState = {
  title: "",
  duration: 0,
  typeId: "",
  levelId: "",
  additional: [],
  locationId: "",
  categoryId: "",
  description: "",
  subcategoryId: "",
  image: undefined,
};

export const optionState = {
  name: "",
};

export const subcategoryState = {
  name: "",
  categoryId: "",
};

export const locationState = {
  name: "",
  address: "",
  geoLocation: "",
};

export const reviewState = {
  rating: 0,
  comment: "",
};

export const markAttendanceState = {
  bookingId: "",
  status: "",
};

export const bookingState = {
  classScheduleId: "",
};

export const createScheduleTemplateState = {
  classId: "",
  capacity: 0,
  dayOfWeek: 0,
  startHour: 0,
  startMinute: 0,
  instructorId: "",
};

export const updateClassScheduleState = {
  capacity: 0,
  endTime: "",
  startTime: "",
};

export const midtransNotificationState = {
  order_id: "",
  payment_type: "",
  fraud_status: "",
  transaction_status: "",
};

export const paymentState = {
  packageId: "",
};

export const instructorState = {
  userId: "",
  experience: 0,
  specialties: "",
  certifications: "",
};

export const packageState = {
  name: "",
  price: 0,
  credit: 0,
  expired: 0,
  discount: 0,
  classIds: [],
  isActive: true,
  additional: [],
  description: "",
  image: undefined,
};

export const createVoucherState = {
  code: "",
  quota: 1,
  discount: 0,
  description: "",
  expiredAt: "",
  maxDiscount: null,
  discountType: "fixed",
  isReusable: true,
};
export const notificationState = {
  title: "",
  message: "",
  typeCode: "",
};

export const openClassState = {
  zoomLink: "",
  verificationCode: "",
};

export const genderOptions = [
  { value: "male", label: "Male" },
  { value: "female", label: "Female" },
];

export const typeCode = [
  { label: "Promo Offer", value: "promo_offer" },
  { label: "System Message", value: "system_message" },
  { label: "Class Reminder", value: "class_reminder" },
];

export const paymentStatusOptions = [
  { value: "all", label: "All" },
  { value: "success", label: "Success" },
  { value: "pending", label: "Pending" },
  { value: "failed", label: "Failed" },
];

export const statusOptions = [
  { value: "all", label: "All" },
  { value: "active", label: "active" },
  { value: "inactive", label: "inactive" },
];

export const roleOptions = [
  { value: "all", label: "All" },
  { value: "admin", label: "admin" },
  { value: "customer", label: "customer" },
  { value: "instructor", label: "instructor" },
];

export const revenueRangeOptions = [
  { value: "daily", label: "Daily" },
  { value: "monthly", label: "monthly" },
  { value: "yearly", label: "yearly" },
];
export const operationMinutes = [0, 15, 30, 45];

export const operationHours = [8, 9, 10, 11, 12, 13, 14, 15, 16, 17];

export const homeTitle = "Home - high-intensity workouts with sweat up";

export const aboutTitle =
  "About Us – Empowering Your Wellness Journey with Sweat Up";

export const scheduleTitle =
  "Discover and book fitness classes that fit your lifestyle. Explore real-time schedules with flexible times, expert instructors, and a variety of wellness programs at FitBook Studio.";

export const classesTitle =
  "Classes - Discover personalized sessions tailored for your needs, from beginner to advanced levels";

export const packagesTitle =
  "Packages - Find the right plan that matches your fitness goals and schedule";
