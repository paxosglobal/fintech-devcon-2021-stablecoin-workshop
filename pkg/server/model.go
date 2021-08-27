package server

import "github.com/shopspring/decimal"

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
	TypeDeposit    Status = "DEPOSIT"
	TypeWithdrawal Status = "WITHDRAWAL"
)

type Activity struct {
	Type               ActivityType
	Asset              string          `json:"asset"`
	Amount             decimal.Decimal `json:"amount"`
	DestinationAddress string          `json:"destinationAddress"`
	ID                 uint64          `json:"id"`
	ChainID            string          `json:"chainId"` // the id on the blockchain
	Status             Status          `json:"status"`  // CREATED|COMPLETED
}
