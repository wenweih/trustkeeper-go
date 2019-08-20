package repository

import (
  "context"
  "trustkeeper-go/app/service/chains_query/pkg/model"
)

func (repo *repo) QueryBalance(ctx context.Context, symbol, address string) (balance string, err error) {
  bal := model.Balance{}
  if err := repo.db.Where("symbol = ? AND address = ?", symbol, address).First(&bal).Error; err != nil {
    return "", err
  }
  return bal.Amount, nil
}
