package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	endpoint "trustkeeper-go/app/gateway/webapi/pkg/endpoint"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeSignupHandler creates the handler logic
func makeSignupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/signup", http1.NewServer(endpoints.SignupEndpoint, decodeSignupRequest, encodeSignupResponse, options...))
}

// decodeSignupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSignupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSigninHandler creates the handler logic
func makeSigninHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/signin", http1.NewServer(endpoints.SigninEndpoint, decodeSigninRequest, encodeSigninResponse, options...))
}

// decodeSigninRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSigninRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SigninRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSigninResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSigninResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSignoutHandler creates the handler logic
func makeSignoutHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/signout", http1.NewServer(endpoints.SignoutEndpoint, decodeSignoutRequest, encodeSignoutResponse, options...))
}

// decodeSignoutRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encodeSignoutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignoutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	switch {
	case strings.HasPrefix(err.Error(), "Fields exist"):
		return http.StatusBadRequest
	case strings.Contains(err.Error(), "token is expired"):
		return http.StatusUnauthorized
	case strings.Contains(err.Error(), "context deadline exceeded"):
		return http.StatusRequestTimeout
	case strings.Contains(err.Error(), "does not exist") ||
		strings.Contains(err.Error(), "no contract code at given address") ||
		strings.Contains(err.Error(), "record not found"):
		return http.StatusNotFound
	case strings.Contains(err.Error(), "duplicate key value"):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

type errorWrapper struct {
	Error string `json:"error"`
}

// makeGetRolesHandler creates the handler logic
func makeGetRolesHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-roles", http1.NewServer(endpoints.GetRolesEndpoint, decodeGetRolesRequest, encodeGetRolesResponse, options...))
}

// decodeGetRolesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRolesRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encodeGetRolesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetRolesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGroupHandler creates the handler logic
func makeGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/groups", http1.NewServer(endpoints.GetGroupsEndpoint, decodeGroupRequest, encodeGroupResponse, options...))
}

// decodeGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encodeGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupsHandler creates the handler logic
func makeGetGroupsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-groups", http1.NewServer(endpoints.GetGroupsEndpoint, decodeGetGroupsRequest, encodeGetGroupsResponse, options...))
}

// decodeGetGroupsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encodeGetGroupsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-group", http1.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...))
}

// decodeCreateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUpdateGroupHandler creates the handler logic
func makeUpdateGroupHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/update-group", http1.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...))
}

// decodeUpdateGroupRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUpdateGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.UpdateGroupRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeUpdateGroupResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUpdateGroupResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeUserInfoHandler creates the handler logic
func makeUserInfoHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/user-info", http1.NewServer(endpoints.UserInfoEndpoint, decodeUserInfoRequest, encodeUserInfoResponse, options...))
}

// decodeUserInfoRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeUserInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// encodeUserInfoResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeUserInfoResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetGroupAssetsHandler creates the handler logic
func makeGetGroupAssetsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-group-assets", http1.NewServer(endpoints.GetGroupAssetsEndpoint, decodeGetGroupAssetsRequest, encodeGetGroupAssetsResponse, options...))
}

// decodeGetGroupAssetsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetGroupAssetsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetGroupAssetsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetGroupAssetsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetGroupAssetsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeChangeGroupAssetsHandler creates the handler logic
func makeChangeGroupAssetsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/change-group-assets", http1.NewServer(endpoints.ChangeGroupAssetsEndpoint, decodeChangeGroupAssetsRequest, encodeChangeGroupAssetsResponse, options...))
}

// decodeChangeGroupAssetsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeChangeGroupAssetsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ChangeGroupAssetsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeChangeGroupAssetsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeChangeGroupAssetsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateWalletHandler creates the handler logic
func makeCreateWalletHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-wallet", http1.NewServer(endpoints.CreateWalletEndpoint, decodeCreateWalletRequest, encodeCreateWalletResponse, options...))
}

// decodeCreateWalletRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateWalletRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateWalletRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateWalletResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateWalletResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetWalletsHandler creates the handler logic
func makeGetWalletsHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-wallets", http1.NewServer(endpoints.GetWalletsEndpoint, decodeGetWalletsRequest, encodeGetWalletsResponse, options...))
}

// decodeGetWalletsRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetWalletsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetWalletsRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetWalletsResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetWalletsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeQueryOmniPropertyHandler creates the handler logic
func makeQueryOmniPropertyHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/query-omni-property", http1.NewServer(endpoints.QueryOmniPropertyEndpoint, decodeQueryOmniPropertyRequest, encodeQueryOmniPropertyResponse, options...))
}

// decodeQueryOmniPropertyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeQueryOmniPropertyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.QueryOmniPropertyRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeQueryOmniPropertyResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeQueryOmniPropertyResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeCreateTokenHandler creates the handler logic
func makeCreateTokenHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-token", http1.NewServer(endpoints.CreateTokenEndpoint, decodeCreateTokenRequest, encodeCreateTokenResponse, options...))
}

// decodeCreateTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeCreateTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeEthTokenHandler creates the handler logic
func makeEthTokenHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/eth-token", http1.NewServer(endpoints.EthTokenEndpoint, decodeEthTokenRequest, encodeEthTokenResponse, options...))
}

// decodeEthTokenRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeEthTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.EthTokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeEthTokenResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeEthTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeSendBTCTxHandler creates the handler logic
func makeSendBTCTxHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/send-btctx", http1.NewServer(endpoints.SendBTCTxEndpoint, decodeSendBTCTxRequest, encodeSendBTCTxResponse, options...))
}

// decodeSendBTCTxRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSendBTCTxRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SendBTCTxRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSendBTCTxResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSendBTCTxResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeQueryBalanceHandler creates the handler logic
func makeQueryBalanceHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/query-balance", http1.NewServer(endpoints.QueryBalanceEndpoint, decodeQueryBalanceRequest, encodeQueryBalanceResponse, options...))
}

// decodeQueryBalanceRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeQueryBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.QueryBalanceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeQueryBalanceResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeQueryBalanceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
