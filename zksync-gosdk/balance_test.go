package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"log"
	"testing"
)

func TestBalance(t *testing.T) {
	var (
		//ETH account
		PrivateKey = "d8611869c1cf0548d412322d5a946b1fa5303d80a9ce48ff0a7b697d1c7f3cd6"
		PublickKey = "2BDDfa90274F14EdeFb750BB8bdDf248e397A95D"

		//L2 account
		//PrivateKey = "3ae3debf21096248d91828fc535f7cd243d817cf96248c0785bc54fb6e61c86f"
		//PublickKey = "52a48cbc7bdF3cd49c747E3dC7e28484Ce52718e"

		ZkSyncProvider   = "http://localhost:3050"
		EthereumProvider = "http://localhost:8545"
	)

	// Connect to zkSync network
	zp, err := clients.NewDefaultProvider(ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// Create singer object from private key for appropriate chain
	chainID, err := zp.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(PrivateKey), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	// Create wallet
	w, err := accounts.NewWallet(es, zp)
	if err != nil {
		log.Panic(err)
	}

	// Connect to Ethereum network
	ethRpc, err := rpc.Dial(EthereumProvider)
	if err != nil {
		log.Panic(err)
	}
	ethCli := ethclient.NewClient(ethRpc)

	_, err = w.CreateEthereumProvider(ethRpc)
	if err != nil {
		log.Panic(err)
	}

	//get balance at l1
	balanceL1, err := ethCli.BalanceAt(context.Background(), common.HexToAddress(PublickKey), nil)
	if err != nil {
		log.Panic(err)
	}

	//get balance at l2
	balanceL2, err := w.GetBalance()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("L1 Account balance: ", balanceL1)
	fmt.Println("L2 Account balance: ", balanceL2)
}
