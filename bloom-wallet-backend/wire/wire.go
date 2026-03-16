//go:build wireinject
// +build wireinject

package wire

import (
	"bloom-wallet/api/handler"
	"bloom-wallet/database"
	"bloom-wallet/repository"
	"bloom-wallet/router"
	"bloom-wallet/services"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// InitializeApp 由 Wire 生成实现，返回组装好的 *gin.Engine（具体实现在 wire_gen.go）
func InitializeApp() (*gin.Engine, error) {
	wire.Build(
		// Database 层
		database.ProvideDB,

		// Repository 层 依赖 Database 层
		repository.NewProjectContractRepository,

		// Service 层 依赖 Repository 层
		services.NewTestService,
		services.NewLoginService,
		services.NewProjectContractService,
		services.NewWalletService,

		// Handler 层（依赖上面的 Service，Wire 会自动注入）
		handler.NewTestHandler,
		handler.NewLoginHandler,
		handler.NewProjectContractHandler,
		handler.NewWalletHandler,
		// 组装路由
		router.SetupRouter,
	)
	return nil, nil
}
