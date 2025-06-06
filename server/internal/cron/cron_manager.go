package cron

import (
	"log"

	"server/internal/services"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	c                   *cron.Cron
	paymentService      services.PaymentService
	scheduleService     services.ScheduleTemplateService
	notificationService services.NotificationService
	bookingService      services.BookingService
}

func NewCronManager(
	payment services.PaymentService,
	schedule services.ScheduleTemplateService,
	notification services.NotificationService,
	attendance services.BookingService,
) *CronManager {
	return &CronManager{
		c:                   cron.New(cron.WithSeconds()),
		paymentService:      payment,
		scheduleService:     schedule,
		notificationService: notification,
		bookingService:      attendance,
	}
}

func (cm *CronManager) RegisterJobs() {
	// Update payment pending → failed (everyday @00:30)
	cm.c.AddFunc("0 30 0 * * *", func() {
		log.Println("Cron: Checking expired pending payments...")
		if err := cm.paymentService.ExpireOldPendingPayments(); err != nil {
			log.Println("Error expiring payments:", err)
		} else {
			log.Println("Payment status updated (pending → failed)")
		}
	})

	// Generate recurring class schedule (everyday @12:00)
	cm.c.AddFunc("0 12 * * *", func() {
		log.Println("Cron: Auto-generating class schedules...")
		if err := cm.scheduleService.AutoGenerateSchedules(); err != nil {
			log.Println("Schedule generation failed:", err)
		} else {
			log.Println("Recurring class schedules generated")
		}
	})

	// Reminder and Mark Absent every 15 minutes
	cm.c.AddFunc("@every 15m", func() {
		log.Println("Cron: Sending class reminders...")
		if err := cm.notificationService.SendClassReminder(); err != nil {
			log.Println(" Reminder failed:", err)
		} else {
			log.Println("Class reminders sent")
		}

		log.Println("Cron: Marking absents for missed bookings...")
		if err := cm.bookingService.MarkAbsentBookings(); err != nil {
			log.Println(" Marking absents failed:", err)
		} else {
			log.Println("Marked absent attendees")
		}
	})

}

func (cm *CronManager) Start() {
	cm.c.Start()
	log.Println("Cron Manager started")
}
