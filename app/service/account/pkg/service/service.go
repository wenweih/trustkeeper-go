package service

import (
	"context"
	"fmt"
	"trustkeeper-go/app/service/account/pkg/configure"
	"trustkeeper-go/app/service/account/pkg/repository"

	"github.com/dgrijalva/jwt-go"

	stdjwt "github.com/go-kit/kit/auth/jwt"

	"github.com/gocraft/work"
	"trustkeeper-go/library/database/redis"
	"trustkeeper-go/library/util"
	"trustkeeper-go/library/database/orm"
)

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password, orgName string) (string, error)
	Signin(ctx context.Context, email, password string) (token string, err error)
	Signout(ctx context.Context) error
	Roles(ctx context.Context) ([]string, error)
	Auth(ctx context.Context) (uuid string, err error)
	Close() error
}

type basicAccountService struct {
	biz  repository.IBiz
	jwtKey string
	jobEnqueuer *work.Enqueuer
}

// Claims jwt clains struct
type Claims struct {
	jwt.StandardClaims
}

func (b *basicAccountService) Close() error{
	return b.biz.Close()
}

// https://www.sohamkamani.com/blog/2018/02/25/golang-password-authentication-and-storage/
func (b *basicAccountService) Create(ctx context.Context, email, password, orgName string) (string, error) {
	uuid, namespaceid, err := b.biz.Signup(email, password, orgName)
	if err != nil{
		return "", err
	}

	if _, err := b.jobEnqueuer.Enqueue(common.WalletMnemonicJob,
		work.Q{"namespaceid": namespaceid}); err != nil {
		return "", err
	}

	// if _, err := b.jobEnqueuer.Enqueue(common.SignUpJobs,
	// 	work.Q{"uuid": uuid,
	// 		"email": email,
	// 		"orgname": orgName}); err != nil {
	// 	return "", err
	// }
	return uuid, nil
}


// https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
func (b *basicAccountService) Signin(ctx context.Context, email string, password string) (string, error) {
	return b.biz.Signin(email, password, b.jwtKey)
}

func (b *basicAccountService) Signout(ctx context.Context) error {
	tokenID, err := extractTOkenIDFromContext(ctx, b.jwtKey)
	if err != nil {
		return err
	}
	return b.biz.Signout(tokenID)
}

func (b *basicAccountService) Roles(ctx context.Context) ([]string, error) {
	tokenID, err := extractTOkenIDFromContext(ctx, b.jwtKey)
	if err != nil {
		return nil, err
	}
	return b.biz.QueryRoles(tokenID)
}


func (b *basicAccountService) Auth(ctx context.Context) (uuid string, err error) {
	tokenID, err := extractTOkenIDFromContext(ctx, b.jwtKey)
	if err != nil {
		return "", err
	}
	return b.biz.Auth(tokenID)
}

func extractTOkenIDFromContext(ctx context.Context, jwtKey string) (string, error) {
	token := ctx.Value(stdjwt.JWTTokenContextKey).(string)
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil || !tkn.Valid {
		return "", fmt.Errorf(err.Error())
	}
	return claims.Id, nil
}

// NewBasicAccountService returns a naive, stateless implementation of AccountService.
func NewBasicAccountService(conf configure.Conf) (AccountService, error) {
	db, err := orm.DB(conf.DBInfo)
	if err != nil {
		return nil, err
	}
	redisPool := redis.NewPool(conf.Redis)
	enqueuer := work.NewEnqueuer(redis.Namespace, redisPool)
	bas := basicAccountService{
		biz: repository.New(redisPool, db, conf.JWTKey),
		jobEnqueuer: enqueuer,
		jwtKey: conf.JWTKey,
	}
	return &bas, nil
}

// New returns a AccountService with all of the expected middleware wired in.
func New(conf configure.Conf, middleware []Middleware) (AccountService, error) {
	srv, err := NewBasicAccountService(conf)
	if err != nil {
		return nil, err
	}
	for _, m := range middleware {
		srv = m(srv)
	}
	var accountSrv AccountService = srv
	return accountSrv, nil
}
