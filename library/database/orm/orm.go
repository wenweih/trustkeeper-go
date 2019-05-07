package orm

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// GormClient gorm client
type GormClient struct {
  *gorm.DB
}

// Connect gorm connect
func (gc *GormClient) Connect(dbInfo string)  {
  var err error
  gc.DB, err = gorm.Open("postgres", dbInfo)
  if err != nil {
    panic("failed to connect database:" + err.Error())
  }
}
