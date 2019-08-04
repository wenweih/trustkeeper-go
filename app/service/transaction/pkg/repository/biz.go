package repository

import (
  "fmt"
  "strings"
  "context"
  "google.golang.org/grpc/metadata"
  "trustkeeper-go/app/service/transaction/pkg/model"
)

// IBiz dashboard service business logic
type IBiz interface {
  AssignAssetsToWallet(ctx context.Context, address string, assets []*SimpleAsset) (err error)
  CreateBalancesForAsset(ctx context.Context, wallets []*Wallet, asset *SimpleAsset) (err error)
  Close() error
}

// What's the fastest way to do a bulk insert into Postgres?
// https://stackoverflow.com/questions/758945/whats-the-fastest-way-to-do-a-bulk-insert-into-postgres
// bulk insert, bulk query with gorm https://zhiruchen.github.io/2017/08/31/bulk-insert-bulk-query-with-gorm/
func (repo *repo) CreateBalancesForAsset(ctx context.Context, wallets []*Wallet, asset *SimpleAsset) (error) {
  valueStrings := []string{}
  valueArgs := []interface{}{}
  for _, w := range wallets {
    valueStrings = append(valueStrings, "(?, ?, ?, ?)")

    valueArgs = append(valueArgs, w.Address)
    valueArgs = append(valueArgs, asset.Symbol)
    valueArgs = append(valueArgs, asset.Identify)
    valueArgs = append(valueArgs, asset.Decimal)
  }
  smt := `INSERT INTO balances(address, symbol, identify, decimal)
    VALUES %s ON CONFLICT (address, symbol) DO UPDATE SET address = excluded.address`
  smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))
  tx := repo.db.Begin()
  err := tx.Exec(smt, valueArgs...).Error
  if err != nil {
    tx.Rollback()
    return err
  }
  return tx.Commit().Error
}

func (repo *repo) AssignAssetsToWallet(ctx context.Context, address string, assets []*SimpleAsset) (err error) {
  _, _, _, err = extractAuthInfoFromContext(ctx)
  if err != nil {
    return err
  }
  tx := repo.db.Begin()
  for _, asset := range assets {
    balance := model.Balance{
      Address: address,
      Symbol: asset.Symbol,
      Identify: asset.Identify,
      Decimal: asset.Decimal}
    tx.Create(&balance)
  }
  return tx.Commit().Error
}

func (repo *repo) Close() error{
  return repo.close()
}

func extractAuthInfoFromContext(ctx context.Context) (string, string, []string, error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !ok {
    return "", "", nil, fmt.Errorf("fail to extract auth info from ctx")
  }
  if len(md["uid"]) < 1 {
    return "", "", nil, fmt.Errorf("uid empty")
  }
  if len(md["nid"]) < 1 {
    return "", "", nil, fmt.Errorf("nid empty")
  }

  if len(md["roles"]) < 1 {
    return "", "", nil, fmt.Errorf("roles empty")
  }

  return md["uid"][0], md["nid"][0], md["roles"], nil
}
