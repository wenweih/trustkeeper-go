package service

import (
	"context"
	"strings"
	"trustkeeper-go/app/service/account/pkg/model"
	"trustkeeper-go/app/service/account/pkg/repository"
	"trustkeeper-go/library/vault"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password string) error
	Sign(ctx context.Context, email, password string) (string, error)
}

type basicAccountService struct {
	repo repository.AccoutRepo
}

func (b *basicAccountService) Create(ctx context.Context, email string, password string) error {
	// Salt and hash the password using the bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	uid := uuid.NewV4()
	acc := &model.Account{
		Email:    email,
		Password: string(hashedPassword),
		UUID:     uid.String()}

	return b.repo.Create(acc)
}

// NewBasicAccountService returns a naive, stateless implementation of AccountService.
func NewBasicAccountService() AccountService {
	vc, err := vault.NewVault()
	if err != nil {
		panic("fail to connect vault" + err.Error())
	}
	// ListSecret
	data, err := vc.Logical().Read("kv1/db_trustkeeper_account")
	if err != nil {
		panic("vaule read error" + err.Error())
	}

	host := strings.Join([]string{"host", data.Data["host"].(string)}, "=")
	port := strings.Join([]string{"port", data.Data["port"].(string)}, "=")
	user := strings.Join([]string{"user", data.Data["username"].(string)}, "=")
	password := strings.Join([]string{"password", data.Data["password"].(string)}, "=")
	dbname := strings.Join([]string{"dbname", data.Data["dbname"].(string)}, "=")
	sslmode := strings.Join([]string{"sslmode", data.Data["sslmode"].(string)}, "=")
	dbInfo := strings.Join([]string{host, port, user, dbname, password, sslmode}, " ")
	return &basicAccountService{
		repo: repository.New(repository.DB(dbInfo)),
	}
}

// New returns a AccountService with all of the expected middleware wired in.
func New(middleware []Middleware) AccountService {
	var svc AccountService = NewBasicAccountService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAccountService) Sign(ctx context.Context, email string, password string) (s0 string, e1 error) {
	// TODO implement the business logic of Sign
	return s0, e1
}
