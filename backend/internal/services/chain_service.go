package services

import (
	"wallet-manager/internal/models"

	"gorm.io/gorm"
)

type ChainService struct {
	db *gorm.DB
}

func NewChainService(db *gorm.DB) *ChainService {
	return &ChainService{db: db}
}

// GetActiveChains 获取所有活跃的区块链网络
func (s *ChainService) GetActiveChains() ([]models.Chain, error) {
	var chains []models.Chain
	err := s.db.Where("is_active = ?", true).
		Preload("BalanceContracts", "is_active = ?", true).
		Order("chain_id ASC").
		Find(&chains).Error
	return chains, err
}

// GetChainByID 根据链ID获取区块链信息
func (s *ChainService) GetChainByID(chainID int) (*models.Chain, error) {
	var chain models.Chain
	err := s.db.Where("chain_id = ? AND is_active = ?", chainID, true).
		Preload("BalanceContracts", "is_active = ?", true).
		First(&chain).Error
	if err != nil {
		return nil, err
	}
	return &chain, nil
}

// GetRPCEndpoints removed - use GetUserRPCs instead

// GetBalanceContracts 获取指定链的余额查询合约
func (s *ChainService) GetBalanceContracts(chainID int) ([]models.BalanceContract, error) {
	var contracts []models.BalanceContract
	err := s.db.Where("chain_id = ? AND is_active = ?", chainID, true).
		Order("priority ASC").
		Find(&contracts).Error
	return contracts, err
}

// GetTokensByChain removed - use GetUserTokens instead

// AddRPCEndpoint removed - use CreateUserRPC instead

// AddBalanceContract 添加余额查询合约
func (s *ChainService) AddBalanceContract(contract *models.BalanceContract) error {
	return s.db.Create(contract).Error
}

// AddToken removed - use CreateUserToken instead

// UpdateToken removed - use UpdateUserToken instead

// GetTokenByID removed - use GetUserTokenByID instead

// UpdateRPCEndpointStatus removed - use UpdateUserRPC instead

// UpdateBalanceContractStatus 更新余额查询合约状态
func (s *ChainService) UpdateBalanceContractStatus(id uint, isActive bool) error {
	return s.db.Model(&models.BalanceContract{}).Where("id = ?", id).Update("is_active", isActive).Error
}

// UpdateTokenStatus removed - use UpdateUserToken instead

// 用户RPC端点管理
// CreateUserRPC 创建用户自定义RPC端点
func (s *ChainService) CreateUserRPC(userRPC *models.UserRPCEndpoint) error {
	return s.db.Create(userRPC).Error
}

// GetUserRPCs 获取用户自定义RPC端点列表
func (s *ChainService) GetUserRPCs(userID uint, chainID int) ([]models.RPCEndpoint, error) {
	var userRPCs []models.UserRPCEndpoint
	err := s.db.Where("user_id = ? AND chain_id = ? AND is_active = ?", userID, chainID, true).
		Find(&userRPCs).Error
	if err != nil {
		return nil, err
	}

	// 转换为标准RPC端点格式
	var endpoints []models.RPCEndpoint
	for _, userRPC := range userRPCs {
		endpoints = append(endpoints, userRPC.ToRPCEndpoint())
	}
	return endpoints, nil
}

// GetUserRPCByID 根据ID获取用户RPC端点
func (s *ChainService) GetUserRPCByID(rpcID uint, userID uint) (*models.UserRPCEndpoint, error) {
	var userRPC models.UserRPCEndpoint
	err := s.db.Where("id = ? AND user_id = ?", rpcID, userID).First(&userRPC).Error
	if err != nil {
		return nil, err
	}
	return &userRPC, nil
}

// UpdateUserRPC 更新用户RPC端点
func (s *ChainService) UpdateUserRPC(rpcID uint, userID uint, updateData *models.UserRPCEndpoint) error {
	return s.db.Model(&models.UserRPCEndpoint{}).
		Where("id = ? AND user_id = ?", rpcID, userID).
		Updates(updateData).Error
}

// DeleteUserRPC 删除用户RPC端点
func (s *ChainService) DeleteUserRPC(rpcID uint, userID uint) error {
	return s.db.Where("id = ? AND user_id = ?", rpcID, userID).Delete(&models.UserRPCEndpoint{}).Error
}

// 用户代币管理
// CreateUserToken 创建用户自定义代币
func (s *ChainService) CreateUserToken(userToken *models.UserToken) error {
	return s.db.Create(userToken).Error
}

// GetUserTokens 获取用户自定义代币列表
func (s *ChainService) GetUserTokens(userID uint, chainID int) ([]models.Token, error) {
	var userTokens []models.UserToken
	err := s.db.Where("user_id = ? AND chain_id = ? AND is_active = ?", userID, chainID, true).
		Find(&userTokens).Error
	if err != nil {
		return nil, err
	}

	// 转换为标准代币格式
	var tokens []models.Token
	for _, userToken := range userTokens {
		tokens = append(tokens, userToken.ToToken())
	}
	return tokens, nil
}

// GetUserTokenByID 根据ID获取用户代币
func (s *ChainService) GetUserTokenByID(tokenID uint, userID uint) (*models.UserToken, error) {
	var userToken models.UserToken
	err := s.db.Where("id = ? AND user_id = ?", tokenID, userID).First(&userToken).Error
	if err != nil {
		return nil, err
	}
	return &userToken, nil
}

// UpdateUserToken 更新用户代币
func (s *ChainService) UpdateUserToken(tokenID uint, userID uint, updateData *models.UserToken) error {
	return s.db.Model(&models.UserToken{}).
		Where("id = ? AND user_id = ?", tokenID, userID).
		Updates(updateData).Error
}

// DeleteUserToken 删除用户代币
func (s *ChainService) DeleteUserToken(tokenID uint, userID uint) error {
	return s.db.Where("id = ? AND user_id = ?", tokenID, userID).Delete(&models.UserToken{}).Error
}
