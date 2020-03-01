package repository

import (
	"fmt"
	"sync"
	"trustkeeper-go/app/service/chains_query/pkg/model"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/coinset"
)

// SimpleCoin implements coinset Coin interface
type SimpleCoin struct {
	TxHash     *chainhash.Hash
	TxIndex    uint32
	TxValue    btcutil.Amount
	TxNumConfs int64
}

// Hash implements coinset Coin interface
func (c *SimpleCoin) Hash() *chainhash.Hash { return c.TxHash }

// Index implements coinset Coin interface
func (c *SimpleCoin) Index() uint32 { return c.TxIndex }

// Value implements coinset Coin interface
func (c *SimpleCoin) Value() btcutil.Amount { return c.TxValue }

// PkScript implements coinset Coin interface
func (c *SimpleCoin) PkScript() []byte { return nil }

// NumConfs implements coinset Coin interface
func (c *SimpleCoin) NumConfs() int64 { return c.TxNumConfs }

// ValueAge implements coinset Coin interface
func (c *SimpleCoin) ValueAge() int64 { return int64(c.TxValue) * c.TxNumConfs }

// CoinSelect btc tx inputs
func CoinSelect(chainHeader int64, txAmount btcutil.Amount, utxos []model.BtcUtxo) ([]model.BtcUtxo, []model.BtcUtxo, coinset.Coins, error) {
	var coins []coinset.Coin
	for _, utxo := range utxos {
		txHash, err := chainhash.NewHashFromStr(utxo.Txid)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Convert utxo hexTxid to txHash %s", err)
		}
		amount, err := btcutil.NewAmount(utxo.Amount)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("Convert utxo amount(float64) to btc amount(int64 as Satoshi) %s", err)
		}
		coins = append(coins, coinset.Coin(&SimpleCoin{TxHash: txHash, TxIndex: utxo.VoutIndex, TxValue: amount, TxNumConfs: chainHeader - utxo.Height + 1}))
	}

	selector := &coinset.MaxValueAgeCoinSelector{
		MaxInputs:       50,
		MinChangeAmount: 10000,
	}

	selectedCoins, err := selector.CoinSelect(txAmount, coins)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("CoinSelect error: %s", err)
	}
	scoins := selectedCoins.Coins()

	var selectedUTXOs, unSelectedUTXOs []model.BtcUtxo
	var selectedUTXOMap = struct {
		sync.RWMutex
		m map[uint]model.BtcUtxo
	}{m: make(map[uint]model.BtcUtxo)}

	for _, coin := range scoins {
		for _, utxo := range utxos {
			if coin.Hash().String() == utxo.Txid && coin.Index() == utxo.VoutIndex {
				selectedUTXOMap.Lock()
				selectedUTXOMap.m[utxo.ID] = utxo
				selectedUTXOMap.Unlock()
			}
		}
	}
	// selectedUTXO
	for _, v := range selectedUTXOMap.m {
		utxo := v
		selectedUTXOs = append(selectedUTXOs, utxo)
	}

	// unSelectedUTXOs
	for _, utxo := range utxos {
		_, found := selectedUTXOMap.m[utxo.ID]
		if !found {
			unSelectedUTXOs = append(unSelectedUTXOs, utxo)
		}
	}
	return selectedUTXOs, unSelectedUTXOs, selectedCoins, nil
}
