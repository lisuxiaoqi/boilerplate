package eth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

var data = "0x02f9011482053912847735940085051f4d5c008405f5e1009446f74a450cc732754dcb814aa221ebf359303afa80b8a482ecf2f60000000000000000000000000000000000000000000000000000000000000000c38863171c193c156db075d09fd076424ee29051fbb398231ca3341743f9370e00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000439fc080a0fd68137949c75a23f4480157a5a914a9a99fb79af57d9e2b4b7b1cda199aca28a017149de4ccd9158c5f91d8142da7b9753d7f94b2bad2a194213ea218a07a0ce7"

// Get Block safe, finalized
func TestSendRawTx(t *testing.T) {
	client, _ := ethclient.Dial(rawRPCURL)

	err := client.Client().CallContext(context.Background(), nil, "eth_sendRawTransaction", data)
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	//b, _ := json.MarshalIndent(block.Header(), "", "  ")
	//fmt.Println(string(b))
}

// Get Tx Receipt
func TestGetTxReceipt(t *testing.T) {
	//v18
	//txHash := "0x897dbdd21a01fcbe92b163d0c77115460836d8aa98d861c7c26e81de2a8ee93f"
	//v20
	txHash := "0x02f3601df69f891dc99bcc05fc34442200338bbdada15b70acf704ef46d62b5c"
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

func TestGetCode(t *testing.T) {
	contractAddr := "0xEf7fCbF4e0740Aa07B9a3024334929402fe3FE45"

	//check contract address
	client, _ := ethclient.Dial(rawRPCURL)
	code, err :=
		client.CodeAt(context.Background(), common.HexToAddress(contractAddr), nil)
	if err != nil {
		fmt.Printf("Get contract code at address:%v error:%v\n", contractAddr, err)
		return
	}

	if len(code) == 0 {
		fmt.Printf("No contract deployed at address::%v\n", contractAddr)
		return
	}

	fmt.Printf("Successfully get contract at address::%v\n", contractAddr)
}
