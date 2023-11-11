package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

// Get Tx Receipt
func TestGetTxReceipt(t *testing.T) {
	rawRPCURL := "http://localhost:8545"
	txHash := "0xb75ba1fec6c309c3b6c828f8a0047f0018a3c3305d2e4814238146cf2e5a32fc"

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
