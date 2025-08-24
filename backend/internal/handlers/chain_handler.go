package handlers

import (
	"net/http"
	"strconv"
	"wallet-manager/internal/models"
	"wallet-manager/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChainHandler struct {
	chainService *services.ChainService
}

func NewChainHandler(db *gorm.DB) *ChainHandler {
	return &ChainHandler{
		chainService: services.NewChainService(db),
	}
}

// GetChains 获取所有活跃的区块链列表
func (h *ChainHandler) GetChains(c *gin.Context) {
	chains, err := h.chainService.GetActiveChains()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取区块链列表失败"})
		return
	}

	// 转换为响应格式
	response := make([]models.ChainInfoResponse, len(chains))
	for i, chain := range chains {
		response[i] = models.ChainInfoResponse{
			ChainID:          chain.ChainID,
			Name:             chain.Name,
			Symbol:           chain.Symbol,
			IsTestnet:        chain.IsTestnet,
			IsActive:         chain.IsActive,
			BlockExplorerURL: chain.BlockExplorerURL,
			RPCEndpoints:     []models.RPCEndpoint{},     // No longer available - use user RPCs
			BalanceContracts: chain.BalanceContracts,
			Tokens:           []models.Token{},           // No longer available - use user tokens
		}
	}

	c.JSON(http.StatusOK, models.ChainListResponse{Chains: response})
}

// GetChainInfo 获取指定链的详细信息
func (h *ChainHandler) GetChainInfo(c *gin.Context) {
	chainIDStr := c.Param("chain_id")
	chainID, err := strconv.Atoi(chainIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的链ID"})
		return
	}

	chain, err := h.chainService.GetChainByID(chainID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "区块链不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取区块链信息失败"})
		}
		return
	}

	response := models.ChainInfoResponse{
		ChainID:          chain.ChainID,
		Name:             chain.Name,
		Symbol:           chain.Symbol,
		IsTestnet:        chain.IsTestnet,
		IsActive:         chain.IsActive,
		BlockExplorerURL: chain.BlockExplorerURL,
		RPCEndpoints:     []models.RPCEndpoint{},     // No longer available - use user RPCs
		BalanceContracts: chain.BalanceContracts,
		Tokens:           []models.Token{},           // No longer available - use user tokens
	}

	c.JSON(http.StatusOK, response)
}

// GetRPCEndpoints removed - use GetUserRPCs instead

