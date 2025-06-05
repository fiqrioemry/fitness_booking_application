package repositories

import (
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	MarkAllNotificationsRead(userID uuid.UUID) error
	InsertNotifications(notifs []models.Notification) error
	CreateNotification(notification *models.Notification) error
	GetAllNotificationTypes() ([]models.NotificationType, error)
	UpdateNotificationSetting(setting *models.NotificationSetting) error
	CreateNotificationSetting(setting *models.NotificationSetting) error
	GetTypeByCode(code string) (*models.NotificationType, error)
	GetAllBrowserNotifications(userID uuid.UUID) ([]models.Notification, error)
	GetBookingsForClassReminder(reminderTime time.Time) ([]models.Booking, error)
	GetNotificationSettingsByUser(userID uuid.UUID) ([]models.NotificationSetting, error)
	GetUsersWithEnabledNotification(typeCode string) ([]models.NotificationSetting, error)
	FindSetting(userID, typeID uuid.UUID, channel string) (*models.NotificationSetting, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) GetAllNotificationTypes() ([]models.NotificationType, error) {
	var types []models.NotificationType
	err := r.db.Find(&types).Error
	return types, err
}
func (r *notificationRepository) MarkAllNotificationsRead(userID uuid.UUID) error {
	return r.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true).Error
}

func (r *notificationRepository) GetNotificationSettingsByUser(userID uuid.UUID) ([]models.NotificationSetting, error) {
	var settings []models.NotificationSetting
	err := r.db.
		Preload("NotificationType").
		Where("user_id = ?", userID).
		Find(&settings).Error
	return settings, err
}

func (r *notificationRepository) FindSetting(userID, typeID uuid.UUID, channel string) (*models.NotificationSetting, error) {
	var setting models.NotificationSetting
	err := r.db.
		Where("user_id = ? AND notification_type_id = ? AND channel = ?", userID, typeID, channel).
		First(&setting).Error
	return &setting, err
}

func (r *notificationRepository) UpdateNotificationSetting(setting *models.NotificationSetting) error {
	return r.db.Save(setting).Error
}

func (r *notificationRepository) InsertNotifications(notifs []models.Notification) error {
	return r.db.Create(&notifs).Error
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) CreateNotificationSetting(setting *models.NotificationSetting) error {
	return r.db.Create(setting).Error
}

func (r *notificationRepository) GetAllBrowserNotifications(userID uuid.UUID) ([]models.Notification, error) {
	var notifs []models.Notification
	err := r.db.
		Where("user_id = ? AND channel = ?", userID, "browser").
		Order("created_at DESC").
		Find(&notifs).Error
	return notifs, err
}

func (r *notificationRepository) GetUsersWithEnabledNotification(typeCode string) ([]models.NotificationSetting, error) {
	var settings []models.NotificationSetting
	err := r.db.
		Preload("NotificationType").
		Preload("User").
		Where("channel IN ?", []string{"browser", "email"}).
		Where("enabled = ? AND notification_type_id IN (SELECT id FROM notification_types WHERE code = ?)", true, typeCode).
		Find(&settings).Error

	return settings, err
}

func (r *notificationRepository) GetTypeByCode(code string) (*models.NotificationType, error) {
	var nt models.NotificationType
	err := r.db.Where("code = ?", code).First(&nt).Error
	return &nt, err
}

func (r *notificationRepository) GetBookingsForClassReminder(reminderTime time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Preload("User").Preload("ClassSchedule.Class").
		Where("status = ? AND DATE(date) = ? AND start_hour = ? AND start_minute = ?",
			"booked",
			reminderTime.Format("2006-01-02"),
			reminderTime.Hour(),
			reminderTime.Minute(),
		).Joins("JOIN class_schedules ON class_schedules.id = bookings.class_schedule_id").
		Find(&bookings).Error
	return bookings, err
}
