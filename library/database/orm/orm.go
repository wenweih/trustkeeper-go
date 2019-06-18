package orm

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect gorm connect
func Connect(dbInfo string) (*gorm.DB, error)  {
  return gorm.Open("postgres", dbInfo)
}

// DB database connect
func DB(dbInfo string) (*gorm.DB, error) {
  db, err := connect(dbInfo)
  if err != nil {
    return nil, err
  }
  return db, nil
}

func connect(dbInfo string) (*gorm.DB, error)  {
  return gorm.Open("postgres", dbInfo)
}
