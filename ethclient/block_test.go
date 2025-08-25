package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

// Get Tx Receipt
func TestGetBlockByHash(t *testing.T) {
	rawRPCURL := "http://localhost:8545"
	hash := "0x71c125191bd6a263706752f84169a9c6952f4254e37a93594ba7626ec4b22d21"

	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	t.Log("origin block hash:", hash)
	t.Log("got block hash:", block.Hash().Hex())
	require.Equal(t, block.Hash(), common.HexToHash(hash))

	b, _ := json.MarshalIndent(block.Header(), "", "  ")
	fmt.Println(string(b))
}
