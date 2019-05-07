package repository

import (
  // "time"
  "regexp"
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

  row := sqlmock.NewRows([]string{"uuid", "email", "password"}).
    AddRow(uid.String(), email, string(hashedPassword))
  query := regexp.QuoteMeta(
    `INSERT INTO account ("uuid", "email", "password") VALUES ($1,$1,$3)`)
  s.mock.ExpectQuery(query).WithArgs(uid.String(), email, string(hashedPassword)).
    WillReturnRows(row)

  err = s.repo.Create(&model.Account{Email: email,
    UUID: uid.String(),
    Password: string(hashedPassword)})
  require.NoError(s.T(), err)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
  s.DB.Close()
}

func TestSuite(t *testing.T) {
  suite.Run(t, new(Suite))
}
