package repository

import (
  "time"
  "errors"
  "strconv"
  uuid "github.com/satori/go.uuid"
  "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
  "trustkeeper-go/app/service/account/pkg/model"
)

// Claims jwt clains struct
type Claims struct {
	jwt.StandardClaims
}

// IBiz repository bussiness logic
type IBiz interface {
  Signup(email, password, orgName string) (uuid string, namespaceID string, err error)
  Signin(email, password, jwtKey string) (token string, err error)
  Signout(tokenID string) error
  QueryRoles(tokenID string) (roles []string, err error)
  Auth(tokenID string) (accountuid string, namespaceid *uint, err error)
  Close() error
}

func (repo *repo)Close() error{
  return repo.close()
}

func (repo *repo) Signup(email, password, orgName string) (uid string, namespaceID string, err error) {
  // Salt and hash the password using the bcrypt algorithm
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
  uuid := uuid.NewV4()
  acc := &model.Account{
    Email:    email,
    Password: string(hashedPassword),
    UUID:     uuid.String()}
  namespace := &model.Namespace{
    CreatorUID: uuid.String(),
    Name: orgName,
  }
  tx := repo.db.Begin()
  if err := repo.iAccountRepo.Create(tx, acc).Error; err != nil {
    tx.Rollback()
    return "", "", err
  }
  if err := repo.iNamespaceRepo.Create(tx, namespace).Error; err != nil {
    tx.Rollback()
    return "", "", err
  }
  if err := tx.Commit().Error; err != nil {
    return "", "", err
  }

  namespaceID = strconv.FormatUint(uint64(namespace.ID), 10)
  repo.iAccountRepo.AddRoleForUserInDomain(acc.UUID, namespaceID, "admin")
  return uuid.String(), namespaceID, nil
}

func (repo *repo)Signin(email, password, jwtKey string) (string, error) {
  accounts, err :=repo.iAccountRepo.Query(repo.db, map[string]interface{}{
    "email": email,
  })
  if err != nil {
    return "", err
  }
  if len(accounts) > 1{
    return "", errors.New("database error")
  }
  acc := accounts[0]
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password)); err != nil {
		return "", errors.New("fail to login")
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
  tokenStr, err := token.SignedString([]byte(jwtKey))
  if err != nil {
    return "", err
  }
  ts := repo.db.Begin()
  if err := repo.iAccountRepo.Update(ts, acc, map[string]interface{}{"token_id": tokenID}); err != nil {
    return "", err
  }
  return tokenStr, nil
}

func (repo *repo)QueryRoles(tokenID string) ([]string, error) {
  return repo.Roles(repo.db, tokenID)
}

func (repo *repo)Signout(tokenID string) error {
  accounts, err := repo.iAccountRepo.Query(repo.db, map[string]interface{}{
    "token_id": tokenID})
  if err != nil {
    return err
  }

  if len(accounts) > 1{
    return errors.New("database error")
  }
  tx := repo.db.Begin()
  if err := repo.iAccountRepo.Update(tx, accounts[0], map[string]interface{}{"token_id": nil}); err != nil {
    return err
  }
  return tx.Commit().Error
}

func (repo *repo)Auth(tokenID string) (string, *uint, error) {
  accounts, err := repo.iAccountRepo.Query(repo.db, map[string]interface{}{
    "token_id": tokenID})
  if err != nil {
    return "", nil, err
  }

  if len(accounts) > 1{
    return "", nil, errors.New("database error")
  }
  return accounts[0].UUID, &accounts[0].Namespace.ID, nil
}
