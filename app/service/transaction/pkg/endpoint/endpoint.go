package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	"trustkeeper-go/app/service/transaction/pkg/repository"
	service "trustkeeper-go/app/service/transaction/pkg/service"
)

// AssignAssetsToWalletRequest collects the request parameters for the AssignAssetsToWallet method.
type AssignAssetsToWalletRequest struct {
	Address string                    `json:"address"`
	Assets  []*repository.SimpleAsset `json:"assets"`
}

// AssignAssetsToWalletResponse collects the response parameters for the AssignAssetsToWallet method.
type AssignAssetsToWalletResponse struct {
	Err error `json:"err"`
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}

// MakeAssignAssetsToWalletEndpoint returns an endpoint that invokes AssignAssetsToWallet on the service.
func MakeAssignAssetsToWalletEndpoint(s service.TransactionService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AssignAssetsToWalletRequest)
		err := s.AssignAssetsToWallet(ctx, req.Address, req.Assets)
		if err != nil {
			return AssignAssetsToWalletResponse{Err: err}, err
		}
		return AssignAssetsToWalletResponse{}, nil
	}
}

// Failed implements Failer.
func (r AssignAssetsToWalletResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// AssignAssetsToWallet implements Service. Primarily useful in a client.
func (e Endpoints) AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error) {
	request := AssignAssetsToWalletRequest{
		Address: address,
		Assets:  assets,
	}
	_, err = e.AssignAssetsToWalletEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return nil
}