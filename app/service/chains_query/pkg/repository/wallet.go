package repository

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"trustkeeper-go/app/service/chains_query/pkg/model"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func (repo *repo) QueryBalance(ctx context.Context, symbol, address string) (balance string, err error) {
	bal := model.Balance{}
	if err := repo.db.Where("symbol = ? AND address = ?", symbol, address).First(&bal).Error; err != nil {
		return "", err
	}
	balanceDecimal, err := decimal.NewFromString(bal.Amount)
	if err != nil {
		return "", err
	}
	withdrawLockDecimal, err := decimal.NewFromString(bal.WithdrawLock)
	if err != nil {
		return "", err
	}
	balanceDecimal = balanceDecimal.Sub(withdrawLockDecimal)

	decimalBig, result := new(big.Int).SetString(strconv.FormatUint(bal.Decimal, 10), 10)
	if !result {
		return "", fmt.Errorf("Fail to convert decimal")
	}
	return balanceDecimal.Div(decimal.NewFromBigInt(decimalBig, 0)).String(), nil
}

func (repo *repo) WalletValidate(ctx context.Context, chainName, address string) (err error) {
	switch chainName {
	case ChainNameBitcoincore:
		_, err := btcutil.DecodeAddress(address, &chaincfg.RegressionNetParams)
		if err != nil {
			return fmt.Errorf("Invalid address: %s", err.Error())
		}
		return nil
	case ChainNameEthereum:
		if !common.IsHexAddress(address) {
			return fmt.Errorf("Invalid address: %s", address)
		}
		return nil
	default:
		return fmt.Errorf("Unsupport Blockchain address")
	}
}
