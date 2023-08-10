package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"testing"
)

// Getter：
// 通过Getter查询freeze信息，Getter代码：contracts/ethereum/contracts/zksync/facets/Getters.sol
// abi: contracts/ethereum/artifacts/cache/solpp-generated-contracts/zksync/facets/Getters.sol/GettersFacet.json
// 合约地址：注意，这里要使用CONTRACTS_DIAMOND_PROXY_ADDR地址，而不是CONTRACTS_GETTERS_FACET_ADDR地址
func TestGetter(t *testing.T) {
	url := "http://localhost:8545"
	//这儿要使用CONTRACTS_DIAMOND_PROXY_ADDR地址作为总入口
	contractAddrStr := "0x27d90116114E5654509716e19777eDd1F7A0165E"

	contractABI := `
	[
		{
		  "inputs": [],
		  "name": "getGovernor",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "isDiamondStorageFrozen",
		  "outputs": [
			{
			  "internalType": "bool",
			  "name": "",
			  "type": "bool"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getTotalBlocksCommitted",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getTotalBlocksExecuted",
		  "outputs": [
			{
			  "internalType": "uint256",
			  "name": "",
			  "type": "uint256"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getTotalBlocksVerified",
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

	getGovernor(client, &parsedABI, &contractAddress)
	isDiamondStorageFrozen(client, &parsedABI, &contractAddress)
	getTotalBlocksCommitted(client, &parsedABI, &contractAddress)
	getTotalBlocksVerified(client, &parsedABI, &contractAddress)
	getTotalBlocksExecuted(client, &parsedABI, &contractAddress)
}

func getGovernor(client *ethclient.Client, parsedABI *abi.ABI, addr *common.Address) {
	functionName := "getGovernor"
	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   addr,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解码返回结果
	var r common.Address
	err = parsedABI.UnpackIntoInterface(&r, functionName, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result of %s: %v\n", functionName, r.String())
}

func isDiamondStorageFrozen(client *ethclient.Client, parsedABI *abi.ABI, addr *common.Address) {
	functionName := "isDiamondStorageFrozen"
	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   addr,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解码返回结果
	var r bool
	err = parsedABI.UnpackIntoInterface(&r, functionName, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result of %s: %v\n", functionName, r)
}

func getTotalBlocksCommitted(client *ethclient.Client, parsedABI *abi.ABI, addr *common.Address) {
	//测试getTotalBlocksCommitted
	functionName := "getTotalBlocksCommitted"
	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   addr,
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

func getTotalBlocksVerified(client *ethclient.Client, parsedABI *abi.ABI, addr *common.Address) {
	functionName := "getTotalBlocksVerified"
	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   addr,
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

func getTotalBlocksExecuted(client *ethclient.Client, parsedABI *abi.ABI, addr *common.Address) {
	functionName := "getTotalBlocksExecuted"
	//封装合约调用函数
	data, err := parsedABI.Pack(functionName)
	if err != nil {
		log.Fatal(err)
	}

	//构建msg
	callMsg := ethereum.CallMsg{
		To:   addr,
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
