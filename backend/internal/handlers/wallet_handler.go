package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wallet-manager/internal/models"
	"wallet-manager/internal/services"
)

type WalletHandler struct {
	walletService *services.WalletService
}

func NewWalletHandler() *WalletHandler {
	walletService := services.NewWalletService()

	return &WalletHandler{
		walletService: walletService,
	}
}

// CreateGroup 创建钱包分组
// @Summary 创建钱包分组
// @Description 为用户创建新的钱包分组
// @Tags 钱包管理
// @Accept json
// @Produce json
// @Param user_id path int true "用户ID"
// @Param request body models.CreateGroupRequest true "创建分组请求"
// @Success 201 {object} models.WalletGroup
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups [post]
func (h *WalletHandler) CreateGroup(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效: " + err.Error()})
		return
	}

	group, err := h.walletService.CreateGroup(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, group)
}

// GetGroups 获取用户的所有分组
// @Summary 获取钱包分组列表
// @Description 获取用户的所有钱包分组
// @Tags 钱包管理
// @Produce json
// @Param user_id path int true "用户ID"
// @Success 200 {array} models.WalletGroup
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups [get]
func (h *WalletHandler) GetGroups(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groups, err := h.walletService.GetUserGroups(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groups)
}

// AddAddress 添加钱包地址
// @Summary 添加钱包地址
// @Description 为用户添加新的钱包地址
// @Tags 钱包管理
// @Accept json
// @Produce json
// @Param user_id path int true "用户ID"
// @Param request body models.AddAddressRequest true "添加地址请求"
// @Success 201 {object} models.WalletAddress
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/addresses [post]
func (h *WalletHandler) AddAddress(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.AddAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效: " + err.Error()})
		return
	}

	address, err := h.walletService.AddAddress(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, address)
}

// GetAddresses 获取用户的所有地址
// @Summary 获取钱包地址列表
// @Description 获取用户的所有钱包地址
// @Tags 钱包管理
// @Produce json
// @Param user_id path int true "用户ID"
// @Success 200 {array} models.WalletAddress
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/addresses [get]
func (h *WalletHandler) GetAddresses(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addresses, err := h.walletService.GetUserAddresses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

// GetGroupAddresses 获取分组中的地址
// @Summary 获取分组地址列表
// @Description 获取指定分组中的所有地址
// @Tags 钱包管理
// @Produce json
// @Param user_id path int true "用户ID"
// @Param group_id path int true "分组ID"
// @Success 200 {array} models.WalletAddress
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/{group_id}/addresses [get]
func (h *WalletHandler) GetGroupAddresses(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := h.getGroupID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addresses, err := h.walletService.GetGroupAddresses(userID, groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, addresses)
}

// DeleteAddress 删除钱包地址
// @Summary 删除钱包地址
// @Description 删除指定的钱包地址
// @Tags 钱包管理
// @Param user_id path int true "用户ID"
// @Param address_id path int true "地址ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/addresses/{address_id} [delete]
func (h *WalletHandler) DeleteAddress(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addressIDStr := c.Param("address_id")
	addressID, err := strconv.ParseUint(addressIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的地址ID"})
		return
	}

	if err := h.walletService.DeleteAddress(userID, uint(addressID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "地址删除成功"})
}

// DeleteGroup 删除分组
// @Summary 删除钱包分组
// @Description 删除指定的钱包分组及其下的所有地址
// @Tags 钱包管理
// @Param user_id path int true "用户ID"
// @Param group_id path int true "分组ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/{group_id} [delete]
func (h *WalletHandler) DeleteGroup(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID, err := h.getGroupID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.walletService.DeleteGroup(userID, groupID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "分组及其下的地址删除成功"})
}

// ReorderGroups updates the sort order of groups
// @Summary 重排序钱包分组
// @Description 更新用户钱包分组的显示顺序
// @Tags 钱包管理
// @Accept json
// @Produce json
// @Param user_id path int true "用户ID"
// @Param request body models.ReorderGroupsRequest true "重排序请求"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /api/users/{user_id}/groups/reorder [put]
func (h *WalletHandler) ReorderGroups(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req models.ReorderGroupsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效: " + err.Error()})
		return
	}

	if err := h.walletService.UpdateGroupsOrder(userID, req.GroupOrders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "分组排序更新成功"})
}

// GetTokens 获取支持的代币列表
// @Summary 获取代币列表
// @Description 获取系统支持的所有代币
// @Tags 钱包管理
// @Produce json
// @Success 200 {array} models.Token
// @Failure 500 {object} gin.H
// @Router /api/tokens [get]
func (h *WalletHandler) GetTokens(c *gin.Context) {
	tokens, err := h.walletService.GetTokens()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// 辅助方法
func (h *WalletHandler) getUserID(c *gin.Context) (uint, error) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}

func (h *WalletHandler) getGroupID(c *gin.Context) (uint, error) {
	groupIDStr := c.Param("group_id")
	groupID, err := strconv.ParseUint(groupIDStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(groupID), nil
}
