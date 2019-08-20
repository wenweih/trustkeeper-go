package endpoint

import (
	"context"
	"trustkeeper-go/app/service/chains_query/pkg/repository"
	service "trustkeeper-go/app/service/chains_query/pkg/service"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	endpoint "github.com/go-kit/kit/endpoint"
)

// BitcoincoreBlockRequest collects the request parameters for the BitcoincoreBlock method.
type BitcoincoreBlockRequest struct {
	BlockHash *chainhash.Hash `json:"block_hash"`
}

// BitcoincoreBlockResponse collects the response parameters for the BitcoincoreBlock method.
type BitcoincoreBlockResponse struct {
	B0 *btcjson.GetBlockVerboseResult `json:"b0"`
	E1 error                          `json:"e1"`
}

// MakeBitcoincoreBlockEndpoint returns an endpoint that invokes BitcoincoreBlock on the service.
func MakeBitcoincoreBlockEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BitcoincoreBlockRequest)
		b0, e1 := s.BitcoincoreBlock(ctx, req.BlockHash)
		return BitcoincoreBlockResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r BitcoincoreBlockResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// BitcoincoreBlock implements Service. Primarily useful in a client.
func (e Endpoints) BitcoincoreBlock(ctx context.Context, blockHash *chainhash.Hash) (b0 *btcjson.GetBlockVerboseResult, e1 error) {
	request := BitcoincoreBlockRequest{BlockHash: blockHash}
	response, err := e.BitcoincoreBlockEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BitcoincoreBlockResponse).B0, response.(BitcoincoreBlockResponse).E1
}

// QueryOmniPropertyRequest collects the request parameters for the QueryOmniProperty method.
type QueryOmniPropertyRequest struct {
	Propertyid int64 `json:"propertyid"`
}

// QueryOmniPropertyResponse collects the response parameters for the QueryOmniProperty method.
type QueryOmniPropertyResponse struct {
	Property *repository.OmniProperty `json:"r0"`
	Err      error                    `json:"e1"`
}

// MakeQueryOmniPropertyEndpoint returns an endpoint that invokes QueryOmniProperty on the service.
func MakeQueryOmniPropertyEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryOmniPropertyRequest)
		resp, err := s.QueryOmniProperty(ctx, req.Propertyid)
		if err != nil {
			return QueryOmniPropertyResponse{Err: err}, err
		}
		return QueryOmniPropertyResponse{Property: resp}, nil
	}
}

// Failed implements Failer.
func (r QueryOmniPropertyResponse) Failed() error {
	return r.Err
}

// QueryOmniProperty implements Service. Primarily useful in a client.
func (e Endpoints) QueryOmniProperty(ctx context.Context, propertyid int64) (r0 *repository.OmniProperty, e1 error) {
	request := QueryOmniPropertyRequest{Propertyid: propertyid}
	response, err := e.QueryOmniPropertyEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(QueryOmniPropertyResponse).Property, nil
}

// ERC20TokenInfoRequest collects the request parameters for the ERC20TokenInfo method.
type ERC20TokenInfoRequest struct {
	TokenHex string `json:"token_hex"`
}

// ERC20TokenInfoResponse collects the response parameters for the ERC20TokenInfo method.
type ERC20TokenInfoResponse struct {
	Token *repository.ERC20Token `json:"token"`
	Err   error                  `json:"err"`
}

// MakeERC20TokenInfoEndpoint returns an endpoint that invokes ERC20TokenInfo on the service.
func MakeERC20TokenInfoEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ERC20TokenInfoRequest)
		token, err := s.ERC20TokenInfo(ctx, req.TokenHex)
		if err != nil {
			return ERC20TokenInfoResponse{Err: err}, err
		}
		return ERC20TokenInfoResponse{
			Token: token,
		}, nil
	}
}

// Failed implements Failer.
func (r ERC20TokenInfoResponse) Failed() error {
	return r.Err
}

// ERC20TokenInfo implements Service. Primarily useful in a client.
func (e Endpoints) ERC20TokenInfo(ctx context.Context, tokenHex string) (token *repository.ERC20Token, err error) {
	request := ERC20TokenInfoRequest{TokenHex: tokenHex}
	response, err := e.ERC20TokenInfoEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(ERC20TokenInfoResponse).Token, nil
}

