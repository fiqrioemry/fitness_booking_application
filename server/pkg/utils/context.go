package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MustGetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		panic("userID not found in context")
	}
	idStr, ok := userID.(string)
	if !ok {
		panic("userID in context is not a string")
	}
	return idStr
}

func MustGetAppID(c *gin.Context) uuid.UUID {
	appID, exists := c.Get("appID")
	if !exists {
		panic("appID not found in context")
	}
	id, ok := appID.(uuid.UUID)
	if !ok {
		panic("appID in context is not uuid.UUID")
	}
	return id
}

func MustGetRole(c *gin.Context) string {
	role, exists := c.Get("role")
	if !exists {
		panic("role not found in context")
	}
	userRole, ok := role.(string)
	if !ok {
		panic("role in context is not a string")
	}
	return userRole
}
