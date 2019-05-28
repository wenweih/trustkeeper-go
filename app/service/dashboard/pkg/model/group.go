package model

import (
  "github.com/jinzhu/gorm"
)

// Group group
type Group struct {
  gorm.Model
  Name      string    `grom:"type:varchar(100);not null;index"`
  Desc      string    `sql:"type:text;"`
  Key       string    `gorm:"type:varchar(255);index"`
  CreatorID string    `gorm:"index;not null"`
  Roles     []Role
}

type Role struct {
  gorm.Model
  GroupID     int
  Group       Group
  Name        string  `gorm:"not null;unique_index"`
  Memo        string  `gorm:"size:200"`
}
