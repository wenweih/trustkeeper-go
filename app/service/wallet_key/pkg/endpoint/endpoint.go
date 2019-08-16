package endpoint

import (
	"context"
	repository "trustkeeper-go/app/service/wallet_key/pkg/repository"
	service "trustkeeper-go/app/service/wallet_key/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

type GenerateMnemonicRequest struct {
	Namespaceid      string  `json:"namespaceid"`
	Bip44ids         []int32 `json:"bip44ids"`
	Bip44accountSize int     `json:"bip44accountSize"`
}

type GenerateMnemonicResponse struct {
	ChainsXpubs []*service.Bip44ThirdXpubsForChain `json:"chainsxpubs"`
	Version     string                             `json:"version"`
	Err         error                              `json:"err"`
}

func MakeGenerateMnemonicEndpoint(s service.WalletKeyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GenerateMnemonicRequest)
		xpubs, version, err := s.GenerateMnemonic(ctx, req.Namespaceid, req.Bip44ids, req.Bip44accountSize)
		if err != nil {
			return GenerateMnemonicResponse{
				Err:         err,
				ChainsXpubs: nil,
			}, err
		}
		return GenerateMnemonicResponse{ChainsXpubs: xpubs, Version: version}, nil
	}
}

func (r GenerateMnemonicResponse) Failed() error {
	return r.Err
}

type Failure interface {
	Failed() error
}

func (e Endpoints) GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (xpubs []*service.Bip44ThirdXpubsForChain, version string, err error) {
	request := GenerateMnemonicRequest{
		Namespaceid:      namespaceID,
		Bip44ids:         bip44ids,
		Bip44accountSize: bip44accountSize,
	}
	response, err := e.GenerateMnemonicEndpoint(ctx, request)
	if err != nil {
		return nil, "", err
	}
	return response.(GenerateMnemonicResponse).ChainsXpubs, response.(GenerateMnemonicResponse).Version, nil
}

func (e Endpoints) Close() error {
	return nil
}

type SignedBitcoincoreTxRequest struct {
	WalletHD repository.WalletHD `json:"wallet_hd"`
	VinAmount   int64  `json:"vin_amount"`
	TxHex    string              `json:"tx_hex"`
}

type SignedBitcoincoreTxResponse struct {
	SignedTxHex string `json:"signed_tx_hex"`
	Err         error  `json:"err"`
}

func MakeSignedBitcoincoreTxEndpoint(s service.WalletKeyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignedBitcoincoreTxRequest)
		signedTxHex, err := s.SignedBitcoincoreTx(ctx, req.WalletHD, req.TxHex, req.VinAmount)
		if err != nil {
			return SignedBitcoincoreTxResponse{Err: err}, err
		}
		return SignedBitcoincoreTxResponse{SignedTxHex: signedTxHex}, nil
	}
}

func (r SignedBitcoincoreTxResponse) Failed() error {
	return r.Err
}

func (e Endpoints) SignedBitcoincoreTx(ctx context.Context,
	walletHD repository.WalletHD, txHex string, vinAmount int64) (signedTxHex string, err error) {
	request := SignedBitcoincoreTxRequest{
		TxHex:    txHex,
		VinAmount: vinAmount,
		WalletHD: walletHD,
	}
	response, err := e.SignedBitcoincoreTxEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	return response.(SignedBitcoincoreTxResponse).SignedTxHex, nil
}
