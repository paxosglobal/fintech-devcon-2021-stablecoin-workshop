package server

import (
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
	s.mu.Lock()
	s.activities = append(s.activities, newActivity)
	s.balances.UsdOnPlatform = s.balances.UsdOnPlatform.Sub(newActivity.Amount) // debut funds from on-platform holdings
	s.balances.UsdInReserve = s.balances.UsdInReserve.Sub(newActivity.Amount)   // put funds in reserve
	s.balances.UsdkMinted = s.balances.UsdkMinted.Add(newActivity.Amount)       // mark as minted on chain
	s.mu.Unlock()

	// TODO: implement kicking off an actual blockchain withdrawal

	return &newActivity, nil
}
