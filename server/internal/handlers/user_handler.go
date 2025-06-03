package handlers

import (
	"net/http"
	"server/internal/dto"
	"server/internal/services"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	response, err := h.service.GetUserDetail(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.UpdateUserDetailRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.UpdateProfile(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *UserHandler) UpdateAvatar(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.UpdateAvatarRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	if req.Avatar != nil && req.Avatar.Filename != "" {
		avatarURL, err := utils.UploadImageWithValidation(req.Avatar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		req.AvatarURL = avatarURL
	}

	if err := h.service.UpdateAvatar(userID, req); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var params dto.UserQueryParam
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	users, pagination, err := h.service.GetAllUsers(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       users,
		"pagination": pagination,
	})
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserDetail(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserStats(c *gin.Context) {
	stats, err := h.service.GetUserStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
