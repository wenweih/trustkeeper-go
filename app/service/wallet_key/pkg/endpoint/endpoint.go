package endpoint

import (
	"context"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GenerateMnemonicRequest collects the request parameters for the GenerateMnemonic method.
type GenerateMnemonicRequest struct {
	Namespaceid  string `json:"namespaceid"`
	Bip44ids     []int32 `json:"bip44ids"`
	Bip44accountSize int  `json:"bip44accountSize"`
}

// GenerateMnemonicResponse collects the response parameters for the GenerateMnemonic method.
type GenerateMnemonicResponse struct {
	ChainsXpubs []*service.Bip44ThirdXpubsForChain `json:"chainsxpubs"`
	Version string `json:"version"`
	Err  error  `json:"err"`
}

// MakeGenerateMnemonicEndpoint returns an endpoint that invokes GenerateMnemonic on the service.
func MakeGenerateMnemonicEndpoint(s service.WalletKeyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateMnemonicRequest)
		xpubs, version, err := s.GenerateMnemonic(ctx, req.Namespaceid, req.Bip44ids, req.Bip44accountSize)
		if err != nil {
			return GenerateMnemonicResponse{
				Err:  err,
				ChainsXpubs: nil,
			}, err
		}
		return GenerateMnemonicResponse{ChainsXpubs: xpubs, Version: version}, nil
	}
}

// Failed implements Failer.
func (r GenerateMnemonicResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GenerateMnemonic implements Service. Primarily useful in a client.
func (e Endpoints) GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (xpubs []*service.Bip44ThirdXpubsForChain, version string, err error) {
	request := GenerateMnemonicRequest{
		Namespaceid: namespaceID,
		Bip44ids: bip44ids,
		Bip44accountSize: bip44accountSize,
	}
	response, err := e.GenerateMnemonicEndpoint(ctx, request)
	if err != nil {
		return nil, "", err
	}
	return response.(GenerateMnemonicResponse).ChainsXpubs, response.(GenerateMnemonicResponse).Version, nil
}

// Auth implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}
