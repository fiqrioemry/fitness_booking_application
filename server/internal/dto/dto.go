package dto

import (
	"mime/multipart"
)

// USER & AUTHENTICATION  ==============
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required,min=5"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type AuthMeResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}

type UpdateAvatarRequest struct {
	Avatar    *multipart.FileHeader `form:"avatar" binding:"required"`
	AvatarURL string                `form:"-"`
}

type UpdateUserDetailRequest struct {
	Fullname string `form:"fullname" binding:"required,min=5"`
	Birthday string `form:"birthday"`
	Gender   string `form:"gender"`
	Phone    string `form:"phone"`
	Bio      string `form:"bio"`
}

type UserQueryParam struct {
	Q     string `form:"q"`
	Role  string `form:"role"`
	Sort  string `form:"sort"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}

type UserListResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	JoinedAt string `json:"joinedAt"`
}

type UserDetailResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday,omitempty"`
	Bio       string `json:"bio"`
	LastLogin string `json:"lastLogin"`
	JoinedAt  string `json:"joinedAt"`
}

type UserStatsResponse struct {
	Total        int64 `json:"total"`
	Customers    int64 `json:"customers"`
	Instructors  int64 `json:"instructors"`
	Admins       int64 `json:"admins"`
	NewThisMonth int64 `json:"newThisMonth"`
}

// USER & AUTHENTICATION  =============

// CLASS  ======================

type ClassQueryParam struct {
	Q             string `form:"q"`
	Status        string `form:"status"`
	Sort          string `form:"sort"`
	Page          int    `form:"page"`
	Limit         int    `form:"limit"`
	TypeID        string `form:"typeId"`
	CategoryID    string `form:"categoryId"`
	SubcategoryID string `form:"subcategoryId"`
	LevelID       string `form:"levelId"`
	LocationID    string `form:"locationId"`
}

type CreateClassRequest struct {
	Title         string                  `form:"title" binding:"required"`
	Description   string                  `form:"description" binding:"required"`
	TypeID        string                  `form:"typeId" binding:"required"`
	LevelID       string                  `form:"levelId" binding:"required"`
	LocationID    string                  `form:"locationId" binding:"required"`
	CategoryID    string                  `form:"categoryId" binding:"required"`
	SubcategoryID string                  `form:"subcategoryId" binding:"required"`
	Duration      int                     `form:"duration" binding:"required,min=15"`
	Additional    []string                `form:"additional"`
	IsActive      bool                    `form:"isActive"`
	Image         *multipart.FileHeader   `form:"image" binding:"required"`
	ImageURL      string                  `form:"-"`
	Images        []*multipart.FileHeader `form:"images" binding:"omitempty"`
	ImageURLs     []string                `form:"-"`
}

type UpdateClassRequest struct {
	Title         string                `form:"title" binding:"required"`
	Duration      int                   `form:"duration" binding:"required,min=15"`
	Description   string                `form:"description" binding:"required"`
	IsActive      bool                  `form:"isActive"`
	Additional    []string              `form:"additional"`
	TypeID        string                `form:"typeId" binding:"required"`
	LevelID       string                `form:"levelId" binding:"required"`
	LocationID    string                `form:"locationId" binding:"required"`
	CategoryID    string                `form:"categoryId" binding:"required"`
	SubcategoryID string                `form:"subcategoryId" binding:"required"`
	Image         *multipart.FileHeader `form:"image"`
	ImageURL      string                `form:"-"`
}

type ClassDetailResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Image       string   `json:"image"`
	IsActive    bool     `json:"isActive"`
	Duration    int      `json:"duration"`
	Description string   `json:"description"`
	Additional  []string `json:"additional"`
	Type        string   `json:"type"`
	Level       string   `json:"level"`
	Location    string   `json:"location"`
	Category    string   `json:"category"`
	Subcategory string   `json:"subcategory"`
	Galleries   []string `json:"galleries"`
	CreatedAt   string   `json:"createdAt"`
}

type GalleryResponse struct {
	ID        string `json:"id"`
	ImageURL  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
}

