package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"trustkeeper-go/app/service/dashboard/pkg/model"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"
)

// IBiz dashboard service business logic
type IBiz interface {
	Signup(uuid, email, name, xpub string) error
	CreateGroup(ctx context.Context, m *model.Group) error
	Close() error
	GetGroups(ctx context.Context, query map[string]interface{}) (groups []*GetGroupsResp, err error)
	UpdateGroup(ctx context.Context, groupID, name, desc string) error
	QueryChainAsset(ctx context.Context, query map[string]interface{}) (chainAssets []*ChainAsset, err error)
	ChangeGroupAssets(ctx context.Context, chainAssets []*ChainAsset, groupid string) (result []*ChainAsset, err error)
	AddAsset(ctx context.Context, groupid, chainid, symbol, identify, decimal string) (asset *SimpleAsset, err error)
}

func (repo *repo) Signup(uuid, email, name, xpub string) error {
	return nil
}

func (repo *repo) CreateGroup(ctx context.Context, m *model.Group) error {
	uid, _, roles, err := extractAuthInfoFromContext(ctx)
	if err != nil {
		return err
	}

	allowRoles := make([]string, 0, len(roles))
	for _, role := range roles {
		allow := repo.iCasbinRepo.HasPolicy([]string{role, groupResource, "create"})
		if allow {
			allowRoles = append(allowRoles, role)
		}
	}

	if len(allowRoles) <= 0 {
		return fmt.Errorf("not allow")
	}

	tx := repo.db.Begin()
	if err := repo.iGroupRepo.Create(tx, m).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	repo.iCasbinRepo.AddReadWriteForRoleInDomain(uid, m.NamespaceID, strconv.FormatUint(uint64(m.ID), 10))
	return nil
}

func (repo *repo) Close() error {
	return repo.close()
}

func (repo *repo) GetGroups(ctx context.Context, query map[string]interface{}) (groupsResp []*GetGroupsResp, err error) {
	uid, nid, _, err := extractAuthInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ids := repo.iCasbinRepo.GetObjForUserInDomain(uid, nid, "read")
	groups, err := repo.iGroupRepo.Query(repo.db, ids, query)
	if err != nil {
		return nil, err
	}
	groupsResp = make([]*GetGroupsResp, len(groups))
	for i, group := range groups {
		groupsResp[i] = &GetGroupsResp{Name: group.Name, Desc: group.Desc, ID: strconv.FormatUint(uint64(group.ID), 10)}
	}
	return groupsResp, nil
}

func (repo *repo) UpdateGroup(ctx context.Context, groupID, name, desc string) error {
	uid, nid, _, err := extractAuthInfoFromContext(ctx)
	if err != nil {
		return err
	}
	allow := repo.iCasbinRepo.HasPolicy([]string{uid, nid, groupID, "write"})
	if !allow {
		return fmt.Errorf("not allow")
	}

	groups, err := repo.iGroupRepo.Query(repo.db, []interface{}{groupID}, nil)
	if err != nil {
		return err
	}
	if len(groups) != 1 {
		return fmt.Errorf("query group error")
	}

	tx := repo.db.Begin()
	group := groups[0]
	group.Name = name
	group.Desc = desc
	if err := repo.iGroupRepo.Update(tx, group).Error; err != nil {
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (repo *repo) QueryChainAsset(ctx context.Context, query map[string]interface{}) (chainAssets []*ChainAsset, err error) {
	uid, nid, _, err := extractAuthInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	ids := repo.iCasbinRepo.GetObjForUserInDomain(uid, nid, "read")
	chains, err := repo.iChainAssetRepo.Query(repo.db, ids, query)
	if err != nil {
		return nil, err
	}

	chainAssets = make([]*ChainAsset, len(chains))
	for i, c := range chains {
		assets := make([]*SimpleAsset, len(c.Assets))
		for it, t := range c.Assets {
			assets[it] = &SimpleAsset{
				AssetID:  strconv.FormatUint(uint64(t.ID), 10),
				Symbol:   t.Symbol,
				Identify: t.Identify,
				Status:   t.Status,
				Decimal:  t.Decimal}
		}
		chainAssets[i] = &ChainAsset{
			ChainID:      strconv.FormatUint(uint64(c.ID), 10),
			Name:         c.Name,
			Coin:         c.Coin,
			Status:       c.Status,
			Decimal:      c.Decimal,
			SimpleAssets: assets}
	}
	return chainAssets, nil
}

func (repo *repo) ChangeGroupAssets(ctx context.Context, chainAssets []*ChainAsset, groupid string) (result []*ChainAsset, err error) {
	uid, nid, roles, err := extractAuthInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := repo.createAuth(roles, chainAssetResource); err != nil {
		return nil, err
	}

	// groupChainAssets := make([]*model.Chain, len(chainAssets))
	tx := repo.db.Begin()
	for _, ca := range chainAssets {
		assets := []model.Asset{}
		if err := copier.Copy(&assets, &ca.SimpleAssets); err != nil {
			return nil, err
		}
		chain := &model.Chain{
			Name:    ca.Name,
			Coin:    ca.Coin,
			Status:  ca.Status,
			GroupID: groupid,
			Decimal: ca.Decimal}
		if len(ca.ChainID) > 1 {
			repo.iChainAssetRepo.Update(tx, chain)
		} else {
			repo.iChainAssetRepo.Create(tx, chain)
		}

		chainID := strconv.FormatUint(uint64(chain.ID), 10)
		if len(assets) <= 0 {
			asset := model.Asset{
				GroupID: groupid,
				Symbol:  ca.Coin,
				Status:  true,
				ChainID: chainID,
				Decimal: ca.Decimal}
			tx.Create(&asset)
			ca.SimpleAssets = append(ca.SimpleAssets, &SimpleAsset{
				AssetID: strconv.FormatUint(uint64(asset.ID), 10),
				Symbol:  asset.Symbol,
				Status:  asset.Status,
				Decimal: asset.Decimal})
		}

		repo.iCasbinRepo.AddReadWriteForRoleInDomain(uid, nid, chainID)
		ca.ChainID = chainID
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return chainAssets, nil
}

func (repo *repo) createAuth(roles []string, resource string) error {
	allowRoles := make([]string, 0, len(roles))
	for _, role := range roles {
		allow := repo.iCasbinRepo.HasPolicy([]string{role, resource, "create"})
		if allow {
			allowRoles = append(allowRoles, role)
		}
	}
	if len(allowRoles) <= 0 {
		return fmt.Errorf("not allow")
	}
	return nil
}

func (repo *repo) AddAsset(ctx context.Context, groupid, chainid, symbol, identify, decimal string) (*SimpleAsset, error) {
	uint64Decimal, err := strconv.ParseUint(decimal, 10, 64)
	if err != nil {
		return nil, err
	}
	asset := model.Asset{
		GroupID:  groupid,
		Symbol:   symbol,
		ChainID:  chainid,
		Status:   true,
		Identify: strings.ToLower(identify),
		Decimal:  uint64Decimal}
	if err := repo.db.Create(&asset).Error; err != nil {
		return nil, err
	}

	return &SimpleAsset{
		AssetID:  strconv.FormatUint(uint64(asset.ID), 10),
		Symbol:   asset.Symbol,
		Status:   asset.Status,
		Identify: asset.Identify,
		Decimal:  asset.Decimal}, nil
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
