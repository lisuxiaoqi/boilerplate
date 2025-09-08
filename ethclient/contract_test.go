package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"
)

/*
* 测试Sample.Sol
 */
func TestGetContractCode(t *testing.T) {
	url := rawRPCURL
	contractAddr := "0x30037F2827B2BCa2118bd8FE66A2b5ef290FD9F4"

	//check contract address
	client, _ := ethclient.Dial(url)
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

func TestGet(t *testing.T) {
	url := rawRPCURL
	contractAddrStr := "0x30037F2827B2BCa2118bd8FE66A2b5ef290FD9F4"
	functionName := "get"

	contractABI := `
	[
		{
		  "inputs": [],
		  "name": "get",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		}
	]
	`
	//连接RPC
	client, _ := ethclient.Dial(url)

	// 合约地址
	contractAddress := common.HexToAddress(contractAddrStr)

	//构建abi
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解码返回结果
	var commited *big.Int
	err = parsedABI.UnpackIntoInterface(&commited, functionName, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result of %s: %v\n", functionName, commited)

}

func TestSet(t *testing.T) {
	url := rawRPCURL
	privateKeyStr := "a18b16c79a875c85a16377735a1fc713d1c90ae59303f2f66aa43b256c5ef41c"
	contractAddrStr := "0x30037F2827B2BCa2118bd8FE66A2b5ef290FD9F4"
	functionName := "set"
	var value int64 = 32

	contractABI := `
	[
		{
		  "inputs": [
			{
			  "internalType": "uint256",
			  "name": "x",
			  "type": "uint256"
			}
		  ],
		  "name": "set",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		}
	]
	`
	//连接RPC
	client, _ := ethclient.Dial(url)

	// 合约地址
	contractAddress := common.HexToAddress(contractAddrStr)

	//构建abi
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	// 私钥和发送者地址
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 设置发送者的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 构建调用数据
	data, err := parsedABI.Pack(functionName, big.NewInt(value))
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易的gas限制和gas价格
	gasLimit := uint64(300000)          // 可以通过estimateGas函数获取
	gasPrice := big.NewInt(20000000000) // 20 Gwei

	// 构建交易
	transaction := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddress,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	// 签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(transaction, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// 等待交易完成
	receipt, err := waitMined(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	formattedReceipt, err := json.MarshalIndent(receipt, "", "    ")
	if err != nil {
		t.Fatalf("Get Tx Receipt error:%v\n", err)
	}

	fmt.Printf("Transaction receipt: %+v\n", string(formattedReceipt))
}

// 等待交易被挖矿并返回收据
func waitMined(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	ctx := context.Background()
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err != nil && err != ethereum.NotFound {
			return nil, err
		}
		if receipt != nil {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}
}