type ClassResponse struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Image         string   `json:"image"`
	IsActive      bool     `json:"isActive"`
	Duration      int      `json:"duration"`
	Description   string   `json:"description"`
	Additional    []string `json:"additional"`
	TypeID        string   `json:"typeId"`
	LevelID       string   `json:"levelId"`
	LocationID    string   `json:"locationId"`
	Galleries     []string `json:"galleries"`
	CategoryID    string   `json:"categoryId"`
	SubcategoryID string   `json:"subcategoryId"`
	CreatedAt     string   `json:"createdAt"`
}

// CLASS  ========================

// CLASS CATEGORY ================
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CLASS CATEGORY ================

// CLASS SUBCATEGORY =============
type CreateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required,min=2"`
	CategoryID string `json:"categoryId" binding:"required"`
}

type UpdateSubcategoryRequest struct {
	Name       string `json:"name" binding:"required,min=2"`
	CategoryID string `json:"categoryId" binding:"required"`
}

type SubcategoryResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"categoryId"`
}

// CLASS SUBCATEGORY =============

// CLASS TYPE ====================
type CreateTypeRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateTypeRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type TypeResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CLASS TYPE ===================

// CLASS LEVEL ==================
type CreateLevelRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type UpdateLevelRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}

type LevelResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CLASS LEVEL ==================

// LOCATION =====================
type CreateLocationRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Address     string `json:"address" binding:"required"`
	GeoLocation string `json:"geoLocation" binding:"required"`
}

type UpdateLocationRequest struct {
	Name        string `json:"name" binding:"required,min=2"`
	Address     string `json:"address" binding:"required"`
	GeoLocation string `json:"geoLocation" binding:"required"`
}

type LocationResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	GeoLocation string `json:"geoLocation"`
}

// LOCATION =========================

// PACKAGE ==========================

type PackageQueryParam struct {
	Q      string `form:"q"`
	Status string `form:"status"`
	Sort   string `form:"sort"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}

type CreatePackageRequest struct {
	Name        string                `form:"name" binding:"required,min=6"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Credit      int                   `form:"credit" binding:"required,gt=0"`
	Expired     int                   `form:"expired" binding:"required,gt=0"`
	Discount    float64               `form:"discount"`
	Additional  []string              `form:"additional"`
	IsActive    bool                  `form:"isActive"`
	ClassIDs    []string              `form:"classIds" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL    string                `form:"-"`
}

