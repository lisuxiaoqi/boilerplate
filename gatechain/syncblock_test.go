package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gatechain/gatemint/protocol"
	"github.com/gatechain/gatemint/rpcs"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestBlockSync(t *testing.T) {
	host := "http://50.116.39.212:7878"
	url := "/v1/mainnet/block"

	//host := "http://127.0.0.1:7878"
	//url := "/v1/gate-66/block"

	//convert to 36 based
	var h int64 = 15180689
	h36 := strconv.FormatInt(h, 36)

	//final url
	url = fmt.Sprintf("%s%s/%s", host, url, h36)
	t.Log(url)

	//Get EncodedBlockCert
	rawBytes, err := httpGet(t, url)
	require.NoError(t, err)
	t.Log("bcertHex", hex.EncodeToString(rawBytes))

	bcert, err := decodeBlockCert(t, rawBytes)
	require.NoError(t, err)

	t.Log("block received",
		"\nheight", bcert.Block.BlockHeader.Round,
		"\nblock hash pretty\t", bcert.Block.Hash().String(),
		"\nblock hash\t", bcert.Block.Digest().String(),
		"\ncert hash\t", bcert.Certificate.Proposal.BlockDigest.String(),
		"\ncert len\t", len(bcert.Certificate.Votes))

}

func httpGet(t *testing.T, url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 100 * time.Second, // 设置请求超时时间
	}
	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应
	data := make([]byte, resp.ContentLength)
	n, err := io.ReadFull(resp.Body, data)
	if err != nil {
		panic(err)
	}

	t.Log("Status Code:", resp.StatusCode)
	t.Log("Response len:", resp.ContentLength, "data len:", n)
	return data, err
}

func decodeBlockCert(t *testing.T, rawBytes []byte) (*rpcs.EncodedBlockCert, error) {
	var decodedEntry rpcs.EncodedBlockCert
	err := protocol.Decode(rawBytes, &decodedEntry)
	return &decodedEntry, err
}

func TestConv36(t *testing.T) {
	var i int64 = 100
	ibase36 := strconv.FormatInt(i, 36)
	t.Log("base32", i, ibase36)

	round, err := strconv.ParseUint(ibase36, 36, 64)
	require.NoError(t, err)
	t.Log("back to base", ibase36, round)
}
