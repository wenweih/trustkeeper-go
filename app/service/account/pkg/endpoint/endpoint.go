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
	Result bool
	E1     error `json:"e1"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e1 := s.Create(ctx, req.Email, req.Password)
		return CreateResponse{
			E1: e1,
		}, e1
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
func (e Endpoints) Create(ctx context.Context, email string, password string) error {
	request := CreateRequest{
		Email:    email,
		Password: password,
	}
	_, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return nil
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
		if e1 != nil {
			return SigninResponse{
				E1: e1,
				S0: "",
			}, e1
		}
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
		return "", err
	}
	return response.(*SigninResponse).S0, response.(*SigninResponse).E1
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
		err := s.Signout(ctx)
		return SignoutResponse{E0: err}, err
	}
}

// Failed implements Failer.
func (r SignoutResponse) Failed() error {
	return r.E0
}

// Signout implements Service. Primarily useful in a client.
func (e Endpoints) Signout(ctx context.Context) (e0 error) {
	request := SignoutRequest{}
	_, err := e.SignoutEndpoint(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

// RolesRequest collects the request parameters for the Roles method.
type RolesRequest struct{}

// RolesResponse collects the response parameters for the Roles method.
type RolesResponse struct {
	S0 []string `json:"s0"`
	E1 error    `json:"e1"`
}

// MakeRolesEndpoint returns an endpoint that invokes Roles on the service.
func MakeRolesEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		roles, err := s.Roles(ctx)
		return RolesResponse{
			S0: roles,
		}, err
	}
}

// Failed implements Failer.
func (r RolesResponse) Failed() error {
	return r.E1
}

// Roles implements Service. Primarily useful in a client.
func (e Endpoints) Roles(ctx context.Context) (s0 []string, e1 error) {
	request := RolesRequest{}
	response, err := e.RolesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*RolesResponse).S0, nil
}

// AuthRequest collects the request parameters for the Auth method.
type AuthRequest struct{}

// AuthResponse collects the response parameters for the Auth method.
type AuthResponse struct {
	Uuid string `json:"uuid"`
	Err  error  `json:"err"`
}

// MakeAuthEndpoint returns an endpoint that invokes Auth on the service.
func MakeAuthEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuid, err := s.Auth(ctx)
		return AuthResponse{
			Err:  err,
			Uuid: uuid,
		}, nil
	}
}

// Failed implements Failer.
func (r AuthResponse) Failed() error {
	return r.Err
}

// Auth implements Service. Primarily useful in a client.
func (e Endpoints) Auth(ctx context.Context) (uuid string, err error) {
	request := AuthRequest{}
	response, err := e.AuthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AuthResponse).Uuid, response.(AuthResponse).Err
}
