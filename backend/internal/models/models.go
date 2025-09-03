package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"uniqueIndex;size:50;not null"`
	Email        string    `json:"email" gorm:"uniqueIndex;size:100;not null"`
	PasswordHash string    `json:"-" gorm:"size:255;not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// 关联关系
	WalletGroups    []WalletGroup   `json:"wallet_groups,omitempty" gorm:"foreignKey:UserID"`
	WalletAddresses []WalletAddress `json:"wallet_addresses,omitempty" gorm:"foreignKey:UserID"`
	UserRPCs        []UserRPCEndpoint `json:"user_rpcs,omitempty" gorm:"foreignKey:UserID"`
	UserTokens      []UserToken     `json:"user_tokens,omitempty" gorm:"foreignKey:UserID"`
}

// WalletGroup 钱包分组模型
type WalletGroup struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"not null;index"`
	Name        string    `json:"name" gorm:"size:100;not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 关联关系
	User            User            `json:"user,omitempty" gorm:"foreignKey:UserID"`
	WalletAddresses []WalletAddress `json:"wallet_addresses,omitempty" gorm:"foreignKey:GroupID"`
}

// WalletAddress 钱包地址模型
type WalletAddress struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	GroupID   *uint     `json:"group_id" gorm:"index"`
	Address   string    `json:"address" gorm:"size:42;not null;uniqueIndex:idx_user_group_address,composite:user_id,group_id"`
	Label     string    `json:"label" gorm:"size:100"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	User  User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Group *WalletGroup `json:"group,omitempty" gorm:"foreignKey:GroupID"`
}

// Token struct removed - using UserToken instead

// Chain 区块链网络模型
type Chain struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	ChainID          int       `json:"chain_id" gorm:"uniqueIndex;not null"`
	Name             string    `json:"name" gorm:"size:50;not null"`
	Symbol           string    `json:"symbol" gorm:"size:10;not null"`
	IsTestnet        bool      `json:"is_testnet" gorm:"default:false"`
	IsActive         bool      `json:"is_active" gorm:"default:true"`
	BlockExplorerURL string    `json:"block_explorer_url" gorm:"size:255"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	// 关联关系
	BalanceContracts []BalanceContract `json:"balance_contracts,omitempty" gorm:"foreignKey:ChainID;references:ChainID"`
}

// RPCEndpoint struct removed - using UserRPCEndpoint instead

// BalanceContract 余额查询合约模型
type BalanceContract struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	ChainID         int       `json:"chain_id" gorm:"not null;index"`
	Name            string    `json:"name" gorm:"size:50;not null"`
	ContractAddress string    `json:"contract_address" gorm:"size:42;not null"`
	IsActive        bool      `json:"is_active" gorm:"default:true"`
	Priority        int       `json:"priority" gorm:"default:0;index"`
	ABIJSON         string    `json:"abi_json,omitempty" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// 关联关系
	Chain Chain `json:"chain,omitempty" gorm:"foreignKey:ChainID;references:ChainID"`
}

// UserRPCEndpoint 用户自定义RPC端点模型
type UserRPCEndpoint struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	ChainID   int       `json:"chain_id" gorm:"not null;index"`
	Name      string    `json:"name" gorm:"size:50;not null"`
	URL       string    `json:"url" gorm:"size:255;not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	User  User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Chain Chain `json:"chain,omitempty" gorm:"foreignKey:ChainID;references:ChainID"`
}

