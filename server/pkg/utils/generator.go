package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomUserAvatar(fullname string) string {
	return fmt.Sprintf("https://api.dicebear.com/6.x/initials/svg?seed=%s", fullname)
}

func GenerateOTP(length int) string {
	digits := "0123456789"
	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteByte(digits[rand.Intn(len(digits))])
	}

	return sb.String()
}

func GenerateSlug(input string) string {

	slug := strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	suffix := strconv.Itoa(rand.Intn(1_000_000))
	slug = slug + "-" + leftPad(suffix, "0", 6)

	return slug
}

func leftPad(s string, pad string, length int) string {
	for len(s) < length {
		s = pad + s
	}
	return s
}

func ParseDeviceType(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "mobile"):
		return "mobile"
	case strings.Contains(ua, "tablet"):
		return "tablet"
	case strings.Contains(ua, "windows"), strings.Contains(ua, "mac"):
		return "desktop"
	default:
		return "unknown"
	}
}

func ParsePlatform(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "android"):
		return "android"
	case strings.Contains(ua, "iphone"):
		return "ios"
	case strings.Contains(ua, "windows"):
		return "windows"
	case strings.Contains(ua, "mac"):
		return "mac"
	default:
		return "unknown"
	}
}

func GenerateInvoiceNumber(paymentID uuid.UUID) string {
	timestamp := time.Now().Format("20060102")
	shortID := paymentID.String()[:8]
	return fmt.Sprintf("INV/%s/%s", timestamp, shortID)
}
