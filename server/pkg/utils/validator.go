package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"slices"

	customErr "server/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BindAndValidateJSON[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON request",
			"error":   err.Error(),
		})
		return false
	}
	return true
}

func BindAndValidateForm[T any](c *gin.Context, req *T) bool {
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid form-data request",
			"error":   err.Error(),
		})
		return false
	}
	return true
}

func GetQueryInt(c *gin.Context, key string, defaultValue int) int {
	valStr := c.Query(key)
	val, err := strconv.Atoi(valStr)
	if err != nil || val <= 0 {
		return defaultValue
	}
	return val
}

func IsDayMatched(currentDay int, allowedDays []int) bool {
	return slices.Contains(allowedDays, currentDay)
}

func IsDiceBear(url string) bool {
	return strings.Contains(url, "dicebear.com")
}

func MustParseUUID(c *gin.Context, input string, fieldName string) uuid.UUID {
	parsed, err := uuid.Parse(input)
	if err != nil || parsed == uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid UUID for field: " + fieldName,
		})
		return uuid.Nil
	}
	return parsed
}

func ParseBoolFormField(c *gin.Context, field string) bool {
	val := c.PostForm(field)
	parsed, err := strconv.ParseBool(val)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Invalid boolean value for field \"%s\": \"%s\"", field, val),
		})
		return false
	}
	return parsed
}

func GetMultipartForm(c *gin.Context) *multipart.Form {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid multipart form",
			"error":   err.Error(),
		})
		c.Abort()
		return nil
	}
	return form
}

func ParseMixedImages(form *multipart.Form) (urls []string, files []*multipart.FileHeader) {
	for _, val := range form.Value["images"] {
		if isURL(val) {
			urls = append(urls, val)
		}
	}
	files = form.File["images"]
	return
}

func isURL(input string) bool {
	return strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://")
}

func HandleBindError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}
func EmptyString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func ContainsInt(slice []int, value int) bool {
	return slices.Contains(slice, value)
}

func SetIfNotEmpty(target *string, source string) {
	if source != "" {
		*target = source
	}
}

func SetIfNotZero(target *int, source int) {
	if source != 0 {
		*target = source
	}
}

func ValidateScheduleNotInPast(date time.Time, hour, minute int) error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	scheduleTime := time.Date(date.Year(), date.Month(), date.Day(), hour, minute, 0, 0, loc)

	if scheduleTime.Before(time.Now().In(loc)) {
		return customErr.NewBadRequest("cannot modify schedule in the past")
	}
	return nil
}
