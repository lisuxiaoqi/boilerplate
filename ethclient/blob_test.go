package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	gokzg4844 "github.com/crate-crypto/go-eth-kzg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"testing"
)

func TestBlob(t *testing.T) {
	for i := 0; i < 1; i++ {
		run()
	}
}

func run() {
	privKey := "a18b16c79a875c85a16377735a1fc713d1c90ae59303f2f66aa43b256c5ef41c"
	toAddr := "0x0000000000000000000000000000000000000000"
	ctx := context.Background()
	rpcURL := rawRPCURL // EIP-4844 测试网
	client, _ := ethclient.Dial(rpcURL)

	privateKey, _ := crypto.HexToECDSA(privKey)
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, _ := client.NonceAt(ctx, from, nil)
	fmt.Println(from.String())
	balance, err := client.BalanceAt(ctx, from, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("before ", balance.String())
	var blobs []kzg4844.Blob
	var comments []kzg4844.Commitment
	var proofs []kzg4844.Proof

	count := mrand.Intn(6)
	if count == 0 {
		count = 1
	}
	for i := 0; i < count; i++ {
		blob := randBlob()
		//	fmt.Println(hex.EncodeToString(blob[:]))
		commentment, err := kzg4844.BlobToCommitment(blob)
		if err != nil {
			panic(err)
		}
		proof, err := kzg4844.ComputeBlobProof(blob, commentment)
		if err != nil {
			panic(err)
		}

		blobs = append(blobs, *blob)
		comments = append(comments, commentment)
		proofs = append(proofs, proof)
	}

	sidecar := &types.BlobTxSidecar{
		Version:     1,
		Blobs:       blobs,
		Commitments: comments,
		Proofs:      proofs,
	}

	chainID := uint64(32382)
	to := common.HexToAddress(toAddr)
	chainId := uint256.NewInt(chainID)

	hashes := sidecar.BlobHashes()

	data, err := hex.DecodeString("c79f8b62")
	if err != nil {
		panic(err)
	}
	blobTx := &types.BlobTx{
		ChainID: chainId, // Sepolia
		Nonce:   nonce,
		To:      to,
		Value:   uint256.NewInt(0),
		Gas:     20000000,
		Data:    data,

		GasFeeCap:  uint256.NewInt(50000000000),
		GasTipCap:  uint256.NewInt(2000000000),
		BlobFeeCap: uint256.NewInt(1000000000),
		Sidecar:    sidecar,
		BlobHashes: hashes,
	}

	signedTx, err := types.SignTx(types.NewTx(blobTx), types.NewPragueSigner(chainId.ToBig()), privateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(signedTx)
	txJson, err := json.MarshalIndent(signedTx.WithoutBlobTxSidecar(), " ", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(txJson))

	binary, _ := signedTx.MarshalBinary()
	fmt.Println("binary ", len(binary))
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		fmt.Println("xxxxx ", err)
		panic(err)
	}
	fmt.Println("hash ", signedTx.Hash(), len(signedTx.BlobHashes()))

	for {
		time.Sleep(time.Second * 10)
		receipt, err := client.TransactionReceipt(ctx, signedTx.Hash())
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				continue
			}
			panic(err)
		}
		price := new(big.Int).Mul(receipt.EffectiveGasPrice, big.NewInt(int64(receipt.GasUsed)))
		fmt.Println("fee ", price.String())

		blk, err := client.BlockByNumber(ctx, receipt.BlockNumber)
		if err != nil {
			panic(err)
		}
		fmt.Println("blk hash expect ", receipt.BlockHash.String(), " actual ", blk.Hash().String(), blk.Header().Number.Int64())
		fmt.Println("blob gas used expect ", 131072*count, "actual", receipt.BlobGasUsed)
		receiptJson, err := json.MarshalIndent(receipt, " ", "\t")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(receiptJson))
		break
	}
	balance, err = client.BalanceAt(ctx, from, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("after ", balance.String())
}

func randFieldElement() [32]byte {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("failed to get random field element")
	}
	var r fr.Element
	r.SetBytes(bytes)

	return gokzg4844.SerializeScalar(r)
}

func randBlob() *kzg4844.Blob {
	var blob kzg4844.Blob
	for i := 0; i < len(blob); i += gokzg4844.SerializedScalarSize {
		fieldElementBytes := randFieldElement()
		copy(blob[i:i+gokzg4844.SerializedScalarSize], fieldElementBytes[:])
	}
	return &blob
}