type UpdatePackageRequest struct {
	Name        string                `form:"name" binding:"required,min=6"`
	Description string                `form:"description" binding:"required"`
	Price       float64               `form:"price" binding:"required,gt=0"`
	Credit      int                   `form:"credit" binding:"required,gt=0"`
	Expired     int                   `form:"expired" binding:"required,gt=0"`
	Discount    float64               `form:"discount"`
	Additional  []string              `form:"additional"`
	IsActive    bool                  `form:"isActive"`
	ClassIDs    []string              `form:"classIds" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL    string                `form:"-"`
}

type PackageListResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Credit      int                    `json:"credit"`
	Expired     int                    `json:"expired"`
	Image       string                 `json:"image"`
	Discount    float64                `json:"discount"`
	IsActive    bool                   `json:"isActive"`
	Additional  []string               `json:"additional"`
	Classes     []ClassSummaryResponse `json:"classes"`
}

type PackageDetailResponse struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Credit      int                    `json:"credit"`
	Expired     int                    `json:"expired"`
	Discount    float64                `form:"discount"`
	Image       string                 `json:"image"`
	IsActive    bool                   `json:"isActive"`
	Additional  []string               `json:"additional"`
	Classes     []ClassSummaryResponse `json:"classes"`
}

type ClassSummaryResponse struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Duration int    `json:"duration"`
}

// PACKAGE ==========================

// INSTRUCTOR ==========================
type CreateInstructorRequest struct {
	UserID         string `json:"userId" binding:"required,uuid"`
	Experience     int    `json:"experience"`
	Specialties    string `json:"specialties" binding:"required"`
	Certifications string `json:"certifications"`
}

type UpdateInstructorRequest struct {
	UserID         string `json:"userId" binding:"required,uuid"`
	Experience     int    `json:"experience"`
	Specialties    string `json:"specialties"`
	Certifications string `json:"certifications"`
}
type InstructorResponse struct {
	ID             string  `json:"id"`
	UserID         string  `json:"userId"`
	Fullname       string  `json:"fullname"`
	Avatar         string  `json:"avatar"`
	Experience     int     `json:"experience"`
	Specialties    string  `json:"specialties"`
	Certifications string  `json:"certifications"`
	Rating         float64 `json:"rating"`
	TotalClass     int     `json:"totalClass"`
}

// INSTRUCTOR ==========================

// PAYMENT ================================
type CreatePaymentRequest struct {
	PackageID   string  `json:"packageId" binding:"required"`
	VoucherCode *string `json:"voucherCode"`
}

type CreatePaymentResponse struct {
	PaymentID string `json:"paymentId"`
	SessionID string `json:"sessionId"`
	SnapURL   string `json:"snapUrl"`
}

type NotificationEvent struct {
	UserID  string `json:"userId"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type PaymentResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	PackageID     string  `json:"packageId"`
	PackageName   string  `json:"packageName"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	Price         float64 `json:"price"`
	PaidAt        string  `json:"paidAt"`
}

type PaymentQueryParam struct {
	Search    string `form:"q"`
	Status    string `form:"status"`
	Sort      string `form:"sort"`
	Page      int    `form:"page"`
	Limit     int    `form:"limit"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
}

type PaymentListResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	InvoiceNumber string  `json:"invoiceNumber"`
	Email         string  `json:"email"`
	Fullname      string  `json:"fullname"`
	PackageID     string  `json:"packageId"`
	PackageName   string  `json:"packageName"`
	Total         float64 `json:"total"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
	PaidAt        string  `json:"paidAt"`
}

type PaymentDetailResponse struct {
	ID              string  `json:"id"`
	UserID          string  `json:"userId"`
	InvoiceNumber   string  `json:"invoiceNumber"`
	Email           string  `json:"email"`
	Fullname        string  `json:"fullname"`
	PackageID       string  `json:"packageId"`
	PackageName     string  `json:"packageName"`
	BasePrice       float64 `json:"basePrice"`
	Tax             float64 `json:"tax"`
	VoucherCode     string  `json:"voucherCode"`
	VoucherDiscount float64 `json:"voucherDiscount"`
	Total           float64 `json:"total"`
	PaymentMethod   string  `json:"paymentMethod"`
	Status          string  `json:"status"`
	PaidAt          string  `json:"paidAt"`
}

// NOTIFICATIONS
type NotificationSettingResponse struct {
	TypeID  string `json:"typeId"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Channel string `json:"channel"`
	Enabled bool   `json:"enabled"`
}

type UpdateNotificationSettingRequest struct {
	TypeID  string `json:"typeId" binding:"required"`
	Channel string `json:"channel" binding:"required,oneof=email browser"`
	Enabled bool   `json:"enabled"`
}

type CreateNotificationRequest struct {
	UserID   string `json:"userId"`
	TypeCode string `json:"typeCode"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Channel  string `json:"channel"` // "email" / "browser"
}

type NotificationResponse struct {
	ID        string `json:"id"`
	TypeCode  string `json:"typeCode"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}

type SendNotificationRequest struct {
	TypeCode string `json:"typeCode" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

// CLASS-SCHEDULE ====================
type CreateScheduleRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	Date         string `json:"date" binding:"required"`
	StartHour    int    `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int    `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int    `json:"capacity" validate:"required,min=1"`
	Color        string `json:"color"`
}

type CreateRecurringScheduleRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	Capacity     int    `json:"capacity" binding:"required"`
	Color        string `json:"color"`
	Date         string `json:"date,omitempty"`
	StartHour    int    `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int    `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	DayOfWeeks   []int  `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	EndDate      string `json:"endDate,omitempty"`
}

type UpdateClassScheduleRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	Date         string `json:"date" binding:"required"`
	StartHour    int    `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int    `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int    `json:"capacity" validate:"required,min=1"`
	Color        string `json:"color"`
}

type ScheduleTemplateResponse struct {
	ID             string `json:"id"`
	ClassID        string `json:"classId"`
	ClassName      string `json:"className"`
	InstructorID   string `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	Instructor     string `json:"instructor"`
	DayOfWeeks     []int  `json:"dayOfWeeks"`
	StartHour      int    `json:"startHour"`
	StartMinute    int    `json:"startMinute"`
	Capacity       int    `json:"capacity"`
	IsActive       bool   `json:"isActive"`
	Frequency      string `json:"frequency"`
	EndDate        string `json:"endDate"`
	CreatedAt      string `json:"createdAt"`
}

type CreateScheduleTemplateRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	DayOfWeeks   []int  `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	StartHour    int    `json:"startHour" binding:"required,min=0,max=23"`
	StartMinute  int    `json:"startMinute" binding:"required,min=0,max=59"`
	Capacity     int    `json:"capacity" binding:"required,gt=0"`
	Color        string `json:"color" binding:"required"`
	EndDate      string `json:"endDate" binding:"required"`
}

type UpdateScheduleTemplateRequest struct {
	ClassID      string `json:"classId" binding:"required"`
	InstructorID string `json:"instructorId" binding:"required"`
	DayOfWeeks   []int  `json:"dayOfWeeks" binding:"required,dive,min=0,max=6"`
	StartHour    int    `json:"startHour" validate:"required,min=8,max=17"`
	StartMinute  int    `json:"startMinute" validate:"required,oneof=0 15 30 45"`
	Capacity     int    `json:"capacity" binding:"required,gt=0"`
	EndDate      string `json:"endDate" binding:"required"`
}

type ScheduleTemplateToggleRequest struct {
	IsActive bool `json:"isActive" binding:"required"`
}

type ClassScheduleResponse struct {
	ID             string `json:"id"`
	ClassID        string `json:"classId"`
	ClassName      string `json:"className"`
	ClassImage     string `json:"classImage"`
	InstructorID   string `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	Location       string `json:"location"`
	Date           string `json:"date"`
	StartHour      int    `json:"startHour"`
	StartMinute    int    `json:"startMinute"`
	Capacity       int    `json:"capacity"`
	BookedCount    int    `json:"bookedCount"`
	Duration       int    `json:"duration"`
	Color          string `json:"color"`
	IsBooked       bool   `json:"isBooked"`
}

type AttendanceWithUserResponse struct {
	Fullname   string `json:"fullname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Status     string `json:"status"`
	CheckedIn  bool   `json:"checkedIn"`
	CheckedOut bool   `json:"checkedOut"`
	CheckedAt  string `json:"checkedAt,omitempty"`
	VerifiedAt string `json:"verifiedAt,omitempty"`
}

type InstructorBrief struct {
	ID       string  `json:"id"`
	Fullname string  `json:"fullname"`
	Rating   float64 `json:"rating"`
}

type ClassScheduleDetailResponse struct {
	ClassScheduleResponse
	Packages []PackageListResponse `json:"packages"`
}

type ClassScheduleQueryParam struct {
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
}

// CLASS-SCHEDULE =====================

// BOOKINGS & ATTENDANCE ===========================

type BookingQueryParam struct {
	Status string `form:"status"`
	Sort   string `form:"sort"`
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1"`
}

type CreateBookingRequest struct {
	PackageID       string `json:"packageId" binding:"required,uuid"`
	ClassScheduleID string `json:"scheduleId" binding:"required,uuid"`
}

