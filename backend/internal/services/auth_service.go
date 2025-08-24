package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims JWT声明结构
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService struct {
	secretKey []byte
}

// NewAuthService 创建认证服务
func NewAuthService() *AuthService {
	// 在生产环境中，这个密钥应该从环境变量或配置文件中读取
	secretKey := []byte("your-secret-key-change-this-in-production")
	return &AuthService{
		secretKey: secretKey,
	}
}

// GenerateToken 生成JWT Token
func (s *AuthService) GenerateToken(userID uint, username string) (string, error) {
	// 设置过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建声明
	claims := &JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "wallet-manager",
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("生成token失败: %v", err)
	}

	return tokenString, nil
}

// ValidateToken 验证JWT Token
func (s *AuthService) ValidateToken(tokenString string) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("无效的签名方法: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token解析失败: %v", err)
	}

	// 验证token是否有效
	if !token.Valid {
		return nil, errors.New("无效的token")
	}

	// 获取声明
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("无效的token声明")
	}

	// 检查是否过期
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token已过期")
	}

	return claims, nil
}

// RefreshToken 刷新JWT Token
func (s *AuthService) RefreshToken(tokenString string) (string, error) {
	// 验证现有token
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("刷新token失败: %v", err)
	}

	// 检查token是否在1小时内过期，如果是则允许刷新
	if time.Until(claims.ExpiresAt.Time) > time.Hour {
		return "", errors.New("token还未到刷新时间")
	}

	// 生成新token
	return s.GenerateToken(claims.UserID, claims.Username)
}
