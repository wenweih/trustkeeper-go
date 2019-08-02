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
	R0 *repository.OmniProperty `json:"r0"`
	E1 error                    `json:"e1"`
}

// MakeQueryOmniPropertyEndpoint returns an endpoint that invokes QueryOmniProperty on the service.
func MakeQueryOmniPropertyEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryOmniPropertyRequest)
		r0, e1 := s.QueryOmniProperty(ctx, req.Propertyid)
		return QueryOmniPropertyResponse{
			E1: e1,
			R0: r0,
		}, nil
	}
}

// Failed implements Failer.
func (r QueryOmniPropertyResponse) Failed() error {
	return r.E1
}

// QueryOmniProperty implements Service. Primarily useful in a client.
func (e Endpoints) QueryOmniProperty(ctx context.Context, propertyid int64) (r0 *repository.OmniProperty, e1 error) {
	request := QueryOmniPropertyRequest{Propertyid: propertyid}
	response, err := e.QueryOmniPropertyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(QueryOmniPropertyResponse).R0, response.(QueryOmniPropertyResponse).E1
}
