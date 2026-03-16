package services

import (
	"bloom-wallet/api/request"
	"bloom-wallet/config"
	"bloom-wallet/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SaleContractService struct {
	projectContractService *ProjectContractService
}

func NewSaleContractService(projectContractService *ProjectContractService) *SaleContractService {
	return &SaleContractService{
		projectContractService: projectContractService,
	}
}

// 合约事件常量
const (
	SaleEventCreated          = "SaleCreated(address,uint256,uint256,uint256)"
	SaleEventStartTime        = "StartTimeSet(uint256)"
	SaleEventRegistrationTime = "RegistrationTimeSet(uint256,uint256)"
	SaleDepolyed              = "SaleDeployed(address)"
)

// 查询销售合约信息
func (s *SaleContractService) QuerySaleInfo(saleAddress string) (*model.ProjectContract, error) {
	// 连接到以太坊节点
	client, err := ethclient.Dial(config.AppConfig.Sales.NetworkUrl)

	if err != nil {
		return nil, err
	}
	defer client.Close()

	// 创建合约实例
	contractAddress := common.HexToAddress(saleAddress)
	contract, err := model.NewC2NSale(contractAddress, client)
	if err != nil {
		return nil, err
	}

	// 调用合约方法
	opts := &bind.CallOpts{
		Context: context.Background(),
	}

	// 获取sale信息
	saleInfo, err := contract.Sale(opts)
	if err != nil {
		return nil, err
	}

	// 获取registration信息
	registrationInfo, err := contract.Registration(opts)
	if err != nil {
		return nil, err
	}

	// 解析数据
	projectPO := &model.ProjectContract{
		SaleContractAddress:    contractAddress.Hex(),
		TokenAddress:           saleInfo.Token.Hex(),
		TokenPriceInPT:         saleInfo.TokenPriceInETH.String(),
		TotalTokensSold:        saleInfo.AmountOfTokensToSell.String(),
		SaleEnd:                time.Unix(saleInfo.SaleEnd.Int64(), 0),
		UnlockTime:             time.Unix(saleInfo.TokensUnlockTime.Int64(), 0),
		SaleStart:              time.Unix(saleInfo.SaleStart.Int64(), 0),
		RegistrationTimeEnds:   time.Unix(registrationInfo.RegistrationTimeEnds.Int64(), 0),
		RegistrationTimeStarts: time.Unix(registrationInfo.RegistrationTimeStarts.Int64(), 0),
	}

	log.Println("获取合约信息成功")
	return projectPO, nil
}

// 启动销售工厂监听
func (s *SaleContractService) StartSaleFactoryListen() error {
	// 1.连接以太坊网络
	client, err := ethclient.Dial(config.AppConfig.Sales.NetworkUrl)
	if err != nil {
		return fmt.Errorf("连接以太坊网络失败: %w", err)
	}

	defer client.Close()

	// 创建过滤器
	saleFactoryContractAddr := common.HexToAddress(config.AppConfig.Sales.SaleFactoryAddress)

	// 计算事件签名哈希
	saleDeploy := s.CalculateEventSignature(SaleDepolyed)

	// 创建过滤查询,查询工厂合约部署代币销售合约事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{saleFactoryContractAddr},
		Topics: [][]common.Hash{
			{saleDeploy},
		},
	}

	fmt.Printf("初始化服务开启销售工厂扫描,扫描销售创建,address=%v\n", config.AppConfig.Sales.SaleFactoryAddress)

	// 3. 创建日志通道并订阅
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 监听所有销售合约
	for _, vLog := range logs {
		saleAddress, err := bytesToAddressFrom20Bytes(vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("监听到销售变更消息,saleFactoryAddress=%v,saleAddress=%v\n", saleFactoryContractAddr, saleAddress.Hex())

		// 监听销售合约变化
		err = s.ListenSaleChange(saleAddress.Hex())
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// 监听销售合约变化
func (s *SaleContractService) ListenSaleChange(saleAddress string) error {
	fmt.Printf("开始扫描合约变化,%v", saleAddress)
	// 1. 连接到以太坊节点
	client, err := ethclient.Dial(config.AppConfig.Sales.NetworkUrl)
	if err != nil {
		return fmt.Errorf("连接以太坊节点失败: %v", err)
	}
	defer client.Close()

	// 2. 创建过滤器查询
	contractAddr := common.HexToAddress(saleAddress)

	// 计算事件签名哈希
	topicCreated := s.CalculateEventSignature(SaleEventCreated)
	topicStartTime := s.CalculateEventSignature(SaleEventStartTime)
	topicRegistration := s.CalculateEventSignature(SaleEventRegistrationTime)

	// 创建过滤查询
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
		Topics: [][]common.Hash{
			{topicCreated, topicStartTime, topicRegistration},
		},
	}

	// 3. 创建日志通道并订阅
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("监听到销售变更消息,address=%v,topic0=%v\n", saleAddress, vLog.Topics[0])
		sale, err := s.QuerySaleInfo(saleAddress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("监听到销售变更消息,组装productPo=%v\n", sale)
		err = s.projectContractService.Update(&request.ProjectContractUpdateRequest{
			ID:                1,
			SaleAddress:       sale.SaleContractAddress,
			SaleToken:         sale.TokenAddress,
			TokenPriceInPT:    sale.TokenPriceInPT,
			TotalTokens:       sale.TotalTokensSold,
			SaleEndTime:       request.TimestampMs(sale.SaleEnd),
			UnlockTime:        request.TimestampMs(sale.UnlockTime),
			RegistrationStart: request.TimestampMs(sale.RegistrationTimeStarts),
			RegistrationEnd:   request.TimestampMs(sale.RegistrationTimeEnds),
			SaleStartTime:     request.TimestampMs(sale.SaleStart),
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("更新成功")
	}

	return nil
}

// 计算事件签名
func (s *SaleContractService) CalculateEventSignature(eventSignature string) common.Hash {
	hash := crypto.Keccak256Hash([]byte(eventSignature))

	return hash
}

// 从 20 个字节中提取地址值
func bytesToAddressFrom20Bytes(input []byte) (common.Address, error) {

	// 直接转换
	address := common.BytesToAddress(input)
	return address, nil
}
