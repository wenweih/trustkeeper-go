package orm

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect gorm connect
func Connect(dbInfo string) (*gorm.DB, error)  {
  return gorm.Open("postgres", dbInfo)
}
