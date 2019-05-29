package repository

import (
  "trustkeeper-go/app/service/dashboard/pkg/model"
)

// Create save repo
func (repo DashboardRepo) Create(group *model.Group) error {
  return repo.db.Create(group).Error
}
