package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"log"
	"math/big"
)

const (
	// The address used to deploy the contract becomes the owner
	// We use the first address in the HD wallet that ganache creates with the seed in the docker-compose
	OwnerPrivateKey    = "0xca3f21d12dbd7a6a301a70414c40c555761dfce4b3f84cb8ce5f9760a783e3f0"
	ContractAddr       = "0xc4680463046E64b10Da390d9049D24b8EC43AaAB"
	GanacheNetworkAddr = "ws://127.0.0.1:8545"
)

type ChainId *big.Int

var (
	OwnerTransactor *bind.TransactOpts
	ChainIDLocal    ChainId = big.NewInt(1337)
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

}

func main() {
	ethClient, err := ethclient.Dial(GanacheNetworkAddr)
	if err != nil {
		panic(err)
	}
	// TODO: fill out this function for exercise 1!
}
