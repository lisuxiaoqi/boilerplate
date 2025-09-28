package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
	"testing"
)

// 生成以太坊地址，私钥
func Test_GenEthAddress(t *testing.T) {
	// privKey
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	//privHex
	privBytes := crypto.FromECDSA(privateKey)
	privHex := hex.EncodeToString(privBytes)

	// pubKey to address
	pubKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pubKey)

	fmt.Println("Private Key (hex, 0x prefixed): 0x" + privHex)
	fmt.Println("Address (0x prefixed):", address.Hex())
}

// 生成以太坊助记词
func Test_OutputMemo(t *testing.T) {
	// 1. 生成 128 位熵 (对应 12 个助记词)
	entropy, err := bip39.NewEntropy(128)
	require.NoError(t, err)

	// 根据熵生成助记词
	mnemonic, err := bip39.NewMnemonic(entropy)
	require.NoError(t, err)

	fmt.Println("Mnemonic (12 words):", mnemonic)

	// 2. 根据助记词生成 seed
	seed := bip39.NewSeed(mnemonic, "") // 空密码

	// 3. 使用 BIP32 从种子生成主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	require.NoError(t, err)

	// 4. 使用 BIP44 标准路径 m/44'/60'/0'/0/0 派生以太坊私钥
	purpose, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) // 44'
	coinType, _ := purpose.NewChildKey(bip32.FirstHardenedChild + 60)  // 60' (ETH)
	account, _ := coinType.NewChildKey(bip32.FirstHardenedChild + 0)   // 0'
	change, _ := account.NewChildKey(0)                                // 0
	addressIndex, _ := change.NewChildKey(0)                           // 0

	// 5. 转换为 ecdsa 私钥
	privKeyECDSA, err := crypto.ToECDSA(addressIndex.Key)
	require.NoError(t, err)

	// 6. 生成地址
	pubKey := privKeyECDSA.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pubKey)

	fmt.Printf("Private Key (hex, 0x prefixed): 0x%x\n", crypto.FromECDSA(privKeyECDSA))
	fmt.Println("Ethereum Address:", address.Hex())
}

func deriveETHAccount(seed []byte, index uint32) (*ecdsa.PrivateKey, string, error) {
	// BIP32: 从种子生成主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, "", err
	}

	// BIP44 标准路径 m/44'/60'/0'/0/index
	purpose, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) // 44'
	coinType, _ := purpose.NewChildKey(bip32.FirstHardenedChild + 60)  // 60' ETH
	account, _ := coinType.NewChildKey(bip32.FirstHardenedChild + 0)   // 0'
	change, _ := account.NewChildKey(0)                                // 0
	addressIndex, _ := change.NewChildKey(index)                       // index

	// 转 ECDSA 私钥
	privKey, err := crypto.ToECDSA(addressIndex.Key)
	if err != nil {
		return nil, "", err
	}

	// 生成地址
	pubKey := privKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*pubKey).Hex()

	return privKey, address, nil
}

// 导入以太坊助记词
func Test_AddressFromMemo(t *testing.T) {
	// 输入助记词（BIP39 12/24 个单词）
	mnemonic := "spice episode shoe wing danger reason sweet beyond dust escape science since"

	// 1. 验证助记词
	require.True(t, bip39.IsMnemonicValid(mnemonic))

	// 2. 助记词 -> 种子
	seed := bip39.NewSeed(mnemonic, "")

	// 3. 派生第 1 个地址（index=0）
	priv1, addr1, err := deriveETHAccount(seed, 0)
	require.NoError(t, err)
	fmt.Printf("Account 1:\n  Private Key: 0x%x\n  Address: %s\n\n", crypto.FromECDSA(priv1), addr1)

	// 4. 派生第 2 个地址（index=1）
	priv2, addr2, err := deriveETHAccount(seed, 1)
	require.NoError(t, err)

	fmt.Printf("Account 2:\n  Private Key: 0x%x\n  Address: %s\n", crypto.FromECDSA(priv2), addr2)
}
