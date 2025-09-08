package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestGetBlockByHash(t *testing.T) {
	hash := "0x6f0eea6a79cc10ce63f38e85755f32b49e3300144f2b8c95e9e5fdf0a36e4302"

	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByHash(context.Background(), common.HexToHash(hash))
	block.BlobGasUsed()
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

// Get Block safe, finalized
func TestGetBlock(t *testing.T) {
	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	b, _ := json.MarshalIndent(block.Header(), "", "  ")
	fmt.Println(string(b))
}

func TestGetBlockByNumber(t *testing.T) {
	number := big.NewInt(int64(3))
	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByNumber(context.Background(), number)
	block.BlobGasUsed()
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	t.Log("got block hash:", block.Hash().Hex())

	b, _ := json.MarshalIndent(block.Header(), "", "  ")
	fmt.Println(string(b))
}
