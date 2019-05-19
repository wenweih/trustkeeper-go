package service

import (
	"context"
	"fmt"
	"time"
	"trustkeeper-go/app/service/account/pkg/model"
	"trustkeeper-go/app/service/account/pkg/repository"
	"trustkeeper-go/app/service/account/pkg/configure"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password string) error
	Signin(ctx context.Context, email, password string) (string, error)
	Signout(ctx context.Context) error
	Roles(ctx context.Context, token string) ([]string, error)
	FindByTokenID(ctx context.Context, tokenID string) (*model.Account, error)
}

type basicAccountService struct {
	repo repository.AccoutRepo
	conf configure.Conf
}

func (b *basicAccountService) FindByTokenID(ctx context.Context, tokenID string) (*model.Account, error) {
	return b.repo.FindByTokenID(tokenID)
}

// https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/
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

type Claims struct {
	jwt.StandardClaims
}

// https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
func (b *basicAccountService) Signin(ctx context.Context, email string, password string) (s0 string, e1 error) {
	acc, err := b.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)); err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(100 * time.Minute)
	tokenID := uuid.NewV4().String()
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Id:        tokenID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(b.conf.JWTKey))
	if err != nil {
		return "", err
	}
	if err := b.repo.Update(acc, map[string]interface{}{"token_id": tokenID}); err != nil {
		return "", err
	}
	return tokenStr, e1
}

func (b *basicAccountService) Signout(ctx context.Context) (e0 error) {
	acc, ok := ctx.Value("account").(*model.Account)
	if !ok {
		return fmt.Errorf("acount not found in context")
	}
	b.repo.Update(acc, map[string]interface{}{"token_id": nil})
	return nil
}

func (b *basicAccountService) Roles(ctx context.Context, token string) (s0 []string, e1 error) {

	return s0, e1
}

// NewBasicAccountService returns a naive, stateless implementation of AccountService.
func NewBasicAccountService(conf configure.Conf) AccountService {
	db := repository.DB(conf.DBInfo)
	bas := basicAccountService{
		repo: repository.New(db),
		conf: conf,
	}
	return &bas
}

// New returns a AccountService with all of the expected middleware wired in.
func New(conf configure.Conf, middleware []Middleware) AccountService {
	var svc AccountService = NewBasicAccountService(conf)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
