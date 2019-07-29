package service

import (
	"context"

	"github.com/btcsuite/btcd/btcjson"
)

// ChainsQueryService describes the service.
type ChainsQueryService interface {
	BitcoincoreBlock(ctx context.Context, blockHash string) (*btcjson.GetBlockVerboseResult, error)
}

type basicChainsQueryService struct{}

func (b *basicChainsQueryService) BitcoincoreBlock(ctx context.Context, blockHash string) (b0 *btcjson.GetBlockVerboseResult, e1 error) {
	// TODO implement the business logic of BitcoincoreBlock
	return b0, e1
}

// NewBasicChainsQueryService returns a naive, stateless implementation of ChainsQueryService.
func NewBasicChainsQueryService() ChainsQueryService {
	return &basicChainsQueryService{}
}

// New returns a ChainsQueryService with all of the expected middleware wired in.
func New(middleware []Middleware) ChainsQueryService {
	var svc ChainsQueryService = NewBasicChainsQueryService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
