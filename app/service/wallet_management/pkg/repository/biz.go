package repository

import (
  "fmt"
  "strconv"
  "context"
  "github.com/gomodule/redigo/redis"
  "trustkeeper-go/app/service/wallet_management/pkg/model"
  libctx "trustkeeper-go/library/context"
  "github.com/btcsuite/btcutil/hdkeychain"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/btcsuite/btcd/chaincfg"
  "github.com/jinzhu/gorm"
  "github.com/jinzhu/copier"
)

type IBiz interface {
  // bip44ThirdXpubsForChains 参数是 slice 的引用传递 [] 里面有具体数字是数组的值传递
  // https://leileiluoluo.com/posts/golang-method-calling-value-copy-or-reference-copy.html
  Signup(uid, nid, version string, bip44ThirdXpubsForChains []*Bip44ThirdXpubsForChain) error
  Close() error
  RedisInstance() *redis.Pool
  GetChains(ctx context.Context, query map[string]interface{}) (chains []*SimpleChain, err error)
  UpdateXpubState(ctx context.Context, from, to, groupid string) error
  CreateWallet(ctx context.Context, groupid, chainname string, bip44change int) (wallet *Wallet, err error)
  GetWallets(ctx context.Context, groupid string, page, limit, bip44change int32) (wallets []*ChainWithWallets, err error)
  QueryWalletsForGroupByChainName(ctx context.Context, groupid, chainName string) (wallets []*Wallet, err error)
}

type Bip44AccountKey struct {
	Account int     `json:"account"`
	Key     string  `json:"key"`
}

type Bip44ThirdXpubsForChain struct {
	Chain   uint           `json:"chain"`
	Xpubs   []*Bip44AccountKey	`json:"xpubs"`
}

func (repo *repo) Signup(uid, nid, version string, bip44ThirdXpubsForChains []*Bip44ThirdXpubsForChain) error {
  tx := repo.db.Begin()
  mnemonicVersion := model.MnemonicVersion{Version: version}
  if err := repo.imnemonicVersionRepo.Create(tx, &mnemonicVersion).Error; err != nil {
    tx.Rollback()
    return err
  }
  for _, bip44ThirdXpubsForChain := range bip44ThirdXpubsForChains {
    for _, xpub := range bip44ThirdXpubsForChain.Xpubs {
      mXpub := model.Xpub{
        Key: xpub.Key,
        Bip44ChainID: &bip44ThirdXpubsForChain.Chain,
        Bip44Account: xpub.Account,
        MnemonicVersionID: mnemonicVersion.ID}
      if err := repo.iXpubRepo.Create(tx, &mXpub).Error; err != nil {
        tx.Rollback()
        return err
      }
      repo.iCasbinRepo.AddActionForRoleInDomain(uid, nid, strconv.FormatUint(uint64(mXpub.ID), 10), "read")
      repo.iCasbinRepo.AddActionForRoleInDomain(uid, nid, strconv.FormatUint(uint64(mXpub.ID), 10), "write")
    }
  }
  if err := tx.Commit().Error; err != nil {
    return err
  }
  repo.iCasbinRepo.AddActionForRoleInDomain(uid, nid, walletResource, "create")
  repo.iCasbinRepo.AddActionForRoleInDomain(uid, nid, strconv.FormatUint(uint64(mnemonicVersion.ID), 10), "read")
  repo.iCasbinRepo.AddActionForRoleInDomain(uid, nid, strconv.FormatUint(uint64(mnemonicVersion.ID), 10), "write")
  return nil
}

func (repo *repo)Close() error{
  return repo.close()
}

func (repo *repo) RedisInstance() *redis.Pool {
  return repo.redisPool
}

func (repo *repo) GetChains(ctx context.Context, query map[string]interface{}) ([]*SimpleChain, error) {
  chains, err := repo.iChainRepo.Query(repo.db, query)
  if err != nil {
    return nil, err
  }
  simpleChains := make([]*SimpleChain, len(chains))
  for i, c := range chains {
    simpleChains[i] = &SimpleChain{
      ID: strconv.FormatUint(uint64(c.ID), 10),
      Name: c.Name,
      Coin: c.Coin,
      Bip44id: c.Bip44id,
      Status: c.Status,
      Decimal: c.Decimal}
  }
  return simpleChains, nil
}

func (repo *repo) UpdateXpubState(ctx context.Context, from, to, groupid string) error {
  _, nid, _, err := libctx.ExtractAuthInfoFromContext(ctx)
  if err != nil {
    return err
  }

  mnemonicVs, err := repo.imnemonicVersionRepo.VersionLikeQuery(repo.db, nid)
  if err != nil {
    return err
  }

  if len(mnemonicVs) != 1 {
    return fmt.Errorf("records error")
  }

  chains, err := repo.iChainRepo.Query(repo.db, map[string]interface{}{})
  if err != nil {
    return err
  }

  tx := repo.db.Begin()
  for _, chain := range chains {
    xpub := model.Xpub{}
    switch from {
    case Idle:
      tx.Where("state = ? AND mnemonic_version_id = ? AND bip44_chain_id = ?", Idle, uint(mnemonicVs[0].ID), chain.Bip44id).First(&xpub)
    case Assigned:
      tx.Where("state = ? AND mnemonic_version_id = ?", Assigned, uint(mnemonicVs[0].ID)).First(&xpub)
    case Abandon:
      tx.Where("state = ? AND mnemonic_version_id = ?", Abandon, uint(mnemonicVs[0].ID)).First(&xpub)
    default:
      return fmt.Errorf("invalid state:" + from)
    }
    if err := repo.iXpubRepo.UpdateState(tx, &xpub, to); err != nil {
      return err
    }
    if err := tx.Model(&xpub).Updates(map[string]interface{}{"state": to, "group_id": groupid}).Error; err != nil {
      return err
    }
  }
  return tx.Commit().Error
}

