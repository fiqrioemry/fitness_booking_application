package services

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	customErr "server/pkg/errors"
	"time"

	"github.com/google/uuid"
)

type NotificationService interface {
	SendClassReminder() error
	MarkAllAsRead(userID string) error
	SendToUser(req dto.NotificationEvent) error
	SendNotificationByType(req dto.SendNotificationRequest) error
	GetAllNotifications(userID string) ([]dto.NotificationResponse, error)
	GetSettingsByUser(userID string) ([]dto.NotificationSettingResponse, error)
	UpdateSetting(userID string, req dto.UpdateNotificationSettingRequest) error
}

type notificationService struct {
	repo repositories.NotificationRepository
}

func NewNotificationService(repo repositories.NotificationRepository) NotificationService {
	return &notificationService{repo}
}

func (s *notificationService) GetSettingsByUser(userID string) ([]dto.NotificationSettingResponse, error) {
	settings, err := s.repo.GetNotificationSettingsByUser(uuid.MustParse(userID))
	if err != nil {
		return nil, customErr.NewInternal("Failed to retrieve notification settings", err)
	}

	var result []dto.NotificationSettingResponse
	for _, s := range settings {
		result = append(result, dto.NotificationSettingResponse{
			TypeID:  s.NotificationTypeID.String(),
			Code:    s.NotificationType.Code,
			Title:   s.NotificationType.Title,
			Channel: s.Channel,
			Enabled: s.Enabled,
		})
	}

	return result, nil
}

func (s *notificationService) UpdateSetting(userID string, req dto.UpdateNotificationSettingRequest) error {
	typeID, err := uuid.Parse(req.TypeID)
	if err != nil {
		return customErr.NewBadRequest("Invalid TypeID")
	}

	setting, err := s.repo.FindSetting(uuid.MustParse(userID), typeID, req.Channel)
	if err != nil {
		return customErr.NewNotFound("Notification setting not found")
	}

	setting.Enabled = req.Enabled
	if err := s.repo.UpdateNotificationSetting(setting); err != nil {
		return customErr.NewInternal("Failed to update notification setting", err)
	}

	return nil
}

func (s *notificationService) GetAllNotifications(userID string) ([]dto.NotificationResponse, error) {
	notifs, err := s.repo.GetAllBrowserNotifications(uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	var result []dto.NotificationResponse
	for _, n := range notifs {
		result = append(result, dto.NotificationResponse{
			ID:        n.ID.String(),
			TypeCode:  n.TypeCode,
			Title:     n.Title,
			Message:   n.Message,
			Channel:   n.Channel,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt.Format(time.RFC3339),
		})
	}

	return result, nil
}

func (s *notificationService) MarkAllAsRead(userID string) error {
	return s.repo.MarkAllNotificationsRead(uuid.MustParse(userID))
}

func (s *notificationService) SendNotificationByType(req dto.SendNotificationRequest) error {
	settings, err := s.repo.GetUsersWithEnabledNotification(req.TypeCode)
	if err != nil {
		return err
	}

	var notifs []models.Notification
	for _, setting := range settings {
		notifs = append(notifs, models.Notification{
			UserID:   setting.UserID,
			TypeCode: req.TypeCode,
			Title:    req.Title,
			Message:  req.Message,
			Channel:  setting.Channel,
		})
		// ! Non aktifin sementara di development biar ga spam email ke dummy email
		// if setting.Channel == "email" && setting.User.Email != "" {
		// 	go func(email, title, msg string) {
		// 		if err := utils.SendNotificationEmail(email, title, msg); err != nil {
		// 			fmt.Printf("failed to send email to %s: %v\n", email, err)
		// 		}
		// 	}(setting.User.Email, req.Title, req.Message)
		// }
	}

	return s.repo.InsertNotifications(notifs)
}

func (s *notificationService) SendClassReminder() error {
	reminderTime := time.Now().UTC().Add(1 * time.Hour).Truncate(time.Minute)

	bookings, err := s.repo.GetBookingsForClassReminder(reminderTime)
	if err != nil {
		return customErr.NewInternal("Failed to fetch bookings for class reminder", err)
	}

	notifType, err := s.repo.GetTypeByCode("class_reminder")
	if err != nil {
		return customErr.NewNotFound("Class reminder notification type not found")
	}

	for _, booking := range bookings {
		userID := booking.UserID
		schedule := booking.ClassSchedule

		setting, err := s.repo.FindSetting(userID, notifType.ID, "email")
		if err != nil || !setting.Enabled {
			continue
		}

		message := fmt.Sprintf("Reminder: You have a class '%s' at %02d:%02d today.",
			schedule.ClassName, schedule.StartHour, schedule.StartMinute)

		notif := models.Notification{
			UserID:   booking.UserID,
			TypeCode: "class_reminder",
			Title:    "Class Reminder",
			Message:  message,
			Channel:  "browser",
		}
		if err := s.repo.CreateNotification(&notif); err != nil {
			return customErr.NewInternal("Failed to send class reminder", err)
		}
	}

	return nil
}

func (s *notificationService) SendToUser(req dto.NotificationEvent) error {
	uid := uuid.MustParse(req.UserID)

	ntype, err := s.repo.GetTypeByCode(req.Type)
	if err != nil {
		return err
	}

	for _, channel := range []string{"browser"} {
		setting, err := s.repo.FindSetting(uid, ntype.ID, channel)
		if err != nil || !setting.Enabled {
			continue
		}

		notif := models.Notification{
			UserID:   uid,
			TypeCode: req.Type,
			Title:    req.Title,
			Message:  req.Message,
			Channel:  channel,
		}
		if err := s.repo.CreateNotification(&notif); err != nil {
			return err
		}

		// ! Non aktifin sementara di development biar ga spam email ke dummy email
		// if channel == "email" && setting.User.Email != "" {
		// 	go func(email, title, msg string) {
		// 		_ = utils.SendNotificationEmail(email, title, msg)
		// 	}(setting.User.Email, req.Title, req.Message)
		// }
	}

	return nil
}
