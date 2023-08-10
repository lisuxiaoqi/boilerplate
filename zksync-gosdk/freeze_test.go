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

var ()

// Freeze：
// 代码：contracts/ethereum/contracts/zksync/facets/DiamondCut.sol
// abi: contracts/ethereum/artifacts/cache/solpp-generated-contracts/zksync/facets/DiamondCut.sol/DiamondCutFacet.json
// 合约地址：注意，这里要使用CONTRACTS_DIAMOND_PROXY_ADDR地址，而不是CONTRACTS_GETTERS_FACET_ADDR地址
func TestFreeze(t *testing.T) {
	url := "http://localhost:8545"
	//这儿要使用CONTRACTS_DIAMOND_PROXY_ADDR地址作为总入口
	contractAddrStr := "0x27d90116114E5654509716e19777eDd1F7A0165E"
	privateKeyStr := "e131bc3f481277a8f73d680d9ba404cc6f959e64296e0914dded403030d4f705"

	contractABI := `
	[
		{
		  "inputs": [],
		  "name": "freezeDiamond",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "unfreezeDiamond",
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

	functionName := "freezeDiamond"

	// 构建调用数据
	data, err := parsedABI.Pack(functionName)
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

func TestUnFreeze(t *testing.T) {
	url := "http://localhost:8545"
	//这儿要使用CONTRACTS_DIAMOND_PROXY_ADDR地址作为总入口
	contractAddrStr := "0x7417822e3ca914A83a24d18CB529E9EfDE2f2b0a"
	privateKeyStr := "e131bc3f481277a8f73d680d9ba404cc6f959e64296e0914dded403030d4f705"

	contractABI := `
	[
		{
		  "inputs": [],
		  "name": "freezeDiamond",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "unfreezeDiamond",
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

	functionName := "unfreezeDiamond"

	// 构建调用数据
	data, err := parsedABI.Pack(functionName)
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
