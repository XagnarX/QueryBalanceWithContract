package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"wallet-manager/internal/database"
	"wallet-manager/internal/handlers"
	"wallet-manager/internal/middleware"
	"wallet-manager/internal/services"
)

func main() {
	// 连接数据库
	config := database.GetDefaultConfig()
	if err := database.Connect(config); err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer database.Close()

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode) // 生产环境使用，开发时可以改为gin.DebugMode

	// 创建路由
	router := gin.Default()

	// 添加CORS中间件
	router.Use(CORSMiddleware())

	// 创建服务和处理器
	authService := services.NewAuthService()
	userHandler := handlers.NewUserHandler()
	walletHandler := handlers.NewWalletHandler()
	chainHandler := handlers.NewChainHandler(database.GetDB())

	// API路由组
	api := router.Group("/api")
	{
		// 公共路由（不需要认证）
		api.POST("/users", userHandler.CreateUser)  // 创建用户
		api.POST("/users/login", userHandler.Login) // 用户登录

		// 区块链和代币相关公共路由
		api.GET("/chains", chainHandler.GetChains)                               // 获取区块链列表
		api.GET("/chains/:chain_id", chainHandler.GetChainInfo)                  // 获取链信息
		// api.GET("/chains/:chain_id/rpc", chainHandler.GetRPCEndpoints)        // 删除 - 使用用户RPC
		api.GET("/chains/:chain_id/contracts", chainHandler.GetBalanceContracts) // 获取余额查询合约
		// api.GET("/chains/:chain_id/tokens", chainHandler.GetTokens)           // 删除 - 使用用户Token
		api.GET("/chains/:chain_id/rpc-endpoints", chainHandler.GetUserRPCs)     // 获取用户RPC端点
		api.GET("/chains/:chain_id/tokens", chainHandler.GetUserTokens)          // 获取用户Token

		// 需要认证的路由
		protected := api.Group("/users")
		protected.Use(middleware.AuthMiddleware(authService))
		protected.Use(middleware.UserIDMiddleware())
		{
			protected.GET("/:user_id", userHandler.GetProfile) // 获取用户信息

			// 钱包分组相关路由
			protected.POST("/:user_id/groups", walletHandler.CreateGroup)                          // 创建分组
			protected.GET("/:user_id/groups", walletHandler.GetGroups)                             // 获取分组列表
			protected.GET("/:user_id/groups/:group_id/addresses", walletHandler.GetGroupAddresses) // 获取分组地址
			protected.DELETE("/:user_id/groups/:group_id", walletHandler.DeleteGroup)              // 删除分组

			// 钱包地址相关路由
			protected.POST("/:user_id/addresses", walletHandler.AddAddress)                  // 添加地址
			protected.GET("/:user_id/addresses", walletHandler.GetAddresses)                 // 获取地址列表
			protected.DELETE("/:user_id/addresses/:address_id", walletHandler.DeleteAddress) // 删除地址

			// 用户代币管理路由
			protected.POST("/:user_id/tokens", chainHandler.CreateUserToken)                    // 用户添加代币
			protected.PUT("/:user_id/tokens/:token_id", chainHandler.UpdateUserToken)        // 用户更新代币
			
			// 用户RPC管理路由
			protected.POST("/:user_id/rpc-endpoints", chainHandler.CreateUserRPC)                    // 用户添加RPC端点
			protected.PUT("/:user_id/rpc-endpoints/:rpc_id", chainHandler.UpdateUserRPC)        // 用户更新RPC端点

			// 管理员路由（添加链相关资源）
			admin := protected.Group("/:user_id/admin")
			{
				// admin.POST("/rpc", chainHandler.AddRPCEndpoint)        // 删除 - 使用用户RPC
				admin.POST("/contracts", chainHandler.AddBalanceContract) // 添加余额查询合约
				// admin.POST("/tokens", chainHandler.AddToken)           // 删除 - 使用用户Token
			}
		}
	}

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "钱包管理服务运行正常",
		})
	})

	// 根路径
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "钱包地址管理API服务",
			"version": "1.0.0",
			"docs":    "访问 /api 获取API文档",
		})
	})

	// 启动服务器
	port := ":8888"
	log.Printf("服务器启动在端口 %s", port)
	log.Printf("访问 http://localhost%s 查看服务状态", port)
	log.Printf("API文档: http://localhost%s/api", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}

// CORSMiddleware CORS中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
