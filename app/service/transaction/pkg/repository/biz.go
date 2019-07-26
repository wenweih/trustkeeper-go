package repository

import (
  "fmt"
  "context"
  "google.golang.org/grpc/metadata"
  "trustkeeper-go/app/service/transaction/pkg/model"
)

// IBiz dashboard service business logic
type IBiz interface {
  AssignAssetsToWallet(ctx context.Context, address string, assets []*SimpleAsset) (err error)
  Close() error
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
