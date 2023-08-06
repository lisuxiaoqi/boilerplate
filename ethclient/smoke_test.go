package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetBlockNumber(t *testing.T) {
	rawRPCURL := "http://localhost:8545"

	client, _ := ethclient.Dial(rawRPCURL)
	h, _ := client.BlockNumber(context.Background())
	fmt.Println(h)
}
