package services

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"wallet-manager/internal/database"
	"wallet-manager/internal/models"
)

type GroupSettingsService struct {
	db *gorm.DB
}

func NewGroupSettingsService() *GroupSettingsService {
	return &GroupSettingsService{
		db: database.GetDB(),
	}
}

// GetGroupSettings retrieves settings for a specific group and chain
func (s *GroupSettingsService) GetGroupSettings(userID, groupID uint, chainID int) (*models.WalletGroupSettings, error) {
	// Verify group belongs to user
	var group models.WalletGroup
	if err := s.db.Where("id = ? AND user_id = ?", groupID, userID).First(&group).Error; err != nil {
		return nil, fmt.Errorf("分组不存在或无权访问")
	}

	var settings models.WalletGroupSettings
	err := s.db.Where("user_id = ? AND group_id = ? AND chain_id = ?", userID, groupID, chainID).First(&settings).Error

	if err != nil {
		// If settings don't exist, return default settings
		if err.Error() == "record not found" {
			return &models.WalletGroupSettings{
				UserID:            userID,
				GroupID:           groupID,
				ChainID:           chainID,
				CountdownEnabled:  false,
				CountdownDuration: 600,
				SelectedRPCID:     nil,
				SelectedTokenIDs:  "[]",
			}, nil
		}
		return nil, err
	}

	return &settings, nil
}

// GetAllGroupSettings retrieves settings for all groups of a user and a specific chain
func (s *GroupSettingsService) GetAllGroupSettings(userID uint, chainID int) ([]models.WalletGroupSettings, error) {
	var settings []models.WalletGroupSettings

	// Get all groups for the user
	var groups []models.WalletGroup
	if err := s.db.Where("user_id = ?", userID).Find(&groups).Error; err != nil {
		return nil, err
	}

	// Get settings for each group and the specified chain
	for _, group := range groups {
		var setting models.WalletGroupSettings
		err := s.db.Where("user_id = ? AND group_id = ? AND chain_id = ?", userID, group.ID, chainID).First(&setting).Error

		if err != nil {
			// If settings don't exist, create default settings
			if err.Error() == "record not found" {
				setting = models.WalletGroupSettings{
					UserID:            userID,
					GroupID:           group.ID,
					ChainID:           chainID,
					CountdownEnabled:  false,
					CountdownDuration: 600,
					SelectedRPCID:     nil,
					SelectedTokenIDs:  "[]",
				}
			} else {
				return nil, err
			}
		}

		settings = append(settings, setting)
	}

	return settings, nil
}

// UpdateGroupSettings updates or creates settings for a group
func (s *GroupSettingsService) UpdateGroupSettings(userID, groupID uint, req *models.UpdateGroupSettingsRequest) (*models.WalletGroupSettings, error) {
	// Verify group belongs to user
	var group models.WalletGroup
	if err := s.db.Where("id = ? AND user_id = ?", groupID, userID).First(&group).Error; err != nil {
		return nil, fmt.Errorf("分组不存在或无权访问")
	}

	// Verify RPC endpoint if provided
	if req.SelectedRPCID != nil && *req.SelectedRPCID > 0 {
		var rpc models.UserRPCEndpoint
		if err := s.db.Where("id = ? AND user_id = ?", *req.SelectedRPCID, userID).First(&rpc).Error; err != nil {
			return nil, fmt.Errorf("RPC端点不存在或无权访问")
		}
	}

	// Verify tokens if provided
	if len(req.SelectedTokenIDs) > 0 {
		var tokens []models.UserToken
		if err := s.db.Where("id IN ? AND user_id = ?", req.SelectedTokenIDs, userID).Find(&tokens).Error; err != nil {
			return nil, fmt.Errorf("Token查询失败")
		}
		if len(tokens) != len(req.SelectedTokenIDs) {
			return nil, fmt.Errorf("部分Token不存在或无权访问")
		}
	}

	// Convert token IDs to JSON
	tokenIDsJSON, err := json.Marshal(req.SelectedTokenIDs)
	if err != nil {
		return nil, fmt.Errorf("Token ID序列化失败")
	}

	// Try to find existing settings for this chain
	var settings models.WalletGroupSettings
	err = s.db.Where("user_id = ? AND group_id = ? AND chain_id = ?", userID, groupID, req.ChainID).First(&settings).Error

	if err != nil {
		// Create new settings if not found
		if err.Error() == "record not found" {
			settings = models.WalletGroupSettings{
				UserID:            userID,
				GroupID:           groupID,
				ChainID:           req.ChainID,
				CountdownEnabled:  req.CountdownEnabled,
				CountdownDuration: req.CountdownDuration,
				SelectedRPCID:     req.SelectedRPCID,
				SelectedTokenIDs:  string(tokenIDsJSON),
			}

			if err := s.db.Create(&settings).Error; err != nil {
				return nil, fmt.Errorf("创建配置失败: %v", err)
			}

			return &settings, nil
		}
		return nil, err
	}

	// Update existing settings
	settings.CountdownEnabled = req.CountdownEnabled
	settings.CountdownDuration = req.CountdownDuration
	settings.SelectedRPCID = req.SelectedRPCID
	settings.SelectedTokenIDs = string(tokenIDsJSON)

	if err := s.db.Save(&settings).Error; err != nil {
		return nil, fmt.Errorf("更新配置失败: %v", err)
	}

	return &settings, nil
}

// DeleteGroupSettings deletes settings for a group and chain
func (s *GroupSettingsService) DeleteGroupSettings(userID, groupID uint, chainID int) error {
	// Verify group belongs to user
	var group models.WalletGroup
	if err := s.db.Where("id = ? AND user_id = ?", groupID, userID).First(&group).Error; err != nil {
		return fmt.Errorf("分组不存在或无权访问")
	}

	result := s.db.Where("user_id = ? AND group_id = ? AND chain_id = ?", userID, groupID, chainID).Delete(&models.WalletGroupSettings{})
	if result.Error != nil {
		return fmt.Errorf("删除配置失败: %v", result.Error)
	}

	return nil
}

// ConvertToResponse converts WalletGroupSettings to GroupSettingsResponse
func (s *GroupSettingsService) ConvertToResponse(settings *models.WalletGroupSettings) (*models.GroupSettingsResponse, error) {
	var tokenIDs []uint
	if err := json.Unmarshal([]byte(settings.SelectedTokenIDs), &tokenIDs); err != nil {
		tokenIDs = []uint{}
	}

	return &models.GroupSettingsResponse{
		ID:                settings.ID,
		GroupID:           settings.GroupID,
		UserID:            settings.UserID,
		ChainID:           settings.ChainID,
		CountdownEnabled:  settings.CountdownEnabled,
		CountdownDuration: settings.CountdownDuration,
		SelectedRPCID:     settings.SelectedRPCID,
		SelectedTokenIDs:  tokenIDs,
		CreatedAt:         settings.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:         settings.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
