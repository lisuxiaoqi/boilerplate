package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

// Get Tx Receipt
func TestGetBlockByHash(t *testing.T) {
	rawRPCURL := "http://localhost:8545"
	hash := "0x71c125191bd6a263706752f84169a9c6952f4254e37a93594ba7626ec4b22d21"

	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	t.Log("origin block hash:", hash)
	t.Log("got block hash:", block.Hash().Hex())
	require.Equal(t, block.Hash(), common.HexToHash(hash))

	b, _ := json.MarshalIndent(block.Header(), "", "  ")
	fmt.Println(string(b))
}

// Get Block safe, finalized
func TestGetBlock(t *testing.T) {
	rawRPCURL := "http://localhost:7545"

	client, _ := ethclient.Dial(rawRPCURL)

	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(rpc.FinalizedBlockNumber)))
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	b, _ := json.MarshalIndent(block.Header(), "", "  ")
	fmt.Println(string(b))
}

type Bytes256 [256]byte
type RPCHeader struct {
	ParentHash  common.Hash      `json:"parentHash"`
	UncleHash   common.Hash      `json:"sha3Uncles"`
	Coinbase    common.Address   `json:"miner"`
	Root        common.Hash      `json:"stateRoot"`
	TxHash      common.Hash      `json:"transactionsRoot"`
	ReceiptHash common.Hash      `json:"receiptsRoot"`
	Bloom       Bytes256         `json:"logsBloom"`
	Difficulty  hexutil.Big      `json:"difficulty"`
	Number      hexutil.Uint64   `json:"number"`
	GasLimit    hexutil.Uint64   `json:"gasLimit"`
	GasUsed     hexutil.Uint64   `json:"gasUsed"`
	Time        hexutil.Uint64   `json:"timestamp"`
	Extra       hexutil.Bytes    `json:"extraData"`
	MixDigest   common.Hash      `json:"mixHash"`
	Nonce       types.BlockNonce `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *hexutil.Big `json:"baseFeePerGas"`

	// WithdrawalsRoot was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsRoot *common.Hash `json:"withdrawalsRoot,omitempty"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *hexutil.Uint64 `json:"blobGasUsed,omitempty"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *hexutil.Uint64 `json:"excessBlobGas,omitempty"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot,omitempty"`

	// RequestsHash was added by EIP-7685 and is ignored in legacy headers.
	RequestsHash *common.Hash `json:"requestsHash,omitempty" rlp:"optional"`

	// untrusted info included by RPC, may have to be checked
	Hash common.Hash `json:"hash"`
}

// 0x02f9011482053912847735940085051f4d5c008405f5e1009446f74a450cc732754dcb814aa221ebf359303afa80b8a482ecf2f60000000000000000000000000000000000000000000000000000000000000000c38863171c193c156db075d09fd076424ee29051fbb398231ca3341743f9370e00000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000439fc080a0fd68137949c75a23f4480157a5a914a9a99fb79af57d9e2b4b7b1cda199aca28a017149de4ccd9158c5f91d8142da7b9753d7f94b2bad2a194213ea218a07a0ce7
// Get Block safe, finalized
func TestGetBlockByContext(t *testing.T) {
	rawRPCURL := "http://localhost:7545"

	client, _ := ethclient.Dial(rawRPCURL)
	var header *RPCHeader
	err := client.Client().CallContext(context.Background(), &header, "eth_getBlockByNumber", "latest", false)
	if err != nil {
		t.Errorf("Get block error:%v\n", err)
		return
	}

	b, _ := json.MarshalIndent(header, "", "  ")
	fmt.Println(string(b))
}
