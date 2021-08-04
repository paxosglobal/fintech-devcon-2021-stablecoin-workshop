package main

import (
	"fmt"
	"github.com/paxosglobal/fintech-devcon-2021-stablecoin-workshop/pkg/server"
	"log"
	"net/http"
)

func main() {
	s := server.Init()
	fmt.Println(s)
	log.Fatal(http.ListenAndServe(":8080", s))
}
