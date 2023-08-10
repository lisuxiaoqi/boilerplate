package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"log"
	"math/big"
	"testing"
)

func TestTransferL2(t *testing.T) {
	var (
		FromPrivateKey       = "d8611869c1cf0548d412322d5a946b1fa5303d80a9ce48ff0a7b697d1c7f3cd6"
		ToAddress            = "0x52a48cbc7bdF3cd49c747E3dC7e28484Ce52718e"
		ZkSyncProvider       = "http://localhost:3050"
		transferAmount int64 = 1_000_000_000_000_000_000
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
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(FromPrivateKey), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	// Create wallet
	w, err := accounts.NewWallet(es, zp)
	if err != nil {
		log.Panic(err)
	}

	// Show balances before transfer for both accounts
	account1Balance, err := w.GetBalance()
	if err != nil {
		log.Panic(err)
	}
	account2Balance, err := zp.BalanceAt(context.Background(), common.HexToAddress(ToAddress), nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Account1 balance before transfer: ", WeiToEther(account1Balance))
	fmt.Println("Account2 balance before transfer: ", WeiToEther(account2Balance))

	// Perform transfer
	hash, err := w.Transfer(
		common.HexToAddress(ToAddress),
		big.NewInt(transferAmount),
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction: ", hash)

	// Wait for transaction to be finalized on L2 network
	_, err = zp.WaitMined(context.Background(), hash)
	if err != nil {
		log.Panic(err)
	}

	// Show balances after transfer for both accounts
	account1Balance, err = w.GetBalance()
	if err != nil {
		log.Panic(err)
	}
	account2Balance, err = zp.BalanceAt(context.Background(), common.HexToAddress(ToAddress), nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Account1 balance after transfer: ", WeiToEther(account1Balance))
	fmt.Println("Account2 balance after transfer: ", WeiToEther(account2Balance))
}
