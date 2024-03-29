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
	OrgName  string `json:"orgname"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	UUID string
	E1   error `json:"e1"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		uuid, err := s.Create(ctx, req.Email, req.Password, req.OrgName)
		return CreateResponse{
			UUID: uuid,
			E1:   err,
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
func (e Endpoints) Create(ctx context.Context, email, password, orgName string) (string, error) {
	request := CreateRequest{
		Email:    email,
		Password: password,
		OrgName:  orgName,
	}
	resp, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return resp.(CreateResponse).UUID, nil
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
		return SigninResponse{E1: e1, S0: s0}, nil
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
	return response.(SigninResponse).S0, nil
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
		return SignoutResponse{E0: err}, nil
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
			E1: err,
		}, nil
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
	return response.(RolesResponse).S0, nil
}

// AuthRequest collects the request parameters for the Auth method.
type AuthRequest struct{}

// AuthResponse collects the response parameters for the Auth method.
type AuthResponse struct {
	Uuid        string   `json:"uuid"`
	NamespaceID string   `json:"namespaceid"`
	Roles       []string `json:"roles"`
	Err         error    `json:"err"`
}

// MakeAuthEndpoint returns an endpoint that invokes Auth on the service.
func MakeAuthEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		uuid, namespaceid, roles, err := s.Auth(ctx)
		return AuthResponse{
			Uuid:        uuid,
			NamespaceID: namespaceid,
			Roles:       roles,
			Err:         err,
		}, nil
	}
}

// Failed implements Failer.
func (r AuthResponse) Failed() error {
	return r.Err
}

// Auth implements Service. Primarily useful in a client.
func (e Endpoints) Auth(ctx context.Context) (string, string, []string, error) {
	request := AuthRequest{}
	response, err := e.AuthEndpoint(ctx, request)
	if err != nil {
		return "", "", nil, err
	}
	namespaceID := response.(AuthResponse).NamespaceID
	uuid := response.(AuthResponse).Uuid
	return uuid, namespaceID, response.(AuthResponse).Roles, nil
}

// Close implements Service. Primarily useful in a client.
func (e Endpoints) Close() error {
	return nil
}

// UserInfoRequest collects the request parameters for the UserInfo method.
type UserInfoRequest struct{}

// UserInfoResponse collects the response parameters for the UserInfo method.
type UserInfoResponse struct {
	Roles   []string `json:"roles"`
	OrgName string   `json:"org_name"`
	Err     error    `json:"err"`
}

// MakeUserInfoEndpoint returns an endpoint that invokes UserInfo on the service.
func MakeUserInfoEndpoint(s service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		roles, orgName, err := s.UserInfo(ctx)
		return UserInfoResponse{
			OrgName: orgName,
			Roles:   roles,
			Err:     err,
		}, nil
	}
}

// Failed implements Failer.
func (r UserInfoResponse) Failed() error {
	return r.Err
}

// UserInfo implements Service. Primarily useful in a client.
func (e Endpoints) UserInfo(ctx context.Context) (roles []string, orgName string, err error) {
	request := UserInfoRequest{}
	response, err := e.UserInfoEndpoint(ctx, request)
	if err != nil {
		return nil, "", err
	}
	return response.(UserInfoResponse).Roles, response.(UserInfoResponse).OrgName, nil
}
