package eth

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestSelector(t *testing.T) {
	signature := "ClockNotExpired()"

	hash := sha3.NewLegacyKeccak256() // Ethereum 使用 Keccak256
	hash.Write([]byte(signature))
	sum := hash.Sum(nil)
	selector := sum[:4] // 前 4 个字节
	fmt.Println("0x" + hex.EncodeToString(selector))
}
