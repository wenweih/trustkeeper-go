// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"
	service "trustkeeper-go/app/service/chains_query/pkg/service"
)

// Endpoints collects all of the endpoints that compose a profile service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	BitcoincoreBlockEndpoint  endpoint.Endpoint
	QueryOmniPropertyEndpoint endpoint.Endpoint
	ERC20TokenInfoEndpoint    endpoint.Endpoint
	ConstructTxBTCEndpoint    endpoint.Endpoint
	SendBTCTxEndpoint         endpoint.Endpoint
	ConstructTxETHEndpoint    endpoint.Endpoint
	QueryBalanceEndpoint      endpoint.Endpoint
	WalletValidateEndpoint    endpoint.Endpoint
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func New(s service.ChainsQueryService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		BitcoincoreBlockEndpoint:  MakeBitcoincoreBlockEndpoint(s),
		ConstructTxBTCEndpoint:    MakeConstructTxBTCEndpoint(s),
		ConstructTxETHEndpoint:    MakeConstructTxETHEndpoint(s),
		ERC20TokenInfoEndpoint:    MakeERC20TokenInfoEndpoint(s),
		QueryBalanceEndpoint:      MakeQueryBalanceEndpoint(s),
		QueryOmniPropertyEndpoint: MakeQueryOmniPropertyEndpoint(s),
		SendBTCTxEndpoint:         MakeSendBTCTxEndpoint(s),
		WalletValidateEndpoint:    MakeWalletValidateEndpoint(s),
	}
	for _, m := range mdw["BitcoincoreBlock"] {
		eps.BitcoincoreBlockEndpoint = m(eps.BitcoincoreBlockEndpoint)
	}
	for _, m := range mdw["QueryOmniProperty"] {
		eps.QueryOmniPropertyEndpoint = m(eps.QueryOmniPropertyEndpoint)
	}
	for _, m := range mdw["ERC20TokenInfo"] {
		eps.ERC20TokenInfoEndpoint = m(eps.ERC20TokenInfoEndpoint)
	}
	for _, m := range mdw["ConstructTxBTC"] {
		eps.ConstructTxBTCEndpoint = m(eps.ConstructTxBTCEndpoint)
	}
	for _, m := range mdw["SendBTCTx"] {
		eps.SendBTCTxEndpoint = m(eps.SendBTCTxEndpoint)
	}
	for _, m := range mdw["ConstructTxETH"] {
		eps.ConstructTxETHEndpoint = m(eps.ConstructTxETHEndpoint)
	}
	for _, m := range mdw["QueryBalance"] {
		eps.QueryBalanceEndpoint = m(eps.QueryBalanceEndpoint)
	}
	for _, m := range mdw["WalletValidate"] {
		eps.WalletValidateEndpoint = m(eps.WalletValidateEndpoint)
	}
	return eps
}
