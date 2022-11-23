package main

import (
	"flag"
	"log"
	"time"

	"github.com/channingduan/bcos-go-sdk-demo/helloworld"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

var isSMCrypto = flag.Bool("isGM", false, "use gm crypto")

func main() {
	flag.Parse()
	confFile := "./configs/bm.toml"
	contractAddress := common.HexToAddress("0xfd8586d547361cae2ba53803809bd3a6aa20ed72")
	if *isSMCrypto {
		confFile = "./configs/gm.toml"
		contractAddress = common.HexToAddress("0xcfd05114d96cfbb508dd20244e63f5e484a943de")
	}

	configs, err := conf.ParseConfigFile(confFile)
	if err != nil {
		log.Panicf("parse config error: %v", err)
		panic(err)
	}
	srv, err := client.Dial(&configs[0])
	if err != nil {
		log.Panicf("client dial error: %v", err)
		panic(err)
	}

	instance, err := helloworld.NewHelloWorld(contractAddress, srv)
	if err != nil {
		log.Panicf("NewHelloWorld error: %v", err)
		return
	}

	helloworldSession := &helloworld.HelloWorldSession{Contract: instance, CallOpts: *srv.GetCallOpts(), TransactOpts: *srv.GetTransactOpts()}

	tx, res, err := helloworldSession.Set(time.Now().String())
	if err != nil {
		log.Fatalf("helloworld Session get: %v", err)
		return
	}

	log.Printf("tx: %v", tx.Hash().String())
	log.Printf("res: %v", res)
	get, err := helloworldSession.Get()
	if err != nil {
		log.Fatalf("helloworld Session get: %v", err)
		return
	}
	log.Println("get", *isSMCrypto, get)
}
