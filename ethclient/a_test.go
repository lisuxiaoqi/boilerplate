package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestSendTransaction(t *testing.T){
	rawRPCURL := "http://localhost:8545"

	client, _ := ethclient.Dial(rawRPCURL)
	//client.SendTransaction(context.Background())
	h,_ := client.BlockNumber(context.Background())
	fmt.Println(h)
}
