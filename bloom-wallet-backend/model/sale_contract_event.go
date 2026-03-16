package model

import (
	"time"
)

type ProjectContractEvent struct {
	ID             uint      `json:"id" gorm:"primarykey"`
	SaleContractId uint      `json:"saleContractId" gorm:"type:int(4);default:0;not null"`
	EventName      string    `json:"eventName" gorm:"type:varchar(80);not null"` // 事件名称
	Content        string    `json:"content" gorm:"type:varchar(2000);"`
	CreateTime     time.Time `json:"createTime" gorm:"type:datetime; not null"` // 创建时间
	UpdateTime     time.Time `json:"updateTime" gorm:"type:datetime; not null"` // 更新时间
}

// TableName 指定表名
func (ProjectContractEvent) TableName() string {
	return "project_contract_event"
}
