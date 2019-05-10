package repository

import (
  // "regexp"
  "testing"
  "database/sql"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
  "github.com/DATA-DOG/go-sqlmock"
  uuid "github.com/satori/go.uuid"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
  "trustkeeper-go/app/service/account/pkg/model"
)

// https://github.com/stretchr/testify/blob/master/suite/suite_test.go


type Suite struct {
  suite.Suite
  DB *gorm.DB
  mock sqlmock.Sqlmock

  repo AccoutRepo
  acc *model.Account
}

func (s *Suite) SetupSuite()  {
  var (
    db *sql.DB
    err error
  )
  db, s.mock, err = sqlmock.New()
  require.NoError(s.T(), err)
  s.DB, err = gorm.Open("postgres", db)
  require.NoError(s.T(), err)
  s.DB.LogMode(true)
  s.repo = New(s.DB)
}

func (s *Suite) TestCreate() {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testing"), bcrypt.DefaultCost)
  require.NoError(s.T(), err)
  var (
    uid = uuid.NewV4()
    email = "tesging@trustkeeper.io"
  )
  // query := regexp.QuoteMeta(
  //   `INSERT INTO ACCOUNTS (uuid,email,password) VALUES ($1,$1,$3)`)
  // prep := s.mock.ExpectPrepare(query)
  // prep.ExpectExec().WithArgs(uid.String(), email, string(hashedPassword)).
  //   WillReturnResult(sqlmock.NewResult(1,2))

  err = s.repo.Create(&model.Account{Email: email,
    UUID: uid.String(),
    Password: string(hashedPassword)})
  // assert.NoError(s.T(), err)
  require.NoError(s.T(), err)

  // s.mock.ExpectQuery(regexp.QuoteMeta(
  //   `SELECT uuid, email FROM "accounts" WHERE (uuid = $1)`)).
  //   WithArgs(uid.String()).
  //   // WillReturnResult(sqlmock.NewResult(1,1))
  //   WillReturnRows(sqlmock.NewRows([]string{"uid", "email"}).
  //   AddRow(uid.String(), email))
  //
  // err = s.mock.ExpectationsWereMet()
  // require.NoError(s.T(),err)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
  s.DB.Close()
}

func TestSuite(t *testing.T) {
  suite.Run(t, new(Suite))
}