// UserToken 用户自定义代币模型
type UserToken struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null;index"`
	ChainID         int       `json:"chain_id" gorm:"not null;index"`
	Symbol          string    `json:"symbol" gorm:"size:10;not null"`
	Name            string    `json:"name" gorm:"size:50;not null"`
	ContractAddress string    `json:"contract_address" gorm:"size:42;not null"`
	Decimals        int       `json:"decimals" gorm:"default:18;not null"`
	IsActive        bool      `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	// 关联关系
	User  User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Chain Chain `json:"chain,omitempty" gorm:"foreignKey:ChainID;references:ChainID"`
}

// 请求/响应结构体

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateGroupRequest 创建分组请求
type CreateGroupRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Description string `json:"description"`
}

// AddAddressRequest 添加地址请求
type AddAddressRequest struct {
	Address string `json:"address" binding:"required,len=42"`
	Label   string `json:"label"`
	GroupID *uint  `json:"group_id"`
}

// CreateUserRPCRequest 创建用户RPC请求
type CreateUserRPCRequest struct {
	ChainID int    `json:"chain_id" binding:"required"`
	Name    string `json:"name" binding:"required,min=1,max=50"`
	URL     string `json:"url" binding:"required,url"`
}

// UpdateUserRPCRequest 更新用户RPC请求
type UpdateUserRPCRequest struct {
	ChainID  int    `json:"chain_id" binding:"required"`
	Name     string `json:"name" binding:"required,min=1,max=50"`
	URL      string `json:"url" binding:"required,url"`
	IsActive bool   `json:"is_active"`
}

// CreateUserTokenRequest 创建用户代币请求
type CreateUserTokenRequest struct {
	ChainID         int    `json:"chain_id" binding:"required"`
	Symbol          string `json:"symbol" binding:"required,min=1,max=10"`
	Name            string `json:"name" binding:"required,min=1,max=50"`
	ContractAddress string `json:"contract_address" binding:"required,len=42"`
	Decimals        int    `json:"decimals" binding:"required,min=0,max=18"`
}

// UpdateUserTokenRequest 更新用户代币请求
type UpdateUserTokenRequest struct {
	ChainID         int    `json:"chain_id" binding:"required"`
	Symbol          string `json:"symbol" binding:"required,min=1,max=10"`
	Name            string `json:"name" binding:"required,min=1,max=50"`
	ContractAddress string `json:"contract_address" binding:"required,len=42"`
	Decimals        int    `json:"decimals" binding:"required,min=0,max=18"`
	IsActive        bool   `json:"is_active"`
}

// AddressBalanceResponse 地址余额响应
type AddressBalanceResponse struct {
	Address       string                 `json:"address"`
	Label         string                 `json:"label"`
	BNBBalance    string                 `json:"bnb_balance"`
	TokenBalances []TokenBalanceResponse `json:"token_balances"`
}

// TokenBalanceResponse 代币余额响应
type TokenBalanceResponse struct {
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	ContractAddress string `json:"contract_address"`
	Balance         string `json:"balance"`
	Decimals        int    `json:"decimals"`
}

// GroupBalanceResponse 分组余额响应
type GroupBalanceResponse struct {
	GroupID     uint                     `json:"group_id"`
	GroupName   string                   `json:"group_name"`
	TotalBNB    string                   `json:"total_bnb"`
	Addresses   []AddressBalanceResponse `json:"addresses"`
	TokenTotals []TokenBalanceResponse   `json:"token_totals"`
}

// ChainInfoResponse 区块链信息响应
type ChainInfoResponse struct {
	ChainID          int               `json:"chain_id"`
	Name             string            `json:"name"`
	Symbol           string            `json:"symbol"`
	IsTestnet        bool              `json:"is_testnet"`
	IsActive         bool              `json:"is_active"`
	BlockExplorerURL string            `json:"block_explorer_url"`
	RPCEndpoints     []RPCEndpoint     `json:"rpc_endpoints"`
	BalanceContracts []BalanceContract `json:"balance_contracts"`
	Tokens           []Token           `json:"tokens"`
}

// ChainListResponse 区块链列表响应
type ChainListResponse struct {
	Chains []ChainInfoResponse `json:"chains"`
}

// 表名设置
func (User) TableName() string {
	return "users"
}

func (WalletGroup) TableName() string {
	return "wallet_groups"
}

func (WalletAddress) TableName() string {
	return "wallet_addresses"
}

// Token TableName removed

func (Chain) TableName() string {
	return "chains"
}

// RPCEndpoint TableName removed

func (BalanceContract) TableName() string {
	return "balance_contracts"
}

func (UserRPCEndpoint) TableName() string {
	return "user_rpc_endpoints"
}

func (UserToken) TableName() string {
	return "user_tokens"
}

// Compatibility structures for frontend
// Token compatible structure for API responses
type Token struct {
	ID              uint   `json:"id"`
	Symbol          string `json:"symbol"`
	Name            string `json:"name"`
	ContractAddress string `json:"contract_address"`
	Decimals        int    `json:"decimals"`
	ChainID         int    `json:"chain_id"`
	IsActive        bool   `json:"is_active"`
}

// RPCEndpoint compatible structure for API responses
type RPCEndpoint struct {
	ID       uint   `json:"id"`
	ChainID  int    `json:"chain_id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	IsActive bool   `json:"is_active"`
	Priority int    `json:"priority"`
}

// Convert UserToken to Token for API response
func (ut UserToken) ToToken() Token {
	return Token{
		ID:              ut.ID,
		Symbol:          ut.Symbol,
		Name:            ut.Name,
		ContractAddress: ut.ContractAddress,
		Decimals:        ut.Decimals,
		ChainID:         ut.ChainID,
		IsActive:        ut.IsActive,
	}
}

// Convert UserRPCEndpoint to RPCEndpoint for API response
func (ur UserRPCEndpoint) ToRPCEndpoint() RPCEndpoint {
	return RPCEndpoint{
		ID:       ur.ID,
		ChainID:  ur.ChainID,
		Name:     ur.Name,
		URL:      ur.URL,
		IsActive: ur.IsActive,
		Priority: 0, // User RPCs have default priority
	}
}

// GORM Hooks
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

func (wg *WalletGroup) BeforeCreate(tx *gorm.DB) error {
	wg.CreatedAt = time.Now()
	wg.UpdatedAt = time.Now()
	return nil
}

func (wg *WalletGroup) BeforeUpdate(tx *gorm.DB) error {
	wg.UpdatedAt = time.Now()
	return nil
}

func (wa *WalletAddress) BeforeCreate(tx *gorm.DB) error {
	wa.CreatedAt = time.Now()
	wa.UpdatedAt = time.Now()
	return nil
}

func (wa *WalletAddress) BeforeUpdate(tx *gorm.DB) error {
	wa.UpdatedAt = time.Now()
	return nil
}
