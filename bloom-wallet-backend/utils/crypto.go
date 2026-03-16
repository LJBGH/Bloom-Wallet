package utils

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"bloom-wallet/model"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

// HD 钱包（分层确定性钱包） 的“派生路径”
// 同一份助记词怎么“派生”出某一个具体的账户私钥/地址
// 不同路径会得到不同的私钥/地址（但都来自同一助记词）
const defaultDerivationPath = "m/44'/60'/0'/0/0"

// CreateMnemonic 生成 BIP39 助记词（默认 12 个词）
func CreateMnemonic() (string, error) {
	entropy, err := bip39.NewEntropy(128) // 12 words
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

// CreateWallet 创建钱包：
// - 传入 mnemonic 为空字符串：自动生成助记词并派生第一个地址
// - 传入 mnemonic 非空：用该助记词派生第一个地址
func CreateWallet(mnemonic string) (*model.WalletInfo, error) {
	if mnemonic == "" {
		var err error
		mnemonic, err = CreateMnemonic()
		if err != nil {
			return nil, err
		}
	}
	return WalletFromMnemonic(mnemonic, defaultDerivationPath)
}

// WalletFromMnemonic 从助记词派生钱包信息（ETH: m/44'/60'/0'/0/0）
func WalletFromMnemonic(mnemonic string, derivationPath string) (*model.WalletInfo, error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic")
	}
	if derivationPath == "" {
		derivationPath = defaultDerivationPath
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	path := hdwallet.MustParseDerivationPath(derivationPath)

	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not *ecdsa.PublicKey")
	}

	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	publicKeyHex := hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &model.WalletInfo{
		Mnemonic:   mnemonic,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}, nil
}

// WalletFromPrivateKeyHex 从私钥(hex)构造钱包信息（不包含助记词）
func WalletFromPrivateKeyHex(privateKeyHex string) (*model.WalletInfo, error) {
	privateKeyHex = strings.TrimSpace(privateKeyHex)
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0X")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not *ecdsa.PublicKey")
	}

	return &model.WalletInfo{
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA).Hex(),
		PrivateKey: hexutil.Encode(crypto.FromECDSA(privateKey)),
		PublicKey:  hexutil.Encode(crypto.FromECDSAPub(publicKeyECDSA)),
	}, nil
}
