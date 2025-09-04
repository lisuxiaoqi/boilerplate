package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

var data = "0x02f9011482053912847735940085051f4d5c008405f5e1009446f74a450cc732754dcb814aa221ebf359303afa80b8a482ecf2f60000000000000000000000000000000000000000000000000000000000000000c38863171c193c156db075d09fd076424ee29051fbb398231ca3341743f9370e00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000439fc080a0fd68137949c75a23f4480157a5a914a9a99fb79af57d9e2b4b7b1cda199aca28a017149de4ccd9158c5f91d8142da7b9753d7f94b2bad2a194213ea218a07a0ce7"

// Get Block safe, finalized
func TestSendRawTx(t *testing.T) {
	rawRPCURL := "http://localhost:7545"

	client, _ := ethclient.Dial(rawRPCURL)

	err := client.Client().CallContext(context.Background(), nil, "eth_sendRawTransaction", data)
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	//b, _ := json.MarshalIndent(block.Header(), "", "  ")
	//fmt.Println(string(b))
}
