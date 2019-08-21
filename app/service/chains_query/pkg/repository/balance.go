package repository

import (
  "fmt"
  "strconv"
  "context"
  "math/big"
  "github.com/shopspring/decimal"
  "trustkeeper-go/app/service/chains_query/pkg/model"
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
  decimalBig, result := new(big.Int).SetString(strconv.FormatUint(bal.Decimal, 10), 10)
  if !result {
    return "", fmt.Errorf("Fail to convert decimal")
  }
  return balanceDecimal.Div(decimal.NewFromBigInt(decimalBig, 0)).String(), nil
}