func (repo *repo) CreateWallet(ctx context.Context, groupid, chainname string, bip44change int) (wallet *Wallet, err error) {
  uid, nid, _, err := libctx.ExtractAuthInfoFromContext(ctx)
  if err != nil {
    return nil, err
  }
  if allow := repo.iCasbinRepo.HasPolicy([]string{uid, nid, walletResource, "create"}); allow != true {
    return nil, fmt.Errorf("not allow")
  }
  chains, err := repo.iChainRepo.Query(repo.db, map[string]interface{}{"name": chainname})
  if err != nil {
    return nil, err
  }
  if len(chains) != 1 {
    return nil, fmt.Errorf("fail to query chain record")
  }
  xpubs := []*model.Xpub{}
  repo.db.Where("group_id = ? AND bip44_chain_id = ?", groupid, uint(chains[0].Bip44id)).Find(&xpubs)

  if len(xpubs) != 1 {
    return nil, fmt.Errorf("fail to query xpub record")
  }

  extendedKey, err := hdkeychain.NewKeyFromString(xpubs[0].Key)
  if err != nil {
    return nil, err
  }
  changeLevel, err := extendedKey.Child(uint32(bip44change))
  if err != nil {
    return nil, err
  }

  maxIndexWallet := model.Wallet{}
  if err := repo.db.Order("bip44_index asc").
    Where("xpub_uid = ? AND bip44_change = ?", xpubs[0].ID, bip44change).
    Find(&maxIndexWallet).Error; err != nil && err != gorm.ErrRecordNotFound {
    return nil, err
  }
  bip44index := uint32(1)
  if maxIndexWallet.ID > 0 {
    bip44index = maxIndexWallet.Bip44Index + 1
  }

  addressIndexLevel, err := changeLevel.Child(bip44index)
  pubKey, err := addressIndexLevel.ECPubKey()
  if err != nil {
    return nil, err
  }
  address := ""
  switch chains[0].Bip44id {
  case 0:
    btcAddress, err := addressIndexLevel.Address(&chaincfg.RegressionNetParams)
    if err != nil {
      return nil, err
    }
    address = btcAddress.String()
  case 60:
    address = crypto.PubkeyToAddress(*pubKey.ToECDSA()).String()
  }
  mWallet := model.Wallet{
    Bip44Change: bip44change,
    Address: address,
    Bip44Index: bip44index,
    XpubUID: xpubs[0].ID,
    Status: true}
  if err := repo.db.Create(&mWallet).Error; err != nil {
    return nil, err
  }
  return &Wallet{
    ID: strconv.FormatUint(uint64(mWallet.ID), 10),
    Address: mWallet.Address,
    Status: mWallet.Status,
    ChainName: chains[0].Name}, nil
}

func (repo *repo) GetWallets(ctx context.Context, groupid string, page, limit, bip44change int32) ([]*ChainWithWallets, error) {
  uid, nid, _, err := libctx.ExtractAuthInfoFromContext(ctx)
  if err != nil {
    return nil, err
  }

  xpubs := []*model.Xpub{}
  if len(groupid) > 0 {
    // if repo.iCasbinRepo.HasPolicy([]string{uid, nid, groupid, "read"}) {
    //   return nil, fmt.Errorf("not allow")
    // }
    err := repo.db.Preload("Chain").Where("group_id = ? AND state = ?", groupid, Assigned).Find(&xpubs).Error
    if err != nil {
      return nil, err
    }
  } else {
    mnemonicVs, err := repo.imnemonicVersionRepo.VersionLikeQuery(repo.db, nid)
    if err != nil {
      return nil, err
    }
    if len(mnemonicVs) != 1 {
      return nil, fmt.Errorf("fail to query mnemonicVersion")
    }
    if err := repo.db.Preload("Chain").Where("state = ? AND mnemonic_version_id = ?", Assigned, uint(mnemonicVs[0].ID)).
      Find(&xpubs).Error; err != nil {
        return nil, err
      }
  }
  chainWithWallets := make([]*ChainWithWallets, len(xpubs))
  // chainWithWallets := []*ChainWithWallets{}
  for i, xpub := range xpubs {
    if !repo.iCasbinRepo.HasPolicy([]string{uid, nid, strconv.FormatUint(uint64(xpub.ID), 10), "read"}) {
      continue
    }
    filterData := []*model.Wallet{}
    xpubWallets := []*model.Wallet{}
    // limit offset usage https://github.com/jinzhu/gorm/issues/1752#issuecomment-454457879
    if err := repo.db.
      Order("created_at desc").
      Model(xpub).
      Where("bip44_change = ?", bip44change).
      Related(&xpubWallets, "Wallets").
      Limit(limit).
      Offset((page - 1) * limit).
      Related(&filterData, "Wallets").
      Error; err != nil {
        return nil, err
      }

    wallets :=[]*Wallet{}
  	if err := copier.Copy(&wallets, &filterData); err != nil {
  		return nil, err
  	}

    chainWithWallets[i] = &ChainWithWallets{
            ChainName: xpub.Chain.Name,
            TotalSize: int32(len(xpubWallets)),
            Wallets: wallets}
  }
  return chainWithWallets, nil
}
