package dto

import (
	"time"

	"github.com/shopspring/decimal"
)

type Asset struct {
	ItemName   string          `json:"item_name"`
	Color      string          `json:"color"`
	Size       string          `json:"size"`
	Price      decimal.Decimal `json:"price"`
	ExpiryDate time.Time       `json:"expiry_date"`
}
