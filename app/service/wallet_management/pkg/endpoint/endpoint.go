package endpoint

import (
	"context"
	"errors"
	"trustkeeper-go/app/service/wallet_management/pkg/repository"
	service "trustkeeper-go/app/service/wallet_management/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateChainRequest collects the request parameters for the CreateChain method.
type CreateChainRequest struct {
	Symbol  string `json:"symbol"`
	Bit44ID string `json:"bit44id"`
	Status  bool   `json:"status"`
}

// CreateChainResponse collects the response parameters for the CreateChain method.
type CreateChainResponse struct {
	Err error `json:"err"`
}

// MakeCreateChainEndpoint returns an endpoint that invokes CreateChain on the service.
func MakeCreateChainEndpoint(s service.WalletManagementService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateChainRequest)
		err := s.CreateChain(ctx, req.Symbol, req.Bit44ID, req.Status)
		return CreateChainResponse{Err: err}, err
	}
}

// Failed implements Failer.
func (r CreateChainResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateChain implements Service. Primarily useful in a client.
func (e Endpoints) CreateChain(ctx context.Context, symbol string, bit44ID string, status bool) error {
	request := CreateChainRequest{
		Bit44ID: bit44ID,
		Status:  status,
		Symbol:  symbol,
	}
	if _, err := e.CreateChainEndpoint(ctx, request); err != nil {
		return err
	}
	return nil
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}

// AssignedXpubToGroupRequest collects the request parameters for the UpdateXpubState method.
type AssignedXpubToGroupRequest struct {
	GroupID string `json:"groupid"`
}

// AssignedXpubToGroupResponse collects the response parameters for the UpdateXpubState method.
type AssignedXpubToGroupResponse struct {
	Err error `json:"err"`
}

// MakeAssignedXpubToGroupEndpoint returns an endpoint that invokes UpdateXpubState on the service.
func MakeAssignedXpubToGroupEndpoint(s service.WalletManagementService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(AssignedXpubToGroupRequest)
		if !ok {
			return nil, errors.New("endpoint AssignedXpubToGroupRequest type assersion error")
		}
		if err := s.AssignedXpubToGroup(ctx, req.GroupID); err != nil {
			return nil, err
		}
		return AssignedXpubToGroupResponse{}, nil
	}
}

// Failed implements Failer.
func (r AssignedXpubToGroupResponse) Failed() error {
	return r.Err
}

// AssignedXpubToGroup implements Service. Primarily useful in a client.
func (e Endpoints) AssignedXpubToGroup(ctx context.Context, groupid string) (err error) {
	request := AssignedXpubToGroupRequest{GroupID: groupid}
	_, err = e.AssignedXpubToGroupEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

// GetChainsRequest collects the request parameters for the GetChains method.
type GetChainsRequest struct{}

// GetChainsResponse collects the response parameters for the GetChains method.
type GetChainsResponse struct {
	Chains []*repository.SimpleChain `json:"chains"`
	Err    error                     `json:"err"`
}

// MakeGetChainsEndpoint returns an endpoint that invokes GetChains on the service.
func MakeGetChainsEndpoint(s service.WalletManagementService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		chains, err := s.GetChains(ctx)
		if err != nil {
			return GetChainsResponse{Err: err}, err
		}
		return GetChainsResponse{Chains: chains}, nil
	}
}

// Failed implements Failer.
func (r GetChainsResponse) Failed() error {
	return r.Err
}

// GetChains implements Service. Primarily useful in a client.
func (e Endpoints) GetChains(ctx context.Context) (chains []*repository.SimpleChain, err error) {
	request := GetChainsRequest{}
	response, err := e.GetChainsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetChainsResponse).Chains, nil
}

// CreateWalletRequest collects the request parameters for the CreateWallet method.
type CreateWalletRequest struct {
	Groupid     string `json:"groupid"`
	Chainname   string `json:"chainname"`
	Bip44change int    `json:"bip44change"`
}

// CreateWalletResponse collects the response parameters for the CreateWallet method.
type CreateWalletResponse struct {
	Wallet *repository.Wallet `json:"wallet"`
	Err    error              `json:"err"`
}

// MakeCreateWalletEndpoint returns an endpoint that invokes CreateWallet on the service.
func MakeCreateWalletEndpoint(s service.WalletManagementService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateWalletRequest)
		if !ok {
			return nil, errors.New("endpoint CreateWalletRequest type assertion error")
		}
		wallet, err := s.CreateWallet(ctx, req.Groupid, req.Chainname, req.Bip44change)
		if err != nil {
			return CreateWalletResponse{Err: err}, err
		}
		return CreateWalletResponse{Wallet: wallet}, nil
	}
}

// Failed implements Failer.
func (r CreateWalletResponse) Failed() error {
	return r.Err
}

// CreateWallet implements Service. Primarily useful in a client.
func (e Endpoints) CreateWallet(ctx context.Context, groupid string, chainname string, bip44change int) (wallet *repository.Wallet, err error) {
	request := CreateWalletRequest{
		Bip44change: bip44change,
		Chainname:   chainname,
		Groupid:     groupid,
	}
	resp, err := e.CreateWalletEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(CreateWalletResponse).Wallet, nil
}

// GetWalletsRequest collects the request parameters for the GetWallets method.
type GetWalletsRequest struct {
	Groupid string `json:"groupid"`
}

// GetWalletsResponse collects the response parameters for the GetWallets method.
type GetWalletsResponse struct {
	Wallets []*repository.Wallet `json:"wallets"`
	Err     error                `json:"err"`
}

// MakeGetWalletsEndpoint returns an endpoint that invokes GetWallets on the service.
func MakeGetWalletsEndpoint(s service.WalletManagementService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetWalletsRequest)
		if !ok {
			e := errors.New("endpoint GetWalletsRequest type assertion error")
			return GetWalletsResponse{Err: e}, e
		}
		wallets, err := s.GetWallets(ctx, req.Groupid)
		if err != nil {
			return GetWalletsResponse{Err: err}, err
		}
		return GetWalletsResponse{
			Wallets: wallets,
		}, nil
	}
}

// Failed implements Failer.
func (r GetWalletsResponse) Failed() error {
	return r.Err
}

// GetWallets implements Service. Primarily useful in a client.
func (e Endpoints) GetWallets(ctx context.Context, groupid string) (wallets []*repository.Wallet, err error) {
	request := GetWalletsRequest{Groupid: groupid}
	response, err := e.GetWalletsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(GetWalletsResponse).Wallets, nil
}
