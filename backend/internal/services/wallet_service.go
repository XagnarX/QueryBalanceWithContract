package services

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"wallet-manager/internal/database"
	"wallet-manager/internal/models"
)

type WalletService struct {
	db           *gorm.DB
	ethClient    *ethclient.Client
	contractAddr common.Address
}

func NewWalletService() *WalletService {
	// 连接BSC主网
	client, err := ethclient.Dial("https://bsc-dataseed1.binance.org/")
	if err != nil {
		panic(fmt.Sprintf("连接BSC主网失败: %v", err))
	}

	return &WalletService{
		db:           database.GetDB(),
		ethClient:    client,
		contractAddr: common.HexToAddress("0x0ab68be1431cd1E6Fd86793C6392181eb4dc636b"),
	}
}

// CreateGroup 创建钱包分组
func (s *WalletService) CreateGroup(userID uint, req *models.CreateGroupRequest) (*models.WalletGroup, error) {
	// 检查分组名是否已存在
	var existingGroup models.WalletGroup
	if err := s.db.Where("user_id = ? AND name = ?", userID, req.Name).First(&existingGroup).Error; err == nil {
		return nil, errors.New("分组名已存在")
	}

	group := &models.WalletGroup{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.db.Create(group).Error; err != nil {
		return nil, fmt.Errorf("创建分组失败: %v", err)
	}

	return group, nil
}

// GetUserGroups 获取用户的所有分组
func (s *WalletService) GetUserGroups(userID uint) ([]models.WalletGroup, error) {
	var groups []models.WalletGroup
	if err := s.db.Where("user_id = ?", userID).Preload("WalletAddresses").Find(&groups).Error; err != nil {
		return nil, fmt.Errorf("查询分组失败: %v", err)
	}
	return groups, nil
}

// AddAddress 添加钱包地址
func (s *WalletService) AddAddress(userID uint, req *models.AddAddressRequest) (*models.WalletAddress, error) {
	// 验证地址格式
	if !common.IsHexAddress(req.Address) {
		return nil, errors.New("无效的钱包地址格式")
	}

	// 标准化地址格式
	address := common.HexToAddress(req.Address).Hex()

	// 检查地址是否已存在
	var existingAddr models.WalletAddress
	if err := s.db.Where("user_id = ? AND address = ?", userID, address).First(&existingAddr).Error; err == nil {
		return nil, errors.New("地址已存在")
	}

	// 如果指定了分组，验证分组是否属于该用户
	if req.GroupID != nil {
		var group models.WalletGroup
		if err := s.db.Where("id = ? AND user_id = ?", *req.GroupID, userID).First(&group).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("指定的分组不存在")
			}
			return nil, fmt.Errorf("验证分组失败: %v", err)
		}
	}

	walletAddr := &models.WalletAddress{
		UserID:  userID,
		GroupID: req.GroupID,
		Address: address,
		Label:   req.Label,
	}

	if err := s.db.Create(walletAddr).Error; err != nil {
		return nil, fmt.Errorf("添加地址失败: %v", err)
	}

	// 预加载关联数据
	if err := s.db.Preload("Group").First(walletAddr, walletAddr.ID).Error; err != nil {
		return walletAddr, nil // 地址已创建，即使预加载失败也返回
	}

	return walletAddr, nil
}

// GetUserAddresses 获取用户的所有地址
func (s *WalletService) GetUserAddresses(userID uint) ([]models.WalletAddress, error) {
	var addresses []models.WalletAddress
	if err := s.db.Where("user_id = ?", userID).Preload("Group").Find(&addresses).Error; err != nil {
		return nil, fmt.Errorf("查询地址失败: %v", err)
	}
	return addresses, nil
}

// GetGroupAddresses 获取分组中的所有地址
func (s *WalletService) GetGroupAddresses(userID, groupID uint) ([]models.WalletAddress, error) {
	// 验证分组属于该用户
	var group models.WalletGroup
	if err := s.db.Where("id = ? AND user_id = ?", groupID, userID).First(&group).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("分组不存在")
		}
		return nil, fmt.Errorf("验证分组失败: %v", err)
	}

	var addresses []models.WalletAddress
	if err := s.db.Where("group_id = ?", groupID).Preload("Group").Find(&addresses).Error; err != nil {
		return nil, fmt.Errorf("查询分组地址失败: %v", err)
	}
	return addresses, nil
}

// DeleteAddress 删除钱包地址
func (s *WalletService) DeleteAddress(userID, addressID uint) error {
	result := s.db.Where("id = ? AND user_id = ?", addressID, userID).Delete(&models.WalletAddress{})
	if result.Error != nil {
		return fmt.Errorf("删除地址失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("地址不存在")
	}
	return nil
}

// DeleteGroup 删除分组（地址会被移出分组但不删除）
func (s *WalletService) DeleteGroup(userID, groupID uint) error {
	// 先将分组中的地址移出分组
	if err := s.db.Model(&models.WalletAddress{}).Where("group_id = ?", groupID).Update("group_id", nil).Error; err != nil {
		return fmt.Errorf("移出分组地址失败: %v", err)
	}

	// 删除分组
	result := s.db.Where("id = ? AND user_id = ?", groupID, userID).Delete(&models.WalletGroup{})
	if result.Error != nil {
		return fmt.Errorf("删除分组失败: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("分组不存在")
	}
	return nil
}

// GetTokens 获取支持的代币列表
func (s *WalletService) GetTokens() ([]models.Token, error) {
	var tokens []models.Token
	if err := s.db.Where("is_active = ?", true).Find(&tokens).Error; err != nil {
		return nil, fmt.Errorf("查询代币列表失败: %v", err)
	}
	return tokens, nil
}

// 辅助函数：格式化余额
func formatBalance(balance *big.Int, decimals int) string {
	if balance == nil {
		return "0"
	}

	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	amount := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(divisor))
	return strings.TrimRight(strings.TrimRight(amount.Text('f', 6), "0"), ".")
}
