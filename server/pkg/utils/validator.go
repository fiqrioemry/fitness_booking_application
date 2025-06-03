package utils

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func ParseBoolFormField(c *gin.Context, field string) (bool, bool) {
	val := c.PostForm(field)
	parsed, err := strconv.ParseBool(val)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid boolean value \"" + field + "\": \"" + val + "\"",
		})
		return false, false
	}
	return parsed, true
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
	for _, d := range allowedDays {
		if d == currentDay {
			return true
		}
	}
	return false
}

func IsDiceBear(url string) bool {
	return strings.Contains(url, "dicebear.com")
}
