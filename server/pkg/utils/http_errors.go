package utils

import (
	"net/http"
	customErr "server/pkg/errors"

	"errors"

	"github.com/gin-gonic/gin"
)

func HandleServiceError(c *gin.Context, err error, fallbackMsg string) {
	switch {
	case errors.Is(err, customErr.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
	case errors.Is(err, customErr.ErrAlreadyExist):
		c.JSON(http.StatusConflict, gin.H{"message": "Resource already exists"})
	case errors.Is(err, customErr.ErrInvalidInput):
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
	case errors.Is(err, customErr.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	case errors.Is(err, customErr.ErrForbidden):
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden Access"})
	case errors.Is(err, customErr.ErrUpdateFailed):
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update resource"})
	case errors.Is(err, customErr.ErrCreateFailed):
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create resource"})
	case errors.Is(err, customErr.ErrDeleteFailed):
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete resource"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": fallbackMsg})
	}
}
