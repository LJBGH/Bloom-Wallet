package test

import (
	"bloom-wallet/config"
	"bloom-wallet/database"
	"bloom-wallet/repository"
	"bloom-wallet/services"
	"testing"

	"github.com/joho/godotenv"
)

// 测试监听销售合约工厂
func TestListenSaleFactory(t *testing.T) {

	// 加载 .env 文件中的环境变量
	godotenv.Load()
	// 加载配置文件
	config.LoadConfig()
	// 初始化数据库
	db := database.ProvideDB()
	// 实例化 repository
	projectContractRepo := repository.NewProjectContractRepository(db)
	// 实例化 service
	peojectContractService := services.NewProjectContractService(projectContractRepo)

	saleContractService := services.NewSaleContractService(peojectContractService)

	saleContractService.StartSaleFactoryListen()
}
