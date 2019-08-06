package endpoint

import (
	"context"
	"trustkeeper-go/app/service/chains_query/pkg/repository"
	service "trustkeeper-go/app/service/chains_query/pkg/service"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	endpoint "github.com/go-kit/kit/endpoint"
)

// BitcoincoreBlockRequest collects the request parameters for the BitcoincoreBlock method.
type BitcoincoreBlockRequest struct {
	BlockHash *chainhash.Hash `json:"block_hash"`
}

// BitcoincoreBlockResponse collects the response parameters for the BitcoincoreBlock method.
type BitcoincoreBlockResponse struct {
	B0 *btcjson.GetBlockVerboseResult `json:"b0"`
	E1 error                          `json:"e1"`
}

// MakeBitcoincoreBlockEndpoint returns an endpoint that invokes BitcoincoreBlock on the service.
func MakeBitcoincoreBlockEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BitcoincoreBlockRequest)
		b0, e1 := s.BitcoincoreBlock(ctx, req.BlockHash)
		return BitcoincoreBlockResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r BitcoincoreBlockResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// BitcoincoreBlock implements Service. Primarily useful in a client.
func (e Endpoints) BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (b0 *btcjson.GetBlockVerboseResult, e1 error) {
	request := BitcoincoreBlockRequest{BlockHash: blockHash}
	response, err := e.BitcoincoreBlockEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BitcoincoreBlockResponse).B0, response.(BitcoincoreBlockResponse).E1
}

// QueryOmniPropertyRequest collects the request parameters for the QueryOmniProperty method.
type QueryOmniPropertyRequest struct {
	Propertyid int64 `json:"propertyid"`
}

// QueryOmniPropertyResponse collects the response parameters for the QueryOmniProperty method.
type QueryOmniPropertyResponse struct {
	Property *repository.OmniProperty `json:"r0"`
	Err      error                    `json:"e1"`
}

// MakeQueryOmniPropertyEndpoint returns an endpoint that invokes QueryOmniProperty on the service.
func MakeQueryOmniPropertyEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryOmniPropertyRequest)
		resp, err := s.QueryOmniProperty(ctx, req.Propertyid)
		if err != nil {
			return QueryOmniPropertyResponse{Err: err}, err
		}
		return QueryOmniPropertyResponse{Property: resp}, nil
	}
}

// Failed implements Failer.
func (r QueryOmniPropertyResponse) Failed() error {
	return r.Err
}

// QueryOmniProperty implements Service. Primarily useful in a client.
func (e Endpoints) QueryOmniProperty(ctx context.Context, propertyid int64) (r0 *repository.OmniProperty, e1 error) {
	request := QueryOmniPropertyRequest{Propertyid: propertyid}
	response, err := e.QueryOmniPropertyEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(QueryOmniPropertyResponse).Property, nil
}

// ERC20TokenInfoRequest collects the request parameters for the ERC20TokenInfo method.
type ERC20TokenInfoRequest struct {
	TokenHex string `json:"token_hex"`
}

// ERC20TokenInfoResponse collects the response parameters for the ERC20TokenInfo method.
type ERC20TokenInfoResponse struct {
	Token *repository.ERC20Token `json:"token"`
	Err   error                  `json:"err"`
}

// MakeERC20TokenInfoEndpoint returns an endpoint that invokes ERC20TokenInfo on the service.
func MakeERC20TokenInfoEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ERC20TokenInfoRequest)
		token, err := s.ERC20TokenInfo(ctx, req.TokenHex)
		if err != nil {
			return ERC20TokenInfoResponse{Err: err}, err
		}
		return ERC20TokenInfoResponse{
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r ERC20TokenInfoResponse) Failed() error {
	return r.Err
}

// ERC20TokenInfo implements Service. Primarily useful in a client.
func (e Endpoints) ERC20TokenInfo(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error) {
	request := ERC20TokenInfoRequest{TokenHex: tokenHex}
	response, err := e.ERC20TokenInfoEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(ERC20TokenInfoResponse).Token, nil
}
