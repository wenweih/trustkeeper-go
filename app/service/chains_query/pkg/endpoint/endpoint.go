package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/chains_query/pkg/service"
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
