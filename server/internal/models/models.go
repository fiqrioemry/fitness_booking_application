package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string     `gorm:"type:text;not null" json:"-"`
	Role      string     `gorm:"type:varchar(255);default:'customer';check:role IN ('customer','admin','instructor')" json:"role"`
	Fullname  string     `gorm:"type:varchar(255);not null" json:"fullname"`
	Birthday  *time.Time `gorm:"birthday,omitempty"`
	Phone     string     `gorm:"type:varchar(20)" json:"phone"`
	Gender    string     `gorm:"type:varchar(10)" json:"gender"`
	Avatar    string     `gorm:"type:varchar(255)" json:"avatar"`
	Bio       string     `gorm:"type:text" json:"bio"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`

	Tokens []Token `gorm:"foreignKey:UserID" json:"-"`
}

type Token struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);index;not null" json:"userId"`
	Token     string    `gorm:"type:text;not null" json:"token"`
	UserAgent string
	IP        string
	ExpiredAt time.Time `json:"expiredAt"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Class struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Title          string         `gorm:"type:varchar(255);not null" json:"title"`
	Image          string         `gorm:"type:varchar(255);not null" json:"image"`
	IsActive       bool           `gorm:"not null;default:true" json:"isActive"`
	Duration       int            `gorm:"not null" json:"duration"`
	Description    string         `gorm:"type:text" json:"description"`
	Additional     string         `gorm:"type:longtext" json:"-"`
	AdditionalList []string       `gorm:"-" json:"additional"`
	TypeID         uuid.UUID      `json:"typeId"`
	LevelID        uuid.UUID      `json:"levelId"`
	LocationID     uuid.UUID      `json:"locationId"`
	CategoryID     uuid.UUID      `json:"categoryId"`
	SubcategoryID  uuid.UUID      `json:"subcategoryId"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// relationship one - to - many
	Type        Type            `gorm:"foreignKey:TypeID"`
	Level       Level           `gorm:"foreignKey:LevelID"`
	Category    Category        `gorm:"foreignKey:CategoryID"`
	Subcategory Subcategory     `gorm:"foreignKey:SubcategoryID"`
	Location    Location        `gorm:"foreignKey:LocationID"`
	Galleries   []*ClassGallery `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;" json:"galleries,omitempty"`
	Reviews     []Review        `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;" json:"reviews,omitempty"`
	// many to many (1 class can be in many packages and 1 package can have many classes)
	Packages []Package `gorm:"many2many:package_classes;" json:"packages,omitempty"`
}

type ClassGallery struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID   uuid.UUID `gorm:"type:char(36);not null;index" json:"classId"`
	URL       string    `gorm:"type:varchar(255);not null" json:"imageUrl"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

type PackageClass struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID   uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	PackageID uuid.UUID `gorm:"type:char(36);not null" json:"packageId"`
}

type UserPackage struct {
	ID              uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          uuid.UUID      `gorm:"type:char(36);not null" json:"userId"`
	PackageID       uuid.UUID      `gorm:"type:char(36);not null" json:"packageId"`
	PackageName     string         `gorm:"type:varchar(255);not null" json:"packageName"`
	RemainingCredit int            `gorm:"not null;default:0" json:"remainingCredit"`
	ExpiredAt       *time.Time     `json:"expiredAt"`
	PurchasedAt     time.Time      `gorm:"autoCreateTime" json:"purchasedAt"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	Package Package `gorm:"foreignKey:PackageID" json:"package"`
}

