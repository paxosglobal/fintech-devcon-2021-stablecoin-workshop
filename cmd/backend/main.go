package main

import (
	"fmt"
	contracts "github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/build"
	"github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/pkg/server"
	"log"
	"net/http"
)

var _ = contracts.USDKABI

func main() {
	s := server.Init()
	fmt.Println(s)

	log.Fatal(http.ListenAndServe(":8080", s))
}
