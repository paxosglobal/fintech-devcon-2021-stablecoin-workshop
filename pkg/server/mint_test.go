package server

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func TestMint(t *testing.T) {
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		OwnerTransactor.From: core.GenesisAccount{
			Balance: decimal.NewFromInt32(100).Mul(DecimalToInt).BigInt(),
		},
	}, 1000000000)
	s := &Server{
		ethClient:  sim,
		mu:         sync.Mutex{},
		balances:   Balances{},
		activities: []Activity{},
	}
	contractAddr, _, usdkBindings, err := contracts.DeployUSDK(OwnerTransactor, sim)
	require.NoError(t, err)
	require.Equal(t, contractAddr.String(), ContractAddr)
	sim.Commit() // write a block to the blockchain

	destination := "0xc8AaE7A85CD634C0808C4aCb250B99DD7980483B"
	amount := decimal.RequireFromString("1.5")
	err = s.mint(destination, amount)
	require.NoError(t, err)
	sim.Commit() // write a block to the blockchain

	options := (*bind.CallOpts)(nil)
	endBalance, err := usdkBindings.BalanceOf(options, common.HexToAddress(destination))
	require.NoError(t, err)
	actual := decimal.NewFromBigInt(endBalance, 0).Div(DecimalToInt)
	require.Equal(t, amount.String(), actual.String(), "make sure to sign and broadcast a mint transaction")
}
