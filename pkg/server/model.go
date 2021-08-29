package server

import (
	"github.com/shopspring/decimal"
	"time"
)

type Balances struct {
	UsdOnPlatform decimal.Decimal `json:"usdOnPlatform"`
	UsdInReserve  decimal.Decimal `json:"usdInReserve"`
	UsdkMinted    decimal.Decimal `json:"usdkMinted"`
}

type Deposit struct {
	Asset  string          `json:"asset"`
	Amount decimal.Decimal `json:"amount"`
}

type Status string

const (
	Created   Status = "CREATED"
	Completed Status = "COMPLETED"
)

type Withdrawal struct {
	Asset              string          `json:"asset"`
	Amount             decimal.Decimal `json:"amount"`
	DestinationAddress string          `json:"destinationAddress"`
}

type ActivityType string

const (
	TypeDeposit    ActivityType = "DEPOSIT"
	TypeWithdrawal ActivityType = "WITHDRAWAL"
)

type Activity struct {
	Type               ActivityType    `json:"type"`
	Asset              string          `json:"asset"`
	Amount             decimal.Decimal `json:"amount"`
	DestinationAddress string          `json:"destinationAddress,omitempty"`
	ID                 uint64          `json:"id"`
	ChainTxID          string          `json:"chainId,omitempty"` // the id on the blockchain
	Status             Status          `json:"status"`            // CREATED|COMPLETED
	CreatedAt          time.Time       `json:"createdAt"`
	CompeletedAt       *time.Time      `json:"compeletedAt,omitempty"`
}
