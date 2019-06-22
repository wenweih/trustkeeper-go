package service

import (
	"context"
	bip39 "github.com/tyler-smith/go-bip39"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"trustkeeper-go/app/service/wallet_key/pkg/repository"
)

type Bip44AccountKey struct {
	Account int     `json:"account"`
	Key     string  `json:"key"`
}

type Bip44ThirdXpubsForChain struct {
	Chain  int           `json:"chain"`
	Xpubs   []*Bip44AccountKey	`json:"xpubs"`
}

// WalletKeyService describes the service.
type WalletKeyService interface {
	GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (chainsWithXpubs []*Bip44ThirdXpubsForChain, version string, err error)
	Close() error
}

type basicWalletKeyService struct{
	biz  repository.IBiz
}


func (b *basicWalletKeyService) Close() error{
	return b.biz.Close()
}

func (b *basicWalletKeyService) GenerateMnemonic(ctx context.Context, namespaceID string, bip44ids []int32, bip44accountSize int) (chainsWithXpubs []*Bip44ThirdXpubsForChain, version string, err error) {
	// https://matthewdowney.github.io/extract-xpub-ethereum-bitcoin-ledger-nano-s.html
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(192)
	if err != nil {
		return nil, "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, "", err
	}

	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, "", err
	}

	// This gives the path: m / 44'
	acc44H, err := masterKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, "", err
	}

	chainsWithXpubs = []*Bip44ThirdXpubsForChain{}
	for _, bip44id := range bip44ids {
		// m / 44' / coin_type'
		coinTypeH, err := acc44H.Child(hdkeychain.HardenedKeyStart + uint32(bip44id))
		if err != nil {
			return nil, "", err
		}
		xpubs := []*Bip44AccountKey{}
		for account := 0;account < bip44accountSize;account++ {
			// m / 44' / coin_type' / account'
			accountH, err := coinTypeH.Child(hdkeychain.HardenedKeyStart + uint32(account))
			if err != nil {
				return nil, "", err
			}
			xpub, err := accountH.Neuter()
			if err != nil {
				return nil, "", err
			}
			xpubs = append(xpubs, &Bip44AccountKey{Account: account, Key: xpub.String()})
		}
		chainsWithXpubs = append(chainsWithXpubs, &Bip44ThirdXpubsForChain{Chain: int(bip44id), Xpubs: xpubs})
	}
	version, err = b.biz.SaveMnemonic(namespaceID, []byte(mnemonic))
	if err != nil {
		return nil, "", err
	}
	return chainsWithXpubs, version, nil
}

// NewBasicWalletKeyService returns a naive, stateless implementation of WalletKeyService.
func newBasicWalletKeyService() (WalletKeyService, error) {
	repo, err := repository.New()
	if err != nil {
		return nil, err
	}
	return &basicWalletKeyService{
		biz: repo,
	}, nil
}

// New returns a WalletKeyService with all of the expected middleware wired in.
func New(middleware []Middleware) (WalletKeyService, error) {
	 ws, err := newBasicWalletKeyService()
	 if err != nil {
		 return nil, err
	 }
	 var svc WalletKeyService = ws
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}
