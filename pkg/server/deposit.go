package server

import (
	"time"
)

func (s *Server) CreateDeposit(req *Deposit) (*Activity, error) {
	newActivity := Activity{
		Type:               TypeDeposit,
		Asset:              req.Asset,
		Amount:             req.Amount,
		DestinationAddress: "",
		ID:                 s.nextActivityID(),
		Status:             Created,
		CreatedAt:          time.Now(),
		CompeletedAt:       nil,
	}
	s.mu.Lock()
	s.activities = append(s.activities, newActivity)
	s.balances.UsdOnPlatform = s.balances.UsdOnPlatform.Add(newActivity.Amount)
	s.mu.Unlock()
	return &newActivity, nil
}
