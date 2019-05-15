package endpoint

import (
	"context"
	service "trustkeeper-go/app/service/account/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Email string `json:"email"`
	E1    error  `json:"e1"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e1 := s.Create(ctx, req.Email, req.Password)
		return CreateResponse{
			E1: e1,
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
func (e Endpoints) Create(ctx context.Context, email string, password string) (string, error) {
	request := CreateRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return response.(CreateResponse).Email, response.(CreateResponse).E1
}

// SignRequest collects the request parameters for the Sign method.
type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignResponse collects the response parameters for the Sign method.
type SigninResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSignEndpoint returns an endpoint that invokes Sign on the service.
func MakeSigninEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SigninRequest)
		s0, e1 := s.Signin(ctx, req.Email, req.Password)
		return SigninResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SigninResponse) Failed() error {
	return r.E1
}

// Signin implements Service. Primarily useful in a client.
func (e Endpoints) Signin(ctx context.Context, email string, password string) (s0 string, e1 error) {
	request := SigninRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.SigninEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SigninResponse).S0, response.(SigninResponse).E1
}

// SignoutRequest collects the request parameters for the Signout method.
type SignoutRequest struct {
	Token string `json:"token"`
}

// SignoutResponse collects the response parameters for the Signout method.
type SignoutResponse struct {
	E0 error `json:"e0"`
}

// MakeSignoutEndpoint returns an endpoint that invokes Signout on the service.
func MakeSignoutEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignoutRequest)
		e0 := s.Signout(ctx, req.Token)
		return SignoutResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SignoutResponse) Failed() error {
	return r.E0
}

// Signout implements Service. Primarily useful in a client.
func (e Endpoints) Signout(ctx context.Context, token string) (e0 error) {
	request := SignoutRequest{Token: token}
	response, err := e.SignoutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignoutResponse).E0
}
