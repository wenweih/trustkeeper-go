package endpoint

import (
	"context"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// GenerateMnemonicRequest collects the request parameters for the GenerateMnemonic method.
type GenerateMnemonicRequest struct {
	Uuid  string `json:"uuid"`
}

// GenerateMnemonicResponse collects the response parameters for the GenerateMnemonic method.
type GenerateMnemonicResponse struct {
	Xpub string `json:"xpub"`
	Err  error  `json:"err"`
}

// MakeGenerateMnemonicEndpoint returns an endpoint that invokes GenerateMnemonic on the service.
func MakeGenerateMnemonicEndpoint(s service.WalletKeyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateMnemonicRequest)
		xpub, err := s.GenerateMnemonic(ctx, req.Uuid)
		return GenerateMnemonicResponse{
			Err:  err,
			Xpub: xpub,
		}, nil
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
func (e Endpoints) GenerateMnemonic(ctx context.Context, email string, uuid string) (xpub string, err error) {
	request := GenerateMnemonicRequest{
		Uuid:  uuid,
	}
	response, err := e.GenerateMnemonicEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GenerateMnemonicResponse).Xpub, response.(GenerateMnemonicResponse).Err
}
