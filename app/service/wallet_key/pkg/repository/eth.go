package repository

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	bip39 "github.com/tyler-smith/go-bip39"
)

func (repo *repo) SignedEthereumTx(ctx context.Context, walletHD WalletHD, txHex string, chainID string) (string, error) {
	chainIDBig, result := new(big.Int).SetString(chainID, 10)
	if !result {
		return "", fmt.Errorf("Fail to extract chain id")
	}

	tx, err := decodeETHTx(txHex)
	if err != nil {
		return "", err
	}
	mnemonic, err := repo.ldb.Get([]byte(walletHD.MnemonicVersion), nil)
	if err != nil {
		return "", fmt.Errorf("Fail to query mnemonic %s", err.Error())
	}
	seed := bip39.NewSeed(string(mnemonic), "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}

	// This gives the path: m / 44'
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return "", err
	}

	coinTypeH, err := acc44H.Child(hdkeychain.HardenedKeyStart + uint32(walletHD.CoinType))
	if err != nil {
		return "", err
	}

	accountH, err := coinTypeH.Child(hdkeychain.HardenedKeyStart + uint32(walletHD.Account))
	if err != nil {
		return "", err
	}

	changeLevel, err := accountH.Child(uint32(walletHD.Change))
	if err != nil {
		return "", err
	}
	addressIndexLevel, err := changeLevel.Child(walletHD.AddressIndex)
	if err != nil {
		return "", err
	}
	privateKey, err := addressIndexLevel.ECPrivKey()
	if err != nil {
		return "", err
	}
	signtx, err := types.SignTx(tx, types.NewEIP155Signer(chainIDBig), privateKey.ToECDSA())
	if err != nil {
		return "", fmt.Errorf("Ethereum transaction signatrue %s", err)
	}
	signedTxHex, err := encodeETHTx(signtx)
	if err != nil {
		return "", err
	}
	return signedTxHex, nil
}

func decodeETHTx(txHex string) (*types.Transaction, error) {
	txc, err := hexutil.Decode(txHex)
	if err != nil {
		return nil, err
	}

	var txde types.Transaction
	t, err := &txde, rlp.Decode(bytes.NewReader(txc), &txde)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// EncodeETHTx encode eth tx
func encodeETHTx(tx *types.Transaction) (string, error) {
	txb, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}
	txHex := hexutil.Encode(txb)
	return txHex, nil
}
