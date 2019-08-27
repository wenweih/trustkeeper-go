package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

	walletKeyGrpcClient "trustkeeper-go/app/service/wallet_key/client"
	walletKeyRepository "trustkeeper-go/app/service/wallet_key/pkg/repository"
	walletKeyService "trustkeeper-go/app/service/wallet_key/pkg/service"

	txGrpcClient "trustkeeper-go/app/service/transaction/client"
	txRepository "trustkeeper-go/app/service/transaction/pkg/repository"
	txService "trustkeeper-go/app/service/transaction/pkg/service"

	chainsqueryGrpcClient "trustkeeper-go/app/service/chains_query/client"
	chainsqueryService "trustkeeper-go/app/service/chains_query/pkg/service"

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
	CreateWallet(ctx context.Context, groupid, chainname string, bip44change int) (id, address, respchainname string, status bool, err error)
	GetWallets(ctx context.Context, groupid string, page, limit, bip44change int) (wallets []*repository.ChainWithWallets, err error)
	QueryOmniProperty(ctx context.Context, identify string) (asset *repository.SimpleAsset, err error)
	EthToken(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error)
	CreateToken(ctx context.Context, groupid, chainid, symbol, identify, decimal, chainName string) (asset *repository.SimpleAsset, err error)
	SendBTCTx(ctx context.Context, from, to, amount string) (txid string, err error)
	SendETHTx(ctx context.Context, from, to, amount string) (txid string, err error)
	SendERC20Tx(ctx context.Context, from, to, amount, symbol string) (txid string, err error)
	QueryBalance(ctx context.Context, symbol, address string) (balance string, err error)
	WalletValidate(ctx context.Context, chainName, address string) (err error)
}

// Credentials Signup Signin params
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgName  string `json:"orgname"`
}

