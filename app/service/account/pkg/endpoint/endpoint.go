package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	gouuid "github.com/satori/go.uuid"
	service "trustkeeper-go/app/service/account/pkg/service"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	U0 gouuid.UUID `json:"u0"`
	E1 error       `json:"e1"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		u0, e1 := s.Create(ctx, req.Email, req.Password)
		return CreateResponse{
			E1: e1,
			U0: u0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, email string, password string) (u0 gouuid.UUID, e1 error) {
	request := CreateRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).U0, response.(CreateResponse).E1
}
