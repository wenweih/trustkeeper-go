package endpoint

import (
	"errors"
	"context"
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
