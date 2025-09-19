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
	"time"
)

func TestGetBlockByHash(t *testing.T) {
	//v19
	//hash := "0x2568d9ce1e80018d13d6d343cddd3c06ecb6b821d6285158e03989eda8bec441"
	//v20
	hash := "0x441b8692ad267597eb5ddf0c52f09e835814790b4e157505f16000b214157c40"

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
	number := big.NewInt(int64(2036))
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

func TestGetBlockByNumberLatest(t *testing.T) {
	client, _ := ethclient.Dial(rawRPCURL)

	for {
		//blockLatest, err := client.BlockByNumber(context.Background(), big.NewInt(int64(rpc.LatestBlockNumber)))
		blockLatest, err := client.BlockByNumber(context.Background(), big.NewInt(int64(rpc.LatestBlockNumber)))
		//blockLatest, err := client.BlockByNumber(context.Background(), big.NewInt(2036))
		if err != nil {
			fmt.Printf("BlockByNumber error:%v\n", err)
			time.Sleep(1 * time.Second)
			continue
		} else {
			fmt.Println("WS get block by Number:", blockLatest.Number(), "Hash:", blockLatest.Hash().Hex())
		}

		//getBlockByHash
		blockByHash, err := client.BlockByHash(context.Background(), blockLatest.Hash())
		if err != nil {
			fmt.Printf("BlockByHash error:%v\n", err)
			panic("Shouldn't happen")
		} else {
			fmt.Println("WS get block by Hash:", blockByHash.Number(), "Hash:", blockByHash.Hash().Hex())
		}
		time.Sleep(1 * time.Second)
	}
}
