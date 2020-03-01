package model

import (
	"github.com/jinzhu/gorm"
)

// Group group
type Group struct {
	gorm.Model
	// UNIQUE constraint across multiple keys https://github.com/jinzhu/gorm/issues/961
	NamespaceID string `gorm:"unique_index:idx_namespace_id_name;not null"`
	Name        string `gorm:"unique_index:idx_namespace_id_name;not null"`
	Desc        string `sql:"type:text;"`
	CreatorID   string `gorm:"index;not null"`
	Chains      []Chain
}
