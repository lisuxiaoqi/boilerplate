package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestGetTxReceipt(t *testing.T) {
	rawRPCURL := "http://localhost:8545"
	txHash := "0xfaa3cb7c64dc20fcf77de55e041ee03ae939ca525594ad759cfe63649540b843"

	client, _ := ethclient.Dial(rawRPCURL)

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		t.Errorf("Get Tx Receipt error:%v\n", err)
		return
	}

	// 将 Receipt 转换为 JSON 字符串
	receiptJSON, err := json.MarshalIndent(receipt, "", "    ")
	if err != nil {
		t.Fatalf("Get Tx Receipt error:%v\n", err)
	}

	fmt.Println("Receipt as JSON string:")
	fmt.Println(string(receiptJSON))
}
