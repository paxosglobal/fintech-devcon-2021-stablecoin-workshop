package server

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strings"
)

const (
	// The address used to deploy the contract becomes the owner
	// We use the first address in the HD wallet that ganache creates with the seed in the docker-compose
	OwnerPrivateKey = "0xca3f21d12dbd7a6a301a70414c40c555761dfce4b3f84cb8ce5f9760a783e3f0"
	ContractAddr    = "0xc4680463046E64b10Da390d9049D24b8EC43AaAB"
)

var (
	OwnerTransactor *bind.TransactOpts
	DecimalToInt    = decimal.New(1, 18) // can also call decimals() on the contract to get the exponent (18)
	UsdkABI         abi.ABI
)

type ChainId *big.Int

var (
	ChainIDMainnet ChainId = big.NewInt(1)
	ChainIDRopsten ChainId = big.NewInt(3)
	ChainIDRinkeby ChainId = big.NewInt(4)
	ChainIDLocal   ChainId = big.NewInt(1337)
)

func init() {
	key, err := crypto.HexToECDSA(OwnerPrivateKey[2:])
	if err != nil {
		panic(err)
	}
	OwnerTransactor, err = bind.NewKeyedTransactorWithChainID(key, ChainIDLocal)
	if err != nil {
		panic(err)
	}
	log.Print("owner address: ", OwnerTransactor.From.String())

	UsdkABI, err = abi.JSON(strings.NewReader(contracts.USDKABI))
	if err != nil {
		panic(err)
	}
}

func (s *Server) mint(destination string, amount decimal.Decimal) error {
	return s.mintWithBindings(destination, amount)
}

func (s *Server) mintWithBindings(destination string, amount decimal.Decimal) error {
	// TODO: fill out this function for exercise 2!
	// test it with `go test ./pkg/server -run TestMint`
	return nil
}
