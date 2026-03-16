package model

import (
	"time"
)

type ProjectContract struct {
	ID                     uint      `json:"id" gorm:"primarykey"`
	Name                   string    `json:"name" gorm:"type:varchar(80);not null"`                                 // 项目名称
	Description            string    `json:"description" gorm:"type:longtext"`                                      // 项目说明
	Img                    string    `json:"img" gorm:"type:varchar(500)"`                                          // 项目图标
	ProjectWebsite         string    `json:"projectWebsite" gorm:"type:varchar(500)"`                               // 项目网站
	Status                 uint      `json:"status" gorm:"type:int(4);default:0;not null"`                          // 状态
	Amount                 string    `json:"amount" gorm:"type:varchar(40)"`                                        // 金额
	SaleContractAddress    string    `json:"saleContractAddress" gorm:"type:varchar(42)"`                           // 合约地址
	TokenAddress           string    `json:"tokenAddress" gorm:"type:varchar(42)"`                                  // 代币地址
	TokenName              string    `json:"tokenName" gorm:"type:varchar(20)"`                                     // 代币名称
	Symbol                 string    `json:"symbol" gorm:"type:varchar(60)"`                                        // 代币缩写
	Decimals               uint      `json:"decimals" gorm:"type:int(8);default:8"`                                 // 代币单位
	Tge                    time.Time `json:"tge" gorm:"type:datetime"`                                              // 代币解锁时间
	CurrentPrice           uint      `json:"currentPrice" gorm:"type:bigint(12)"`                                   // 当前价格
	RegistrationTimeStarts time.Time `json:"registrationTimeStarts" gorm:"type:datetime"`                           // 注册开始时间
	RegistrationTimeEnds   time.Time `json:"registrationTimeEnds" gorm:"type:datetime"`                             // 注册结束时间
	NumberOfRegistrations  uint      `json:"numberOfRegistrations" gorm:"type:int(8);column:number_of_registrants"` // 注册人数
	SaleStart              time.Time `json:"saleStart" gorm:"type:datetime"`                                        // 销售开始时间
	SaleEnd                time.Time `json:"saleEnd" gorm:"type:datetime"`                                          // 销售结束时间
	TokenPriceInPT         string    `json:"tokenPriceInPT" gorm:"type:varchar(40);column:token_price_in_PT"`       // 代币价格
	MaxParticipation       string    `json:"maxParticipation" gorm:"type:varchar(40)"`                              // 单人最大可购代币数量
	TotalTokensSold        string    `json:"totalTokensSold" gorm:"type:varchar(40)"`                               // 已卖出的代币数量
	AmountOfTokensToSell   string    `json:"amountOfTokensToSell" gorm:"type:varchar(60)"`                          // 计划出售的代币数量
	TotalRaised            string    `json:"totalRaised" gorm:"type:varchar(60)"`                                   // 筹集总额
	EarningsWithdrawn      bool      `json:"earningsWithdrawn" gorm:"default:true"`                                 // 是否提取收益
	LeftoverWithdrawn      bool      `json:"leftoverWithdrawn" gorm:"default:true"`                                 // 是否提取剩余代币
	UnlockTime             time.Time `json:"unlockTime" gorm:"type:datetime"`                                       // 用户可领取代币的最早时间
	CreateTime             time.Time `json:"createTime" gorm:"type:datetime; not null"`                             // 创建时间
	UpdateTime             time.Time `json:"updateTime" gorm:"type:datetime; not null"`                             // 更新时间
}

// TableName 指定表名
func (ProjectContract) TableName() string {
	return "project_contract"
}