type Package struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name           string         `gorm:"type:varchar(255);not null" json:"name"`
	Description    string         `gorm:"type:text;not null" json:"description"`
	IsActive       bool           `gorm:"not null;default:true" json:"isActive"`
	Image          string         `gorm:"type:varchar(255)" json:"image"`
	Price          float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Credit         int            `gorm:"not null" json:"credit"`
	Discount       float64        `gorm:"not null;default:0" json:"discount"`
	Expired        int            `json:"expired"`
	Additional     string         `gorm:"type:longtext" json:"-"`
	AdditionalList []string       `gorm:"-" json:"additional"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	Classes []Class `gorm:"many2many:package_classes;" json:"classes,omitempty"`
}

type Payment struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	PackageID       uuid.UUID `gorm:"type:char(36);not null" json:"packageId"`
	InvoiceNumber   string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	PackageName     string    `gorm:"type:varchar(255);not null" json:"packageName"`
	UserID          uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	Fullname        string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Email           string    `gorm:"type:varchar(255);not null" json:"email"`
	PaymentMethod   string    `gorm:"type:varchar(50);not null" json:"paymentMethod"`
	PaymentLink     string    `gorm:"type:text;not null" json:"paymentLink"`
	Status          string    `gorm:"type:varchar(20);default:'pending';check:status IN ('success', 'pending', 'failed')" json:"status"`
	PaidAt          time.Time `gorm:"autoCreateTime" json:"paidAt"`
	BasePrice       float64   `gorm:"type:decimal(10,2);not null"`
	Tax             float64   `gorm:"type:decimal(10,2);not null"`
	Total           float64   `gorm:"type:decimal(10,2);not null"`
	VoucherCode     *string   `gorm:"type:varchar(100)" json:"voucherCode,omitempty"`
	VoucherDiscount float64   `gorm:"default:0" json:"voucherDiscount"`
}

type ClassSchedule struct {
	ID               uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID          uuid.UUID `gorm:"type:char(36);not null" json:"classId"`
	ClassImage       string    `gorm:"type:varchar(255);not null" json:"classImage"`
	ClassName        string    `gorm:"type:varchar(255);not null" json:"className"`
	Location         string    `gorm:"type:varchar(255);not null" json:"location"`
	InstructorID     uuid.UUID `gorm:"type:char(36);not null" json:"instructorId"`
	InstructorName   string    `gorm:"type:varchar(255);not null" json:"instructorName"`
	Capacity         int       `gorm:"not null" json:"capacity"`
	Color            string    `gorm:"type:varchar(20)" json:"color"`
	Date             time.Time `gorm:"not null" json:"date"`
	Booked           int       `gorm:"not null;default:0" json:"booked"`
	StartHour        int       `gorm:"not null" json:"startHour"`
	StartMinute      int       `gorm:"not null" json:"startMinute"`
	Duration         int       `gorm:"not null" json:"duration"`
	ZoomLink         *string   `gorm:"type:varchar(255)" json:"zoomLink,omitempty"`
	IsOpened         bool      `gorm:"default:false" json:"isOpened"`
	VerificationCode *string   `gorm:"type:varchar(10)" json:"verificationCode,omitempty"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Bookings  []Booking      `gorm:"foreignKey:ClassScheduleID" json:"bookings"`
}

type Booking struct {
	ID              uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID          uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:idx_user_schedule" json:"userId"`
	ClassScheduleID uuid.UUID `gorm:"type:char(36);not null;uniqueIndex:idx_user_schedule" json:"classScheduleId"`
	Status          string    `gorm:"type:varchar(20);not null;default:'booked';check:status IN ('booked','canceled')" json:"status"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt"`

	User          User          `gorm:"foreignKey:UserID" json:"user"`
	ClassSchedule ClassSchedule `gorm:"foreignKey:ClassScheduleID" json:"classSchedule"`
	Attendance    Attendance    `gorm:"foreignKey:BookingID" json:"attendance"`
}

type Attendance struct {
	ID         uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	BookingID  uuid.UUID  `gorm:"type:char(36);not null;uniqueIndex" json:"bookingId"`
	Status     string     `gorm:"type:varchar(20);not null;default:'not-join';check:status IN ('not-join','entered', 'attended', 'absent')" json:"status"`
	IsReviewed bool       `gorm:"default:false" json:"isReviewed"`
	CheckedIn  bool       `gorm:"default:false" json:"checkedIn"`
	CheckedOut bool       `gorm:"default:false" json:"checkedOut"`
	CheckedAt  *time.Time `json:"checkedAt"`
	VerifiedAt *time.Time `json:"verifiedAt"`
}

type ScheduleTemplate struct {
	ID              uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	ClassID         uuid.UUID      `gorm:"type:char(36);not null" json:"classId"`
	ClassImage      string         `gorm:"type:varchar(255);not null" json:"classImage"`
	ClassName       string         `gorm:"type:varchar(255);not null" json:"className"`
	Location        string         `gorm:"type:varchar(255);not null" json:"location"`
	InstructorID    uuid.UUID      `gorm:"type:char(36);not null" json:"instructorId"`
	InstructorName  string         `gorm:"type:varchar(255);not null" json:"instructorName"`
	DayOfWeeks      datatypes.JSON `gorm:"type:json" json:"dayOfWeeks"`
	StartHour       int            `gorm:"not null" json:"startHour"`
	StartMinute     int            `gorm:"not null" json:"startMinute"`
	Capacity        int            `gorm:"not null" json:"capacity"`
	IsActive        bool           `gorm:"default:true" json:"isActive"`
	Color           string         `gorm:"type:varchar(20)" json:"color"`
	LastGeneratedAt *time.Time     `gorm:"column:last_generated_at" json:"lastGeneratedAt"`
	EndDate         time.Time      `gorm:"not null" json:"endDate"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	Class      Class      `gorm:"foreignKey:ClassID" json:"class"`
	Instructor Instructor `gorm:"foreignKey:InstructorID" json:"instructor"`
}

type Location struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(255);not null" json:"name"`
	Address     string         `gorm:"type:varchar(255);not null" json:"address"`
	GeoLocation string         `gorm:"type:varchar(255);not null" json:"geoLocation"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Category struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Subcategory struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID uuid.UUID      `gorm:"type:char(36);not null" json:"categoryId"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// relasi
	Category Category `gorm:"foreignKey:CategoryID"`
}

