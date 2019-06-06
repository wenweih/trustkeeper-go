package service

import (
	"context"
	"fmt"
	"time"
	"trustkeeper-go/app/service/account/pkg/configure"
	"trustkeeper-go/app/service/account/pkg/model"
	"trustkeeper-go/app/service/account/pkg/repository"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	stdjwt "github.com/go-kit/kit/auth/jwt"
)

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password, orgName string) (string, error)
	Signin(ctx context.Context, email, password string) (string, error)
	Signout(ctx context.Context) error
	Roles(ctx context.Context) ([]string, error)
	Auth(ctx context.Context) (uuid string, err error)
}

type basicAccountService struct {
	repo repository.AccoutRepo
	conf configure.Conf
}

func (b *basicAccountService) findByTokenID(ctx context.Context) (*model.Account, error) {
	token := ctx.Value(stdjwt.JWTTokenContextKey).(string)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(b.conf.JWTKey), nil
	})
	if err != nil || !tkn.Valid {
		return nil, fmt.Errorf(err.Error())
	}
	return b.repo.FindByTokenID(claims.Id)
}

// https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/
func (b *basicAccountService) Create(ctx context.Context, email, password, orgName string) (string, error) {
	// Salt and hash the password using the bcrypt algorithm
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	uid := uuid.NewV4()
	acc := &model.Account{
		Email:    email,
		Password: string(hashedPassword),
		UUID:     uid.String()}

		if err := b.repo.Create(acc); err != nil {
			return "", err
		}
		// TODO: 如果是注册用户，则需新建默认群组和分配生成助记词
		return acc.UUID, nil
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

func (b *basicAccountService) Signout(ctx context.Context) error {
	acc, err := b.findByTokenID(ctx)
	if err != nil {
		return err
	}
	b.repo.Update(acc, map[string]interface{}{"token_id": nil})
	return nil
}

func (b *basicAccountService) Roles(ctx context.Context) ([]string, error) {
	acc, err := b.findByTokenID(ctx)
	if err != nil {
		return nil, err
	}
	roles := b.repo.GetRoles(acc)
	return roles, nil
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

func (b *basicAccountService) Auth(ctx context.Context) (uuid string, err error) {
	acc, err := b.findByTokenID(ctx)
	if err != nil {
		return "", err
	}
	return acc.UUID, nil
}
