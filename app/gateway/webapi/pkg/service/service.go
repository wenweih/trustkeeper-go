package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	// "github.com/jinzhu/gorm"
	accountService "trustkeeper-go/app/service/account/pkg/service"
	dashboardService "trustkeeper-go/app/service/dashboard/pkg/service"

	"trustkeeper-go/app/gateway/webapi/pkg/repository"
	accountGrpcClient "trustkeeper-go/app/service/account/client"
	dashboardGrpcClient "trustkeeper-go/app/service/dashboard/client"
	dashboardRepository "trustkeeper-go/app/service/dashboard/pkg/repository"

	walletManagementGrpcClient "trustkeeper-go/app/service/wallet_management/client"
	walletManagementService "trustkeeper-go/app/service/wallet_management/pkg/service"

	"github.com/caarlos0/env"
	log "github.com/go-kit/kit/log"
	"github.com/jinzhu/copier"
)

// WebapiService describes the service.
type WebapiService interface {
	Signup(ctx context.Context, user Credentials) (result bool, err error)
	Signin(ctx context.Context, user Credentials) (token string, err error)
	Signout(ctx context.Context) (result bool, err error)
	GetRoles(ctx context.Context) ([]string, error)
	UserInfo(ctx context.Context) (roles []string, orgName string, err error)
	GetGroups(ctx context.Context) (groups []*repository.GetGroupsResp, err error)
	CreateGroup(ctx context.Context, name, desc string) (group *repository.GetGroupsResp, err error)
	UpdateGroup(ctx context.Context, groupid, name, desc string) (err error)
	GetGroupAssets(ctx context.Context, groupid string) (groupAssets []*repository.GroupAsset, err error)
	ChangeGroupAssets(ctx context.Context, chainAssets []*repository.GroupAsset, groupid string) (result []*repository.GroupAsset, err error)
	CreateWallet(ctx context.Context, groupid, chainname string, bip44change int) (id, address string, status bool, err error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"orgname"`
}

type basicWebapiService struct {
	accountSrv   accountService.AccountService
	dashboardSrv dashboardService.DashboardService
	WalletSrv    walletManagementService.WalletManagementService
}

func makeError(ctx context.Context, err error) error {
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "violates unique constraint"):
		re := regexp.MustCompile(`[(](.*?)[)]`)
		result := re.FindAll([]byte(err.Error()), -1)
		return errors.New("Fields exist: " + string(result[0]))
	default:
		return err
	}
}

func (b *basicWebapiService) auth(ctx context.Context) (accountUID string, namespaceID string, roles []string, err error) {
	uid, nid, roles, err := b.accountSrv.Auth(ctx)
	if err != nil {
		return "", "", nil, err
	}
	namespaceID = nid
	accountUID = uid
	return
}

func (b *basicWebapiService) Signup(ctx context.Context, user Credentials) (bool, error) {
	if _, err := b.accountSrv.Create(ctx,
		user.Email,
		user.Password,
		user.OrgName); err != nil {
		e := makeError(ctx, err)
		return false, e
	}
	return true, nil
}

func (b *basicWebapiService) Signin(ctx context.Context, user Credentials) (token string, err error) {
	token, err = b.accountSrv.Signin(ctx, user.Email, user.Password)
	return
}
func (b *basicWebapiService) Signout(ctx context.Context) (result bool, err error) {
	if err := b.accountSrv.Signout(ctx); err != nil {
		return false, err
	}
	return true, nil
}

func (b *basicWebapiService) GetRoles(ctx context.Context) (s0 []string, e1 error) {
	roles, err := b.accountSrv.Roles(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (b *basicWebapiService) UserInfo(ctx context.Context) (roles []string, orgName string, err error) {
	return b.accountSrv.UserInfo(ctx)
}

func (b *basicWebapiService) GetGroups(ctx context.Context) ([]*repository.GetGroupsResp, error) {
	accountUID, namespaceID, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, accountUID, namespaceID)
	groups, err := b.dashboardSrv.GetGroups(ctxWithAuthInfo, namespaceID)
	if err != nil {
		return nil, err
	}
	lgroups := []*repository.GetGroupsResp{}
	if err := copier.Copy(&lgroups, &groups); err != nil {
		return nil, err
	}
	return lgroups, nil
}

// NewBasicWebapiService returns a naive, stateless implementation of WebapiService.
func NewBasicWebapiService(logger log.Logger) (WebapiService, error) {
	type config struct {
		ConsulAddr string `end:"consuladdr"`
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("fail to parse env: " + err.Error())
	}

	accountClient, err := accountGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("accountGrpcClient: %s", err.Error())
	}

	dashboardClient, err := dashboardGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("dashboardGrpcClient: %s", err.Error())
	}

	wmClient, err := walletManagementGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("walletManagementGrpcClient: %s", err.Error())
	}

	return &basicWebapiService{
		accountSrv:   accountClient,
		dashboardSrv: dashboardClient,
		WalletSrv:    wmClient,
	}, nil
}

// New returns a WebapiService with all of the expected middleware wired in.
func New(logger log.Logger, middleware []Middleware) (WebapiService, error) {
	service, err := NewBasicWebapiService(logger)
	if err != nil {
		return nil, err
	}
	var svc WebapiService = service
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}

func (b *basicWebapiService) CreateGroup(ctx context.Context, name string, desc string) (*repository.GetGroupsResp, error) {
	accountUID, namespaceid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := context.WithValue(ctx, "auth",
		struct {
			Roles []string
			UID   string
			NID   string
		}{roles, accountUID, namespaceid})

	resp, err := b.dashboardSrv.CreateGroup(ctxWithAuthInfo, accountUID, name, desc, namespaceid)
	if err != nil {
		return nil, err
	}

	if err := b.WalletSrv.AssignedXpubToGroup(ctxWithAuthInfo, resp.ID); err != nil {
		return nil, err
	}

	return &repository.GetGroupsResp{Name: resp.Name, Desc: resp.Desc, ID: resp.ID}, nil
}

func (b *basicWebapiService) UpdateGroup(ctx context.Context, groupid string, name string, desc string) (err error) {
	accountUID, namespaceid, roles, err := b.auth(ctx)
	if err != nil {
		return err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, accountUID, namespaceid)
	err = b.dashboardSrv.UpdateGroup(ctxWithAuthInfo, groupid, name, desc)
	return err
}

func (b *basicWebapiService) GetGroupAssets(ctx context.Context, groupid string) ([]*repository.GroupAsset, error) {
	accountUID, namespaceid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, accountUID, namespaceid)
	groupAssets, err := b.dashboardSrv.GetGroupAssets(ctxWithAuthInfo, groupid)
	if err != nil && !strings.Contains(err.Error(), "Empty records") {
		return nil, err
	}
	defaultChains, err := b.WalletSrv.GetChains(ctxWithAuthInfo)
	if err != nil {
		return nil, err
	}

	groupAssetsResp := make([]*repository.GroupAsset, len(defaultChains))
	for di, defaultChain := range defaultChains {
		if defaultChain.Status {
			groupAssetsResp[di] = &repository.GroupAsset{
				Name:   defaultChain.Name,
				Coin:   defaultChain.Coin,
				Status: false}
			for _, ga := range groupAssets {
				if groupAssetsResp[di].Name != ga.Name && groupAssetsResp[di].Coin != ga.Coin {
					continue
				}
				tokens := make([]*repository.SimpleToken, 0, len(ga.SimpleTokens))
				for it, token := range ga.SimpleTokens {
					tokens[it] = &repository.SimpleToken{TokenID: token.TokenID,
						Symbol: token.Symbol,
						Status: token.Status}
				}
				groupAssetsResp[di].Chainid = ga.Chainid
				groupAssetsResp[di].Status = ga.Status
				groupAssetsResp[di].SimpleTokens = tokens
			}
		}
	}
	return groupAssetsResp, nil
}

func constructAuthInfoContext(ctx context.Context, roles []string, uid, nid string) (ctxWithAuthInfo context.Context) {
	ctxWithAuthInfo = context.WithValue(ctx, "auth",
		struct {
			Roles []string
			UID   string
			NID   string
		}{roles, uid, nid})
	return
}

func (b *basicWebapiService) ChangeGroupAssets(ctx context.Context, chainAssets []*repository.GroupAsset, groupid string) (result []*repository.GroupAsset, err error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	dashboardRepo := []*dashboardRepository.ChainAsset{}
	if err := copier.Copy(&dashboardRepo, &chainAssets); err != nil {
		return nil, err
	}

	groupChainAssets, err := b.dashboardSrv.ChangeGroupAssets(ctxWithAuthInfo, dashboardRepo, groupid)
	if err != nil {
		return nil, err
	}
	result = []*repository.GroupAsset{}
	if err := copier.Copy(&result, &groupChainAssets); err != nil {
		return nil, err
	}
	return result, nil
}

func (b *basicWebapiService) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (id string, address string, status bool, err error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", "", false, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	wallet, err := b.WalletSrv.CreateWallet(ctxWithAuthInfo, groupid, chainname, bip44change)
	if err != nil {
		return "", "", false, err
	}
	return wallet.ID, wallet.Address, wallet.Status, nil
}
