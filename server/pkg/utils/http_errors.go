package utils

import (
	"net/http"
	"server/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HandleServiceError(c *gin.Context, err error, fallbackMsg string) {
	logger := GetLogger()

	appErr, ok := err.(*errors.AppError)
	if !ok {
		logger.Warn("⚠️ Unrecognized error",
			zap.String("path", c.FullPath()),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("fallback", fallbackMsg),
			zap.String("raw_error", err.Error()),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": fallbackMsg, "error": err.Error()})
		return
	}

	if appErr.Code >= 500 {
		logger.Error("🔥 Internal error",
			zap.String("path", c.FullPath()),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("message", appErr.Message),
			zap.Error(appErr.Err),
		)
	}

	c.JSON(appErr.Code, gin.H{"message": appErr.Message, "error": err.Error()})
}
