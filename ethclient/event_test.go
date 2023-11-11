package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type RollupEventWithdraw struct {
	Idx         uint64
	NumExitRoot uint64
	TxHash      common.Hash // Hash of the transaction that generated this event
}

var (
	logHermezWithdrawEvent = crypto.Keccak256Hash([]byte(
		"WithdrawEvent(uint48,uint32)"))
)

// Scan Event
func TestEventScan(t *testing.T) {
	rawRPCURL := "https://meteora-evm.gatenode.cc"
	cli, _ := ethclient.Dial(rawRPCURL)
	var startNum int64 = 2960000

	endNum, _ := cli.BlockNumber(context.Background())

	var gap int64 = 10000
	fmt.Printf("Scan from %v to %v\n", startNum, endNum)
	for i := startNum; i < int64(endNum)+gap; i = i + gap {
		bNum := big.NewInt(i)
		eNum := big.NewInt(i + gap)
		fmt.Println(bNum, eNum)
		query := ethereum.FilterQuery{
			FromBlock: bNum,
			ToBlock:   eNum,
			Addresses: []common.Address{
				common.HexToAddress("0xE81f31dCb52681Eea7d180096d9879e780151224"),
			},
			Topics: [][]common.Hash{},
		}
		logs, err := cli.FilterLogs(context.Background(), query)
		require.NoError(t, err)

		for _, vLog := range logs {
			switch vLog.Topics[0] {
			case logHermezWithdrawEvent:
				var withdraw RollupEventWithdraw
				withdraw.Idx = new(big.Int).SetBytes(vLog.Topics[1][:]).Uint64()
				withdraw.NumExitRoot = new(big.Int).SetBytes(vLog.Topics[2][:]).Uint64()
				withdraw.TxHash = vLog.TxHash
				fmt.Printf("WithdrawEvent Received,blockNum:%v, batchNum:%v, idx:%v\n", vLog.BlockNumber, withdraw.NumExitRoot, withdraw.Idx)
			}
		}
	}
}