type Type struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Level struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Voucher struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Code         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	Description  string         `gorm:"type:text" json:"description"`
	DiscountType string         `gorm:"type:varchar(20);not null" json:"discountType"`
	Discount     float64        `gorm:"not null" json:"discount"`
	MaxDiscount  *float64       `gorm:"maxDiscount,omitempty"`
	Quota        int            `gorm:"not null" json:"quota"`
	IsReusable   bool           `gorm:"default:false" json:"isReusable"`
	ExpiredAt    time.Time      `gorm:"not null" json:"expiredAt"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type UsedVoucher struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);index;not null"`
	VoucherID uuid.UUID      `gorm:"type:char(36);index;not null"`
	UsedAt    time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Review struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null" json:"userId"`
	ClassID   uuid.UUID      `gorm:"type:char(36);not null" json:"classId"`
	Rating    int            `gorm:"not null" json:"rating"`
	Comment   string         `gorm:"type:text" json:"comment"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	User  User  `gorm:"foreignKey:UserID" json:"user"`
	Class Class `gorm:"foreignKey:ClassID" json:"class"`
}

type Instructor struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID         uuid.UUID      `gorm:"type:char(36);not null;uniqueIndex" json:"userId"`
	Experience     int            `gorm:"not null;default:0" json:"experience"`
	Specialties    string         `gorm:"type:text" json:"specialties"`
	Rating         float64        `gorm:"type:decimal(2,1);default:5.0" json:"rating"`
	TotalClass     int            `gorm:"not null;default:0" json:"totalClass"`
	Certifications string         `gorm:"type:text" json:"certifications"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID"`
}

type NotificationType struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Code           string         `gorm:"unique;not null"`
	Title          string         `gorm:"type:varchar(255);not null"`
	Category       string         `gorm:"type:varchar(100)"`
	DefaultEnabled bool           `gorm:"default:true"`
	IsRequired     bool           `gorm:"default:false"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type NotificationSetting struct {
	ID                 uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID             uuid.UUID      `gorm:"type:char(36);index;not null"`
	NotificationTypeID uuid.UUID      `gorm:"type:char(36);index;not null"`
	Channel            string         `gorm:"type:varchar(20);not null;check:channel IN ('email','browser')"`
	Enabled            bool           `gorm:"default:true"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`

	NotificationType NotificationType `gorm:"foreignKey:NotificationTypeID"`
	User             User             `gorm:"foreignKey:UserID"`
}

type Notification struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;index"`
	TypeCode  string         `gorm:"type:varchar(100);not null"`
	Title     string         `gorm:"type:varchar(255);not null"`
	Message   string         `gorm:"type:text;not null"`
	Channel   string         `gorm:"type:varchar(50);not null"`
	IsRead    bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (up *UserPackage) BeforeCreate(tx *gorm.DB) (err error) {
	if up.ID == uuid.Nil {
		up.ID = uuid.New()
	}
	return
}
func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (cs *ClassSchedule) BeforeCreate(tx *gorm.DB) (err error) {
	if cs.ID == uuid.Nil {
		cs.ID = uuid.New()
	}
	return
}

func (s *ScheduleTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}
func (s *Subcategory) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}
func (i *Instructor) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return
}

func (n *Notification) BeforeCreate(tx *gorm.DB) (err error) {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return
}

func (l *Level) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}

func (t *Type) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}
func (p *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (p *Package) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return
}

func (c *Class) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}

func (c *Voucher) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return
}
func (pc *PackageClass) BeforeCreate(tx *gorm.DB) (err error) {
	if pc.ID == uuid.Nil {
		pc.ID = uuid.New()
	}
	return
}

func (m *ClassGallery) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}

func (p *Package) BeforeSave(tx *gorm.DB) error {
	if p.AdditionalList != nil {
		data, err := json.Marshal(p.AdditionalList)
		if err != nil {
			return err
		}
		p.Additional = string(data)
	}
	return nil
}

func (c *Class) BeforeSave(tx *gorm.DB) (err error) {
	if c.AdditionalList != nil {
		jsonBytes, err := json.Marshal(c.AdditionalList)
		if err != nil {
			return err
		}
		c.Additional = string(jsonBytes)
	}
	return nil
}

func (c *Class) AfterFind(tx *gorm.DB) (err error) {
	if c.Additional != "" {
		var tags []string
		if err := json.Unmarshal([]byte(c.Additional), &tags); err != nil {
			return err
		}
		c.AdditionalList = tags
	}
	return nil
}

func (p *Package) AfterFind(tx *gorm.DB) error {
	if p.Additional != "" {
		var data []string
		if err := json.Unmarshal([]byte(p.Additional), &data); err != nil {
			return err
		}
		p.AdditionalList = data
	}
	return nil
}
