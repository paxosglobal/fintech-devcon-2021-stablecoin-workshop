package server

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"sync"
)

// nothing much to see in server.go. Just basic rpc to make the UI work with minimal code

type Server struct {
	// dependencies
	// TODO: blockchain connection

	// data layer
	mu         sync.Mutex
	balances   Balances
	activities []Activity
}

func Init() *Server {
	return &Server{
		mu:         sync.Mutex{},
		balances:   Balances{},
		activities: []Activity{},
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	type req struct {
		method string
		path   string
	}
	request := req{r.Method, r.URL.Path}
	switch request {
	case req{"GET", "/balances"}:
		s.getBalances(w, r)
	case req{"GET", "/activities"}:
		s.getAcvitities(w, r)
	case req{"POST", "/deposits"}:
		s.deposit(w, r)
	case req{"POST", "/withdrawals"}:
		s.withdraw(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *Server) do(w http.ResponseWriter, r *http.Request, requestDest interface{}, doFunc func() (interface{}, error)) {
	if requestDest != nil {
		d := json.NewDecoder(r.Body)
		err := d.Decode(&requestDest)
		if err != nil {
			fmt.Println(err.Error(), "decoding json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		spew.Dump("requestDest", requestDest)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	resp, err := doFunc()
	if err != nil {
		fmt.Println(err.Error(), "doFunc")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err.Error(), "encoding json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		fmt.Println(err.Error(), "doFunc")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getBalances(w http.ResponseWriter, r *http.Request) {
	s.do(w, r, nil, func() (interface{}, error) {
		return s.balances, nil
	})
}

func (s *Server) getAcvitities(w http.ResponseWriter, r *http.Request) {
	s.do(w, r, nil, func() (interface{}, error) {
		return s.activities, nil
	})
}

func (s *Server) deposit(w http.ResponseWriter, r *http.Request) {
	req := &Deposit{}
	s.do(w, r, req, func() (interface{}, error) {
		return s.CreateDeposit(req)
	})
}

func (s *Server) withdraw(w http.ResponseWriter, r *http.Request) {
	req := &Withdrawal{}
	s.do(w, r, req, func() (interface{}, error) {
		return s.CreateWithdrawal(req)
	})
}

func (s *Server) nextActivityID() uint64 {
	return uint64(len(s.activities) + 1)
}
