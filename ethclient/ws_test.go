package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
	"time"
)

func TestWSSubScribe(t *testing.T) {
	//wsURL := "ws://localhost:7545"
	wsURL := "wss://gatelayer-ws-testnet.gatenode.cc"
	loop := 5
	cnt := 0

	client, err := ethclient.Dial(wsURL)
	if err != nil {
		utils.Fatalf("Failed to connect to Ethereum node: %v", err)
		return
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		utils.Fatalf("Failed to sucscribe newHeads: %v", err)
		return
	}

	quitChan := make(chan struct{})

	errChan := make(chan error)
	go func() {
		for {
			select {
			case err := <-sub.Err():
				utils.Fatalf("SubscribeNewHead error:%v", err)
				errChan <- err
				return
			case header := <-headers:
				if cnt >= loop {
					quitChan <- struct{}{}
					return
				} else {
					cnt++
				}

				fmt.Println(time.Now(), "WS Get header:", header.Number, "Hash:", header.Hash().Hex())
				//gteBlockByNum
				blockByNum, err := client.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					fmt.Printf("BlockByNumber error:%v\n", err)
					continue
				} else {
					fmt.Println("WS get block by Number:", blockByNum.Number(), "Hash:", header.Hash().Hex())
				}

				//getBlockByHash
				block, err := client.BlockByHash(context.Background(), header.Hash())
				if err != nil {
					fmt.Printf("BlockByHash error:%v\n", err)
					continue
				} else {
					fmt.Println("WS get block by Hash:", block.Number(), "txs:", block.Transactions().Len())
				}

				if block.Transactions() == nil || block.Transactions().Len() == 0 {
					if (header.TxHash != common.Hash{}) {
						fmt.Println("Get Block with empty tx, blockNumber:", block.Number())
						continue
					}
				}

				//getTxReceipt
				for _, tx := range block.Transactions() {
					//fmt.Println("Raw txHash", tx.Hash().Hex())
					receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
					if err != nil {
						fmt.Printf("Get Tx Receipt error:%v\n", err)
					} else {
						fmt.Println("tx Hash in receipt", receipt.TxHash.Hex())
					}

				}
			}
		}
	}()

	for {
		select {
		case <-errChan:
			return
		case <-quitChan:
			return
		}
	}
}
