package service

import (
	"fmt"
	"context"
	bip39 "github.com/tyler-smith/go-bip39"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

type Bip44AccountKey struct {
	Account int     `json:"account"`
	Key     string  `json:"key"`
}

type Bip44ThirdXpubsForChain struct {
	Chain  string           `json:"chain"`
	Xpubs   []*Bip44AccountKey	`json:"xpubs"`
}

// WalletKeyService describes the service.
type WalletKeyService interface {
	GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (xpubs []*Bip44ThirdXpubsForChain, err error)
}

type basicWalletKeyService struct{}

func (b *basicWalletKeyService) GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (xpubs []*Bip44ThirdXpubsForChain, err error) {
	// https://matthewdowney.github.io/extract-xpub-ethereum-bitcoin-ledger-nano-s.html
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(192)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}
	masterXpub, err :=  masterKey.Neuter()
	if err != nil {
		return nil, err
	}
	fmt.Println("masterXpub: ", masterXpub)
	return nil, nil
	// return masterPubKey.String(), nil
}

// NewBasicWalletKeyService returns a naive, stateless implementation of WalletKeyService.
func NewBasicWalletKeyService() WalletKeyService {
	return &basicWalletKeyService{}
}

// New returns a WalletKeyService with all of the expected middleware wired in.
func New(middleware []Middleware) WalletKeyService {
	var svc WalletKeyService = NewBasicWalletKeyService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
