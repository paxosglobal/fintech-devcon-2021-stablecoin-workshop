package server

import (
	"fmt"
	"time"
)

func (s *Server) CreateWithdrawal(req *Withdrawal) (*Activity, error) {
	newActivity := Activity{
		Type:               TypeWithdrawal,
		Asset:              req.Asset,
		Amount:             req.Amount,
		DestinationAddress: req.DestinationAddress,
		ID:                 s.nextActivityID(),
		Status:             Created,
		CreatedAt:          time.Now(),
		CompeletedAt:       nil,
	}
	s.activities = append(s.activities, newActivity)
	if s.balances.UsdOnPlatform.LessThan(newActivity.Amount) {
		return nil, fmt.Errorf("not enough funds")
	}
	s.balances.UsdOnPlatform = s.balances.UsdOnPlatform.Sub(newActivity.Amount) // debit funds from on-platform holdings
	s.balances.UsdInReserve = s.balances.UsdInReserve.Add(newActivity.Amount)   // put funds in reserve

	err := s.mint(req.DestinationAddress, req.Amount)
	if err != nil {
		s.balances.UsdOnPlatform = s.balances.UsdOnPlatform.Add(newActivity.Amount) // reverse debit funds from on-platform holdings
		s.balances.UsdInReserve = s.balances.UsdInReserve.Sub(newActivity.Amount)   // reverse put funds in reserve
		return nil, err
	}

	s.balances.UsdkMinted = s.balances.UsdkMinted.Add(newActivity.Amount) // mark as minted on chain

	return &newActivity, nil
}
