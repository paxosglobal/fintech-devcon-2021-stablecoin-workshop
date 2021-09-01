package server

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
)

const (
	// same as the deployer address
	SupplyControllerPrivateKey = "0xca3f21d12dbd7a6a301a70414c40c555761dfce4b3f84cb8ce5f9760a783e3f0"
	GanacheNetworkAddr         = "ws://127.0.0.1:8545"
	ContractAddr               = "0xc4680463046E64b10Da390d9049D24b8EC43AaAB"
)

var (
	GanacheChainID  = big.NewInt(1337)
	OwnerTransactor *bind.TransactOpts
	EthToWei        = decimal.New(1, 18)
)

func init() {
	key, err := crypto.HexToECDSA(SupplyControllerPrivateKey[2:])
	if err != nil {
		panic(err)
	}
	OwnerTransactor, err = bind.NewKeyedTransactorWithChainID(key, GanacheChainID)
	if err != nil {
		panic(err)
	}
	log.Print("owner address: ", OwnerTransactor.From.String())
}

func (s *Server) mint(destination string, amount decimal.Decimal) error {
	usdkBindings, err := contracts.NewUSDK(common.HexToAddress(ContractAddr), s.ethClient)
	if err != nil {
		return err
	}
	_, err = usdkBindings.Mint(OwnerTransactor, common.HexToAddress(destination), amount.Mul(EthToWei).BigInt())
	return err
}
