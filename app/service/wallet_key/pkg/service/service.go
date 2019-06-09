package service

import (
	"context"
	bip39 "github.com/tyler-smith/go-bip39"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

// WalletKeyService describes the service.
type WalletKeyService interface {
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GenerateMnemonic(ctx context.Context, xpubUUID string) (xpub string, err error)
}

type basicWalletKeyService struct{}

func (b *basicWalletKeyService) GenerateMnemonic(ctx context.Context, xpubUUID string) (xpub string, err error) {
	// https://matthewdowney.github.io/extract-xpub-ethereum-bitcoin-ledger-nano-s.html
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(192)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	masterPubKey, err :=  masterKey.Neuter()
	if err != nil {
		return "", err
	}
	return masterPubKey.String(), nil
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
