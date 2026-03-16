package router

import (
	"bloom-wallet/api/handler"
	docs "bloom-wallet/docs"

	"bloom-wallet/middleware"
	exception "bloom-wallet/middleware/exception"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 初始化所有路由
func SetupRouter(
	testHandler *handler.TestHandler,
	loginHandler *handler.LoginHandler,
	projectContractHandler *handler.ProjectContractHandler,
	walletHandler *handler.WalletHandler) *gin.Engine {
	r := gin.New()

	// 使用自定义的 Recovery 中间件
	r.Use(exception.CustomRecovery())

	// 记录接口请求日志
	r.Use(gin.Logger())

	// 全局使用跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 配置 Swagger 基础路径
	docs.SwaggerInfo.BasePath = "/api"

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 业务路由
	// 业务路由加统一前缀
	api := r.Group("/api")
	{
		api.GET("/test/ping", testHandler.Ping)
		api.GET("/test/time", testHandler.GetTime)

		api.POST("/auth/login", loginHandler.Login)

		api.GET("/project/:id", projectContractHandler.GetById)
		api.GET("/project/all", projectContractHandler.GetAll)
		api.PUT("/project/:id", projectContractHandler.Update)

		// 钱包
		api.GET("/wallet/create/mnemonic", walletHandler.CreateMnemonic)
		api.GET("/wallet/create/wallet", walletHandler.CreateWallet)
		api.POST("/wallet/import", walletHandler.Import)
		api.GET("/wallet/backup", walletHandler.Backup)
		api.GET("/wallet/restore", walletHandler.Restore)
		api.GET("/wallet/restore/privateKey", walletHandler.RestoreWithPrivateKey)
		api.GET("/wallet/restore/mnemonic", walletHandler.RestoreWithMnemonic)

		// 账户
		api.GET("/wallet/account/current", walletHandler.CurrentAccount)
		api.GET("/wallet/account/all", walletHandler.AllAccounts)
		api.GET("/wallet/account/switch", walletHandler.SwitchCurrentAccount)
	}
	return r
}
