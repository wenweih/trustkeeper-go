package service

import (
	"fmt"
	"context"
	"strings"
	"time"
	"trustkeeper-go/app/service/account/pkg/model"
	"trustkeeper-go/app/service/account/pkg/repository"
	"trustkeeper-go/library/vault"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var conf Conf

type Conf struct {
	dbinfo string
	jwtkey string
}

// AccountService describes the service.
type AccountService interface {
	Create(ctx context.Context, email, password string) error
	Signin(ctx context.Context, email, password string) (string, error)
	Signout(ctx context.Context, token string) error
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

type Claims struct {
	jwt.StandardClaims
}

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
	tokenStr, err := token.SignedString([]byte(conf.jwtkey))
	if err != nil {
		return "", err
	}
	if err := b.repo.Update(acc, map[string]interface{}{"token_id": tokenID}); err != nil {
		return "", err
	}
	return tokenStr, e1
}


func (b *basicAccountService) Signout(ctx context.Context, token string) (e0 error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func (token *jwt.Token) (interface{}, error){
		return []byte(conf.jwtkey), nil
	})

	if err != nil  || !tkn.Valid{
		return fmt.Errorf(err.Error())
	}

	acc, err := b.repo.FindByTokenID(claims.Id)
	if err != nil {
		return fmt.Errorf("token was reset" + err.Error())
	}
	b.repo.Update(acc, map[string]interface{}{"token_id": nil})
	return nil
}
// NewBasicAccountService returns a naive, stateless implementation of AccountService.
func NewBasicAccountService() AccountService {
	return &basicAccountService{
		repo: repository.New(repository.DB(conf.dbinfo)),
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

func init() {
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
	jwtkey := data.Data["jwtkey"].(string)
	conf = Conf{
		dbinfo: dbInfo,
		jwtkey: jwtkey,
	}
}
