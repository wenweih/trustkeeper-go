package service

import (
	"context"
	"trustkeeper-go/app/service/wallet_key/pkg/repository"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/tyler-smith/go-bip39"
)

type Bip44AccountKey struct {
	Account int    `json:"account"`
	Key     string `json:"key"`
}

type Bip44ThirdXpubsForChain struct {
	Chain uint               `json:"chain"`
	Xpubs []*Bip44AccountKey `json:"xpubs"`
}

// WalletKeyService describes the service.
type WalletKeyService interface {
	GenerateMnemonic(ctx context.Context,
		namespaceID string, bip44ids []int32, bip44accountSize int) (chainsWithXpubs []*Bip44ThirdXpubsForChain, version string, err error)
	Close() error
	SignedBitcoincoreTx(ctx context.Context, walletHD repository.WalletHD, txHex string, vinAmount int64) (signedTxHex string, err error)
	SignedEthereumTx(ctx context.Context, walletHD repository.WalletHD, txHex string, chainID string) (signedTxHex string, err error)
}

type basicWalletKeyService struct {
	biz repository.IBiz
}

func (b *basicWalletKeyService) Close() error {
	return b.biz.Close()
}

// GenerateMnemonic return chainsWithXpubs is slice of Bip44ThirdXpubsForChain pointer
func (b *basicWalletKeyService) GenerateMnemonic(ctx context.Context,
	namespaceID string, bip44ids []int32, bip44accountSize int) (chainsWithXpubs []*Bip44ThirdXpubsForChain, version string, err error) {
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

	// Defining a function that returns a slice of variable size in golang
	// https://stackoverflow.com/questions/22317329/defining-a-function-that-returns-a-slice-of-variable-size-in-golang
	chainsWithXpubs = make([]*Bip44ThirdXpubsForChain, len(bip44ids))
	for i, bip44id := range bip44ids {
		// m / 44' / coin_type'
		coinTypeH, err := acc44H.Child(hdkeychain.HardenedKeyStart + uint32(bip44id))
		if err != nil {
			return nil, "", err
		}
		xpubs := make([]*Bip44AccountKey, bip44accountSize)
		for account := 0; account < bip44accountSize; account++ {
			// m / 44' / coin_type' / account'
			accountH, err := coinTypeH.Child(hdkeychain.HardenedKeyStart + uint32(account))
			if err != nil {
				return nil, "", err
			}
			xpub, err := accountH.Neuter()
			if err != nil {
				return nil, "", err
			}
			xpubs[account] = &Bip44AccountKey{Account: account, Key: xpub.String()}
		}
		chainsWithXpubs[i] = &Bip44ThirdXpubsForChain{Chain: uint(bip44id), Xpubs: xpubs}
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

// NewBasicWalletKeyService returns a naive, stateless implementation of WalletKeyService.
func NewBasicWalletKeyService() WalletKeyService {
	return &basicWalletKeyService{}
}

func (b *basicWalletKeyService) SignedBitcoincoreTx(ctx context.Context,
	walletHD repository.WalletHD, txHex string, vinAmount int64) (signedTxHex string, err error) {
	signedTxHex, err = b.biz.SignedBitcoincoreTx(ctx, walletHD, txHex, vinAmount)
	return
}

func (b *basicWalletKeyService) SignedEthereumTx(ctx context.Context, walletHD repository.WalletHD, txHex string, chainID string) (string, error) {
	return b.biz.SignedEthereumTx(ctx, walletHD, txHex, chainID)
}
