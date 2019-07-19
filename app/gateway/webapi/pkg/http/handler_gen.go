// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package http

import (
	http "github.com/go-kit/kit/transport/http"
	http1 "net/http"
	endpoint "trustkeeper-go/app/gateway/webapi/pkg/endpoint"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]http.ServerOption) http1.Handler {
	m := http1.NewServeMux()
	makeSignupHandler(m, endpoints, options["Signup"])
	makeSigninHandler(m, endpoints, options["Signin"])
	makeSignoutHandler(m, endpoints, options["Signout"])
	makeGetRolesHandler(m, endpoints, options["GetRoles"])
	makeUserInfoHandler(m, endpoints, options["UserInfo"])
	makeGetGroupsHandler(m, endpoints, options["GetGroups"])
	makeCreateGroupHandler(m, endpoints, options["CreateGroup"])
	makeUpdateGroupHandler(m, endpoints, options["UpdateGroup"])
	makeGetGroupAssetsHandler(m, endpoints, options["GetGroupAssets"])
	makeChangeGroupAssetsHandler(m, endpoints, options["ChangeGroupAssets"])
	makeCreateWalletHandler(m, endpoints, options["CreateWallet"])
	return m
}
