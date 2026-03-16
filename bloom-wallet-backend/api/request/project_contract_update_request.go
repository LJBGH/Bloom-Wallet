package request

import (
	"time"
)

// TimestampMs 自定义时间类型，用于处理毫秒时间戳字符串
type TimestampMs time.Time

func (t TimestampMs) Time() time.Time {
	return time.Time(t)
}

type ProjectContractUpdateRequest struct {
	ID                uint        `json:"id" binding:"required"`
	SaleAddress       string      `json:"saleAddress"`
	SaleToken         string      `json:"saleToken"`
	SaleOwner         string      `json:"saleOwner"`
	TokenPriceInPT    string      `json:"tokenPriceInPT"`
	PaymentToken      string      `json:"paymentToken"`
	TotalTokens       string      `json:"totalTokens"`
	SaleEndTime       TimestampMs `json:"saleEndTime"`
	UnlockTime        TimestampMs `json:"unlockTime"`
	RegistrationStart TimestampMs `json:"registrationStart"`
	RegistrationEnd   TimestampMs `json:"registrationEnd"`
	SaleStartTime     TimestampMs `json:"saleStartTime"`
}