// ConstructTxBTCRequest collects the request parameters for the ConstructTxBTC method.
type ConstructTxBTCRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

// ConstructTxBTCResponse collects the response parameters for the ConstructTxBTC method.
type ConstructTxBTCResponse struct {
	UnsignedTxHex string `json:"unsigned_tx_hex"`
	VinAmount     int64  `json:"vin_amount"`
	Err           error  `json:"err"`
}

// MakeConstructTxBTCEndpoint returns an endpoint that invokes ConstructTxBTC on the service.
func MakeConstructTxBTCEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConstructTxBTCRequest)
		unsignedTxHex, vinAmount, err := s.ConstructTxBTC(ctx, req.From, req.To, req.Amount)
		if err != nil {
			return ConstructTxBTCResponse{Err: err}, err
		}
		return ConstructTxBTCResponse{UnsignedTxHex: unsignedTxHex, VinAmount: vinAmount}, nil
	}
}

// Failed implements Failer.
func (r ConstructTxBTCResponse) Failed() error {
	return r.Err
}

// ConstructTxBTC implements Service. Primarily useful in a client.
func (e Endpoints) ConstructTxBTC(ctx context.Context, from string, to string, amount string) (unsignedTxHex string, vinAmount int64, err error) {
	request := ConstructTxBTCRequest{
		Amount: amount,
		From:   from,
		To:     to,
	}
	response, err := e.ConstructTxBTCEndpoint(ctx, request)
	if err != nil {
		return "", 0, err
	}
	return response.(ConstructTxBTCResponse).UnsignedTxHex, response.(ConstructTxBTCResponse).VinAmount, nil
}

// SendBTCTxRequest collects the request parameters for the SendBTCTx method.
type SendBTCTxRequest struct {
	SignedTxHex string `json:"signed_tx_hex"`
}

// SendBTCTxResponse collects the response parameters for the SendBTCTx method.
type SendBTCTxResponse struct {
	TxID string `json:"tx_id"`
	Err  error  `json:"err"`
}

// MakeSendBTCTxEndpoint returns an endpoint that invokes SendBTCTx on the service.
func MakeSendBTCTxEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendBTCTxRequest)
		txID, err := s.SendBTCTx(ctx, req.SignedTxHex)
		if err != nil {
			return SendBTCTxResponse{Err: err}, err
		}
		return SendBTCTxResponse{TxID: txID}, nil
	}
}

// Failed implements Failer.
func (r SendBTCTxResponse) Failed() error {
	return r.Err
}

// SendBTCTx implements Service. Primarily useful in a client.
func (e Endpoints) SendBTCTx(ctx context.Context, signedTxHex string) (txID string, err error) {
	request := SendBTCTxRequest{SignedTxHex: signedTxHex}
	response, err := e.SendBTCTxEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return response.(SendBTCTxResponse).TxID, nil
}

// QueryBalanceRequest collects the request parameters for the QueryBalance method.
type QueryBalanceRequest struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
}

// QueryBalanceResponse collects the response parameters for the QueryBalance method.
type QueryBalanceResponse struct {
	Balance string `json:"balance"`
	Err     error  `json:"err"`
}

// MakeQueryBalanceEndpoint returns an endpoint that invokes QueryBalance on the service.
func MakeQueryBalanceEndpoint(s service.ChainsQueryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryBalanceRequest)
		balance, err := s.QueryBalance(ctx, req.Symbol, req.Address)
		if err != nil {
			return QueryBalanceResponse{
				Err: err,
			}, err
		}
		return QueryBalanceResponse{
			Balance: balance,
		}, nil
	}
}

// Failed implements Failer.
func (r QueryBalanceResponse) Failed() error {
	return r.Err
}

// QueryBalance implements Service. Primarily useful in a client.
func (e Endpoints) QueryBalance(ctx context.Context, symbol string, address string) (balance string, err error) {
	request := QueryBalanceRequest{
		Address: address,
		Symbol:  symbol,
	}
	response, err := e.QueryBalanceEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return response.(QueryBalanceResponse).Balance, nil
}
