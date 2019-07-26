package service

import (
	"context"
	"trustkeeper-go/app/service/transaction/pkg/configure"
	"trustkeeper-go/app/service/transaction/pkg/repository"
	"trustkeeper-go/library/database/orm"
)

// TransactionService describes the service.
type TransactionService interface {
	Close() error
	AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error)
}

type basicTransactionService struct{
	biz repository.IBiz
}

func (b *basicTransactionService) Close() error {
	return b.biz.Close()
}

func (b *basicTransactionService) AssignAssetsToWallet(ctx context.Context, address string, assets []*repository.SimpleAsset) (err error) {
	err = b.biz.AssignAssetsToWallet(ctx, address, assets)
	return
}

// NewBasicTransactionService returns a naive, stateless implementation of TransactionService.
func NewBasicTransactionService(conf configure.Conf) (TransactionService, error) {
	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}

	bas := basicTransactionService{
		biz: repository.New(db),
	}
	return &bas, nil
}

// New returns a TransactionService with all of the expected middleware wired in.
func New(conf configure.Conf, middleware []Middleware) (TransactionService, error) {
	bs, err := NewBasicTransactionService(conf)
	if err != nil {
		return nil, err
	}

	var svc TransactionService = bs
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc, nil
}
