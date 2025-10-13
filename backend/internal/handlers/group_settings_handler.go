package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wallet-manager/internal/models"
	"wallet-manager/internal/services"
)

type GroupSettingsHandler struct {
	service *services.GroupSettingsService
}

func NewGroupSettingsHandler() *GroupSettingsHandler {
	return &GroupSettingsHandler{
		service: services.NewGroupSettingsService(),
	}
}

// GetGroupSettings gets settings for a specific group
// @Summary Get group settings
// @Description Get configuration settings for a specific wallet group
// @Tags Group Settings
// @Produce json
// @Param user_id path int true "User ID"
// @Param group_id path int true "Group ID"
// @Success 200 {object} models.GroupSettingsResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/users/{user_id}/groups/{group_id}/settings [get]
func (h *GroupSettingsHandler) GetGroupSettings(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := getGroupID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settings, err := h.service.GetGroupSettings(userID, groupID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.ConvertToResponse(settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "转换响应失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllGroupSettings gets settings for all groups of a user
// @Summary Get all group settings
// @Description Get configuration settings for all wallet groups of a user
// @Tags Group Settings
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} models.GroupSettingsResponse
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/settings [get]
func (h *GroupSettingsHandler) GetAllGroupSettings(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	settingsList, err := h.service.GetAllGroupSettings(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to response format
	var responses []models.GroupSettingsResponse
	for _, settings := range settingsList {
		response, err := h.service.ConvertToResponse(&settings)
		if err != nil {
			continue // Skip invalid settings
		}
		responses = append(responses, *response)
	}

	c.JSON(http.StatusOK, responses)
}

// UpdateGroupSettings updates or creates settings for a group
// @Summary Update group settings
// @Description Update or create configuration settings for a wallet group
// @Tags Group Settings
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param group_id path int true "Group ID"
// @Param request body models.UpdateGroupSettingsRequest true "Settings update request"
// @Success 200 {object} models.GroupSettingsResponse
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/{group_id}/settings [put]
func (h *GroupSettingsHandler) UpdateGroupSettings(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := getGroupID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.UpdateGroupSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效: " + err.Error()})
		return
	}

	settings, err := h.service.UpdateGroupSettings(userID, groupID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.ConvertToResponse(settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "转换响应失败"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteGroupSettings deletes settings for a group
// @Summary Delete group settings
// @Description Delete configuration settings for a wallet group
// @Tags Group Settings
// @Param user_id path int true "User ID"
// @Param group_id path int true "Group ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/{group_id}/settings [delete]
func (h *GroupSettingsHandler) DeleteGroupSettings(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := getGroupID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DeleteGroupSettings(userID, groupID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "配置删除成功"})
}

// Helper functions
func getUserID(c *gin.Context) (uint, error) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}

func getGroupID(c *gin.Context) (uint, error) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(groupID), nil
}
