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