// GetBalanceContracts 获取指定链的余额查询合约列表
func (h *ChainHandler) GetBalanceContracts(c *gin.Context) {
	chainIDStr := c.Param("chain_id")
	chainID, err := strconv.Atoi(chainIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的链ID"})
		return
	}

	contracts, err := h.chainService.GetBalanceContracts(chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取余额查询合约失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance_contracts": contracts})
}

// GetTokens removed - use GetUserTokens instead

// AddRPCEndpoint removed - use CreateUserRPC instead

// AddBalanceContract 添加余额查询合约
func (h *ChainHandler) AddBalanceContract(c *gin.Context) {
	var req struct {
		ChainID         int    `json:"chain_id" binding:"required"`
		Name            string `json:"name" binding:"required"`
		ContractAddress string `json:"contract_address" binding:"required"`
		Priority        int    `json:"priority"`
		ABIJSON         string `json:"abi_json"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误"})
		return
	}

	contract := &models.BalanceContract{
		ChainID:         req.ChainID,
		Name:            req.Name,
		ContractAddress: req.ContractAddress,
		Priority:        req.Priority,
		ABIJSON:         req.ABIJSON,
		IsActive:        true,
	}

	if err := h.chainService.AddBalanceContract(contract); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加余额查询合约失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "余额查询合约添加成功", "contract": contract})
}

// AddToken removed - use CreateUserToken instead

// UpdateToken removed - use UpdateUserToken instead

// CreateUserRPC 用户添加自定义RPC端点
func (h *ChainHandler) CreateUserRPC(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var req models.CreateUserRPCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	userRPC := &models.UserRPCEndpoint{
		UserID:   uint(userID),
		ChainID:  req.ChainID,
		Name:     req.Name,
		URL:      req.URL,
		IsActive: true,
	}

	if err := h.chainService.CreateUserRPC(userRPC); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加RPC端点失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "RPC端点添加成功", "rpc": userRPC})
}

// UpdateUserRPC 用户更新自定义RPC端点
func (h *ChainHandler) UpdateUserRPC(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	rpcIDStr := c.Param("rpc_id")
	rpcID, err := strconv.ParseUint(rpcIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的RPC ID"})
		return
	}

	var req models.UpdateUserRPCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证RPC是否属于该用户
	_, err = h.chainService.GetUserRPCByID(uint(rpcID), uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "RPC端点不存在或无权限访问"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取RPC端点信息失败"})
		}
		return
	}

	// 更新字段
	updateData := &models.UserRPCEndpoint{
		ChainID:  req.ChainID,
		Name:     req.Name,
		URL:      req.URL,
		IsActive: req.IsActive,
	}

	if err := h.chainService.UpdateUserRPC(uint(rpcID), uint(userID), updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新RPC端点失败: " + err.Error()})
		return
	}

	// 获取更新后的RPC信息
	updatedRPC, err := h.chainService.GetUserRPCByID(uint(rpcID), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取更新后的RPC信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "RPC端点更新成功", "rpc": updatedRPC})
}

// GetUserRPCs 获取用户自定义的RPC端点列表
func (h *ChainHandler) GetUserRPCs(c *gin.Context) {
	chainIDStr := c.Param("chain_id")
	chainID, err := strconv.Atoi(chainIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的链ID"})
		return
	}

	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		// 没有用户ID，返回空列表
		c.JSON(http.StatusOK, gin.H{"rpc_endpoints": []models.RPCEndpoint{}})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 获取用户自定义RPC
	userRPCs, err := h.chainService.GetUserRPCs(uint(userID), chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户RPC端点失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rpc_endpoints": userRPCs})
}

// CreateUserToken 用户添加自定义代币
func (h *ChainHandler) CreateUserToken(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	var req models.CreateUserTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	userToken := &models.UserToken{
		UserID:          uint(userID),
		ChainID:         req.ChainID,
		Symbol:          req.Symbol,
		Name:            req.Name,
		ContractAddress: req.ContractAddress,
		Decimals:        req.Decimals,
		IsActive:        true,
	}

	if err := h.chainService.CreateUserToken(userToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加代币失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "代币添加成功", "token": userToken})
}

// UpdateUserToken 用户更新自定义代币
func (h *ChainHandler) UpdateUserToken(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	tokenIDStr := c.Param("token_id")
	tokenID, err := strconv.ParseUint(tokenIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的代币ID"})
		return
	}

	var req models.UpdateUserTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数错误: " + err.Error()})
		return
	}

	// 验证代币是否属于该用户
	_, err = h.chainService.GetUserTokenByID(uint(tokenID), uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "代币不存在或无权限访问"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取代币信息失败"})
		}
		return
	}

	// 更新字段
	updateData := &models.UserToken{
		ChainID:         req.ChainID,
		Symbol:          req.Symbol,
		Name:            req.Name,
		ContractAddress: req.ContractAddress,
		Decimals:        req.Decimals,
		IsActive:        req.IsActive,
	}

	if err := h.chainService.UpdateUserToken(uint(tokenID), uint(userID), updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新代币失败: " + err.Error()})
		return
	}

	// 获取更新后的代币信息
	updatedToken, err := h.chainService.GetUserTokenByID(uint(tokenID), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取更新后的代币信息失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "代币更新成功", "token": updatedToken})
}

// GetUserTokens 获取用户代币列表（包含系统和用户的）
func (h *ChainHandler) GetUserTokens(c *gin.Context) {
	chainIDStr := c.Param("chain_id")
	chainID, err := strconv.Atoi(chainIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的链ID"})
		return
	}

	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		// 如果没有用户ID，返回空列表（不再有系统代币）
		c.JSON(http.StatusOK, gin.H{"tokens": []models.Token{}})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	// 只获取用户自定义代币
	userTokens, err := h.chainService.GetUserTokens(uint(userID), chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户代币失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": userTokens})
}
