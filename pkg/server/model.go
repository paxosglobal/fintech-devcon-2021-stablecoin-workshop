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
	// forbidden in request
	ID      uint64 `json:"id"`
	ChainID string `json:"chainId"`
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
	// forbidden in request
	ID      uint64 `json:"id"`
	ChainID string `json:"chainId"` // the id on the blockchain
	Status  Status `json:"status"`  // CREATED|COMPLETED
}
