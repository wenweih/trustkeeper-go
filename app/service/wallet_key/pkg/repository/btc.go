package repository

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/tyler-smith/go-bip39"
)

func (repo *repo) SignedBitcoincoreTx(ctx context.Context, walletHD WalletHD, txHex string, vinAmount int64) (string, error) {
	if txHex == "" {
		return "", fmt.Errorf("signtx 交易签名参数错误")
	}
	txByte, err := hex.DecodeString(txHex)
	if err != nil {
		return "", err
	}

	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(bytes.NewReader(txByte)); err != nil {
		return "", err
	}

	tx := btcutil.NewTx(&msgTx)
	mnemonic, err := repo.ldb.Get([]byte(walletHD.MnemonicVersion), nil)
	if err != nil {
		return "", fmt.Errorf("Fail to query mnemonic %s", err.Error())
	}
	seed := bip39.NewSeed(string(mnemonic), "")
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.RegressionNetParams)
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

	fromAddress, err := addressIndexLevel.Address(&chaincfg.RegressionNetParams)
	if err != nil {
		return "", err
	}

	subscript, _ := txscript.PayToAddrScript(fromAddress)
	for i, txIn := range tx.MsgTx().TxIn {
		if txIn.SignatureScript, err = txscript.SignatureScript(tx.MsgTx(), i, subscript, txscript.SigHashAll, privateKey, true); err != nil {
			return "", fmt.Errorf("SignatureScript %s", err)
		}
	}

	//Validate signature
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(subscript, tx.MsgTx(), 0, flags, nil, nil, vinAmount)
	if err != nil {
		return "", fmt.Errorf("Txscript.NewEngine %s", err)
	}
	if err := vm.Execute(); err != nil {
		return "", fmt.Errorf("Fail to sign tx %s", err)
	}

	// txToHex
	buf := bytes.NewBuffer(make([]byte, 0, tx.MsgTx().SerializeSize()))
	tx.MsgTx().Serialize(buf)
	signedTxHex := hex.EncodeToString(buf.Bytes())
	return signedTxHex, nil
}
