package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/pkg/server"
	"log"
)

func main() {
	ethClient, err := ethclient.Dial(server.GanacheNetworkAddr)
	if err != nil {
		panic(err)
	}
	addr, _, _, err := contracts.DeployUSDK(server.OwnerTransactor, ethClient)
	if err != nil {
		panic(err)
	}
	log.Print("contract address: ", addr) // 0xc4680463046E64b10Da390d9049D24b8EC43AaAB
}