type BookingResponse struct {
	ID             string `json:"id"`
	BookingStatus  string `json:"bookingStatus"`
	ClassID        string `json:"classId"`
	ClassName      string `json:"className"`
	ClassImage     string `json:"classImage"`
	InstructorName string `json:"instructorName"`
	Duration       int    `json:"duration"`
	Date           string `json:"date"`
	StartHour      int    `json:"startHour"`
	StartMinute    int    `json:"startMinute"`
	Location       string `json:"location"`
	BookedAt       string `json:"bookedAt"`
	IsOpened       bool   `json:"isOpen"`
}

type BookingDetailResponse struct {
	ID               string `json:"id"`
	ScheduleID       string `json:"scheduleId"`
	ClassID          string `json:"classId"`
	ClassName        string `json:"className"`
	ClassImage       string `json:"classImage"`
	InstructorName   string `json:"instructorName"`
	Date             string `json:"date"`
	StartHour        int    `json:"startHour"`
	StartMinute      int    `json:"startMinute"`
	Duration         int    `json:"duration"`
	CheckedIn        bool   `json:"checkedIn"`
	CheckedOut       bool   `json:"checkedOut"`
	ZoomLink         string `json:"zoomLink"`
	AttendanceStatus string `json:"attendanceStatus"`
	IsReviewed       bool   `json:"isReviewed"`
	IsOpened         bool   `json:"isOpen"`
	CheckedAt        string `json:"checkedAt"`
	VerifiedAt       string `json:"verifiedAt,omitempty"`
}

type ValidateCheckoutRequest struct {
	VerificationCode string `json:"verificationCode" binding:"required"`
}

type MarkAttendanceRequest struct {
	BookingID string `json:"bookingId" binding:"required"`
	Status    string `json:"status" binding:"required,oneof=attended absent cancelled"`
}

type AttendanceResponse struct {
	ID             string `json:"id"`
	ScheduleID     string `json:"scheduleId"`
	ClassName      string `json:"className"`
	ClassImage     string `json:"classImage"`
	InstructorID   string `json:"instructorId"`
	InstructorName string `json:"instructorName"`
	Date           string `json:"date"`
	StartHour      int    `json:"startHour"`
	StartMinute    int    `json:"startMinute"`
	Status         string `json:"status"`
	CheckedAt      string `json:"checkedAt"`
	Reviewed       bool   `json:"reviewed"`
	Verified       bool   `json:"verified"`
}

type AttendanceDetailResponse struct {
	ID        string `json:"id"`
	Fullname  string `json:"fullname"`
	Avatar    string `json:"avatar"`
	CheckedAt string `json:"checkedAt"`
}

// ATTENDANCE ==========================

// REVIEWS ==========================
type CreateReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment" binding:"required,min=10"`
}

type ReviewResponse struct {
	ID        string `json:"id"`
	Avatar    string `json:"avatar"`
	UserName  string `json:"userName"`
	ClassName string `json:"className"`
	Rating    int    `json:"rating"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"createdAt"`
}

// ATTENDANCE ==========================

type UserPackageResponse struct {
	ID              string `json:"id"`
	PackageID       string `json:"packageId"`
	PackageName     string `json:"packageName"`
	RemainingCredit int    `json:"remainingCredit"`
	ExpiredAt       string `json:"expiredAt,omitempty"`
	ExpiredInDays   int    `json:"expiredInDays,omitempty"`
	PurchasedAt     string `json:"purchasedAt"`
}

type BookingListResponse struct {
	Bookings []BookingResponse `json:"bookings"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

// VOUCHER
type CreateVoucherRequest struct {
	Code         string   `json:"code" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	DiscountType string   `json:"discountType" binding:"required,oneof=fixed percentage"`
	Discount     float64  `json:"discount" binding:"required,gt=0"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota" binding:"required,gt=0"`
	IsReusable   bool     `json:"isReusable"`
	ExpiredAt    string   `json:"expiredAt" binding:"required"`
}

type UpdateVoucherRequest struct {
	Description  string   `json:"description" binding:"required"`
	DiscountType string   `json:"discountType" binding:"required,oneof=fixed percentage"`
	Discount     float64  `json:"discount" binding:"required,gt=0"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota" binding:"required,gt=0"`
	IsReusable   bool     `json:"isReusable"`
	ExpiredAt    string   `json:"expiredAt" binding:"required"`
}