type basicWebapiService struct {
	accountSrv     accountService.AccountService
	dashboardSrv   dashboardService.DashboardService
	WalletSrv      walletManagementService.WalletManagementService
	txSrv          txService.TransactionService
	chainsQuerySrv chainsqueryService.ChainsQueryService
	KeySrv         walletKeyService.WalletKeyService
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

	txClient, err := txGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("txGrpcClient: %s", err.Error())
	}

	chainsqueryClient, err := chainsqueryGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("txGrpcClient: %s", err.Error())
	}

	wkClient, err := walletKeyGrpcClient.New(cfg.ConsulAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("walletKeyGrpcClient: %s", err.Error())
	}

	return &basicWebapiService{
		accountSrv:     accountClient,
		dashboardSrv:   dashboardClient,
		WalletSrv:      wmClient,
		txSrv:          txClient,
		chainsQuerySrv: chainsqueryClient,
		KeySrv:         wkClient,
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
				Name:    defaultChain.Name,
				Coin:    defaultChain.Coin,
				Decimal: defaultChain.Decimal,
				Status:  false}
			for _, ga := range groupAssets {
				if groupAssetsResp[di].Name != ga.Name &&
					groupAssetsResp[di].Coin != ga.Coin {
					continue
				}
				assets := make([]*repository.SimpleAsset, len(ga.SimpleAssets))
				for it, asset := range ga.SimpleAssets {
					assets[it] = &repository.SimpleAsset{
						AssetID:  asset.AssetID,
						Symbol:   asset.Symbol,
						Status:   asset.Status,
						Decimal:  asset.Decimal,
						Identify: asset.Identify}
				}
				groupAssetsResp[di].ChainID = ga.ChainID
				groupAssetsResp[di].Status = ga.Status
				groupAssetsResp[di].SimpleAssets = assets
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

func (b *basicWebapiService) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (string, string, string, bool, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", "", "", false, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	wallet, err := b.WalletSrv.CreateWallet(ctxWithAuthInfo, groupid, chainname, bip44change)
	if err != nil {
		return "", "", "", false, err
	}
	groupAssets, err := b.dashboardSrv.GetGroupAssets(ctxWithAuthInfo, groupid)
	if err != nil {
		return "", "", "", false, err
	}
	txRepoAssets := make([]*txRepository.SimpleAsset, 0)
	for _, groupAsset := range groupAssets {
		if groupAsset.Name == chainname {
			if err := copier.Copy(&txRepoAssets, &groupAsset.SimpleAssets); err != nil {
				return "", "", "", false, err
			}
			break
		}
	}
	if len(txRepoAssets) >= 1 {
		err := b.txSrv.AssignAssetsToWallet(ctxWithAuthInfo, wallet.Address, txRepoAssets)
		if err != nil {
			return "", "", "", false, err
		}
	}
	return wallet.ID, wallet.Address, wallet.ChainName, wallet.Status, nil
}

func (b *basicWebapiService) GetWallets(ctx context.Context, groupid string, page int, limit, bip44change int) ([]*repository.ChainWithWallets, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	wallets, err := b.WalletSrv.GetWallets(ctxWithAuthInfo, groupid, int32(page), int32(limit), int32(bip44change))
	if err != nil {
		return nil, err
	}

	respwallets := []*repository.ChainWithWallets{}
	if err := copier.Copy(&respwallets, &wallets); err != nil {
		return nil, err
	}
	return respwallets, nil
}

func (b *basicWebapiService) QueryOmniProperty(ctx context.Context, identify string) (asset *repository.SimpleAsset, err error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	propertyid, err := strconv.ParseInt(identify, 10, 64)
	if err != nil {
		return nil, err
	}
	omniProperty, err := b.chainsQuerySrv.QueryOmniProperty(ctxWithAuthInfo, propertyid)
	if err != nil {
		return nil, err
	}
	return &repository.SimpleAsset{
		Symbol:   omniProperty.Name,
		Identify: strconv.FormatInt(omniProperty.Propertyid, 10),
		Decimal:  100000000}, err
}

func (b *basicWebapiService) EthToken(ctx context.Context, tokenHex string) (*repository.ERC20Token, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	token, err := b.chainsQuerySrv.ERC20TokenInfo(ctxWithAuthInfo, tokenHex)
	if err != nil {
		return nil, err
	}
	respToken := repository.ERC20Token{}
	if err := copier.Copy(&respToken, token); err != nil {
		return nil, err
	}
	return &respToken, err
}

func (b *basicWebapiService) CreateToken(
	ctx context.Context, groupid string,
	chainid string, symbol string,
	identify string, decimal string, chainName string) (*repository.SimpleAsset, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return nil, err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	asset, err := b.dashboardSrv.AddAsset(ctxWithAuthInfo, groupid, chainid, symbol, identify, decimal)
	if err != nil {
		return nil, err
	}
	wallets, err := b.WalletSrv.QueryWalletsForGroupByChainName(ctxWithAuthInfo, groupid, chainName)
	if err != nil {
		return nil, err
	}

	txSrvAsset := txRepository.SimpleAsset{}
	if err := copier.Copy(&txSrvAsset, asset); err != nil {
		return nil, err
	}

	respAsset := repository.SimpleAsset{}
	if err := copier.Copy(&respAsset, asset); err != nil {
		return nil, err
	}

	txSrvWallets := make([]*txRepository.Wallet, 0)
	if err := copier.Copy(&txSrvWallets, &wallets); err != nil {
		return nil, err
	}

	err = b.txSrv.CreateBalancesForAsset(ctxWithAuthInfo, txSrvWallets, &txSrvAsset)
	if err != nil {
		return nil, err
	}
	return &respAsset, nil
}

func (b *basicWebapiService) SendBTCTx(ctx context.Context, from string, to string, amount string) (string, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	hd, err := b.WalletSrv.QueryWalletHD(ctxWithAuthInfo, from)
	if err != nil {
		return "", err
	}
	unsignedTxHex, vinAmount, err := b.chainsQuerySrv.ConstructTxBTC(ctxWithAuthInfo, from, to, amount)
	if err != nil {
		return "", err
	}
	walletHD := walletKeyRepository.WalletHD{}
	if err := copier.Copy(&walletHD, hd); err != nil {
		return "", err
	}
	signedTxHex, err := b.KeySrv.SignedBitcoincoreTx(ctxWithAuthInfo, walletHD, unsignedTxHex, vinAmount)
	if err != nil {
		return "", err
	}
	txid, err := b.chainsQuerySrv.SendBTCTx(ctxWithAuthInfo, signedTxHex)
	if err != nil {
		return "", err
	}
	return txid, err
}

func (b *basicWebapiService) QueryBalance(ctx context.Context, symbol string, address string) (balance string, err error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	if _, err := b.WalletSrv.QueryWalletHD(ctxWithAuthInfo, address); err != nil {
		return "", err
	}
	balance, err = b.chainsQuerySrv.QueryBalance(ctxWithAuthInfo, symbol, address)
	return
}

func (b *basicWebapiService) WalletValidate(ctx context.Context, chainName string, address string) (err error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	return b.chainsQuerySrv.WalletValidate(ctxWithAuthInfo, chainName, address)
}

func (b *basicWebapiService) SendETHTx(ctx context.Context, from string, to string, amount string) (string, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	hd, err := b.WalletSrv.QueryWalletHD(ctxWithAuthInfo, from)
	if err != nil {
		return "", err
	}
	unsignedTxHex, chainID, err := b.chainsQuerySrv.ConstructTxETH(ctxWithAuthInfo, from, to, amount)
	if err != nil {
		return "", err
	}
	walletHD := walletKeyRepository.WalletHD{}
	if err := copier.Copy(&walletHD, hd); err != nil {
		return "", err
	}
	signedTxHex, err := b.KeySrv.SignedEthereumTx(ctxWithAuthInfo, walletHD, unsignedTxHex, chainID)
	if err != nil {
		return "", err
	}
	txid, err := b.chainsQuerySrv.SendETHTx(ctxWithAuthInfo, signedTxHex)
	if err != nil {
		return "", err
	}
	return txid, err
}

func (b *basicWebapiService) SendERC20Tx(ctx context.Context, from string, to string, amount string, symbol string) (string, error) {
	uid, nid, roles, err := b.auth(ctx)
	if err != nil {
		return "", err
	}
	ctxWithAuthInfo := constructAuthInfoContext(ctx, roles, uid, nid)
	hd, err := b.WalletSrv.QueryWalletHD(ctxWithAuthInfo, from)
	if err != nil {
		return "", err
	}
	unsignedTxHex, chainID, err := b.chainsQuerySrv.ConstructTxERC20(ctxWithAuthInfo, from, to, amount, symbol)
	if err != nil {
		return "", err
	}
	walletHD := walletKeyRepository.WalletHD{}
	if err := copier.Copy(&walletHD, hd); err != nil {
		return "", err
	}
	signedTxHex, err := b.KeySrv.SignedEthereumTx(ctxWithAuthInfo, walletHD, unsignedTxHex, chainID)
	if err != nil {
		return "", err
	}
	txid, err := b.chainsQuerySrv.SendETHTx(ctxWithAuthInfo, signedTxHex)
	if err != nil {
		return "", err
	}
	return txid, err
}
