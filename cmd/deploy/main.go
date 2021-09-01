package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"log"
	"math/big"
)

const (
	// one of the first addresses that ganache has available
	deployerPrivateKey = "0xca3f21d12dbd7a6a301a70414c40c555761dfce4b3f84cb8ce5f9760a783e3f0"
	ganacheNetworkAddr = "ws://127.0.0.1:8545"
)

var ganacheChainID = big.NewInt(1337)

func main() {
	key, err := crypto.HexToECDSA(deployerPrivateKey[2:])
	if err != nil {
		panic(err)
	}
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(key, ganacheChainID)
	if err != nil {
		panic(err)
	}
	log.Print("deployer address: ", ownerAuth.From.String())
	ethClient, err := ethclient.Dial(ganacheNetworkAddr)
	if err != nil {
		panic(err)
	}
	addr, _, _, err := contracts.DeployUSDK(ownerAuth, ethClient)
	if err != nil {
		panic(err)
	}
	log.Print("contract address: ", addr) // 0xc4680463046E64b10Da390d9049D24b8EC43AaAB
}