type VoucherResponse struct {
	ID           string   `json:"id"`
	Code         string   `json:"code"`
	Description  string   `json:"description"`
	DiscountType string   `json:"discountType"`
	Discount     float64  `json:"discount"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota"`
	IsReusable   bool     `json:"isReusable"`
	ExpiredAt    string   `json:"expiredAt"`
	CreatedAt    string   `json:"createdAt"`
}

type ApplyVoucherRequest struct {
	UserID *string `json:"userId"`
	Code   string  `json:"code" binding:"required"`
	Total  float64 `json:"total" binding:"required"`
}

type ApplyVoucherResponse struct {
	Code          string   `json:"code"`
	DiscountType  string   `json:"discountType"`
	Discount      float64  `json:"discount"`
	MaxDiscount   *float64 `json:"maxDiscount,omitempty"`
	DiscountValue float64  `json:"discountValue"`
	FinalTotal    float64  `json:"finalTotal"`
}

type GoogleSignInRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

type DashboardSummaryResponse struct {
	TotalUsers         int     `json:"totalUsers"`
	TotalInstructors   int     `json:"totalInstructors"`
	TotalClasses       int     `json:"totalClasses"`
	TotalBookings      int     `json:"totalBookings"`
	TotalPayments      int     `json:"totalPayments"`
	TotalRevenue       float64 `json:"totalRevenue"`
	ActivePackages     int     `json:"activePackages"`
	TotalAttendance    int     `json:"totalAttendance"`
	AbsentAttendance   int     `json:"absentAttendance"`
	AttendedAttendance int     `json:"attendedAttendance"`
}
type InstructorStatisticResponse struct {
	InstructorID        string  `json:"instructorId"`
	Fullname            string  `json:"fullname"`
	Rating              float64 `json:"rating"`
	Experience          int     `json:"experience"`
	TotalClass          int     `json:"totalClass"`
	TotalSchedule       int     `json:"totalSchedule"`
	ParticipantsCount   int     `json:"participantsCount"`
	TotalAttendance     int     `json:"totalAttendance"`
	ReviewCount         int     `json:"reviewCount"`
	AverageReviewRating float64 `json:"averageReviewRating"`
}

type RevenueStatRequest struct {
	Range string `form:"range" binding:"omitempty,oneof=daily monthly yearly"`
}

type RevenueStat struct {
	Date  string  `json:"date"`
	Total float64 `json:"total"`
}

type RevenueStatsResponse struct {
	Range         string        `json:"range"`
	TotalRevenue  float64       `json:"totalRevenue"`
	RevenueSeries []RevenueStat `json:"revenueSeries"`
}

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}

// INSTRUCTOR SCHEDULE MANAGEMENT

type InstructorScheduleQueryParam struct {
	Status string `form:"status"`
	Sort   string `form:"sort"`
	Page   int    `form:"page" binding:"omitempty,min=1"`
	Limit  int    `form:"limit" binding:"omitempty,min=1"`
}

type OpenClassScheduleRequest struct {
	VerificationCode string `json:"verificationCode" binding:"required,len=6"`
	ZoomLink         string `json:"zoomLink" binding:"omitempty"`
}

type InstructorScheduleResponse struct {
	ID               string `json:"id"`
	ClassID          string `json:"classId"`
	ClassName        string `json:"className"`
	ClassImage       string `json:"classImage"`
	InstructorID     string `json:"instructorId"`
	InstructorName   string `json:"instructorName"`
	Location         string `json:"location"`
	Date             string `json:"date"`
	StartHour        int    `json:"startHour"`
	StartMinute      int    `json:"startMinute"`
	Capacity         int    `json:"capacity"`
	BookedCount      int    `json:"bookedCount"`
	Duration         int    `json:"duration"`
	IsOpened         bool   `json:"isOpen"`
	VerificationCode string `json:"verificationCode"`
	ZoomLink         string `json:"zoomLink"`
}
