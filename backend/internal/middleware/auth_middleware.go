package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"wallet-manager/internal/services"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header中获取Authorization token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少认证token"})
			c.Abort()
			return
		}

		// 检查Bearer格式
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token格式，应为: Bearer <token>"})
			c.Abort()
			return
		}

		// 提取token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 验证token
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的token: " + err.Error()})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// UserIDMiddleware 验证URL中的user_id是否与token中的user_id匹配
func UserIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从token中获取用户ID
		tokenUserID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未找到用户认证信息"})
			c.Abort()
			return
		}

		// 从URL参数中获取user_id
		userIDParam := c.Param("user_id")
		if userIDParam == "" {
			c.Next()
			return
		}

		// 转换URL中的user_id
		urlUserID, err := strconv.ParseUint(userIDParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID格式"})
			c.Abort()
			return
		}

		// 验证是否匹配
		if uint(urlUserID) != tokenUserID.(uint) {
			c.JSON(http.StatusForbidden, gin.H{"error": "无权限访问其他用户的数据"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) (uint, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, errors.New("未找到用户认证信息")
	}

	return userID.(uint), nil
}

// GetUsernameFromContext 从上下文中获取用户名
func GetUsernameFromContext(c *gin.Context) (string, error) {
	username, exists := c.Get("username")
	if !exists {
		return "", errors.New("未找到用户认证信息")
	}

	return username.(string), nil
}
