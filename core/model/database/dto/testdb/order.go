package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	OID        int             `json:"o_id"`
	UserID     string          `json:"user_id"`
	OrderTime  time.Time       `json:"order_time"`
	Money      decimal.Decimal `json:"money"`
	Remark     string          `json:"remark"`
	UpdateTime time.Time       `json:"update_time"`
	CreateTime time.Time       `json:"create_time"`
}
